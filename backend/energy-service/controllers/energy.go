package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/beego/beego/v2/server/web"
)

type EnergyController struct {
	web.Controller
}

type EnergyRecord struct {
	ID         int64     `json:"id"`
	DeviceID   int64     `json:"device_id"`
	RecordTime time.Time `json:"record_time"`
	PowerKW    float64   `json:"power_kw"`
	EnergyKWH  float64   `json:"energy_kwh"`
	VoltageV   float64   `json:"voltage_v"`
	CurrentA   float64   `json:"current_a"`
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

func (c *EnergyController) List() {
	page, _ := strconv.Atoi(c.GetString("page", "1"))
	size, _ := strconv.Atoi(c.GetString("size", "20"))
	deviceID := c.GetString("device_id")
	startDate := c.GetString("start_date")
	endDate := c.GetString("end_date")

	fmt.Printf("List energy records: page=%d, size=%d, device=%s, start=%s, end=%s\n", page, size, deviceID, startDate, endDate)

	records := []EnergyRecord{
		{ID: 1, DeviceID: 1, RecordTime: time.Now(), PowerKW: 125.5, EnergyKWH: 125.5, VoltageV: 380.2, CurrentA: 190.3},
		{ID: 2, DeviceID: 2, RecordTime: time.Now(), PowerKW: 75.2, EnergyKWH: 75.2, VoltageV: 380.1, CurrentA: 110.5},
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

func (c *EnergyController) Get() {
	id, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)
	fmt.Printf("Get energy record: id=%d\n", id)

	c.Data["json"] = APIResponse{
		Code: 0,
		Msg:  "success",
		Data: EnergyRecord{ID: id, DeviceID: 1, RecordTime: time.Now(), PowerKW: 125.5, EnergyKWH: 125.5, VoltageV: 380.2, CurrentA: 190.3},
	}
	c.ServeJSON()
}

func (c *EnergyController) Create() {
	var req map[string]interface{}
	if err := c.ParseForm(&req); err != nil {
		c.Data["json"] = APIResponse{Code: 400, Msg: "参数错误"}
		c.ServeJSON()
		return
	}

	fmt.Printf("Create energy record: %+v\n", req)

	c.Data["json"] = APIResponse{
		Code: 0,
		Msg:  "创建成功",
		Data: map[string]interface{}{"id": 3},
	}
	c.ServeJSON()
}

func (c *EnergyController) Update() {
	var req map[string]interface{}
	if err := c.ParseForm(&req); err != nil {
		c.Data["json"] = APIResponse{Code: 400, Msg: "参数错误"}
		c.ServeJSON()
		return
	}

	fmt.Printf("Update energy record: %+v\n", req)

	c.Data["json"] = APIResponse{Code: 0, Msg: "更新成功"}
	c.ServeJSON()
}

func (c *EnergyController) Delete() {
	id, _ := strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64)
	fmt.Printf("Delete energy record: id=%d\n", id)

	c.Data["json"] = APIResponse{Code: 0, Msg: "删除成功"}
	c.ServeJSON()
}

func (c *EnergyController) GetStats() {
	period := c.GetString("period", "day")
	deviceID := c.GetString("device_id")

	fmt.Printf("Get energy stats: period=%s, device=%s\n", period, deviceID)

	stats := map[string]interface{}{
		"total_energy_kwh": 15420.5,
		"avg_power_kw":     125.3,
		"max_power_kw":     185.6,
		"min_power_kw":     45.2,
		"total_cost":       8491.28,
		"peak_hours":       5,
		"off_peak_hours":   19,
		"period":           period,
	}

	c.Data["json"] = APIResponse{Code: 0, Msg: "success", Data: stats}
	c.ServeJSON()
}