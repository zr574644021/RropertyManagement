package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["PowerManage/controllers:MainController"] = append(beego.GlobalControllerRouter["PowerManage/controllers:MainController"],
		beego.ControllerComments{
			Method: "DoLogin",
			Router: `/login`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["PowerManage/controllers:MainController"] = append(beego.GlobalControllerRouter["PowerManage/controllers:MainController"],
		beego.ControllerComments{
			Method: "DoLogout",
			Router: `/login_out`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["PowerManage/controllers:ResourceController"] = append(beego.GlobalControllerRouter["PowerManage/controllers:ResourceController"],
		beego.ControllerComments{
			Method: "ResourceList",
			Router: `/resource/resource_list`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["PowerManage/controllers:ResourceController"] = append(beego.GlobalControllerRouter["PowerManage/controllers:ResourceController"],
		beego.ControllerComments{
			Method: "DoUserResource",
			Router: `/resource/user_resource`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["PowerManage/controllers:RoleController"] = append(beego.GlobalControllerRouter["PowerManage/controllers:RoleController"],
		beego.ControllerComments{
			Method: "DoRoleAdd",
			Router: `/role/role_add`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["PowerManage/controllers:RoleController"] = append(beego.GlobalControllerRouter["PowerManage/controllers:RoleController"],
		beego.ControllerComments{
			Method: "DoRoleDelete",
			Router: `/role/role_delete`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["PowerManage/controllers:RoleController"] = append(beego.GlobalControllerRouter["PowerManage/controllers:RoleController"],
		beego.ControllerComments{
			Method: "RoleList",
			Router: `/role/role_list`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["PowerManage/controllers:RoleController"] = append(beego.GlobalControllerRouter["PowerManage/controllers:RoleController"],
		beego.ControllerComments{
			Method: "DoRoleQuery",
			Router: `/role/role_query`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["PowerManage/controllers:RoleController"] = append(beego.GlobalControllerRouter["PowerManage/controllers:RoleController"],
		beego.ControllerComments{
			Method: "DoRoleUpdate",
			Router: `/role/role_update`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["PowerManage/controllers:UserController"] = append(beego.GlobalControllerRouter["PowerManage/controllers:UserController"],
		beego.ControllerComments{
			Method: "DoUserAdd",
			Router: `/user/user_add`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["PowerManage/controllers:UserController"] = append(beego.GlobalControllerRouter["PowerManage/controllers:UserController"],
		beego.ControllerComments{
			Method: "DoUserDelete",
			Router: `/user/user_delete`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["PowerManage/controllers:UserController"] = append(beego.GlobalControllerRouter["PowerManage/controllers:UserController"],
		beego.ControllerComments{
			Method: "UserList",
			Router: `/user/user_list`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["PowerManage/controllers:UserController"] = append(beego.GlobalControllerRouter["PowerManage/controllers:UserController"],
		beego.ControllerComments{
			Method: "UserLoginList",
			Router: `/user/user_login_list`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["PowerManage/controllers:UserController"] = append(beego.GlobalControllerRouter["PowerManage/controllers:UserController"],
		beego.ControllerComments{
			Method: "DoUserLoginQuery",
			Router: `/user/user_login_query`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["PowerManage/controllers:UserController"] = append(beego.GlobalControllerRouter["PowerManage/controllers:UserController"],
		beego.ControllerComments{
			Method: "DoUserOneGet",
			Router: `/user/user_one_get`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["PowerManage/controllers:UserController"] = append(beego.GlobalControllerRouter["PowerManage/controllers:UserController"],
		beego.ControllerComments{
			Method: "DoUserOneUpdate",
			Router: `/user/user_one_update`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["PowerManage/controllers:UserController"] = append(beego.GlobalControllerRouter["PowerManage/controllers:UserController"],
		beego.ControllerComments{
			Method: "DoUserQuery",
			Router: `/user/user_query`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["PowerManage/controllers:UserController"] = append(beego.GlobalControllerRouter["PowerManage/controllers:UserController"],
		beego.ControllerComments{
			Method: "DoUserUpdate",
			Router: `/user/user_update`,
			AllowHTTPMethods: []string{"*"},
			MethodParams: param.Make(),
			Params: nil})

}
