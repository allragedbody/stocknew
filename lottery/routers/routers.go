package routers

import (
	"stocknew/lottery/controllers"

	"github.com/astaxie/beego"
)

func Init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/lotterydata/:size", &controllers.LotteryController{}, "get:LotteryData")
	beego.Router("/missdata/", &controllers.LotteryController{}, "get:MissData")
	beego.Router("/putdata/", &controllers.LotteryController{}, "get:PutData")

}
