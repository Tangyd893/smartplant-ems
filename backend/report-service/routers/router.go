package routers

import (
	"smartplant-ems/report-service/controllers"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	reportNs := web.NewNamespace("/api/report",
		web.NSRouter("/list", &controllers.ReportController{}, "Get:List"),
		web.NSRouter("/:id", &controllers.ReportController{}, "Get:Get"),
		web.NSRouter("", &controllers.ReportController{}, "Post:Create"),
		web.NSRouter("/:id/download", &controllers.ReportController{}, "Get:Download"),
		web.NSRouter("/:id", &controllers.ReportController{}, "Delete:Delete"),
	)

	web.AddNamespace(reportNs)
}