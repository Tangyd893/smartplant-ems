package routers

import (
<<<<<<< HEAD
	"smartplant-ems/common/filters"
	"smartplant-ems/energy-service/controllers"

=======
	"smartplant-ems/energy-service/controllers"
>>>>>>> 395aa7cad5570b4b699b4c029768d04fab946652
	"github.com/beego/beego/v2/server/web"
)

func init() {
<<<<<<< HEAD
	web.InsertFilter("*", web.BeforeRouter, filters.CORS())
	web.InsertFilter("/api/*", web.BeforeRouter, filters.Logger())

=======
>>>>>>> 395aa7cad5570b4b699b4c029768d04fab946652
	energyNs := web.NewNamespace("/api/energy",
		web.NSRouter("/list", &controllers.EnergyController{}, "Get:List"),
		web.NSRouter("/stats", &controllers.EnergyController{}, "Get:GetStats"),
		web.NSRouter("/:id", &controllers.EnergyController{}, "Get:Get;Put:Update;Delete:Delete"),
		web.NSRouter("", &controllers.EnergyController{}, "Post:Create"),
	)

	web.AddNamespace(energyNs)
}
