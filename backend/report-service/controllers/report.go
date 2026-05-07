package controllers

import (
	"fmt"
	"strconv"
<<<<<<< HEAD
	"strings"
	"time"

	"smartplant-ems/common/utils"
	"smartplant-ems/report-service/models"

	"github.com/beego/beego/v2/client/orm"
=======
>>>>>>> 395aa7cad5570b4b699b4c029768d04fab946652
	"github.com/beego/beego/v2/server/web"
)

type ReportController struct {
	web.Controller
}

func (c *ReportController) List() {
	page, _ := strconv.Atoi(c.GetString("page", "1"))
	size, _ := strconv.Atoi(c.GetString("size", "20"))
<<<<<<< HEAD
	reportType := c.GetString("report_type")
	periodType := c.GetString("period_type")

	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 20
	}

	o := orm.NewOrm()
	qs := o.QueryTable(new(models.Report))

	if reportType != "" {
		qs = qs.Filter("report_type", reportType)
	}
	if periodType != "" {
		qs = qs.Filter("period_type", periodType)
	}

	var reports []models.Report
	total, _ := qs.Count()
	qs.OrderBy("-created_at").Offset((page - 1) * size).Limit(size).All(&reports)

	utils.ReturnSuccess(&c.Controller, map[string]interface{}{
		"records": reports,
		"total":   total,
		"page":    page,
		"size":    size,
	})
}

func (c *ReportController) Get() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.ReturnBadRequest(&c.Controller, "无效的报表ID")
		return
	}

	o := orm.NewOrm()
	report := models.Report{ID: id}
	if err := o.Read(&report); err != nil {
		utils.ReturnNotFound(&c.Controller, "报表不存在")
		return
	}

	utils.ReturnSuccess(&c.Controller, report)
}

func (c *ReportController) Create() {
	var req models.ReportCreate
	if err := c.ParseForm(&req); err != nil {
		utils.ReturnBadRequest(&c.Controller, "无效的请求参数")
		return
	}

	if req.ReportName == "" || req.ReportType == "" || req.PeriodType == "" ||
		req.PeriodStart == "" || req.PeriodEnd == "" {
		utils.ReturnBadRequest(&c.Controller, "缺少必填字段")
		return
	}

	periodStart, err := time.Parse("2006-01-02", req.PeriodStart)
	if err != nil {
		utils.ReturnBadRequest(&c.Controller, "无效的开始日期格式")
		return
	}

	periodEnd, err := time.Parse("2006-01-02", req.PeriodEnd)
	if err != nil {
		utils.ReturnBadRequest(&c.Controller, "无效的结束日期格式")
		return
	}

	userID := int64(1)
	if uid := c.GetSession("user_id"); uid != nil {
		userID = uid.(int64)
	}

	o := orm.NewOrm()
	report := models.Report{
		ReportName:  req.ReportName,
		ReportType:  req.ReportType,
		PeriodType:  req.PeriodType,
		PeriodStart: periodStart,
		PeriodEnd:   periodEnd,
		CreatedBy:   userID,
		Status:      0,
	}

	if _, err := o.Insert(&report); err != nil {
		utils.ReturnServerError(&c.Controller, "创建报表失败")
		return
	}

	utils.ReturnSuccess(&c.Controller, map[string]interface{}{"id": report.ID})
}

func (c *ReportController) Download() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.ReturnBadRequest(&c.Controller, "无效的报表ID")
		return
	}

	o := orm.NewOrm()
	report := models.Report{ID: id}
	if err := o.Read(&report); err != nil {
		utils.ReturnNotFound(&c.Controller, "报表不存在")
		return
	}

	energyRecords, totalCount, err := fetchEnergyDataForReport(o, report.PeriodStart, report.PeriodEnd)
	if err != nil {
		utils.ReturnServerError(&c.Controller, "获取报表数据失败")
		return
	}

	csvContent := generateCSVContent(report, energyRecords, totalCount)

	filename := fmt.Sprintf("%s_%s_%s.csv", report.ReportName,
		report.PeriodStart.Format("2006-01-02"),
		report.PeriodEnd.Format("2006-01-02"))

	c.Ctx.ResponseWriter.Header().Set("Content-Type", "text/csv; charset=utf-8")
	c.Ctx.ResponseWriter.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	c.Ctx.ResponseWriter.Write([]byte(csvContent))
}

func fetchEnergyDataForReport(o orm.Ormer, start, end time.Time) ([]map[string]interface{}, int64, error) {
	var results []orm.Params
	statsSql := `SELECT
		device_id,
		SUM(energy_kwh) as total_energy_kwh,
		AVG(power_kw) as avg_power_kw,
		MAX(power_kw) as max_power_kw,
		MIN(power_kw) as min_power_kw,
		COUNT(*) as record_count
	FROM energy_records
	WHERE record_time >= ? AND record_time < ?
	GROUP BY device_id
	ORDER BY device_id`

	_, err := o.Raw(statsSql, start, end.AddDate(0, 0, 1)).Values(&results)
	if err != nil {
		return nil, 0, err
	}

	var total int64
	countSql := "SELECT COUNT(*) FROM energy_records WHERE record_time >= ? AND record_time < ?"
	if err := o.Raw(countSql, start, end.AddDate(0, 0, 1)).QueryRow(&total); err != nil {
		return nil, 0, err
	}

	return results, total, nil
}

func generateCSVContent(report models.Report, records []orm.Params, totalCount int64) string {
	var lines []string
	lines = append(lines, fmt.Sprintf("报表名称,%s", report.ReportName))
	lines = append(lines, fmt.Sprintf("报表类型,%s", report.ReportType))
	lines = append(lines, fmt.Sprintf("统计周期,%s 至 %s", report.PeriodStart.Format("2006-01-02"), report.PeriodEnd.Format("2006-01-02")))
	lines = append(lines, fmt.Sprintf("生成时间,%s", time.Now().Format("2006-01-02 15:04:05")))
	lines = append(lines, "")

	lines = append(lines, "设备ID,总能耗(KWH),平均功率(KW),最大功率(KW),最小功率(KW),记录数")

	for _, row := range records {
		deviceID := getMapValue(row, "device_id")
		totalEnergy := getMapValue(row, "total_energy_kwh")
		avgPower := getMapValue(row, "avg_power_kw")
		maxPower := getMapValue(row, "max_power_kw")
		minPower := getMapValue(row, "min_power_kw")
		recordCount := getMapValue(row, "record_count")

		lines = append(lines, fmt.Sprintf("%v,%v,%v,%v,%v,%v",
			deviceID, totalEnergy, avgPower, maxPower, minPower, recordCount))
	}

	lines = append(lines, "")
	lines = append(lines, fmt.Sprintf("总记录数,%d", totalCount))

	return strings.Join(lines, "\n")
}

func getMapValue(m orm.Params, key string) interface{} {
	if v, ok := m[key]; ok && v != nil {
		switch val := v.(type) {
		case float64:
			return fmt.Sprintf("%.4f", val)
		case int64:
			return fmt.Sprintf("%d", val)
		default:
			return fmt.Sprintf("%v", val)
		}
	}
	return ""
}

func (c *ReportController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.ReturnBadRequest(&c.Controller, "无效的报表ID")
		return
	}

	o := orm.NewOrm()
	report := models.Report{ID: id}
	if _, err := o.Delete(&report); err != nil {
		utils.ReturnServerError(&c.Controller, "删除报表失败")
		return
	}

	utils.ReturnSuccess(&c.Controller, nil)
}
=======

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
>>>>>>> 395aa7cad5570b4b699b4c029768d04fab946652
