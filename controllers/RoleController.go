package controllers

import (
	"PowerManage/models/class"
	"encoding/json"
)

type RoleController struct {
	class.Role
	MainController
}

type queryName struct {
	Name string
}

//@router /role/role_list [*]
func (c *RoleController)RoleList() {
	if LoginCheck(&c.MainController) {
		flag,list := class.RoleDataList()
		if flag {
			result := ResultJson(200)
			result["record"] = list
			c.Data["json"] = result
		}
		c.ServeJSON()
		return
	}
}

//@router /role/role_update [*]
func (c *RoleController)DoRoleUpdate() {
	if LoginCheck(&c.MainController) {
		json.Unmarshal(c.Ctx.Input.RequestBody, &c)
		if c.RoleUpdate() {
			c.Data["json"] = ResultJson(200)
		} else {
			c.Data["json"] = ResultJson(400)
		}
		c.ServeJSON()
		return
	}
}

//@router /role/role_add [*]
func (c *RoleController)DoRoleAdd() {
	if LoginCheck(&c.MainController) {
		json.Unmarshal(c.Ctx.Input.RequestBody, &c)
		flag := c.RoleAdd()
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
		}
		c.ServeJSON()
		return
	}
}

//@router /role/role_delete [*]
func (c *RoleController)DoRoleDelete() {
	if LoginCheck(&c.MainController) {
		json.Unmarshal(c.Ctx.Input.RequestBody,&c)
		if class.RoleDelete(c.Id) {
			c.Data["json"] = ResultJson(200)
		}else {
			c.Data["json"] = ResultJson(400)
		}
		c.ServeJSON()
		return
	}
}

//@router /role/role_query [*]
func (c *RoleController)DoRoleQuery() {
	if LoginCheck(&c.MainController) {
		var roleName queryName
		json.Unmarshal(c.Ctx.Input.RequestBody,&roleName)
		if roleName.Name == "" {
			flag,list := class.RoleDataList()
			if flag {
				result := ResultJson(200)
				result["record"] = list
				c.Data["json"] = result
			}else {
				c.Data["json"] = ResultJson(400)
			}
		}else {
			if role := class.RoleQuery(roleName.Name); role != nil{
				result := ResultJson(200)
				result["record"] = role
				c.Data["json"] = result
			}else {
				c.Data["json"] = ResultJson(400)
			}
		}
		c.ServeJSON()
		return
	}
}
