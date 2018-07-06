package controllers

import (
	"PowerManage/models/class"
	"encoding/json"
	"fmt"
)

type UserController struct {
	class.User
	MainController
}

type queryUser struct {
	Name string
	UserName string
}

//@router /user/user_list [*]
func (c *UserController)UserList() {
	if LoginCheck(&c.MainController) {
		flag,list := class.UserDataList()
		if flag {
			result := ResultJson(200)
			result["record"] = list
			c.Data["json"] = result
		}
		c.ServeJSON()
		return
	}
}

//@router /user/user_login_list [*]
func (c *UserController)UserLoginList() {
	if LoginCheck(&c.MainController) {
		flag,list := class.UserLoginDataList()
		if flag {
			result := ResultJson(200)
			result["record"] = list
			c.Data["json"] = result
		}
		c.ServeJSON()
		return
	}
}

//@router /user/user_login_query [*]
func (c *UserController)DoUserLoginQuery() {
	if LoginCheck(&c.MainController) {
		var user queryUser
		json.Unmarshal(c.Ctx.Input.RequestBody, &user)
		fmt.Println(user.Name)
		flag,userOne := class.QueryLoginUser(user.Name, user.UserName)
		if flag {
			result := ResultJson(200)
			result["record"] = userOne
			c.Data["json"] = result
		}else {
			result := ResultJson(400)
			c.Data["json"] = result
		}
		c.ServeJSON()
		return
	}
}

//@router /user/user_query [*]
func (c *UserController)DoUserQuery() {
	if LoginCheck(&c.MainController) {
		var user queryUser
		json.Unmarshal(c.Ctx.Input.RequestBody, &user)
		flag,userOne := class.QueryUser(user.Name, user.UserName)
		if flag {
			result := ResultJson(200)
			result["record"] = userOne
			c.Data["json"] = result
		}else {
			result := ResultJson(400)
			c.Data["json"] = result
		}
		c.ServeJSON()
		return
	}
}

//@router /user/user_delete [*]
func (c *UserController)DoUserDelete() {
	if LoginCheck(&c.MainController) {
		json.Unmarshal(c.Ctx.Input.RequestBody, &c)
		flag := class.DeleteUser(c.Id)
		if flag {
			result := ResultJson(200)
			c.Data["json"] = result
		}else {
			result := ResultJson(400)
			c.Data["json"] = result
		}
		c.ServeJSON()
		return
	}
}

//@router /user/user_add [*]
func (c *UserController)DoUserAdd() {
	if LoginCheck(&c.MainController) {
		var addUser class.AddUser
		json.Unmarshal(c.Ctx.Input.RequestBody, &addUser)
		flag := addUser.AddUser()
		switch flag {
		case 0:
			result := ResultJson(200)
			c.Data["json"] = result
			break
		case 1:
			result := ResultJson(201)
			c.Data["json"] = result
			break
		case 2:
			result := ResultJson(202)
			c.Data["json"] = result
			break
		case 3:
			result := ResultJson(203)
			c.Data["json"] = result
			break
		}
		c.ServeJSON()
		return
	}
}

//@router /user/user_one_get [*]
func (c *UserController)DoUserOneGet()  {
	//var user class.User
	if LoginCheck(&c.MainController) {
		username := c.GetSession("username")
		user,_ := class.UserOneByUserName(username.(string))
		result := ResultJson(200)
		result["record"] = user
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
}

//@router /user/user_one_update [*]
func (c *UserController)DoUserOneUpdate()  {
	//var user class.User
	if LoginCheck(&c.MainController) {
		json.Unmarshal(c.Ctx.Input.RequestBody, &c)
		flag := c.UserOneUpdate()
		if flag {
			result := ResultJson(200)
			c.Data["json"] = result
		}else {
			result := ResultJson(400)
			c.Data["json"] = result
		}
		c.ServeJSON()
		return
	}
}

//@router /user/user_update [*]
func (c *UserController)DoUserUpdate()  {
	//var user class.User
	if LoginCheck(&c.MainController) {
		var adduser class.AddUser
		json.Unmarshal(c.Ctx.Input.RequestBody, &adduser)
		flag := adduser.UpdateUser()
		if flag {
			result := ResultJson(200)
			c.Data["json"] = result
		}else {
			result := ResultJson(400)
			c.Data["json"] = result
		}
		c.ServeJSON()
		return
	}
}