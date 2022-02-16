package routers

import (
    "beego_start/controllers"
    "github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    // get请求访问到Hi方法
    beego.Router("/hi", &controllers.MainController{}, "get:Hi")
}
