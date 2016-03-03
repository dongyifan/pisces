package routers

import (
	"github.com/wishdream/pisces/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
