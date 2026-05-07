package utils

import (
	"github.com/beego/beego/v2/server/web"
)

func ReturnSuccess(c *web.Controller, data interface{}) {
	c.Data["json"] = map[string]interface{}{
		"code": 0,
		"msg":  "success",
		"data": data,
	}
	c.ServeJSON()
}

func ReturnError(c *web.Controller, code int, msg string) {
	c.Data["json"] = map[string]interface{}{
		"code": code,
		"msg":  msg,
	}
	c.ServeJSON()
}

func ReturnUnauthorized(c *web.Controller, msg string) {
	ReturnError(c, 401, msg)
}

func ReturnBadRequest(c *web.Controller, msg string) {
	ReturnError(c, 400, msg)
}

func ReturnNotFound(c *web.Controller, msg string) {
	ReturnError(c, 404, msg)
}

func ReturnServerError(c *web.Controller, msg string) {
	ReturnError(c, 500, msg)
}