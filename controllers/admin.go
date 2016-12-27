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
	c.Data["admin"] = true
	c.Data["user"] = user

	hd := &models.HomeData{}
	hd = hd.GetData()

	tmp := make([]models.TripDes, 0)

	c.Data["first"] = hd.First
	firstTmp := append(hd.First, models.TripDes{Index: -1})
	FirstList := make([][]models.TripDes, 0)
	tmp = make([]models.TripDes, 0)
	for index, v := range firstTmp {
		if v.Index == -1 {

		} else {
			v.Index = index
		}

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
	SecondTmp := append(hd.Second, models.TripDes{Index: -1})
	SecondList := make([][]models.TripDes, 0)
	tmp = make([]models.TripDes, 0)
	for index, v := range SecondTmp {
		if v.Index == -1 {

		} else {
			v.Index = index
		}

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
	thirdTmp := append(hd.Third, models.TripDes{Index: -1})
	ThirdList := make([][]models.TripDes, 0)
	tmp = make([]models.TripDes, 0)
	for index, v := range thirdTmp {
		if v.Index == -1 {

		} else {
			v.Index = index
		}

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
				if u.Index == -1 {
					hd.First = append(hd.First, u)
				} else {
					hd.First[u.Index] = u
				}
			}
			if u.Type == 1 {
				if u.Index == -1 {
					hd.Second = append(hd.Second, u)
				} else {
					hd.Second[u.Index] = u
				}
			}
			if u.Type == 2 {
				if u.Index == -1 {
					hd.Third = append(hd.Third, u)
				} else {
					hd.Third[u.Index] = u
				}
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

//Delete xx
func (this *AdminController) Delete() {

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
				hd.First = append(hd.First[:u.Index], hd.First[u.Index+1:]...)
			}
			if u.Type == 1 {
				hd.Second = append(hd.Second[:u.Index], hd.First[u.Index+1:]...)
			}
			if u.Type == 2 {
				hd.Third = append(hd.Third[:u.Index], hd.First[u.Index+1:]...)
			}
			err = hd.Save()
			if err != nil {
				result["err"] = err.Error()
			} else {
				result["data"] = u.Index
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

//UploadImg xxx
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
