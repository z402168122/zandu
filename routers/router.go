package routers

import (
	"zhangjl/zanadu/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin", &controllers.AdminController{})
	beego.Router("/admin/json", &controllers.AdminController{}, "get:Json")
	beego.Router("/admin/upload_img", &controllers.AdminController{}, "post:UploadImg")
}
