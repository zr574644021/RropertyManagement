package class

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"time"
)

type User struct {
	Id 			int
	Name 		string
	UserName 	string		`orm:"unique"`
	PassWord 	string
	Message 	string
	OnlineState string							//在线状态
	LoginTime 	string							//上线时间
	DownTime 	string							//下线时间
	LoginIp		string							//登录Ip
	RoleUser []*RoleUser	`orm:"reverse(many)"`
}

type AddUser struct{
	User
	OptionsValues []int
}

//登录
func Login(username , password, loginIp string) (bool,*User){
	o := orm.NewOrm()
	var user User
	if err := o.QueryTable("user").Filter("user_name",username).One(&user); err != nil {
		beego.Error("username error is ",err)
		return false, nil
	}
	if password == user.PassWord {
		LoginTime(username, loginIp)
		return true, &user
	}else {
		return false, nil
	}
}

//登录记录
func LoginTime(username, loginIp string) {
	o := orm.NewOrm()
	nowTime := time.Now().Format("2006-01-02 03:04:05 PM")
	if _, err := o.QueryTable("user").Filter("user_name", username).
		Update(orm.Params{"login_time": nowTime,
		"login_ip": loginIp, "online_state": "在线"}); err != nil {
		beego.Error("username logintime error is ",err)
	}
	return
}

//下线记录
func LoginDown(username string) {
	o := orm.NewOrm()
	nowTime := time.Now().Format("2006-01-02 03:04:05 PM")
	if _, err := o.QueryTable("user").Filter("user_name", username).
		Update(orm.Params{"down_time": nowTime,
		"online_state": "离线"}); err != nil {
		beego.Error("username downtime error is ",err)
	}
	return
}

//新增用户
func (c *AddUser)AddUser() int{
	o := orm.NewOrm()
	if nameC, _ := UserOneByName(c.Name); nameC == nil {
		if usernameC, _ := UserOneByUserName(c.UserName); usernameC == nil {
			var user *User
			user = &c.User
			if id, err := o.Insert(user); err != nil {
				beego.Error("add user error ",err)
				return 3
			}else {
				AddRoleUser(int(id), c.OptionsValues)
			}
			return 0
		}else {
			return 2
		}
	}else {
		return 1
	}
}

//编辑用户
func (c *AddUser)UpdateUser() bool{
	o := orm.NewOrm()
	var user *User
	user = &c.User
	if _, err := o.Update(user); err != nil {
		beego.Error("updateUser error ",err)
		return false
	}else {
		DeleteRoleUser(c.Id)
		AddRoleUser(int(c.Id), c.OptionsValues)
		return true
	}
}

//获取用户列表
func UserDataList() (bool, []User) {
	o := orm.NewOrm()
	var users []User
	if  _,err := o.QueryTable("user").All(&users); err != nil {
		return false, nil
	}else {
		//var a error
		for i:=0 ; i<len(users) ; i++  {
			_,err := o.LoadRelated(&users[i],"RoleUser")
			if err != nil {
				break
			}
			for j:=0; j<len(users[i].RoleUser); j++ {
				users[i].RoleUser[j].Role = RoleByUserId(users[i].RoleUser[j].Role.Id)
			}
		}
		return true, users
	}
}

//获取在线用户列表
func UserLoginDataList() (bool, []User) {
	o := orm.NewOrm()
	var users []User
	if  _,err := o.QueryTable("user").Filter("online_state","在线").All(&users); err != nil {
		return false, nil
	}else {
		return true, users
	}
}

//删除用户
func DeleteUser(id int) bool{
	o := orm.NewOrm()
	if _,err := o.QueryTable("user").Filter("id",id).Delete(); err != nil {
		beego.Error("delete user error is", err)
		return false
	}else {
		DeleteRoleUser(id)
		return true
	}
}

//查询用户
func QueryUser(name, username string) (bool, []User){
	o := orm.NewOrm()
	var user []User
	var err error
	if name != "" && username == "" {
		err = o.QueryTable("user").Filter("name",name).One(&user)
	} else if name == "" && username != "" {
		err = o.QueryTable("user").Filter("user_name",username).One(&user)
	} else if name != "" && username != ""  {
		err = o.QueryTable("user").Filter("user_name",username).Filter("name",name).One(&user)
	} else {
		var flag bool
		flag,user = UserDataList()
		return flag, user
	}
	if err != nil {
		beego.Error("query user error is ",err)
		return false, nil
	}else {
		_,err := o.LoadRelated(&user[0],"RoleUser")
		if err != nil {
			return false, nil
		}
		for j:=0; j<len(user[0].RoleUser); j++ {
			user[0].RoleUser[j].Role = RoleByUserId(user[0].RoleUser[j].Role.Id)
		}
		return true, user
	}
}

//查询在线用户
func QueryLoginUser(name, username string) (bool, []User){
	o := orm.NewOrm()
	var user []User
	var err error
	if name != "" && username == "" {
		err = o.QueryTable("user").Filter("name",name).Filter("online_state","在线").One(&user)
	} else if name == "" && username != "" {
		err = o.QueryTable("user").Filter("user_name",username).Filter("online_state","在线").One(&user)
	} else if name != "" && username != ""  {
		err = o.QueryTable("user").Filter("user_name",username).Filter("name",name).Filter("online_state","在线").One(&user)
	} else {
		var flag bool
		flag,user = UserLoginDataList()
		return flag, user
	}
	if err != nil {
		beego.Error("query user error is ",err)
		return false, nil
	}else {
		return true, user
	}
}

//更新个人信息
func (c *User)UserOneUpdate()  bool{
	o := orm.NewOrm()
	if _, err := UserOneByUserName(c.UserName); err != nil {
		beego.Error("updateUser get role error ",err)
		return false
	}else {
		if _, err := o.Update(c); err != nil {
			beego.Error("updateUser error ",err)
			return false
		}
	}
	return true
}

func UserOneByName(name string) (*User, error) {
	o := orm.NewOrm()
	var m User
	err := o.QueryTable("user").Filter("name",name).One(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func UserOneByUserName(username string) (*User, error) {
	o := orm.NewOrm()
	var m User
	err := o.QueryTable("user").Filter("user_name",username).One(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}