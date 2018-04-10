package main

import (
	"fmt"
	"stocknew/lottery/db"
	"stocknew/lottery/process"
	"stocknew/lottery/routers"
	"strconv"
	"strings"
	"time"

	"stocknew/lottery/model"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	//	"github.com/astaxie/beego/logs"
)

//var FirstPeriods = 669754
//var firstNumber int
var puttolettery []int
var missdatalettery []int

func initLog() error {
	jsonConfig := `{
	        "filename" : "d:/lottery.log",
			"level":7
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
		time.Sleep(time.Second * 5)
		logs.Info("获取彩票历史数据。")
		data, err := process.GetHistoryData()
		if err != nil {
			logs.Error("获取彩票历史数据失败。错误：%v", err)
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
	lotterPlans := make([]model.LotterPlan, 0)
	var lastmysqlplan int
	for {
		time.Sleep(time.Second * 10)
		logs.Info("获取彩票数据库数据。")

		data, err := process.GetDBData(lastmysqlplan)
		if err != nil {
			logs.Error("获取彩票数据库数据失败。", err)
			continue
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

		if len(lotterPlans) == 0 {
			putdata := process.CalculatePut(missdata)
			plan := model.LotterPlan{}
			nextPeriodNum, _ := strconv.Atoi(data[0][0])
			plan.CurrentPierod = strconv.Itoa(nextPeriodNum + 1)
			plan.NumberList = putdata
			plan.GetReward = false
			plan.PutTime = 1
			plan.Status = "等开"
			plan.RealPutTime = 0
			plan.CreateTime = time.Now().Format("2006-01-02 15:04:05")
			process.PuttoLottery = putdata

			logs.Info("第一次计算下注数据为 %v ", plan)
			//存数据库
			err := process.RestorePlan(plan)
			if err != nil {
				logs.Error("RestorePlanToDB [%v] err: %v", plan, err)
			}
			lotterPlans = append(lotterPlans, plan)
		} else {
			//是否已经变更期数
			l := len(lotterPlans)
			lastplan := lotterPlans[l-1]
			logs.Info("最后一次计算下注数据为 %v ", lastplan)

			nextPeriodNum, _ := strconv.Atoi(data[0][0])
			c, _ := strconv.Atoi(lastplan.CurrentPierod)

			if nextPeriodNum == c {
				//如果不中 则等开1 变更为倍投 2 ，行数不增加
				//如果中了 则变为中，增加一行。
				for _, i := range lastplan.NumberList {
					rewardNum, _ := strconv.Atoi(data[0][1])
					if rewardNum == i {
						lastplan.GetReward = true
						lastplan.Status = "中"
						lotterPlans = lotterPlans[0 : l-1]
						lotterPlans = append(lotterPlans, lastplan)
						putdata := process.CalculatePut(missdata)
						plan := model.LotterPlan{}
						nextPeriodNum, _ := strconv.Atoi(data[0][0])
						plan.CurrentPierod = strconv.Itoa(nextPeriodNum + 1)
						plan.NumberList = putdata
						plan.GetReward = false
						plan.PutTime = 1
						plan.RealPutTime = 0
						plan.Status = "等开"
						plan.CreateTime = time.Now().Format("2006-01-02 15:04:05")
						process.PuttoLottery = putdata
						lotterPlans = append(lotterPlans, plan)
						logs.Info("中奖了，计算下注数据为 %v ", plan)
						//存数据库
						err := process.RestorePlan(plan)
						if err != nil {
							logs.Error("RestorePlanToDB [%v] err: %v", plan, err)
						}
						break
					}
				}
				if lastplan.GetReward != true {
					lastplan.PutTime += 1
					if lastplan.PutTime > 2 {
						lastplan.RealPutTime = lastplan.PutTime - 2
					}

					lastplan.Status = "倍投"
					nextPeriodNum, _ := strconv.Atoi(data[0][0])
					lastplan.CurrentPierod = strconv.Itoa(nextPeriodNum + 1)
					lotterPlans = lotterPlans[0 : l-1]
					process.PuttoLottery = lastplan.NumberList
					lastplan.CreateTime = time.Now().Format("2006-01-02 15:04:05")
					lotterPlans = append(lotterPlans, lastplan)
					logs.Info("未中奖,计算下注数据为 %v ", lastplan)
					//存数据库
					err := process.RestorePlan(lastplan)
					if err != nil {
						logs.Error("RestorePlanToDB [%v] err: %v", lastplan, err)
					}
				}
			}
		}
		logs.Info("计算最终下注数据为: ")
		for _, plans := range lotterPlans {
			logs.Info("数据列表： %v ", plans)
		}
		n := len(lotterPlans)

		lastmysqlplan, _ = strconv.Atoi(lotterPlans[n-1].CurrentPierod)
		process.RestoreLotterResult(lotterPlans)
	}
}

