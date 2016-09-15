package routers

import (
	"github.com/winstar/kindlebooks/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
