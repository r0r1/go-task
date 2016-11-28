package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["go-task/controllers:TaskController"] = append(beego.GlobalControllerRouter["go-task/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["go-task/controllers:TaskController"] = append(beego.GlobalControllerRouter["go-task/controllers:TaskController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["go-task/controllers:TaskController"] = append(beego.GlobalControllerRouter["go-task/controllers:TaskController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["go-task/controllers:TaskController"] = append(beego.GlobalControllerRouter["go-task/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["go-task/controllers:TaskController"] = append(beego.GlobalControllerRouter["go-task/controllers:TaskController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["go-task/controllers:TaskStatusController"] = append(beego.GlobalControllerRouter["go-task/controllers:TaskStatusController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["go-task/controllers:TaskStatusController"] = append(beego.GlobalControllerRouter["go-task/controllers:TaskStatusController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["go-task/controllers:TaskStatusController"] = append(beego.GlobalControllerRouter["go-task/controllers:TaskStatusController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["go-task/controllers:TaskStatusController"] = append(beego.GlobalControllerRouter["go-task/controllers:TaskStatusController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["go-task/controllers:TaskStatusController"] = append(beego.GlobalControllerRouter["go-task/controllers:TaskStatusController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params: nil})

	beego.GlobalControllerRouter["go-task/controllers:UserController"] = append(beego.GlobalControllerRouter["go-task/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
