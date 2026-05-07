package controllers

import (
	"strconv"

	"smartplant-ems/common/utils"
	"smartplant-ems/device-service/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

type DeviceController struct {
	web.Controller
}

func (c *DeviceController) List() {
	page, _ := strconv.Atoi(c.GetString("page", "1"))
	size, _ := strconv.Atoi(c.GetString("size", "20"))
	deviceType := c.GetString("device_type")
	status, _ := c.GetInt("status", -1)

	if page < 1 {
		page = 1
	}
	if size < 1 || size > 100 {
		size = 20
	}

	o := orm.NewOrm()
	qs := o.QueryTable(new(models.Device))

	if deviceType != "" {
		qs = qs.Filter("device_type", deviceType)
	}
	if status >= 0 {
		qs = qs.Filter("status", status)
	}

	var devices []models.Device
	total, err := qs.Count()
	if err != nil {
		utils.ReturnServerError(&c.Controller, "查询设备列表失败")
		return
	}
	_, err = qs.Offset((page-1)*size).Limit(size).All(&devices)
	if err != nil {
		utils.ReturnServerError(&c.Controller, "查询设备列表失败")
		return
	}

	utils.ReturnSuccess(&c.Controller, map[string]interface{}{
		"records": devices,
		"total":   total,
		"page":    page,
		"size":    size,
	})
}

func (c *DeviceController) Get() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.ReturnBadRequest(&c.Controller, "无效的设备ID")
		return
	}

	o := orm.NewOrm()
	device := models.Device{ID: id}
	if err := o.Read(&device); err != nil {
		utils.ReturnNotFound(&c.Controller, "设备不存在")
		return
	}

	utils.ReturnSuccess(&c.Controller, device)
}

func (c *DeviceController) Post() {
	var req models.DeviceCreate
	if err := c.ParseForm(&req); err != nil {
		utils.ReturnBadRequest(&c.Controller, "无效的请求参数")
		return
	}

	if req.DeviceCode == "" || req.DeviceName == "" || req.DeviceType == "" {
		utils.ReturnBadRequest(&c.Controller, "设备编号、名称和类型不能为空")
		return
	}

	o := orm.NewOrm()
	device := models.Device{
		DeviceCode:  req.DeviceCode,
		DeviceName:  req.DeviceName,
		DeviceType:  req.DeviceType,
		Location:    req.Location,
		PowerRating: req.PowerRating,
		Status:      1,
	}

	if _, err := o.Insert(&device); err != nil {
		utils.ReturnServerError(&c.Controller, "创建设备失败")
		return
	}

	utils.ReturnSuccess(&c.Controller, map[string]interface{}{"id": device.ID})
}

func (c *DeviceController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.ReturnBadRequest(&c.Controller, "无效的设备ID")
		return
	}

	var req struct {
		DeviceName  string  `json:"device_name"`
		Location    string  `json:"location"`
		Status      int     `json:"status"`
		PowerRating float64 `json:"power_rating"`
	}

	if err := c.ParseForm(&req); err != nil {
		utils.ReturnBadRequest(&c.Controller, "无效的请求参数")
		return
	}

	o := orm.NewOrm()
	device := models.Device{ID: id}
	if err := o.Read(&device); err != nil {
		utils.ReturnNotFound(&c.Controller, "设备不存在")
		return
	}

	if req.DeviceName != "" {
		device.DeviceName = req.DeviceName
	}
	if req.Location != "" {
		device.Location = req.Location
	}
	if req.Status > 0 {
		device.Status = req.Status
	}
	if req.PowerRating > 0 {
		device.PowerRating = req.PowerRating
	}

	if _, err := o.Update(&device); err != nil {
		utils.ReturnServerError(&c.Controller, "更新设备失败")
		return
	}

	utils.ReturnSuccess(&c.Controller, nil)
}

func (c *DeviceController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.ReturnBadRequest(&c.Controller, "无效的设备ID")
		return
	}

	o := orm.NewOrm()
	device := models.Device{ID: id}
	if _, err := o.Delete(&device); err != nil {
		utils.ReturnServerError(&c.Controller, "删除设备失败")
		return
	}

	utils.ReturnSuccess(&c.Controller, nil)
}

func (c *DeviceController) GetRealTimeData() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.ReturnBadRequest(&c.Controller, "无效的设备ID")
		return
	}

	o := orm.NewOrm()
	device := models.Device{ID: id}
	if err := o.Read(&device); err != nil {
		utils.ReturnNotFound(&c.Controller, "设备不存在")
		return
	}

	data := map[string]interface{}{
		"device_id":    device.ID,
		"device_name":   device.DeviceName,
		"power_kw":      device.PowerRating * 0.85,
		"voltage_v":     380.5,
		"current_a":     device.PowerRating * 0.85 * 1.52,
		"power_factor":  0.92,
		"record_time":   "2026-05-07T10:00:00+08:00",
	}

	utils.ReturnSuccess(&c.Controller, data)
}
