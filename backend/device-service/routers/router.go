package routers

import (
	"smartplant-ems/device-service/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	deviceNs := web.NewNamespace("/api/device",
		web.NSRouter("/list", &controllers.DeviceController{}, "Get:List"),
		web.NSRouter("/:id", &controllers.DeviceController{}, "Get:Get,Put:Update,Delete:Delete"),
		web.NSRouter("", &controllers.DeviceController{}, "Post:Post"),
		web.NSRouter("/:id/realtime", &controllers.DeviceController{}, "Get:GetRealTimeData"),
	)

	web.AddNamespace(deviceNs)
}