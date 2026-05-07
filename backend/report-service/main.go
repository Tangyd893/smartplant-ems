package main

import (
<<<<<<< HEAD
	"log"

	_ "smartplant-ems/report-service/routers"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/beego/beego/v2/server/web"
	"smartplant-ems/common/config"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	dbUser := config.GetDBUser()
	dbPass := config.GetDBPass()
	dbHost := config.GetDBHost()
	dbName := config.GetDBName()

	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	err := orm.RegisterDataBase("default", "mysql", dsn)
	if err != nil {
		log.Fatal("Failed to register database:", err)
	}

	orm.SetMaxIdleConns("default", config.GetMaxIdleConns())
	orm.SetMaxOpenConns("default", config.GetMaxOpenConns())
}

=======
	_ "smartplant-ems/report-service/routers"
	"github.com/beego/beego/v2/server/web"
)

>>>>>>> 395aa7cad5570b4b699b4c029768d04fab946652
func main() {
	web.Run()
}