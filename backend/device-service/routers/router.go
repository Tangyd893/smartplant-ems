package routers

import (
<<<<<<< HEAD
	"smartplant-ems/common/filters"
	"smartplant-ems/device-service/controllers"

=======
	"smartplant-ems/device-service/controllers"
>>>>>>> 395aa7cad5570b4b699b4c029768d04fab946652
	"github.com/beego/beego/v2/server/web"
)

func init() {
<<<<<<< HEAD
	web.InsertFilter("*", web.BeforeRouter, filters.CORS())
	web.InsertFilter("/api/*", web.BeforeRouter, filters.Logger())

	authNs := web.NewNamespace("/api/auth",
		web.NSRouter("/login", &controllers.AuthController{}, "Post:Login"),
		web.NSRouter("/logout", &controllers.AuthController{}, "Post:Logout"),
		web.NSRouter("/me", &controllers.AuthController{}, "Get:GetCurrentUser"),
	)

=======
>>>>>>> 395aa7cad5570b4b699b4c029768d04fab946652
	deviceNs := web.NewNamespace("/api/device",
		web.NSRouter("/list", &controllers.DeviceController{}, "Get:List"),
		web.NSRouter("/:id", &controllers.DeviceController{}, "Get:Get;Put:Put;Delete:Delete"),
		web.NSRouter("", &controllers.DeviceController{}, "Post:Post"),
		web.NSRouter("/:id/realtime", &controllers.DeviceController{}, "Get:GetRealTimeData"),
	)

<<<<<<< HEAD
	web.AddNamespace(authNs)
=======
>>>>>>> 395aa7cad5570b4b699b4c029768d04fab946652
	web.AddNamespace(deviceNs)
}
