package controllers

func ResultJson(code int)  (result map[string]interface{}){
	result = make(map[string]interface{})
	result["status"] = code
	return result
}

func LoginCheck(c *MainController) bool{
	username := c.GetSession("username")
	user_id := c.GetSession("user_id")
	if username == nil || user_id == nil {
		c.Data["json"] = ResultJson(4001)//未登录
		c.ServeJSON()
		return false
	}else {
		return true
	}
}
