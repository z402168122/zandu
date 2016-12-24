package controllers

import (
	"zhangjl/zanadu/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	hd := &models.HomeData{}
	hd = hd.GetData()
	c.Data["first"] = hd.First
	c.Data["second"] = hd.Second
	c.Data["third"] = hd.Third
	c.TplName = "admin.html"
}
