package controllers

import (
	"fmt"
	"strconv"
	"github.com/beego/beego/v2/server/web"
)

type ReportController struct {
	web.Controller
}

func (c *ReportController) List() {
	page, _ := strconv.Atoi(c.GetString("page", "1"))
	size, _ := strconv.Atoi(c.GetString("size", "20"))

	reports := []map[string]interface{}{
		{"id": 1, "report_name": "日能耗报表-2026-04-20", "report_type": "daily", "period_type": "daily", "status": 1, "created_at": "2026-04-20T18:00:00+08:00"},
		{"id": 2, "report_name": "月能耗报表-2026-03", "report_type": "monthly", "period_type": "monthly", "status": 1, "created_at": "2026-04-01T09:00:00+08:00"},
	}

	c.Data["json"] = map[string]interface{}{
		"code": 0,
		"msg":  "success",
		"data": map[string]interface{}{
			"records": reports,
			"total":   2,
			"page":    page,
			"size":    size,
		},
	}
	c.ServeJSON()
}

func (c *ReportController) Get() {
	id := c.Ctx.Input.Param(":id")
	report := map[string]interface{}{
		"id": 1, "report_name": "日能耗报表-2026-04-20", "report_type": "daily",
		"period_type": "daily", "period_start": "2026-04-20", "period_end": "2026-04-20",
		"status": 1, "file_path": "/reports/daily_2026-04-20.xlsx",
		"created_at": "2026-04-20T18:00:00+08:00",
	}
	_ = id
	c.Data["json"] = map[string]interface{}{"code": 0, "msg": "success", "data": report}
	c.ServeJSON()
}

func (c *ReportController) Create() {
	var req map[string]interface{}
	c.ParseForm(&req)
	fmt.Printf("创建报表: %+v\n", req)
	c.Data["json"] = map[string]interface{}{"code": 0, "msg": "创建成功", "data": map[string]interface{}{"id": 5}}
	c.ServeJSON()
}

func (c *ReportController) Download() {
	id := c.Ctx.Input.Param(":id")
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Ctx.ResponseWriter.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=report_%s.xlsx", id))
	c.Data["json"] = map[string]interface{}{"code": 0, "msg": "下载成功"}
	c.ServeJSON()
}

func (c *ReportController) Delete() {
	id := c.Ctx.Input.Param(":id")
	fmt.Printf("删除报表: id=%s\n", id)
	c.Data["json"] = map[string]interface{}{"code": 0, "msg": "删除成功"}
	c.ServeJSON()
}
