package class

import (
	//"fmt"
	"github.com/astaxie/beego/orm"
	"fmt"
	"github.com/astaxie/beego"
)

type Resource struct {
	Id 		int
	Name 	string
	Parent *Resource 	`orm:"null;rel(fk)"`
	Rtype   int
	Sons   	[]*Resource `orm:"reverse(many)"` // fk 的反向关系
	SonNum          int                `orm:"-"`
	Url 	string
	Level           int                `orm:"-"`             //第几级，从0开始
	RoleResource []*RoleResource `orm:"reverse(many)"`

}


//根据用户获取有权管理的资源列表
func ResourceTreeByUserId(id int)  []*Resource{
	o := orm.NewOrm()
	var list []*Resource

	sql := "SELECT DISTINCT T0.resource_id,T2.id,T2.name,T2.parent_id,T2.rtype,T2.url "+
		"FROM role_resource AS T0 INNER JOIN role_user AS T1 ON T0.role_id = T1.role_id "+
		"INNER JOIN resource AS T2 ON T2.id = T0.resource_id WHERE T1.user_id = ? Order By T2.id asc"
	o.Raw(sql,id).QueryRows(&list)
	result := resourceList2TreeGrid(list)
	return result

}


//获取所有资源
func ResourceAllTree() []*Resource{
	o := orm.NewOrm()
	var list []*Resource
	if _,err := o.QueryTable("resource").All(&list); err != nil {
		beego.Error(" all tree err ",err)
		return nil
	}
	result := resourceList2TreeGrid(list)
	return result
}

//将资源列表转成treegrid格式
func resourceList2TreeGrid(list []*Resource) []*Resource {
	result := make([]*Resource, 0)
	for _, item := range list {
		if item.Parent == nil || item.Parent.Id == 0 {
			item.Level = 0
			result = append(result, item)
			result = resourceAddSons(item, list, result)
		}
		fmt.Println("1:"+item.Name)
	}
	return result
}

//resourceAddSons 添加子菜单
func resourceAddSons(cur *Resource, list, result []*Resource) []*Resource {
	for _, item := range list {
		if item.Parent != nil && item.Parent.Id == cur.Id {
			cur.SonNum++
			item.Level = cur.Level + 1
			result = append(result, item)
			result = resourceAddSons(item, list, result)
		}
	}
	return result
}

//根据id获取资源
func ResourceByUserId(id int)  *Resource{
	o := orm.NewOrm()
	var resource Resource
	if err := o.QueryTable("resource").Filter("Id",id).One(&resource);
		err != nil {
		return nil
	}else {
		return &resource
	}
}


