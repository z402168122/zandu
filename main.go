package main

import (
	"fmt"
	"zhangjl/zanadu/basic"
	"zhangjl/zanadu/models"
	_ "zhangjl/zanadu/routers"

	"github.com/astaxie/beego"
)

func main() {
	basic.InitLogger()

	filename := "data/db"

	hd := models.HomeData{}
	hd.Path = filename
	err := hd.Init()
	if err != nil {
		basic.Log.Error(err.Error())
		return
	}
	fmt.Println("conifg :", hd)
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
