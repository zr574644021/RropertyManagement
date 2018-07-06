package class

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type Role struct {
	Id int
	Name string
	RoleUser []*RoleUser 			`orm:"reverse(many)"`
	RoleResource []*RoleResource 	`orm:"reverse(many)"`
}

//根据用户id获取角色
func RoleByUserId(id int) *Role {
	o := orm.NewOrm()
	var role Role
	if err := o.QueryTable("role").Filter("Id",id).One(&role);
	err != nil {
		return nil
	}else {
		return &role
	}
}

//获取角色列表
func RoleDataList() (bool,[]Role){
	o := orm.NewOrm()
	var roles []Role
	if  _,err := o.QueryTable("role").All(&roles); err != nil {
		return false, nil
	}else {
		for i := 0; i<len(roles) ; i++  {
			_,err := o.LoadRelated(&roles[i],"RoleResource")
			if err != nil {
				break
			}
			for j:=0; j<len(roles[i].RoleResource); j++ {
				roles[i].RoleResource[j].Resource = ResourceByUserId(roles[i].RoleResource[j].Resource.Id)
			}
		}
		return true, roles
	}
}

//查询
func RoleQuery(name string) []Role{
	o := orm.NewOrm()
	var role []Role
	if  _, err := o.QueryTable("role").Filter("Name",name).All(&role); err != nil {
		return nil
	}else {
		for j:=0; j<len(role[0].RoleResource); j++ {
			role[0].RoleResource[j].Resource = ResourceByUserId(role[0].RoleResource[j].Resource.Id)
		}
		return role
	}
}

//新增
func (c *Role)RoleAdd() int{
	o := orm.NewOrm()
	if nameC, _ := RoleOneByName(c.Name); nameC == nil {
			if id, err := o.Insert(c); err != nil {
				beego.Error("add user error ",err)
				return 2
			}else {
				AddRoleResource(int(id), c.RoleResource)
			}
			return 0
	}else {
		return 1
	}
}

//修改
func (c *Role)RoleUpdate()  bool{
	o := orm.NewOrm()
	if _, err := RoleOne(c.Id); err != nil {
		beego.Error("updateRole get role error ",err)
		return false
	}else {
		if _, err := o.Update(c); err != nil {
			beego.Error("updateRole error ",err)
			return false
		}
	}
	err := DeleyedRoleOneResource(c.Id)
	AddRoleResource(int(c.Id), c.RoleResource)
	if err != nil {
		return false
	}
	return true
}

//删除
func RoleDelete(id int) bool {
	o := orm.NewOrm()
	var roleUser RoleUser
	if err := o.QueryTable("role_user").Filter("user_id",id).One(&roleUser);
	err == nil && roleUser.Id > 0 {
		return false
	}else {
		if _, err := o.QueryTable("role").Filter("id",id).Delete(); err != nil {
			beego.Error("delete role error ",err)
			return false
		}else {
			DeleteRoleResource(id)
			return true
		}
	}
}

func RoleOneByName(name string) (*Role, error) {
	o := orm.NewOrm()
	var m Role
	err := o.QueryTable("role").Filter("name",name).One(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

func RoleOne(id int) (*Role, error) {
	o := orm.NewOrm()
	m := Role{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}