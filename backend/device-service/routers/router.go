package routers

import (
	"smartplant-ems/common/filters"
	"smartplant-ems/device-service/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.InsertFilter("*", web.BeforeRouter, filters.CORS())
	web.InsertFilter("/api/*", web.BeforeRouter, filters.Logger())

	authNs := web.NewNamespace("/api/auth",
		web.NSRouter("/login", &controllers.AuthController{}, "Post:Login"),
		web.NSRouter("/logout", &controllers.AuthController{}, "Post:Logout"),
		web.NSRouter("/me", &controllers.AuthController{}, "Get:GetCurrentUser"),
	)

	deviceNs := web.NewNamespace("/api/device",
		web.NSRouter("/list", &controllers.DeviceController{}, "Get:List"),
		web.NSRouter("/:id", &controllers.DeviceController{}, "Get:Get;Put:Put;Delete:Delete"),
		web.NSRouter("", &controllers.DeviceController{}, "Post:Post"),
		web.NSRouter("/:id/realtime", &controllers.DeviceController{}, "Get:GetRealTimeData"),
	)

	web.AddNamespace(authNs)
	web.AddNamespace(deviceNs)
}
