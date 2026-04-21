package controllers

import (
	"fmt"
	"strconv"
	"github.com/beego/beego/v2/server/web"
)

type DeviceController struct {
	web.Controller
}

func (c *DeviceController) List() {
	page, _ := strconv.Atoi(c.GetString("page", "1"))
	size, _ := strconv.Atoi(c.GetString("size", "20"))

	devices := []map[string]interface{}{
		{"id": 1, "device_code": "AC-001", "device_name": "空压机A1", "device_type": "air_compressor", "location": "车间A", "status": 1, "power_rating": 150.0},
		{"id": 2, "device_code": "AC-002", "device_name": "空压机A2", "device_type": "air_compressor", "location": "车间A", "status": 1, "power_rating": 150.0},
		{"id": 3, "device_code": "IM-001", "device_name": "注塑机B1", "device_type": "injection_machine", "location": "车间B", "status": 1, "power_rating": 75.0},
	}

	c.Data["json"] = map[string]interface{}{
		"code": 0,
		"msg":  "success",
		"data": map[string]interface{}{
			"records": devices,
			"total":   3,
			"page":    page,
			"size":    size,
		},
	}
	c.ServeJSON()
}

func (c *DeviceController) Get() {
	id := c.Ctx.Input.Param(":id")
	device := map[string]interface{}{
		"id":           1,
		"device_code":   "AC-001",
		"device_name":   "空压机A1",
		"device_type":   "air_compressor",
		"location":      "车间A",
		"status":        1,
		"power_rating":  150.0,
		"created_at":    "2026-04-01T10:00:00+08:00",
	}
	_ = id
	c.Data["json"] = map[string]interface{}{"code": 0, "msg": "success", "data": device}
	c.ServeJSON()
}

func (c *DeviceController) Post() {
	var req map[string]interface{}
	c.ParseForm(&req)
	fmt.Printf("创建设备: %+v\n", req)
	c.Data["json"] = map[string]interface{}{"code": 0, "msg": "创建设备成功", "data": map[string]interface{}{"id": 4}}
	c.ServeJSON()
}

func (c *DeviceController) Put() {
	id := c.Ctx.Input.Param(":id")
	var req map[string]interface{}
	c.ParseForm(&req)
	fmt.Printf("更新设备 id=%s: %+v\n", id, req)
	c.Data["json"] = map[string]interface{}{"code": 0, "msg": "更新成功"}
	c.ServeJSON()
}

func (c *DeviceController) Delete() {
	id := c.Ctx.Input.Param(":id")
	fmt.Printf("删除设备: id=%s\n", id)
	c.Data["json"] = map[string]interface{}{"code": 0, "msg": "删除成功"}
	c.ServeJSON()
}

func (c *DeviceController) GetRealTimeData() {
	id := c.Ctx.Input.Param(":id")
	_ = id
	data := map[string]interface{}{
		"device_id":     1,
		"power_kw":      125.5,
		"voltage_v":     380.2,
		"current_a":     190.3,
		"power_factor":  0.92,
		"record_time":   "2026-04-21T11:50:00+08:00",
	}
	c.Data["json"] = map[string]interface{}{"code": 0, "msg": "success", "data": data}
	c.ServeJSON()
}
