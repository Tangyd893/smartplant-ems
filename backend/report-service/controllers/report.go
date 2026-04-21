package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/beego/beego/v2/server/web"
)

type ReportController struct {
	web.Controller
}

type Report struct {
	ID          int64     `json:"id"`
	ReportName  string    `json:"report_name"`
	ReportType  string    `json:"report_type"`
	PeriodType  string    `json:"period_type"`
	PeriodStart time.Time `json:"period_start"`
	PeriodEnd   time.Time `json:"period_end"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

type PageResult struct {
	Records interface{} `json:"records"`
	Total   int64       `json:"total"`
	Page    int         `json:"page"`
	Size    int         `json:"size"`
}

type APIResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func (c *ReportController) List() {
	page, _ := strconv.Atoi(c.GetString("page", "1"))
	size, _ := strconv.Atoi(c.GetString("size", "20"))
	reportType := c.GetString("report_type")

	fmt.Printf("List reports: page=%d, size=%d, type=%s\n", page, size, reportType)

	records := []Report{
		{ID: 1, ReportName: "日能耗报表-2026-04-20", ReportType: "daily", PeriodType: "daily", PeriodStart: time.Now().AddDate(0, 0, -1), PeriodEnd: time.Now(), Status: 1},
		{ID: 2, ReportName: "月能耗报表-2026-04", ReportType: "monthly", PeriodType: "monthly", PeriodStart: time.Date(2026, 4, 1, 0, 0, 0, 0, time.Local), PeriodEnd: time.Now(), Status: 1},
	}

	c.Data["json"] = APIResponse{
		Code: 0,
		Msg:  "success",
		Data: PageResult{
			Records: records,
			Total:   2,
			Page:    page,
			Size:    size,
		},
	}
	c.ServeJSON()
}

func (c *ReportController) Get() {
	id, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)
	fmt.Printf("Get report: id=%d\n", id)

	c.Data["json"] = APIResponse{
		Code: 0,
		Msg:  "success",
		Data: Report{ID: id, ReportName: "日能耗报表-2026-04-20", ReportType: "daily", Status: 1},
	}
	c.ServeJSON()
}

func (c *ReportController) Create() {
	var req map[string]interface{}
	if err := c.ParseForm(&req); err != nil {
		c.Data["json"] = APIResponse{Code: 400, Msg: "参数错误"}
		c.ServeJSON()
		return
	}

	fmt.Printf("Create report: %+v\n", req)

	c.Data["json"] = APIResponse{
		Code: 0,
		Msg:  "报表生成任务已创建",
		Data: map[string]interface{}{"id": 3},
	}
	c.ServeJSON()
}

func (c *ReportController) Download() {
	id, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)
	fmt.Printf("Download report: id=%d\n", id)

	c.Data["json"] = APIResponse{
		Code: 0,
		Msg:  "success",
		Data: map[string]interface{}{
			"download_url": "/api/report/download?id=" + strconv.FormatInt(id, 10),
			"filename":     "report_" + strconv.FormatInt(id, 10) + ".xlsx",
		},
	}
	c.ServeJSON()
}

func (c *ReportController) Delete() {
	id, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)
	fmt.Printf("Delete report: id=%d\n", id)

	c.Data["json"] = APIResponse{Code: 0, Msg: "删除成功"}
	c.ServeJSON()
}