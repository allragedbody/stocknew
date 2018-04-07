package controllers

import (
	//	"encoding/json"
	//	"fmt"
	//	"strconv"
	//	"bufio"
	//	"encoding/json"
	//	"encoding/base64"
	//	"io/ioutil"
	"stocknew/lotterytools/db"
	"stocknew/lotterytools/process"
	//	"strings"

	"github.com/astaxie/beego"
	//	"github.com/astaxie/beego/logs"
	//	"github.com/gorilla/schema"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "numberview.tpl"
}

type LotteryController struct {
	beego.Controller
}
type Response struct {
	HistoryDatas [][]string `json:"historydatas"`
}

func (this *LotteryController) LotteryData() {
	this.AllowCross() //允许跨域
	rsp := &Response{}

	size, err := this.GetInt(":size")
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
	dbconn := db.GetDB()
	data, err := dbconn.GetLotterData(0, size)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
	rsp.HistoryDatas = data

	this.Data["json"] = rsp

	this.ServeJSON()
	//	this.Ctx.WriteString(info)
}

func (this *LotteryController) MissDataList() {
	this.AllowCross() //允许跨域
	rsp := &Response{}

	size, err := this.GetInt(":size")
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
	dbconn := db.GetDB()
	data, err := dbconn.GetMissData(0, size)
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
	this.AllowCross() //允许跨域

	this.Data["json"] = process.MissDataLottery
	this.ServeJSON()

}
func (this *LotteryController) PutData() {
	this.AllowCross() //允许跨域

	this.Data["json"] = process.PuttoLotteryMax4
	this.ServeJSON()
}
func (this *LotteryController) GetPlan() {
	this.AllowCross() //允许跨域

	this.Data["json"] = process.LotterPlans
	this.ServeJSON()
}

func (this *LotteryController) GetDateData() {
	this.AllowCross() //允许跨域

	this.Data["json"] = process.DateData
	this.ServeJSON()
}

func (this *LotteryController) ImportantMiss() {
        this.AllowCross() //允许跨域

        this.Data["json"] = process.ImportantMiss
        this.ServeJSON()
}


type BaseController struct {
	beego.Controller
}

func (c *BaseController) Options() {
	c.AllowCross() //允许跨域
	c.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo": ""}
	c.ServeJSON()
}

func (c *BaseController) AllowCross() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")                           //允许访问源
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")    //允许post访问
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization") //header的类型
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Max-Age", "1728000")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Ctx.ResponseWriter.Header().Set("content-type", "application/json") //返回数据格式是json
}

func (c *LotteryController) AllowCross() {
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")                           //允许访问源
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")    //允许post访问
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization") //header的类型
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Max-Age", "1728000")
	c.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Ctx.ResponseWriter.Header().Set("content-type", "application/json") //返回数据格式是json
}
