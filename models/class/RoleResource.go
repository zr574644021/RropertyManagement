package class

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
)

type RoleResource struct {
	Id int
	Role *Role `orm:"rel(fk)"`
	Resource *Resource `orm:"rel(fk)"`
}

//删除角色时删除所有对应权限
func DeleteRoleResource(id int) {
	o := orm.NewOrm()
	if _,err := o.QueryTable("role_resource").Filter("role_id",id).Delete(); err != nil {
		beego.Error("delete all resource is ",err)
	}
	return
}

//新增角色资源
func AddRoleResource(id int,resourceId []*RoleResource) {
	o := orm.NewOrm()
	m := Role{Id : id}
	var relations []RoleResource
	for _,item := range resourceId {
		r := Resource{Id: item.Resource.Id}
		relation := RoleResource{Role:&m, Resource: &r}
		relations = append(relations, relation)
	}
	if _, err := o.InsertMulti(len(relations), relations); err != nil {
		return
	}
}

//删除指定用户资源
func DeleyedRoleOneResource(id int) error{
	o := orm.NewOrm()
	if _,err := o.QueryTable("role_resource").Filter("user_id",id).Delete(); err != nil {
		beego.Error("delete one role resource error " ,err)
		return err
	}
	return nil
}