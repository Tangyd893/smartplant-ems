package filters

import (
	"github.com/beego/beego/v2/server/web"
	"smartplant-ems/common/config"
)

func CORS() web.FilterFunc {
	return func(c *web.Context) {
		c.SetHeader("Access-Control-Allow-Origin", config.GetCORSAllowedOrigin())
		c.SetHeader("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		c.SetHeader("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")
		c.SetHeader("Access-Control-Max-Age", "86400")
		c.SetHeader("Access-Control-Allow-Credentials", "true")
	}
}