package routers

import (
	"smartplant-ems/energy-service/controllers"
	"github.com/beego/beego/v2/server/web"
)

func init() {
	energyNs := web.NewNamespace("/api/energy",
		web.NSRouter("/list", &controllers.EnergyController{}, "Get:List"),
		web.NSRouter("/stats", &controllers.EnergyController{}, "Get:GetStats"),
		web.NSRouter("/:id", &controllers.EnergyController{}, "Get:Get;Put:Update;Delete:Delete"),
		web.NSRouter("", &controllers.EnergyController{}, "Post:Create"),
	)

	web.AddNamespace(energyNs)
}
