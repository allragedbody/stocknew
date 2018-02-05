package controllers

import (
	"fmt"
	"strconv"
	//	"encoding/json"
	"stocknew/fortune/db"

	"stocknew/fortune/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "getStockInfo.tpl"
}

//type TrendController struct {
//	beego.Controller
//}

//func (this *TrendController) Get() {
//	this.Ctx.WriteString("hello")
//	//12345676545678987654567
//}

type StockController struct {
	beego.Controller
}

func (this *StockController) StockDateData() {

	rsp := &models.Response{}
	code := this.GetString(":code")
	datesize, err := this.GetInt(":datesize")
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
	dbconn := db.GetDB()
	data, err := dbconn.GetStockDateData(code, datesize)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}

	dataReverse := make([]*models.DayInfo, 0)

	//	info := fmt.Sprintf("获取代码为：%v %v日的数据。\n", code, datesize)

	//	for _, v := range data {
	//		info = fmt.Sprintf("%v\n%v", info, v)
	//	}
	for i := len(data) - 1; i >= 0; i-- {
		d := &models.DayInfo{}
		d.Code = data[i].Code
		d.Date = data[i].Date
		f1, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", data[i].OpenPx), 64)
		f2, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", data[i].ClosePx), 64)
		f3, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", data[i].HighPx), 64)
		f4, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", data[i].LowPx), 64)
		d.OpenPx = f1
		d.ClosePx = f2
		d.HighPx = f3
		d.LowPx = f4
		d.BusinessAmount = data[i].BusinessAmount
		d.BusinessBalance = data[i].BusinessBalance
		dataReverse = append(dataReverse, data[i])
	}

	rsp.HistoryDatas = dataReverse
	//	rsp.MaxPoints = rsp.GetMaxPoints(dataReverse)
	//	rsp.MinPoints = rsp.GetMinPoints(dataReverse)

	this.Data["json"] = rsp

	this.ServeJSON()
	//	this.Ctx.WriteString(info)

}

type DrawStockController struct {
	beego.Controller
}

func (c *DrawStockController) Get() {
	code := c.GetString("code")
	datesize, _ := c.GetInt("datesize")

	c.Data["code"] = code
	c.Data["datesize"] = datesize
	c.TplName = "stockdraw.html"
	c.Render()
}

type PushMLDataController struct {
	beego.Controller
}

func (c *PushMLDataController) Post() {
	c.Render()
}
