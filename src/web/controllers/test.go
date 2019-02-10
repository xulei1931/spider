package controllers

import (
	"github.com/astaxie/beego"
)

type TestController struct {
	// 继承
	beego.Controller
}
type User struct {
	Username string
	Password string
}

func (this *TestController) Get() {
	// 接受get参数
	/*
		id := this.GetString("id")

		//this.input.
		name := this.Input().Get("name")

		this.Ctx.WriteString("<html>" + id + "<br/>")
		this.Ctx.WriteString(name + "<br/>")
		this.Ctx.WriteString("hello this is test ......</html>")
	*/
	this.Ctx.WriteString(`<html>
	<form action="http://0.0.0.0:8080/test" method="post">
	<input type="text" name="Username"/>
	<input type="password" name="Password"/>
	<input type="submit" value="提交"/>
	               </form></html>`)

}

func (this *TestController) Post() {
	//	var ob models.Object
	//	json.Unmarshal(this.Ctx.Input.RequestBody,&ob)
	//	this.Ctx.WriteString(ob)

	//	u := User{}
	//	if err := this.  (&u); err != nil {
	//		// error
	//	}

	//	this.Ctx.WriteString("username:" + u.Username + "password" + u.Password)

}