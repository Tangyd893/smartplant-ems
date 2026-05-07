package routers

import (
	"smartplant-ems/common/filters"
	"smartplant-ems/report-service/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	web.InsertFilter("*", web.BeforeRouter, filters.CORS())
	web.InsertFilter("/api/*", web.BeforeRouter, filters.Logger())

	reportNs := web.NewNamespace("/api/report",
		web.NSRouter("/list", &controllers.ReportController{}, "Get:List"),
		web.NSRouter("/:id/download", &controllers.ReportController{}, "Get:Download"),
		web.NSRouter("/:id", &controllers.ReportController{}, "Get:Get;Delete:Delete"),
		web.NSRouter("", &controllers.ReportController{}, "Post:Create"),
	)

	web.AddNamespace(reportNs)
}
