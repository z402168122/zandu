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

	user := c.GetSession("user")

	hd := &models.HomeData{}
	hd = hd.GetData()
	c.Data["user"] = user
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

	user := this.GetSession("user")
	if user == nil {
		result["err"] = "没有登陆"
	} else {
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
	result := make(map[string]interface{}, 0)
	result["url"] = path

	user := this.GetSession("user")
	if user == nil {
		result["err"] = "没有登陆"
	} else {

		err := this.SaveToFile("temp_img", fmt.Sprintf(".%s", path))

		if err != nil {
			result["err"] = err.Error()
		} else {
			result["err"] = ""
		}
	}
	v, err := json.Marshal(result)
	if err == nil {

	} else {
		fmt.Println(err.Error())
	}
	this.Ctx.WriteString(string(v))
}

//Login xx
func (this *AdminController) Login() {

	result := make(map[string]interface{}, 0)
	result["err"] = ""
	user := this.GetString("user")
	password := this.GetString("password")
	if user == "2826060360@qq.com" && password == "tangwan123" {
		this.SetSession("user", user)
	} else {
		result["err"] = "账号密码错误"
	}

	v, err := json.Marshal(result)
	if err == nil {

	} else {
		fmt.Println(err.Error())
	}
	this.Ctx.WriteString(string(v))

}
