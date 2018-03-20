package controllers

import (
	//	"encoding/json"
	//	"fmt"
	//	"strconv"
	//	"bufio"
	//	"encoding/json"
	//	"encoding/base64"
	//	"io/ioutil"
	"stocknew/lottery/db"
	"stocknew/lottery/process"
	//	"strings"

	"github.com/astaxie/beego"
	//	"github.com/astaxie/beego/logs"
	//	"github.com/gorilla/schema"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "numberview.tpl"
}

type LotteryController struct {
	beego.Controller
}
type Response struct {
	HistoryDatas [][]string `json:"historydatas"`
}

func (this *LotteryController) LotteryData() {
	rsp := &Response{}

	size, err := this.GetInt(":size")
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
	dbconn := db.GetDB()
	data, err := dbconn.GetLotterData(size)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
	rsp.HistoryDatas = data

	this.Data["json"] = rsp

	this.ServeJSON()
	//	this.Ctx.WriteString(info)
}

func (this *LotteryController) MissData() {
	this.Data["json"] = process.MissDataLottery
	this.ServeJSON()

}
func (this *LotteryController) PutData() {
	this.Data["json"] = process.PuttoLottery
	this.ServeJSON()
}
func (this *LotteryController) GetPlan() {
	this.Data["json"] = process.LotterPlans
	this.ServeJSON()
}

//type DrawStockController struct {
//	beego.Controller
//}

//func (c *DrawStockController) Get() {
//	code := c.GetString("code")
//	datesize, _ := c.GetInt("datesize")

//	c.Data["code"] = code
//	c.Data["datesize"] = datesize
//	c.TplName = "stockdraw.html"
//	logs.Info("画出曲线图，股票代码：%v 天数：%v", code, datesize)
//	c.Render()
//}

//type PushMLDataController struct {
//	beego.Controller
//}

//type Item struct {
//	Date   string  `json:"date"`
//	Highpx float64 `json:"highpx"` //对应表单中的name值,字段名首字母也必须大写，否则无法解析该参数的值
//	Lowpx  float64 `json:"lowpx"`
//}

//type Mlobject struct {
//	MyItems   []Item `json:"myItems"`
//	PushPoint []int  `json:"pushPoint"` //对应表单中的name值,字段名首字母也必须大写，否则无法解析该参数的值
//}

//func (this *PushMLDataController) PushDataToDB() {
//	ob := &Mlobject{}
//	//	err := this.ParseForm(ob)
//	//	if err != nil {
//	//		fmt.Printf("解析表单数据失败! error:%v", err)
//	//	}
//	//	var decoder = schema.NewDecoder()
//	//	err = decoder.Decode(&ob, this.Post())
//	//	if err != nil {
//	//		fmt.Println("解码表单数据失败!")
//	//		fmt.Println(err)
//	//	}

//	//	for k, v := range ob.MyItems {
//	//		fmt.Println(k, v)
//	//	}

//	//	logs.Info("数据 %v", string(this.Ctx.Input.RequestBody))

//	s := strings.NewReader(string(this.Ctx.Input.RequestBody))
//	result, err := ioutil.ReadAll(s)
//	if err != nil {
//		logs.Info("解析数据失败，原因为 %v", err)
//	}

//	err = json.Unmarshal(result, &ob)
//	if err != nil {
//		logs.Info("解析json数据失败，原因为 %v", err)
//	}
//	logs.Info("后端拿到的数据为：%v", ob)
//}
