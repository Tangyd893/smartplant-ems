package main

import (
	_ "smartplant-ems/device-service/routers"
	"github.com/beego/beego/v2/server/web"
)

func main() {
	web.Run()
}