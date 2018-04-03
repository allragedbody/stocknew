package main

import (
	"fmt"
	"stocknew/lotterytools/db"
	"stocknew/lotterytools/process"
	"stocknew/lotterytools/routers"
	"strconv"
	//	"strings"
	//	"sort"
	"stocknew/lotterytools/model"
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

	logs.Info("这是一个辅助工具。作者 邓云飞。")
	err = db.Init()
	if err != nil {
		return
	}
	process.Init()
	logs.Info("開始加載數據。")
	//	go caculateData()
	go caculateDataMax4()
	routers.Init()
	beego.SetStaticPath("/views", "views")
	beego.Run()
}

var currentPierod string
var putTime int

//func caculateData() {
//	lotterPlans := make([]model.LotterPlan, 0)
//	sendtime := 0
//	var lastmysqlplan int
//	var isadd bool
//	finalnums := make([]string, 0)

//	for {
//		time.Sleep(time.Second * 10)
//		logs.Info("获取彩票数据库数据。")

//		data, err := process.GetDBData(lastmysqlplan)
//		if err != nil {
//			logs.Error("获取彩票数据库数据失败。", err)
//			continue
//		}

//		alldata, err := process.CalculateMiss(data)
//		if err != nil {
//			logs.Error("获取彩票数据库数据失败。", err)
//			continue
//		}
//		missdata := make([]int, 0)
//		missdata = append(missdata, alldata["1"].MissTime)
//		missdata = append(missdata, alldata["2"].MissTime)
//		missdata = append(missdata, alldata["3"].MissTime)
//		missdata = append(missdata, alldata["4"].MissTime)
//		missdata = append(missdata, alldata["5"].MissTime)
//		missdata = append(missdata, alldata["6"].MissTime)
//		missdata = append(missdata, alldata["7"].MissTime)
//		missdata = append(missdata, alldata["8"].MissTime)
//		missdata = append(missdata, alldata["9"].MissTime)
//		missdata = append(missdata, alldata["10"].MissTime)

//		process.MissDataLottery = missdata
//		logs.Info("计算遗漏数据为 %v", process.MissDataLottery)

//		if len(lotterPlans) == 0 {
//			putdata := process.CalculatePut(missdata)
//			plan := model.LotterPlan{}
//			nextPeriodNum, _ := strconv.Atoi(data[0][0])
//			plan.CurrentPierod = strconv.Itoa(nextPeriodNum + 1)
//			plan.NumberList = putdata
//			plan.GetReward = false
//			plan.PutTime = 1
//			plan.Status = "等开"
//			plan.RealPutTime = 0
//			plan.CreateTime = time.Now().Format("2006-01-02 15:04:05")
//			process.PuttoLottery = putdata

//			logs.Info("第一次计算下注数据为 %v ", plan)
//			//存数据库
//			err := process.RestorePlan(plan)
//			if err != nil {
//				logs.Error("RestorePlanToDB [%v] err: %v", plan, err)
//			}
//			lotterPlans = append(lotterPlans, plan)
//		} else {
//			//是否已经变更期数
//			l := len(lotterPlans)
//			lastplan := lotterPlans[l-1]
//			logs.Info("最后一次计算下注数据为 %v ", lastplan)

//			nextPeriodNum, _ := strconv.Atoi(data[0][0])
//			c, _ := strconv.Atoi(lastplan.CurrentPierod)

//			if nextPeriodNum == c {
//				//如果不中 则等开1 变更为倍投 2 ，行数不增加
//				//如果中了 则变为中，增加一行。
//				for _, i := range lastplan.NumberList {
//					rewardNum, _ := strconv.Atoi(data[0][1])
//					if rewardNum == i {
//						lastplan.GetReward = true
//						lastplan.Status = "中"

//						lotterPlans = lotterPlans[0 : l-1]
//						lotterPlans = append(lotterPlans, lastplan)
//						putdata := process.CalculatePut(missdata)
//						plan := model.LotterPlan{}
//						nextPeriodNum, _ := strconv.Atoi(data[0][0])
//						plan.CurrentPierod = strconv.Itoa(nextPeriodNum + 1)
//						plan.NumberList = putdata
//						plan.GetReward = false
//						plan.PutTime = 1
//						plan.RealPutTime = 0
//						plan.Status = "等开"
//						plan.CreateTime = time.Now().UTC().Format("2006-01-02 15:04:05")
//						process.PuttoLottery = putdata
//						lotterPlans = append(lotterPlans, plan)
//						logs.Info("中奖了，计算下注数据为 %v ", plan)
//						finalnums = make([]string, 0)
//						isadd = false
//						//存数据库
//						err := process.RestorePlan(plan)
//						if err != nil {
//							logs.Error("RestorePlanToDB [%v] err: %v", plan, err)
//						}
//						break
//					}
//				}
//				if lastplan.GetReward != true {
//					lastplan.PutTime += 1
//					if lastplan.PutTime > 2 {
//						lastplan.RealPutTime = lastplan.PutTime - 2
//					}

//					lastplan.Status = "倍投"
//					nextPeriodNum, _ := strconv.Atoi(data[0][0])
//					lastplan.CurrentPierod = strconv.Itoa(nextPeriodNum + 1)
//					lotterPlans = lotterPlans[0 : l-1]

//					if lastplan.PutTime > 5 {
//						//补充两个号
//						logs.Debug("倒数第1个第一%v", data[0][1])
//						logs.Debug("倒数第2个第一%v", data[1][1])
//						logs.Debug("倒数第3个第一%v", data[2][1])
//						logs.Debug("倒数第4个第一%v", data[3][1])
//						logs.Debug("倒数第5个第一%v", data[4][1])

//						ld := make(map[string]int, 0)
//						_, ok := ld[data[0][1]]
//						if ok {
//							ld[data[0][1]] += 1
//						} else {
//							ld[data[0][1]] = 1
//						}
//						_, ok = ld[data[1][1]]
//						if ok {
//							ld[data[1][1]] += 1
//						} else {
//							ld[data[1][1]] = 1
//						}
//						_, ok = ld[data[2][1]]
//						if ok {
//							ld[data[2][1]] += 1
//						} else {
//							ld[data[2][1]] = 1
//						}
//						_, ok = ld[data[3][1]]
//						if ok {
//							ld[data[3][1]] += 1
//						} else {
//							ld[data[3][1]] = 1
//						}
//						_, ok = ld[data[4][1]]
//						if ok {
//							ld[data[4][1]] += 1
//						} else {
//							ld[data[4][1]] = 1
//						}

//						logs.Debug("计数 %v", ld)

//						addnum := make([]string, 0)

//						for k, v := range ld {
//							if v == 5 {
//								addnum = append(addnum, k)
//								continue
//							}
//						}
//						if len(addnum) < 2 {
//							for k, v := range ld {
//								if v == 4 {
//									addnum = append(addnum, k)
//									continue
//								}
//							}
//						}
//						if len(addnum) < 2 {
//							for k, v := range ld {
//								if v == 3 {
//									addnum = append(addnum, k)
//									continue
//								}
//							}
//						}
//						if len(addnum) < 2 {
//							for k, v := range ld {
//								if v == 2 {
//									addnum = append(addnum, k)
//									continue
//								}
//							}
//						}
//						if len(addnum) < 2 {
//							for k, v := range ld {
//								if v == 1 {
//									addnum = append(addnum, k)
//									continue
//								}
//							}
//						}

//						if len(finalnums) == 0 {
//							finalnums = addnum[0:2]
//							logs.Debug("增加两个号码为 %v", finalnums)
//						}
//						if !isadd {
//							for _, fl := range finalnums {
//								fln, _ := strconv.Atoi(fl)
//								lastplan.NumberList = append(lastplan.NumberList, fln)
//							}

//							isadd = true
//						}
//						process.PuttoLottery = lastplan.NumberList

//					} else {
//						process.PuttoLottery = lastplan.NumberList

//					}

//					lastplan.CreateTime = time.Now().UTC().Format("2006-01-02 15:04:05")
//					if lastplan.RealPutTime > 0 {
//						if sendtime > 1 {
//							logs.Info("超过1次不再提醒")
//						} else {
//							logs.Info("发短信给企业号 内容为 %v ", lastplan)
//							err := process.SendWeChat(lastplan)
//							if err != nil {
//								logs.Info("发送计划失败： %v ", err)
//							} else {
//								sendtime += 1
//							}
//						}

//					} else {
//						sendtime = 0
//					}

//					lotterPlans = append(lotterPlans, lastplan)
//					logs.Info("未中奖,计算下注数据为 %v ", lastplan)
//					//存数据库
//					err := process.RestorePlan(lastplan)
//					if err != nil {
//						logs.Error("RestorePlanToDB [%v] err: %v", lastplan, err)
//					}
//				}
//			}
//		}
//		logs.Info("计算最终下注数据为: ")
//		for _, plans := range lotterPlans {
//			logs.Info("数据列表： %v ", plans)
//		}
//		n := len(lotterPlans)

//		lastmysqlplan, _ = strconv.Atoi(lotterPlans[n-1].CurrentPierod)
//		process.RestoreLotterResult(lotterPlans)
//	}
//}

func caculateDataMax4() {
	var isadd bool
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
		//存数据库
		err = process.RestoreMissData(data[0][0], missdata)
		if err != nil {
			logs.Error("RestoreMissData [%v-%v] err: %v", data[0][0], missdata, err)
		}

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
			process.PuttoLotteryMax4 = putdata

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
						isadd = false
						plan.CreateTime = time.Now().UTC().Format("2006-01-02 15:04:05")
						process.PuttoLotteryMax4 = putdata
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
					if lastplan.PutTime > 3 {
						newputdata := process.NewCalculatePut(lastplan.NumberList, missdata)
						lpn := lastplan.NumberList
						if !isadd {
							lastplan.NumberList = getNewPlan(lpn, newputdata)
							isadd = true
						}

						process.PuttoLotteryMax4 = lastplan.NumberList
					} else {
						process.PuttoLotteryMax4 = lastplan.NumberList

					}

					lastplan.CreateTime = time.Now().UTC().Format("2006-01-02 15:04:05")

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
		process.GetDateData()
	}
}

func getNewPlan(a, b []int) []int {
	m := make(map[int]int, 0)
	l := make([]int, 0)
	for _, i := range a {
		_, ok := m[i]
		if !ok {
			m[i] = i
		}
	}
	for _, j := range b {
		_, ok := m[j]
		if !ok {
			m[j] = j
		}
	}
	for k, _ := range m {
		l = append(l, k)
	}
	return l
}
