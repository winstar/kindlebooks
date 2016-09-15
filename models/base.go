package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Register() {
	kindleHost := beego.AppConfig.String("mysql.kindle.host")
	kindlePort := beego.AppConfig.String("mysql.kindle.port")
	kindleUser := beego.AppConfig.String("mysql.kindle.user")
	kindlePassword := beego.AppConfig.String("mysql.kindle.password")
	kindleDatabase := beego.AppConfig.String("mysql.kindle.database")
	if kindlePort == "" {
		kindlePort = "3306"
	}
	dataSource := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8", kindleUser, kindlePassword, kindleHost, kindlePort, kindleDatabase)
	fmt.Println(dataSource)
	orm.RegisterDataBase("default", "mysql", dataSource)
	orm.RegisterModel(new(Member))
}
