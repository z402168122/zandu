package routers

import (
	"zhangjl/zanadu/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/admin", &controllers.AdminController{})
	beego.Router("/admin/login", &controllers.AdminController{}, "post:Login")

	beego.Router("/admin/upload_img", &controllers.AdminController{}, "post:UploadImg")
	beego.Router("/admin/update", &controllers.AdminController{}, "post:Update")

}
