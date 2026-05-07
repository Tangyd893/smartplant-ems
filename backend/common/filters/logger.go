package filters

import (
	"log"
	"time"

	"github.com/beego/beego/v2/server/web"
)

func Logger() web.FilterFunc {
	return func(c *web.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		latency := time.Since(start)
		status := c.Response.StatusCode

		log.Printf("[%s] %s %s %d %v",
			time.Now().Format("2006-01-02 15:04:05"),
			method,
			path,
			status,
			latency,
		)
	}
}