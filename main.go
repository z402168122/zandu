package main

import (
	"fmt"
	"zhangjl/zanadu/models"
	_ "zhangjl/zanadu/routers"

	"github.com/astaxie/beego"
)

func main() {

	filename := "data/db"

	hd := models.HomeData{}
	hd.Path = filename
	err := hd.Init()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("conifg :", hd)
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
}
