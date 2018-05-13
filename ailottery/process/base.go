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
	"os"
	"time"
	//	"encoding/json"
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

func (oridata *OriData) StoreData(traninglist [][]int) error {
	logs.Debug("KNN存储数据 %v ", traninglist)
	f, err := os.Create(AIDIR + "/knnlist.tmp") //创建文件
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

	_, err = f.WriteString(str) //写入文件(字节数组)
	if err != nil {
		return err
	}
	f.Sync()

	err = os.Rename(AIDIR+"/knnlist.tmp", AIDIR+"/knnlist.txt")
	if err != nil {
		return err
	}

	return nil
}

func (oridata *OriData) CalculateData(size int) error {
	dbconn := db.GetDB()
	data, err := dbconn.GetLotterData(0, 20)
	rdata := reverse(data)

	param := ""
	for _, v := range rdata {
		logs.Debug("第 %v 期中奖数据为 %v", v[0], v[1])
		param += " " + v[1]
	}

	next, _ := strconv.Atoi(rdata[len(rdata)-1][0])

	//推测中奖号码为：
	calculateNumbers, err := oridata.KNNCalculate(param)
	if err != nil {
		logs.Error("第 %v 期中奖数据推测计算错误 %v", next+1, err)
		return err
	}

	logs.Debug("第 %v 期中奖数据推测为 %v", next+1, calculateNumbers)
	return nil
}

func Init() {
	OD = &OriData{}

}

func Running() {
	for {
		time.Sleep(time.Second * 2)
		logs.Info("获取彩票数据库数据。")
		oridata := OD.GetData(20)
		OD.HandleData(oridata)
		err := OD.StoreData(oridata)
		if err != nil {
			continue
		}
		OD.CalculateData(20)

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

