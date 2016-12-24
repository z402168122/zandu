package controllers

import (
	"fmt"
	"time"

	"encoding/json"

	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Get() {
	c.Data["admin"] = true
	c.TplName = "admin.html"
}

func (this *AdminController) Json() {
	result := make(map[string]interface{}, 0)
	result["url"] = "0000"
	this.Data["json"] = result
	this.ServeJSON(false)
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
