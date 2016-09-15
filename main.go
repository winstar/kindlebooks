package main

import (
	_ "github.com/winstar/kindlebooks/routers"
	"github.com/astaxie/beego"
	"github.com/winstar/kindlebooks/models"
)

func init() {
	models.Register()
}

func main() {
	beego.Run()
}

