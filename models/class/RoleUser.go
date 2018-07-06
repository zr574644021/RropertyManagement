package class

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type RoleUser struct {
	Id int
	User *User `orm:"rel(fk)"`
	Role *Role `orm:"rel(fk)"`
}

/*//获取所有用户角色列表
func RoleUserDataList() (bool,[]RoleUser){
	o := orm.NewOrm()
	var roles []RoleUser
	if  _,err := o.QueryTable("role_user").RelatedSel().All(&roles); err != nil {
		return false, nil
	}else {
		return true, roles
	}
}*/

//删除用户时删除所有角色关系
func DeleteRoleUser(id int) {
	o := orm.NewOrm()
	if _,err := o.QueryTable("role_user").Filter("user_id",id).Delete(); err != nil {
		beego.Error("delete all role is ",err)
	}
	return
}


//新增用户时新增角色关系
func AddRoleUser(id int,roleId []int)  {
	o := orm.NewOrm()
	m := User{Id : id}
	var relations []RoleUser
	for _,item := range roleId {
		r := Role{Id: item}
		relation := RoleUser{User:&m, Role: &r}
		relations = append(relations, relation)
		//fmt.Println(len(relations))

	}
	if _, err := o.InsertMulti(len(relations), relations); err != nil {
		return
	}
}
