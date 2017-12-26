package routers

import (
	"fortune/controllers"

	"github.com/astaxie/beego"
)

func Init() {
	beego.Router("/", &controllers.MainController{})
	//	beego.Router("/gettrend", &controllers.TrendController{})
	beego.Router("/stockdatedata/:code/:datesize", &controllers.StockController{}, "get:StockDateData")
	//示例：http://127.0.0.1/stockdatedata/000158.sz/60
	beego.Router("/draw", &controllers.DrawStockController{})
}
