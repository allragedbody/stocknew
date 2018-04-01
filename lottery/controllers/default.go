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
	c.TplName = "index.html"
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

