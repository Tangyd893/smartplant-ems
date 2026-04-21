package main

import (
	_ "smartplant-ems/report-service/routers"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	web.Run()
}