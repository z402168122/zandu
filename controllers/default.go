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

	tmp := make([]models.TripDes, 0)

	c.Data["first"] = hd.First

	FirstList := make([][]models.TripDes, 0)
	tmp = make([]models.TripDes, 0)
	for index, v := range hd.First {
		v.Index = index
		tmp = append(tmp, v)
		if (index+1)%3 == 0 {
			FirstList = append(FirstList, tmp)
			tmp = make([]models.TripDes, 0)
		}

	}
	if len(tmp) > 0 {
		FirstList = append(FirstList, tmp)
		tmp = make([]models.TripDes, 0)
	}
	c.Data["first_list"] = FirstList
	// ----------
	c.Data["second"] = hd.Second
	SecondList := make([][]models.TripDes, 0)
	tmp = make([]models.TripDes, 0)
	for index, v := range hd.Second {
		v.Index = index
		tmp = append(tmp, v)
		if (index+1)%3 == 0 {
			SecondList = append(SecondList, tmp)
			tmp = make([]models.TripDes, 0)
		}

	}
	if len(tmp) > 0 {
		SecondList = append(SecondList, tmp)
		tmp = make([]models.TripDes, 0)
	}
	c.Data["second_list"] = SecondList

	// ------------

	c.Data["third"] = hd.Third
	ThirdList := make([][]models.TripDes, 0)
	tmp = make([]models.TripDes, 0)
	for index, v := range hd.Third {
		v.Index = index
		tmp = append(tmp, v)
		if (index+1)%3 == 0 {
			ThirdList = append(ThirdList, tmp)
			tmp = make([]models.TripDes, 0)
		}

	}
	if len(tmp) > 0 {
		ThirdList = append(ThirdList, tmp)
		tmp = make([]models.TripDes, 0)
	}
	c.Data["third_list"] = ThirdList

	c.TplName = "index.html"
}
