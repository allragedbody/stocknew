package main

import (
	"fmt"
	"stocknew/ailottery/db"
	"stocknew/ailottery/process"
	"stocknew/ailottery/routers"
	//	"strconv"
	//	"strings"
	//	"sort"
	//	"stocknew/ailottery/model"
	//	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	//	"github.com/astaxie/beego/logs"
)

func initLog() error {
	jsonConfig := `{
	        "filename" : "./AILottery.log",
			"level":6
	    }`

	logs.SetLogger("file", jsonConfig)
	logs.SetLogFuncCall(true)
	return nil
}

func main() {
	err := initLog()
	if err != nil {
		fmt.Println("init log error:%v", err)
		return
	}

	logs.Info("这是一个辅助工具。作者 邓云飞。")
	err = db.Init()
	if err != nil {
		return
	}
	process.Init()
	logs.Info("开始运行准备。")
	recordSize, _ := beego.AppConfig.Int("knnsize20")
	err = process.DataPrepare(recordSize)
	if err != nil {
		logs.Error("数据未准备好，重试。")
		return
	}
	logs.Info("开始运行Reload。")
	go process.DataReload(recordSize)
	logs.Info("开始运行计算。")
	go process.Running(recordSize)
	logs.Info("开始运行Webserver。")
	routers.Init()
	beego.SetStaticPath("/views", "views")
	beego.Run()
}

