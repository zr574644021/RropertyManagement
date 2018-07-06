package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"PowerManage/models/class"
)

type MainController struct {
	beego.Controller
}

type UserLogin struct {
	Username string
	Password string
}

//@router /login [*]
func (c *MainController) DoLogin() {
	var login UserLogin
	//var user User
	json.Unmarshal(c.Ctx.Input.RequestBody,&login)
	username := login.Username
	password := login.Password
	loginIp := c.Ctx.Request.RemoteAddr
	if flag,user := class.Login(username, password, loginIp); !flag && user == nil {
		c.Data["json"] = ResultJson(4000)//账号或密码错误
	}else {
		c.SetSession("user_id",user.Id)
		c.SetSession("name",user.Name)
		c.SetSession("username",username)
		class.ResourceTreeByUserId(user.Id)
		c.Data["json"] = ResultJson(4002)//登录成功
	}
	c.ServeJSON()
	return
}

//@router /login_out [*]
func (c *MainController) DoLogout() {
	user_id := c.GetSession("user_id")
	username := c.GetSession("username")

	if username == nil || user_id == nil {
		c.Data["json"] = ResultJson(4001)//未登录
	} else {
		c.DelSession("username")
		c.DelSession("password")
		c.Data["json"] = ResultJson(4003)//注销成功
		class.LoginDown(username.(string))
	}
	c.ServeJSON()
	return
}


