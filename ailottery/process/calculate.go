package process

import (
	"os/exec"
	"sort"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func (oridata *OriData) KNNCalculate(param string, size int) (string, error) {
	cmdstr := "cd /var/source/src/stocknew/ailottery/aidir; sh knn" + strconv.Itoa(size) + ".sh " + param
	logs.Debug("推测命令 %v", cmdstr)
	cmd := exec.Command("/bin/sh", "-c", cmdstr)
	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}
	str := string(stdout)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	switch str {
	case "1":
		return "1 2 3 4 5", nil
	case "2":
		return "6 7 8 9 10", nil
	}

	return "未知", nil
}
func (oridata *OriData) KNNMissCalculate(param string, size int, lastlist []int) ([]int, error) {

	boundary, _ := beego.AppConfig.Int("boundary")
	list := make([]int, 0)
	cmdstr := "cd /var/source/src/stocknew/ailottery/aidir; sh knnmiss" + strconv.Itoa(size) + ".sh " + param
	logs.Debug("miss数据推测 %v", cmdstr)
	cmd := exec.Command("/bin/sh", "-c", cmdstr)
	stdout, err := cmd.Output()
	if err != nil {
		return list, err
	}
	str := string(stdout)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	logs.Debug("最后一行的Miss数据为 %v", lastlist)
	switch str {
	// 多数里出
	case "1":
		gttenums := 0
		mtmpvs1 := make(map[int]int)
		for i, tmiss := range lastlist {
			mtmpvs1[tmiss] = i
		}
		var keys1 []int
		for _, k := range lastlist {
			keys1 = append(keys1, k)
		}

		sort.Ints(keys1)
		//获取分界线信息
		for _, k := range keys1 {
			logs.Debug("Key:", k, "Value:", mtmpvs1[k])
			if k >= boundary {
				gttenums += 1
			}
		}
		//当大于9的miss值大于等于5个的时候,从小数里面出就是少数，为2状态
		if gttenums >= 5 {
			list = append(list, mtmpvs1[keys1[5]]+1)
			list = append(list, mtmpvs1[keys1[6]]+1)
			list = append(list, mtmpvs1[keys1[7]]+1)
			list = append(list, mtmpvs1[keys1[8]]+1)
			list = append(list, mtmpvs1[keys1[9]]+1)
		} else {
			//当大于9的miss值小于5个的时候，如果从小数里面出就是多数，为1状态。
			list = append(list, mtmpvs1[keys1[0]]+1)
			list = append(list, mtmpvs1[keys1[1]]+1)
			list = append(list, mtmpvs1[keys1[2]]+1)
			list = append(list, mtmpvs1[keys1[3]]+1)
			list = append(list, mtmpvs1[keys1[4]]+1)
		}

	case "2":
		gttenums := 0
		//少梳里出
		mtmpvs2 := make(map[int]int)

		for i, tmiss := range lastlist {
			mtmpvs2[tmiss] = i
		}
		var keys2 []int
		for _, k := range lastlist {
			keys2 = append(keys2, k)
		}
		sort.Ints(keys2)
		//获取分界线信息
		for _, k := range keys2 {
			logs.Debug("44444444444444Key:", k, "Value:", mtmpvs2[k])
			if k >= boundary {
				gttenums += 1
			}
		}
		//当大于9的miss值大于等于5个的时候,从小数里面出就是少数，为2状态
		if gttenums >= 5 {
			list = append(list, mtmpvs2[keys2[0]]+1)
			list = append(list, mtmpvs2[keys2[1]]+1)
			list = append(list, mtmpvs2[keys2[2]]+1)
			list = append(list, mtmpvs2[keys2[3]]+1)
			list = append(list, mtmpvs2[keys2[4]]+1)
		} else {
			//当大于9的miss值小于5个的时候，如果从小数里面出就是多数，为1状态。
			list = append(list, mtmpvs2[keys2[5]]+1)
			list = append(list, mtmpvs2[keys2[6]]+1)
			list = append(list, mtmpvs2[keys2[7]]+1)
			list = append(list, mtmpvs2[keys2[8]]+1)
			list = append(list, mtmpvs2[keys2[9]]+1)
		}

	}

	return list, nil
}

