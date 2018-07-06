package controllers

import "PowerManage/models/class"

type ResourceController struct {
	class.Resource
	MainController
}

//@router /resource/user_resource
func (c *ResourceController)DoUserResource() {
	if LoginCheck(&c.MainController) {
		id := c.GetSession("user_id")
		list := class.ResourceTreeByUserId(id.(int))
		result := ResultJson(200)
		result["record"] = list
		c.Data["json"] = result
		c.ServeJSON()
	}
}

//@router /resource/resource_list
func (c *ResourceController)ResourceList()  {
	if LoginCheck(&c.MainController) {
		list := class.ResourceAllTree()
		if list != nil {
			result := ResultJson(200)
			result["record"] = list
			c.Data["json"] = result
		}
		c.ServeJSON()
		return
	}
}


