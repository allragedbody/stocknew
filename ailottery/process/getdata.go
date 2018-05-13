package process

import (
	"stocknew/ailottery/db"

	//	"stocknew/ailottery/nettools"
	"strconv"
	//	"github.com/axgle/mahonia"
	//	"bytes"
	//	"errors"
	//	"fmt"
	//	"io/ioutil"
	//	"net/http"
	//	"sort"
	//	"encoding/json"

	"github.com/astaxie/beego/logs"
)

type KNN struct {
	Trainning
	K int
}

func (knn *KNN) HandleData() {
	logs.Debug("KNN HandleData")
}

func (knn *KNN) StoreData() {
	logs.Debug("KNN StoreData")
}

func (knn *KNN) CalculateData() {
	logs.Debug("KNN CalculateData")
}

func (knn *KNN) Running() {
	logs.Debug("KNN Running")
}

func (knn *KNN) GetData(size int) [][]int {
	logs.Debug("KNN GetData %v", size)
	knndata := make([][]int, 0)
	dbconn := db.GetDB()
	data, err := dbconn.GetLotterData(0, 100000)
	if err != nil {
		logs.Error("KNN GetData error %v", err)
		return knndata
	}

	rdata := reverse(data)
	stidata := strToInt(rdata)

	for i, _ := range stidata {
                if i< len(stidata)-size+10{
		tmpdata := stidata[i : i+size+1]
                logs.Debug(i,tmpdata)
		if isTraningData(tmpdata, size) {
			td := getTraningData(tmpdata)
			knndata = append(knndata, td)
		}
}
	}

	return knndata
}

func reverse(data [][]string) [][]string {
	d := make([][]string, 0)
	lenth := len(data)
	for i := lenth - 1; i > 0; i-- {
		d = append(d, data[i])
	}
	return d
}

func strToInt(data [][]string) [][]int {
	d := make([][]int, 0)
	for _, v := range data {
		slist := make([]int, 0)
		for _, vv := range v {
			tmpv, _ := strconv.Atoi(vv)
			slist = append(slist, tmpv)
		}
		d = append(d, slist)

	}
	return d
}

func isTraningData(data [][]int, size int) bool {
	tmpdata := make([][]int, 0)

	for _, v := range data {
		if len(v) != 0 {
			tmpdata = append(tmpdata, v)
		}
	}

	if len(tmpdata) != size+1 {
		return false
	}

	t := 0
	for _, v := range tmpdata {
		if t == 0 {
			t = v[0]
		} else {
			t = t + 1
			if t != v[0] {
				return false
			}

		}
	}
	return true
}
func getTraningData(data [][]int) []int {
	l := len(data)

	traningSet := make([]int, 0)
	for i, v := range data {
		if i != l-1 {
			traningSet = append(traningSet, v[1])
		} else {
			if getType(v[1]) != 0 {
				traningSet = append(traningSet, getType(v[1]))
			}
		}
	}
	return traningSet
}

func getType(num int) int {
	switch num {
	case 1, 2, 3, 4, 5:
		return 1
	case 6, 7, 8, 9, 10:
		return 2
	default:
		return 0
	}
}

