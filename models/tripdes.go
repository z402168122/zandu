package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//G_HomeData 主页数据
var g_HomeData *HomeData

//TripDes xx
type TripDes struct {
	Type     int    `json:"type" form:"type"`
	Index    int    `json:"index" form:"index"`
	Img      string `json:"img" form:"img"`
	Code     string `json:"code" form:"code"`
	Title    string `json:"title" form:"title"`
	SubTitle string `json:"sub_title" form:"sub_title"`
	Note     string `json:"note" form:"note"`
}

//HomeData xx
type HomeData struct {
	First  []TripDes `json:"first"`
	Second []TripDes `json:"second"`
	Third  []TripDes `json:"third"`
	Path   string    `json:"path"`
}

//GetData xxx
func (rec *HomeData) GetData() *HomeData {
	return g_HomeData
}

//Save xx
func (rec *HomeData) Save() error {
	var err error
	data, err := json.Marshal(rec)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(rec.Path, data, 0666)
	return err
}

//Init xxx
func (rec *HomeData) Init() error {
	var err error
	_, err = os.Stat(rec.Path)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return err
	}
	bytes, err := ioutil.ReadFile(rec.Path)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return err
	}
	if err := json.Unmarshal(bytes, rec); err != nil {
		err = fmt.Errorf("unmarshal error :%s", string(bytes))
		fmt.Println(err.Error())
		First := make([]TripDes, 0)
		Second := make([]TripDes, 0)
		Third := make([]TripDes, 0)
		rec.First = First
		rec.Second = Second
		rec.Third = Third
		err = rec.Save()
		if err != nil {
			fmt.Println(err.Error())
			return err
		} else {
			fmt.Println("初始化数据成功")
		}
	}
	g_HomeData = rec
	return err
}
