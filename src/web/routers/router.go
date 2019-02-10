package routers

import (
	"web/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Test;post:Post")
	//beego.Router("/test", &controllers.TestController{}, "get:Get;post:Post")
	//beego.Router("/test_model", &controllers.TestModelController{}, "get:Get;post:Post")
	beego.Router("/movie", &controllers.MovieController{}, "*:Create")

}
