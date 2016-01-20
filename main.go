package main

import (
	"log"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"myMessage/controllers"
	"myMessage/models"
)

const (
	configPath = "conf/app.conf"
	db         = "mysql"
)

var MyConfig config.Configer

func init() {
	myConfig, err := config.NewConfig("ini", configPath)
	if err != nil {
		log.Panicf("Can't find config file at '%v'.", configPath)
	} else {
		MyConfig = myConfig
	}

	orm.RegisterModel(new(models.User))
	orm.RegisterModel(new(models.Message))

	dbUser := myConfig.String("db_user")
	dbPwd := myConfig.String("db_pwd")
	dbPort := myConfig.String("db_port")
	dbHost := myConfig.String("db_host")
	dbName := myConfig.String("db_name")

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dbUser+":"+dbPwd+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8")
	orm.RunSyncdb("default", false, true)
}

func main() {
	beego.Router("/messages/", &controllers.MessageController{}, "get:List;post:Post")
	beego.Router("/messages/:id", &controllers.MessageController{}, "get:Get;delete:Delete;put:Put")
	beego.Run(MyConfig.String("http_port"))
}
