package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	//	"bufio"
	//	"encoding/json"
	//	"encoding/base64"
	"io/ioutil"
	"stocknew/fortune/db"
	"stocknew/fortune/models"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	//	"github.com/gorilla/schema"
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
	logs.Info("画出曲线图，股票代码：%v 天数：%v", code, datesize)
	c.Render()
}

type PushMLDataController struct {
	beego.Controller
}

type Item struct {
	Date   string  `json:"date"`
	Highpx float64 `json:"highpx"` //对应表单中的name值,字段名首字母也必须大写，否则无法解析该参数的值
	Lowpx  float64 `json:"lowpx"`
}

type Mlobject struct {
	MyItems   []Item `json:"myItems"`
	PushPoint []int  `json:"pushPoint"` //对应表单中的name值,字段名首字母也必须大写，否则无法解析该参数的值
}

func (this *PushMLDataController) PushDataToDB() {
	ob := &Mlobject{}
	//	err := this.ParseForm(ob)
	//	if err != nil {
	//		fmt.Printf("解析表单数据失败! error:%v", err)
	//	}
	//	var decoder = schema.NewDecoder()
	//	err = decoder.Decode(&ob, this.Post())
	//	if err != nil {
	//		fmt.Println("解码表单数据失败!")
	//		fmt.Println(err)
	//	}

	//	for k, v := range ob.MyItems {
	//		fmt.Println(k, v)
	//	}

	//	logs.Info("数据 %v", string(this.Ctx.Input.RequestBody))

	s := strings.NewReader(string(this.Ctx.Input.RequestBody))
	result, err := ioutil.ReadAll(s)
	if err != nil {
		logs.Info("解析数据失败，原因为 %v", err)
	}

	err = json.Unmarshal(result, &ob)
	if err != nil {
		logs.Info("解析json数据失败，原因为 %v", err)
	}
	logs.Info("后端拿到的数据为：%v", ob)
}
