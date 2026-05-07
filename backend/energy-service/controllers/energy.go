package controllers

import (
	"strconv"
	"time"

	"smartplant-ems/common/utils"
	"smartplant-ems/energy-service/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

type EnergyController struct {
	web.Controller
}

func (c *EnergyController) List() {
	page, _ := strconv.Atoi(c.GetString("page", "1"))
	size, _ := strconv.Atoi(c.GetString("size", "20"))
	deviceID, _ := c.GetInt("device_id", 0)
	startTime := c.GetString("start_time")
	endTime := c.GetString("end_time")

	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 20
	}

	o := orm.NewOrm()
	qs := o.QueryTable(new(models.EnergyRecord))

	if deviceID > 0 {
		qs = qs.Filter("device_id", deviceID)
	}

	if startTime != "" {
		if t, err := time.Parse("2006-01-02", startTime); err == nil {
			qs = qs.Filter("record_time__gte", t)
		}
	}
	if endTime != "" {
		if t, err := time.Parse("2006-01-02", endTime); err == nil {
			endOfDay := t.Add(24 * time.Hour)
			qs = qs.Filter("record_time__lt", endOfDay)
		}
	}

	var records []models.EnergyRecord
	total, err := qs.Count()
	if err != nil {
		utils.ReturnServerError(&c.Controller, "查询能源记录列表失败")
		return
	}
	_, err = qs.OrderBy("-record_time").Offset((page - 1) * size).Limit(size).All(&records)
	if err != nil {
		utils.ReturnServerError(&c.Controller, "查询能源记录列表失败")
		return
	}

	utils.ReturnSuccess(&c.Controller, map[string]interface{}{
		"records": records,
		"total":   total,
		"page":    page,
		"size":    size,
	})
}

func (c *EnergyController) Get() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.ReturnBadRequest(&c.Controller, "无效的记录ID")
		return
	}

	o := orm.NewOrm()
	record := models.EnergyRecord{ID: id}
	if err := o.Read(&record); err != nil {
		utils.ReturnNotFound(&c.Controller, "记录不存在")
		return
	}

	utils.ReturnSuccess(&c.Controller, record)
}

func (c *EnergyController) Create() {
	var req struct {
		DeviceID    int64   `json:"device_id"`
		RecordTime  string  `json:"record_time"`
		PowerKW     float64 `json:"power_kw"`
		EnergyKWH   float64 `json:"energy_kwh"`
		VoltageV    float64 `json:"voltage_v"`
		CurrentA    float64 `json:"current_a"`
		PowerFactor float64 `json:"power_factor"`
	}

	if err := c.ParseForm(&req); err != nil {
		utils.ReturnBadRequest(&c.Controller, "无效的请求参数")
		return
	}

	if req.DeviceID <= 0 {
		utils.ReturnBadRequest(&c.Controller, "设备ID不能为空")
		return
	}

	recordTime := time.Now()
	if req.RecordTime != "" {
		if t, err := time.Parse(time.RFC3339, req.RecordTime); err == nil {
			recordTime = t
		}
	}

	o := orm.NewOrm()
	record := models.EnergyRecord{
		DeviceID:    req.DeviceID,
		RecordTime:  recordTime,
		PowerKW:     req.PowerKW,
		EnergyKWH:   req.EnergyKWH,
		VoltageV:    req.VoltageV,
		CurrentA:    req.CurrentA,
		PowerFactor: req.PowerFactor,
	}

	if _, err := o.Insert(&record); err != nil {
		utils.ReturnServerError(&c.Controller, "创建记录失败")
		return
	}

	utils.ReturnSuccess(&c.Controller, map[string]interface{}{"id": record.ID})
}

func (c *EnergyController) Update() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.ReturnBadRequest(&c.Controller, "无效的记录ID")
		return
	}

	var req struct {
		PowerKW     float64 `json:"power_kw"`
		EnergyKWH   float64 `json:"energy_kwh"`
		VoltageV    float64 `json:"voltage_v"`
		CurrentA    float64 `json:"current_a"`
		PowerFactor float64 `json:"power_factor"`
	}

	if err := c.ParseForm(&req); err != nil {
		utils.ReturnBadRequest(&c.Controller, "无效的请求参数")
		return
	}

	o := orm.NewOrm()
	record := models.EnergyRecord{ID: id}
	if err := o.Read(&record); err != nil {
		utils.ReturnNotFound(&c.Controller, "记录不存在")
		return
	}

	if req.PowerKW > 0 {
		record.PowerKW = req.PowerKW
	}
	if req.EnergyKWH > 0 {
		record.EnergyKWH = req.EnergyKWH
	}
	if req.VoltageV > 0 {
		record.VoltageV = req.VoltageV
	}
	if req.CurrentA > 0 {
		record.CurrentA = req.CurrentA
	}
	if req.PowerFactor > 0 {
		record.PowerFactor = req.PowerFactor
	}

	if _, err := o.Update(&record); err != nil {
		utils.ReturnServerError(&c.Controller, "更新记录失败")
		return
	}

	utils.ReturnSuccess(&c.Controller, nil)
}

func (c *EnergyController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.ReturnBadRequest(&c.Controller, "无效的记录ID")
		return
	}

	o := orm.NewOrm()
	record := models.EnergyRecord{ID: id}
	if _, err := o.Delete(&record); err != nil {
		utils.ReturnServerError(&c.Controller, "删除记录失败")
		return
	}

	utils.ReturnSuccess(&c.Controller, nil)
}

func (c *EnergyController) GetStats() {
	deviceID, _ := c.GetInt("device_id", 0)
	startTime := c.GetString("start_time")
	endTime := c.GetString("end_time")

	o := orm.NewOrm()

	whereClause := "1=1"
	args := []interface{}{}

	if deviceID > 0 {
		whereClause += " AND device_id = ?"
		args = append(args, deviceID)
	}

	if startTime != "" {
		if t, err := time.Parse("2006-01-02", startTime); err == nil {
			whereClause += " AND record_time >= ?"
			args = append(args, t)
		}
	}
	if endTime != "" {
		if t, err := time.Parse("2006-01-02", endTime); err == nil {
			endOfDay := t.Add(24 * time.Hour)
			whereClause += " AND record_time < ?"
			args = append(args, endOfDay)
		}
	}

	var total int64
	countSql := "SELECT COUNT(*) FROM energy_records WHERE " + whereClause
	if err := o.Raw(countSql, args).QueryRow(&total); err != nil {
		utils.ReturnServerError(&c.Controller, "查询统计数据失败")
		return
	}

	var results []orm.Params
	statsSql := `SELECT
		device_id,
		SUM(energy_kwh) as total_energy_kwh,
		AVG(power_kw) as avg_power_kw,
		MAX(power_kw) as max_power_kw,
		MIN(power_kw) as min_power_kw,
		COUNT(*) as record_count
	FROM energy_records
	WHERE ` + whereClause + `
	GROUP BY device_id
	ORDER BY device_id`

	if _, err := o.Raw(statsSql, args).Values(&results); err != nil {
		utils.ReturnServerError(&c.Controller, "查询统计数据失败")
		return
	}

	stats := make([]models.EnergyStats, 0, len(results))
	for _, row := range results {
		stats = append(stats, models.EnergyStats{
			DeviceID:        toInt64(row["device_id"]),
			DeviceName:      "Device-" + strconv.FormatInt(toInt64(row["device_id"]), 10),
			TotalEnergyKWH: toFloat64(row["total_energy_kwh"]),
			AvgPowerKW:      toFloat64(row["avg_power_kw"]),
			MaxPowerKW:      toFloat64(row["max_power_kw"]),
			MinPowerKW:      toFloat64(row["min_power_kw"]),
			RecordCount:     toInt64(row["record_count"]),
		})
	}

	utils.ReturnSuccess(&c.Controller, map[string]interface{}{
		"stats":        stats,
		"total_count":  total,
		"device_count": len(stats),
	})
}

func toInt64(v interface{}) int64 {
	if v == nil {
		return 0
	}
	switch val := v.(type) {
	case int64:
		return val
	case int:
		return int64(val)
	case float64:
		return int64(val)
	}
	return 0
}

func toFloat64(v interface{}) float64 {
	if v == nil {
		return 0
	}
	switch val := v.(type) {
	case float64:
		return val
	case int64:
		return float64(val)
	case int:
		return float64(val)
	}
	return 0
}