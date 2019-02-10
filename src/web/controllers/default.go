package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	// 继承
	beego.Controller
}

func (c *MainController) Test() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
func (c *MainController) Post() {
	c.Ctx.WriteString("hello this is post ......")

}
