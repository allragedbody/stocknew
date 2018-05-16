package process

import (
	"stocknew/ailottery/db"
	"stocknew/ailottery/model"
	"stocknew/ailottery/nettools"
	"strconv"
	//	"github.com/axgle/mahonia"
	//	"bytes"
	"errors"
	//	"fmt"
	//	"io/ioutil"
	//	"net/http"
	//	"sort"
	"bufio"
	"os"
	"time"
	//	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

var (
	AIDIR = "aidir"
)
var OD *OriData
var lotterPlans []model.LotterPlan

type OriData struct {
}

func (oridata *OriData) GetData(size int) [][]int {
	logs.Debug("oridata GetData")
	T := &KNN{}
	T.K = 3
	T.TrainingSet.KT = &model.KNNTrainingSet{}
	T.TrainingSet.KT.Size = size

	knnlist := T.GetData(size)

	logs.Debug("KNN获取数据 %v ", knnlist)
	return knnlist
}

func (oridata *OriData) GetMissData(size int) [][]int {
	logs.Debug("oridata GetMissData")
	T := &KNN{}
	T.K = 3
	T.TrainingSet.KT = &model.KNNTrainingSet{}
	T.TrainingSet.KT.Size = size

	knnlist := T.GetMissData(size)

	logs.Debug("KNN获取数据 %v ", knnlist)
	return knnlist
}

func (oridata *OriData) HandleData(traninglist [][]int) error {
	logs.Debug("KNN处理数据 %v ", traninglist)
	return nil
}

func (oridata *OriData) StoreData(traninglist [][]int, size int) error {
	logs.Debug("KNN存储数据 %v ", traninglist)
	f, err := os.Create(AIDIR + "/knnlist" + strconv.Itoa(size) + ".tmp") //创建文件
	if err != nil {
		return err
	}
	defer f.Close()

	str := ""

	for i, v := range traninglist {
		if i == len(traninglist)-1 {
			for ii, vv := range v {
				if ii == len(v)-1 {
					str += strconv.Itoa(vv)
				} else {
					str += strconv.Itoa(vv) + ","
				}
			}
		} else {
			for ii, vv := range v {
				if ii == len(v)-1 {
					str += strconv.Itoa(vv)
				} else {
					str += strconv.Itoa(vv) + ","
				}
			}
			str += "\n"
		}

	}
	w := bufio.NewWriter(f)
	_, err = w.WriteString(str)

	if err != nil {
		return err
	}
	w.Flush()
	f.Close()

	err = os.Rename(AIDIR+"/knnlist"+strconv.Itoa(size)+".tmp", AIDIR+"/knnlist"+strconv.Itoa(size)+".txt")
	if err != nil {
		return err
	}

	return nil
}

func (oridata *OriData) StoreMissData(traninglist [][]int, size int) error {
	logs.Debug("KNN存储数据 %v ", traninglist)
	f, err := os.Create(AIDIR + "/missknnlist" + strconv.Itoa(size) + ".tmp") //创建文件
	if err != nil {
		return err
	}
	defer f.Close()

	str := ""

	for i, v := range traninglist {
		if i == len(traninglist)-1 {
			for ii, vv := range v {
				if ii == len(v)-1 {
					str += strconv.Itoa(vv)
				} else {
					str += strconv.Itoa(vv) + ","
				}
			}
		} else {
			for ii, vv := range v {
				if ii == len(v)-1 {
					str += strconv.Itoa(vv)
				} else {
					str += strconv.Itoa(vv) + ","
				}
			}
			str += "\n"
		}

	}
	w := bufio.NewWriter(f)
	_, err = w.WriteString(str)

	if err != nil {
		return err
	}
	w.Flush()
	f.Close()

	err = os.Rename(AIDIR+"/missknnlist"+strconv.Itoa(size)+".tmp", AIDIR+"/missknnlist"+strconv.Itoa(size)+".txt")
	if err != nil {
		return err
	}

	return nil
}

func (oridata *OriData) CalculateData(size int) error {
	dbconn := db.GetDB()
	data, err := dbconn.GetLotterData(0, size+1)
	rdata := reverse(data)

	param := ""
	for _, v := range rdata {
		logs.Debug("第 %v 期数据为 %v", v[0], v[1])
		param += " " + v[1]
	}

	next, _ := strconv.Atoi(rdata[len(rdata)-1][0])

	//推测号码为：
	calculateNumbers, err := oridata.KNNCalculate(param, size)
	if err != nil {
		logs.Error("%v 数据推测计算错误 %v", next+1, err)
		return err
	}

	logs.Debug("%v 数据推测为 %v", next+1, calculateNumbers)
	return nil
}

func (oridata *OriData) CalculateMissData(size int) error {
	dbconn := db.GetDB()
	data, err := dbconn.GetMissData(0, size+1)
	if err != nil {
		return err
	}
	if len(data) == 0 {
		return errors.New("CalculateMissData get data zero.")
	}

	rdata := reverse(data)
	logs.Debug("KNN GetMissData compare %v %v", data, rdata)

	param := ""
	for index, v := range rdata {
		if index != 0 {
			logs.Debug("第 %v 期miss数据为 %v", v[0], v[1:])
			param += " " + v[1] + " " + v[2] + " " + v[3] + " " + v[4] + " " + v[5] + " " + v[6] + " " + v[7] + " " + v[8] + " " + v[9] + " " + v[10]
		}

	}

	next, _ := strconv.Atoi(rdata[len(rdata)-1][0])
	lastliststr := rdata[len(rdata)-1][1:]
	logs.Debug("最後，第 %v 期miss数据为 %v", next, lastliststr)
	lastlist := make([]int, 0)
	for _, v := range lastliststr {
		t, _ := strconv.Atoi(v)
		lastlist = append(lastlist, t)
	}

	//推测号码为：
	calculateNumbers, err := oridata.KNNMissCalculate(param, size, lastlist)
	if err != nil {
		logs.Error("%v miss数据推测计算错误 %v", next+1, err)
		return err
	}
	if len(calculateNumbers) == 0 {
		logs.Debug("%v miss数据推测为 %v", next+1, "未知")
		KnnMissCount = calculateNumbers
	} else {
		logs.Debug("%v miss数据推测为 %v", next+1, calculateNumbers)
		KnnMissCount = calculateNumbers
	}

	return nil
}

func Init() {
	OD = &OriData{}
	c = nettools.CreateClient()
	lotterPlans = make([]model.LotterPlan, 0)
}

func Running() {
	recordSize, _ := beego.AppConfig.Int("knnsize20")
	err := DataPrepare(recordSize)
	if err != nil {
		logs.Error("数据未准备好，重试。")
		return
	}
	go DataReload(recordSize)
	for {
		time.Sleep(time.Second * 10)
		logs.Info("进行对k临近算法的运算。")
		OD.CalculateMissData(recordSize)
		CaculateDataByAI()
	}
}

func DataPrepare(recordSize int) error {
	orimissdata := OD.GetMissData(recordSize)
	OD.HandleData(orimissdata)
	err := OD.StoreMissData(orimissdata, recordSize)
	if err != nil {
		return err
	}

	return nil
}

func DataReload(recordSize int) {
	for {
		time.Sleep(time.Second * 3600)
		orimissdata := OD.GetMissData(recordSize)
		OD.HandleData(orimissdata)
		err := OD.StoreMissData(orimissdata, recordSize)
		if err != nil {
			logs.Error("数据未准备好，重试。")
			continue
		}
	}
}

type GetData interface {
	GetData()
}

type HandleData interface {
	HandleData()
}

type StoreData interface {
	StoreData()
}

type CalculateData interface {
	CalculateData()
}

type Trainning struct {
	TrainingSet    model.TrainingSet
	CalculationSet model.CalculationSet
}

func CaculateDataByAI() {

	sendtime := 0
	var lastmysqlplan int

	data, err := GetDBData(lastmysqlplan)
	if err != nil {
		logs.Error("获取数据库数据失败。", err)
		return
	}

	alldata, err := CalculateMiss(data)
	if err != nil {
		logs.Error("获取数据库数据失败。", err)
		return
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

	if fiveTenAndOneNine(missdata) || sixTen(missdata) {
		RestoreImportantMiss(missdata)
		if sendtime > 0 {
			logs.Info("超过1次不再提醒")
		} else {
			logs.Info("发短信给企业号 内容为 %v ", missdata)
			err := SendWeChat(missdata)
			if err != nil {
				logs.Info("发送计划失败： %v ", err)
			} else {
				sendtime += 1
			}
		}
	} else {
		tmpdata := make([]int, 10)
		RestoreImportantMiss(tmpdata)
		sendtime = 0
	}

	MissDataLottery = missdata
	logs.Info("计算遗漏数据为 %v", MissDataLottery)
	//存数据库
	err = RestoreMissData(data[0][0], missdata)
	if err != nil {
		logs.Error("RestoreMissData [%v-%v] err: %v", data[0][0], missdata, err)
	}

	if len(lotterPlans) == 0 {
		putdata := KnnMissCount
		plan := model.LotterPlan{}
		nextPeriodNum, _ := strconv.Atoi(data[0][0])
		plan.CurrentPierod = strconv.Itoa(nextPeriodNum + 1)
		plan.NumberList = putdata
		plan.GetReward = false
		plan.PutTime = 1
		plan.Status = "等开"
		plan.RealPutTime = 0
		plan.CreateTime = time.Now().Format("2006-01-02 15:04:05")
		PuttoLotteryKnnMiss = putdata

		logs.Info("第一次计算数据为 %v ", plan)
		//存数据库
		err := RestorePlan(plan)
		if err != nil {
			logs.Error("RestorePlanToDB [%v] err: %v", plan, err)
		}
		lotterPlans = append(lotterPlans, plan)
	} else {
		//是否已经变更期数
		l := len(lotterPlans)
		lastplan := lotterPlans[l-1]
		logs.Info("最后一次计算数据为 %v ", lastplan)

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
					putdata := KnnMissCount
					plan := model.LotterPlan{}
					nextPeriodNum, _ := strconv.Atoi(data[0][0])
					plan.CurrentPierod = strconv.Itoa(nextPeriodNum + 1)
					plan.NumberList = putdata
					plan.GetReward = false
					plan.PutTime = 1
					plan.RealPutTime = 0
					plan.Status = "等开"
					plan.CreateTime = time.Now().Format("2006-01-02 15:04:05")
					PuttoLotteryKnnMiss = putdata
					lotterPlans = append(lotterPlans, plan)
					logs.Info("中了，计算数据为 %v ", plan)

					//存数据库
					err := RestorePlan(plan)
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

				lastplan.Status = "不中"
				lotterPlans = lotterPlans[0 : l-1]
				lotterPlans = append(lotterPlans, lastplan)
				putdata := KnnMissCount
				plan := model.LotterPlan{}
				nextPeriodNum, _ := strconv.Atoi(data[0][0])
				plan.CurrentPierod = strconv.Itoa(nextPeriodNum + 1)
				plan.NumberList = putdata
				plan.GetReward = false
				plan.PutTime = 1
				plan.RealPutTime = 0
				plan.Status = "等开"
				plan.CreateTime = time.Now().Format("2006-01-02 15:04:05")
				PuttoLotteryKnnMiss = putdata
				lotterPlans = append(lotterPlans, plan)
				logs.Info("没中，计算数据为 %v ", plan)

				//存数据库
				err := RestorePlan(plan)
				if err != nil {
					logs.Error("RestorePlanToDB [%v] err: %v", plan, err)
				}
			}
		}
	}
	logs.Info("计算最终数据为: ")
	for _, plans := range lotterPlans {
		logs.Info("数据列表： %v ", plans)
	}
	n := len(lotterPlans)

	lastmysqlplan, _ = strconv.Atoi(lotterPlans[n-1].CurrentPierod)
	RestoreLotterResult(lotterPlans)

}

