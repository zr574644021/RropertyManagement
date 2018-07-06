package models

import (
	"github.com/astaxie/beego/orm"
	"PowerManage/models/class"
)

func init()  {
	orm.RegisterModel(new(class.User),new(class.Role), new(class.Resource), new(class.RoleResource), new(class.RoleUser))
}


