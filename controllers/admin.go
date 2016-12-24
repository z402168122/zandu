package controllers

import (
	"fmt"
	"time"

	"encoding/json"
	"zhangjl/zanadu/models"

	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

//Get xx
func (c *AdminController) Get() {
	hd := &models.HomeData{}
	hd = hd.GetData()
	c.Data["admin"] = true
	c.Data["first"] = hd.First
	c.Data["second"] = hd.Second
	c.Data["third"] = hd.Third
	c.TplName = "admin.html"
}

//Update xx
func (this *AdminController) Update() {

	result := make(map[string]interface{}, 0)
	result["err"] = ""
	hd := &models.HomeData{}
	hd = hd.GetData()
	var err error
	u := models.TripDes{}
	if err := this.ParseForm(&u); err != nil {
		fmt.Println(err.Error())
		result["err"] = err.Error()
	} else {
		if u.Type == 0 {
			hd.First[u.Index] = u
		}
		if u.Type == 1 {
			hd.Second[u.Index] = u
		}
		if u.Type == 2 {
			hd.Third[u.Index] = u
		}
		err = hd.Save()
		if err != nil {
			result["err"] = err.Error()
		} else {
			result["data"] = u
		}
	}

	v, err := json.Marshal(result)
	if err == nil {

	} else {
		fmt.Println(err.Error())
	}
	this.Ctx.WriteString(string(v))
}

func (this *AdminController) UploadImg() {

	path := fmt.Sprintf("/static/upload/%d.jpg", time.Now().Unix())
	err := this.SaveToFile("temp_img", fmt.Sprintf(".%s", path))

	result := make(map[string]interface{}, 0)
	result["url"] = path
	if err != nil {
		result["err"] = err.Error()
	} else {
		result["err"] = ""
	}
	// this.Data["json"] = result
	v, err := json.Marshal(result)
	if err == nil {

	} else {
		fmt.Println(err.Error())
	}
	this.Ctx.WriteString(string(v))
}
