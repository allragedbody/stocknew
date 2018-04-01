package routers

import (
	"stocknew/lotterytools/controllers"

	"github.com/astaxie/beego"
)

func Init() {
	beego.NSNamespace("/*",
		//Options用于跨域复杂请求预检
		beego.NSRouter("/*", &controllers.BaseController{}, "options:Options"),
	)

	beego.Router("/", &controllers.MainController{})
	beego.Router("/lotterydata/:size", &controllers.LotteryController{}, "get:LotteryData")
	beego.Router("/missdata/", &controllers.LotteryController{}, "get:MissData")
	beego.Router("/putdata/", &controllers.LotteryController{}, "get:PutData")
	beego.Router("/newplan/", &controllers.LotteryController{}, "get:GetPlan")
	beego.Router("/getdatedata/", &controllers.LotteryController{}, "get:GetDateData")

}

