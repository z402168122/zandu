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
	tmp := make([]models.TripDes, 0)
	FirstList := make([][]models.TripDes, 0)
	for _, v := range hd.First {
		tmp = append(tmp, v)
		FirstList = append(FirstList, tmp)
	}
	c.Data["first_list"] = FirstList

	c.Data["second"] = hd.Second
	tmp = make([]models.TripDes, 0)
	SecondList := make([][]models.TripDes, 0)
	for _, v := range hd.First {
		tmp = append(tmp, v)
		SecondList = append(SecondList, tmp)
	}
	c.Data["second_list"] = SecondList

	tmp = make([]models.TripDes, 0)
	c.Data["third"] = hd.Third
	ThirdList := make([][]models.TripDes, 0)
	for _, v := range hd.Third {
		tmp = append(tmp, v)
		ThirdList = append(ThirdList, tmp)
	}
	c.Data["third_list"] = ThirdList
	tmp = make([]models.TripDes, 0)

	c.TplName = "admin.html"
}
