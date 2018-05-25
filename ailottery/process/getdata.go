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
	"sort"
	//	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
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
		//此处有一个bug
		if i < len(stidata)-size {
			tmpdata := stidata[i : i+size+1]
			logs.Debug(i, tmpdata)
			if isTraningData(tmpdata, size) {
				td := getTraningData(tmpdata)
				knndata = append(knndata, td)
			}
		}
	}

	return knndata
}

func (knn *KNN) GetMissData(size int) [][]int {
	logs.Debug("KNN GetMissData %v", size)
	knndata := make([][]int, 0)
	dbconn := db.GetDB()
	data, err := dbconn.GetMissData(0, 100000)
	if err != nil {
		logs.Error("KNN GetData error %v", err)
		return knndata
	}

	rdata := reverse(data)

	stidata := strToInt(rdata)

	for i, _ := range stidata {
		//此处有一个bug
		if i < len(stidata)-size {
			tmpdata := stidata[i : i+size+1]
			logs.Debug(i, tmpdata)
			if isTraningData(tmpdata, size) {
				td, isadd := getTraningMissData(tmpdata)
				if isadd {
					knndata = append(knndata, td)
				}
			}
		}
	}

	return knndata
}

func reverse(data [][]string) [][]string {
	d := make([][]string, 0)
	lenth := len(data)
	for i := lenth - 1; i >= 0; {
		d = append(d, data[i])
		i--
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

func getTraningMissData(data [][]int) ([]int, bool) {
	l := len(data)
	var tmpv []int
	traningSet := make([]int, 0)
	for i, v := range data {
		if i != l-1 {
			traningSet = append(traningSet, v[1])
			traningSet = append(traningSet, v[2])
			traningSet = append(traningSet, v[3])
			traningSet = append(traningSet, v[4])
			traningSet = append(traningSet, v[5])
			traningSet = append(traningSet, v[6])
			traningSet = append(traningSet, v[7])
			traningSet = append(traningSet, v[8])
			traningSet = append(traningSet, v[9])
			traningSet = append(traningSet, v[10])
			if i == l-2 {
				tmpv = v
			}
		} else {
			if getMissType(tmpv[1:], v[1:]) != 0 {
				traningSet = append(traningSet, getMissType(tmpv[1:], v[1:]))
			} else {
				return traningSet, false
			}
		}
	}
	return traningSet, true
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

func getMissType(tmpv []int, numlist []int) int {
	boundary, _ := beego.AppConfig.Int("boundary")
	findindex := 0
	for i, v := range numlist {
		if v == 0 {
			findindex = i
		}
	}

	gttenums := 0

	mtmpvs := make(map[int]int)

	for i, tmiss := range tmpv {
		mtmpvs[tmiss] = i
	}

	var keys []int
	for k := range mtmpvs {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	//获取分界线信息
	for _, k := range keys {
		logs.Debug("Key:", k, "Value:", mtmpvs[k])
		if k >= boundary {
			gttenums += 1
		}
	}

	missupbondary, _ := beego.AppConfig.Int("missupbondary")

	upbondarystr := beego.AppConfig.String("upbondary")
	underbondarystr := beego.AppConfig.String("underbondary")
	upbondary := stringsToList(upbondarystr)
	underbondary := stringsToList(underbondarystr)
	logs.Debug("upbondary %v", upbondary)
	logs.Debug("underbondary %v", underbondary)
	//当大于9的miss值大于等于5个的时候,从小数里面出就是少数，为2状态
	if gttenums >= missupbondary {
		for _, i := range upbondary {
			logs.Debug("對比情況 %v", i)
			if findindex == mtmpvs[keys[i]] {
				return 2
			} else {
				return 1
			}
		}

	} else {
		//当大于9的miss值小于5个的时候，如果从小数里面出就是多数，为1状态。
		for _, i := range underbondary {
			logs.Debug("對比情況 %v", i)
			if findindex == mtmpvs[keys[i]] {
				return 1
			} else {
				return 2
			}
		}
	}

	return 0

}

func stringsToList(str string) []int {
	list := make([]int, 0)
	for _, s := range strings.Split(str, " ") {
		ii, _ := strconv.Atoi(s)
		list = append(list, ii)
	}
	return list
}

