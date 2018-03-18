package main

import (
	"fmt"
	"stocknew/lottery/db"
	"stocknew/lottery/process"
	"stocknew/lottery/routers"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	//	"github.com/astaxie/beego/logs"
)

//var FirstPeriods = 669754
//var firstNumber int
var puttolettery []int
var missdatalettery []int

func initLog() error {
	//	jsonConfig := `{
	//        "filename" : "test.log",
	//		"level":7
	//    }`

	//	logs.SetLogger("file", jsonConfig)
	//	logs.SetLogFuncCall(true)
	return nil
}

func main() {
	err := initLog()
	if err != nil {
		fmt.Println("init log error:%v", err)
		return
	}

	logs.Info("这是一个彩票游戏。作者 邓云飞。")
	err = db.Init()
	if err != nil {
		return
	}
	process.Init()
	logs.Info("開始加載數據。")
	go reloadLotteryData()
	go caculateData()
	routers.Init()
	beego.SetStaticPath("/views", "views")
	beego.Run()
}

//func runLotteryFirst(firstNumber int, n1 *model.PKTen) {
//	fmt.Println("开奖。")
//	fmt.Printf("本次开出第一名的数字是 %v。\n", firstNumber)
//	fmt.Println("存储上一期开奖信息。")
//	fmt.Println("当前信息修正。")

//	rewardNumbers := make([]int, 0)
//	rewardNumbers = append(rewardNumbers, firstNumber)
//	n1.ChangeInfos(FirstPeriods, rewardNumbers, n1)
//	fmt.Printf("当前期次 %v：\n", n1.Periods)
//	for _, nu := range n1.FirstInfos {
//		fmt.Printf("号码 %v 遗漏次数 %v 最大遗漏次数 %v 平均遗漏次数%v。\n", nu.Number, nu.Miss, nu.MAXMiss, nu.AvarageMiss)
//	}
//}

func reloadLotteryData() {
	for {
		time.Sleep(time.Second * 10)
		logs.Info("获取彩票历史数据。")
		data, err := process.GetHistoryData()
		if err != nil {
			logs.Error("获取彩票历史数据失败。", err)
			continue
		}
		logs.Info("data %v", data)
		for k, v := range data {
			numlist := strings.Split(v.Number, ",")
			logs.Info("期数 %v,号码 %v", k, numlist)
			err = process.RestoreData(k, numlist)
			if err != nil {
				fmt.Printf("第 %v 期彩票数据失败 %v", k, err)
				continue
			}
			fmt.Printf("插入彩票数据成功。data %v.", data)
		}

	}
}

var currentPierod string
var putTime int

func caculateData() {
	for {
		time.Sleep(time.Second * 10)
		logs.Info("获取彩票数据库数据。")
		data, err := process.GetDBData()
		if err != nil {
			logs.Error("获取彩票数据库数据失败。", err)
			continue
		}
		if currentPierod != data[0][0] {
			currentPierod = data[0][0]
		}

		alldata, err := process.CalculateMiss(data)
		if err != nil {
			logs.Error("获取彩票数据库数据失败。", err)
			continue
		}
		missdata := make([]int, 0)
		missdata = append(missdata, alldata["1"].MissTime)
		missdata = append(missdata, alldata["2"].MissTime)
		missdata = append(missdata, alldata["3"].MissTime)
		missdata = append(missdata, alldata["4"].MissTime)
		missdata = append(missdata, alldata["5"].MissTime)
		missdata = append(missdata, alldata["6"].MissTime)
		missdata = append(missdata, alldata["7"].MissTime)
		missdata = append(missdata, alldata["8"].MissTime)
		missdata = append(missdata, alldata["9"].MissTime)
		missdata = append(missdata, alldata["10"].MissTime)

		process.MissDataLottery = missdata
		logs.Info("计算遗漏数据为 %v", process.MissDataLottery)

		putdata := process.CalculatePut(missdata)

		if len(puttolettery) == 0 {
			process.PuttoLottery = putdata
			if currentPierod != data[0][0] {
				putTime += 1
			} else {
				putTime = 1
			}
		} else {
			for _, i := range puttolettery {
				ne, _ := strconv.Atoi(data[0][1])
				if ne == i {
					process.PuttoLottery = putdata
					if currentPierod != data[0][0] {
						putTime = 1
					} else {
						putTime = 1
					}
				}
			}
		}

		logs.Info("计算下注数据为 %v 次数为 %v", process.PuttoLottery, putTime)
	}
}
