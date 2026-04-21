package controllers

import (
	"fmt"
	"strconv"
	"github.com/beego/beego/v2/server/web"
)

type EnergyController struct {
	web.Controller
}

func (c *EnergyController) List() {
	page, _ := strconv.Atoi(c.GetString("page", "1"))
	size, _ := strconv.Atoi(c.GetString("size", "20"))

	records := []map[string]interface{}{
		{"id": 1, "device_id": 1, "record_time": "2026-04-21T10:00:00+08:00", "power_kw": 125.5, "energy_kwh": 38520.5, "voltage_v": 380.2, "current_a": 190.3},
		{"id": 2, "device_id": 2, "record_time": "2026-04-21T10:00:00+08:00", "power_kw": 118.2, "energy_kwh": 36210.3, "voltage_v": 379.8, "current_a": 180.5},
	}

	c.Data["json"] = map[string]interface{}{
		"code": 0,
		"msg":  "success",
		"data": map[string]interface{}{
			"records": records,
			"total":   2,
			"page":    page,
			"size":    size,
		},
	}
	c.ServeJSON()
}

func (c *EnergyController) Get() {
	id := c.Ctx.Input.Param(":id")
	record := map[string]interface{}{
		"id": 1, "device_id": 1, "record_time": "2026-04-21T10:00:00+08:00",
		"power_kw": 125.5, "energy_kwh": 38520.5, "voltage_v": 380.2, "current_a": 190.3,
	}
	_ = id
	c.Data["json"] = map[string]interface{}{"code": 0, "msg": "success", "data": record}
	c.ServeJSON()
}

func (c *EnergyController) Create() {
	var req map[string]interface{}
	c.ParseForm(&req)
	fmt.Printf("创建能耗记录: %+v\n", req)
	c.Data["json"] = map[string]interface{}{"code": 0, "msg": "创建成功", "data": map[string]interface{}{"id": 10}}
	c.ServeJSON()
}

func (c *EnergyController) Update() {
	id := c.Ctx.Input.Param(":id")
	var req map[string]interface{}
	c.ParseForm(&req)
	fmt.Printf("更新能耗记录 id=%s: %+v\n", id, req)
	c.Data["json"] = map[string]interface{}{"code": 0, "msg": "更新成功"}
	c.ServeJSON()
}

func (c *EnergyController) Delete() {
	id := c.Ctx.Input.Param(":id")
	fmt.Printf("删除能耗记录: id=%s\n", id)
	c.Data["json"] = map[string]interface{}{"code": 0, "msg": "删除成功"}
	c.ServeJSON()
}

func (c *EnergyController) GetStats() {
	period := c.GetString("period", "day")
	data := map[string]interface{}{
		"total_energy_kwh": 15420.5,
		"avg_power_kw":     125.3,
		"max_power_kw":     185.6,
		"min_power_kw":     45.2,
		"total_cost":       8491.28,
		"peak_hours":       5,
		"off_peak_hours":   19,
		"period":           period,
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "msg": "success", "data": data}
	c.ServeJSON()
}
