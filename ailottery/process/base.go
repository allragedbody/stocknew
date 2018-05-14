package process

import (
	"stocknew/ailottery/model"
	//	"stocknew/ailottery/nettools"
	"stocknew/ailottery/db"
	"strconv"
	//	"github.com/axgle/mahonia"
	//	"bytes"
	//	"errors"
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

func Init() {
	OD = &OriData{}

}

func Running() {
	for {
		time.Sleep(time.Second * 2)
		recordSize, _ := beego.AppConfig.Int("knnsize20")
		logs.Info("进行对k临近算法的运算。")

		err := DataPrepare(recordSize)
		if err != nil {
			logs.Error("数据未准备好，重试。")
			continue
		}

		go DataReload(recordSize)

		OD.CalculateData(recordSize)

	}
}

func DataPrepare(recordSize int) error {
	oridata := OD.GetData(recordSize)
	OD.HandleData(oridata)
	err := OD.StoreData(oridata, recordSize)
	if err != nil {
		return err
	}
	return nil
}

func DataReload(recordSize int) {
	for {
		time.Sleep(time.Second * 120)
		oridata := OD.GetData(recordSize)
		OD.HandleData(oridata)
		err := OD.StoreData(oridata, recordSize)
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

