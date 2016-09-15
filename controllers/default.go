package controllers

import (
	"github.com/winstar/kindlebooks/models"
	"fmt"
)

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	member := models.Member{Id:1}
	if err := member.Read(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(member)
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
