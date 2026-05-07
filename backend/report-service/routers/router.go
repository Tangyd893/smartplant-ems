package routers

import (
<<<<<<< HEAD
	"smartplant-ems/common/filters"
	"smartplant-ems/report-service/controllers"

=======
	"smartplant-ems/report-service/controllers"
>>>>>>> 395aa7cad5570b4b699b4c029768d04fab946652
	"github.com/beego/beego/v2/server/web"
)

func init() {
<<<<<<< HEAD
	web.InsertFilter("*", web.BeforeRouter, filters.CORS())
	web.InsertFilter("/api/*", web.BeforeRouter, filters.Logger())

=======
>>>>>>> 395aa7cad5570b4b699b4c029768d04fab946652
	reportNs := web.NewNamespace("/api/report",
		web.NSRouter("/list", &controllers.ReportController{}, "Get:List"),
		web.NSRouter("/:id/download", &controllers.ReportController{}, "Get:Download"),
		web.NSRouter("/:id", &controllers.ReportController{}, "Get:Get;Delete:Delete"),
		web.NSRouter("", &controllers.ReportController{}, "Post:Create"),
	)

	web.AddNamespace(reportNs)
}
