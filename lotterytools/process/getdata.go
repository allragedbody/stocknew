package process

import (
	"encoding/json"
	"stocknew/lotterytools/db"
	"stocknew/lotterytools/model"
	"stocknew/lotterytools/nettools"
	"strconv"
	//	"github.com/axgle/mahonia"
	//	"bytes"
	//	"errors"
	"fmt"
	//	"io/ioutil"
	//	"net/http"
	"sort"
	"time"

	"github.com/astaxie/beego/logs"
)

var PuttoLottery []int
var PuttoLotteryMax4 []int
var MissDataLottery []int
var LotterPlans []model.LotterPlan
var DateData map[int]int

var c *nettools.HttpClient

func Init() {
	c = nettools.CreateClient()
}

//{"status":{"code":"403","text":"请求超频,违规3次"}}

func SendWeChat(lotteryPlan model.LotterPlan) error {
	tkurl := "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=ww3bc839d9990b1e89&corpsecret=eF05WFY7SRfK1hpjjb3UxSzGZnFnyREtKK47PvMloN8"

	tkbody, err := c.HttpDoGet(tkurl)
	if err != nil {
		logs.Error("获取token数据失败，错误: %v", err)
		return err
	}

	tr := &model.TokenResp{}

	err = json.Unmarshal(tkbody, tr)
	if err != nil {
		logs.Error("获取token数据失败，错误：%v", err)
		return err
	}

	if tr.ErrCode != 0 {
		logs.Error("获取token数据失败，错误：%v", tr.ErrMsg)
		return err
	}

	access_token := tr.AccessToken

	pushData := &model.PushData{}
	now := time.Now().UTC().Format("2006-01-02 15:04:05")

	localtext := fmt.Sprintf("(%v)(%v)(%v)(%v)", lotteryPlan.CurrentPierod, lotteryPlan.NumberList, lotteryPlan.RealPutTime, now)
	pushData.Touser = "@all"
	pushData.Msgtype = "text"
	pushData.Agentid = 1000002
	pushData.Text = model.Content{Content: localtext}

	body, err := c.HttpDoPost(pushData, access_token)
	if err != nil {
		logs.Error("发送短信失败，错误: %v", err)
		return err
	}

	logs.Info("Send message body：%v", string(body))
	return nil
}

func RestorePlan(lp model.LotterPlan) error {
	dbconn := db.GetDB()
	err := dbconn.RestorePlanToDB(lp.CurrentPierod, lp.NumberList, lp.PutTime, lp.RealPutTime, lp.Status, lp.GetReward, lp.CreateTime)
	if err != nil {
		return err
	}
	return nil
}

////http://hq.sinajs.cn/list=sz000158,sh601766
//func convertToString(src string, srcCode string, tagCode string) string {
//	srcCoder := mahonia.NewDecoder(srcCode)
//	srcResult := srcCoder.ConvertString(src)
//	tagCoder := mahonia.NewDecoder(tagCode)
//	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
//	result := string(cdata)
//	return result
//}

func GetDBData(cur int) ([][]string, error) {
	dbconn := db.GetDB()
	data, err := dbconn.GetLotterData(cur, 1000)
	if err != nil {
		return nil, err
	}
	return data, nil
}

type NumberMiss struct {
	HitPeriod    string
	CurentPeriod string
	MissTime     int
	GetMissData  bool
}

func (nums *NumberMiss) getMiss() int {
	a, _ := strconv.Atoi(nums.CurentPeriod)
	b, _ := strconv.Atoi(nums.HitPeriod)

	return a - b
}

func CalculateMiss(data [][]string) (map[string]*NumberMiss, error) {
	cPeriod := data[0][0]
	alldata := make(map[string]*NumberMiss, 0)
	alldata["1"] = &NumberMiss{CurentPeriod: cPeriod, GetMissData: false}
	alldata["2"] = &NumberMiss{CurentPeriod: cPeriod, GetMissData: false}
	alldata["3"] = &NumberMiss{CurentPeriod: cPeriod, GetMissData: false}
	alldata["4"] = &NumberMiss{CurentPeriod: cPeriod, GetMissData: false}
	alldata["5"] = &NumberMiss{CurentPeriod: cPeriod, GetMissData: false}
	alldata["6"] = &NumberMiss{CurentPeriod: cPeriod, GetMissData: false}
	alldata["7"] = &NumberMiss{CurentPeriod: cPeriod, GetMissData: false}
	alldata["8"] = &NumberMiss{CurentPeriod: cPeriod, GetMissData: false}
	alldata["9"] = &NumberMiss{CurentPeriod: cPeriod, GetMissData: false}
	alldata["10"] = &NumberMiss{CurentPeriod: cPeriod, GetMissData: false}

	for _, numbers := range data {
		logs.Debug("第 %v 期 的第一名为 %v.", numbers[0], numbers[1])
		if numbers[1] == "1" {
			if alldata["1"].GetMissData == false {
				if alldata["1"].HitPeriod == "" {
					alldata["1"].HitPeriod = numbers[0]
				} else {
					alldata["1"].GetMissData = true
					alldata["1"].MissTime = alldata["1"].getMiss()
					logs.Debug("第一名是01的遗漏情况 %v.", alldata["1"])
				}
			}
		}
		if numbers[1] == "2" {
			if alldata["2"].GetMissData == false {
				if alldata["2"].HitPeriod == "" {
					alldata["2"].HitPeriod = numbers[0]
				} else {
					alldata["2"].GetMissData = true
					alldata["2"].MissTime = alldata["2"].getMiss()
					logs.Debug("第一名是02的遗漏情况 %v.", alldata["2"])
				}
			}
		}
		if numbers[1] == "3" {
			if alldata["3"].GetMissData == false {
				if alldata["3"].HitPeriod == "" {
					alldata["3"].HitPeriod = numbers[0]
				} else {
					alldata["3"].GetMissData = true
					alldata["3"].MissTime = alldata["3"].getMiss()
					logs.Debug("第一名是03的遗漏情况 %v.", alldata["3"])
				}
			}
		}

		if numbers[1] == "4" {
			if alldata["4"].GetMissData == false {
				if alldata["4"].HitPeriod == "" {
					alldata["4"].HitPeriod = numbers[0]
				} else {
					alldata["4"].GetMissData = true
					alldata["4"].MissTime = alldata["4"].getMiss()
					logs.Debug("第一名是04的遗漏情况 %v.", alldata["4"])
				}
			}
		}

		if numbers[1] == "5" {
			if alldata["5"].GetMissData == false {
				if alldata["5"].HitPeriod == "" {
					alldata["5"].HitPeriod = numbers[0]
				} else {
					alldata["5"].GetMissData = true
					alldata["5"].MissTime = alldata["5"].getMiss()
					logs.Debug("第一名是05的遗漏情况 %v.", alldata["5"])
				}
			}
		}
		if numbers[1] == "6" {
			if alldata["6"].GetMissData == false {
				if alldata["6"].HitPeriod == "" {
					alldata["6"].HitPeriod = numbers[0]
				} else {
					alldata["6"].GetMissData = true
					alldata["6"].MissTime = alldata["6"].getMiss()
					logs.Debug("第一名是06的遗漏情况 %v.", alldata["6"])
				}
			}
		}
		if numbers[1] == "7" {
			if alldata["7"].GetMissData == false {
				if alldata["7"].HitPeriod == "" {
					alldata["7"].HitPeriod = numbers[0]
				} else {
					alldata["7"].GetMissData = true
					alldata["7"].MissTime = alldata["7"].getMiss()
					logs.Debug("第一名是07的遗漏情况 %v.", alldata["7"])
				}
			}
		}
		if numbers[1] == "8" {
			if alldata["8"].GetMissData == false {
				if alldata["8"].HitPeriod == "" {
					alldata["8"].HitPeriod = numbers[0]
				} else {
					alldata["8"].GetMissData = true
					alldata["8"].MissTime = alldata["8"].getMiss()
					logs.Debug("第一名是08的遗漏情况 %v.", alldata["8"])
				}
			}
		}
		if numbers[1] == "9" {
			if alldata["9"].GetMissData == false {
				if alldata["9"].HitPeriod == "" {
					alldata["9"].HitPeriod = numbers[0]
				} else {
					alldata["9"].GetMissData = true
					alldata["9"].MissTime = alldata["9"].getMiss()
					logs.Debug("第一名是09的遗漏情况 %v.", alldata["9"])
				}
			}
		}
		if numbers[1] == "10" {
			if alldata["10"].GetMissData == false {
				if alldata["10"].HitPeriod == "" {
					alldata["10"].HitPeriod = numbers[0]
				} else {
					alldata["10"].GetMissData = true
					alldata["10"].MissTime = alldata["10"].getMiss()
					logs.Debug("第一名是10的遗漏情况 %v.", alldata["10"])
				}
			}
		}

	}

	return alldata, nil
}
func CalculatePut(tenNums []int) []int {
	nums := make(map[int]int, 0)
	for i, v := range tenNums {
		nums[i] = v
	}
	sort.Ints(tenNums)

	selectNums := make([]int, 0)

	for k, v := range nums {
		if v == tenNums[1] {
			logs.Info("选择第1个号码是 %v", k+1)
			selectNums = append(selectNums, k+1)
			continue
		}
		if v == tenNums[2] {
			logs.Info("选择第2个号码是 %v", k+1)
			selectNums = append(selectNums, k+1)
			continue
		}
		if v == tenNums[3] {
			logs.Info("选择第3个号码是 %v", k+1)
			selectNums = append(selectNums, k+1)
			continue
		}
		if v == tenNums[8] {
			logs.Info("选择第4个号码是 %v", k+1)
			selectNums = append(selectNums, k+1)
			continue
		}
		if v == tenNums[9] {
			logs.Info("选择第5个号码是 %v", k+1)
			selectNums = append(selectNums, k+1)
			continue
		}
	}
	sort.Ints(selectNums)
	logs.Debug("本期采用号码 %v", selectNums)
	return selectNums
}

func NewCalculatePut(numbers, tenNums []int) []int {
	nums := make(map[int]int, 0)

	for i, v := range tenNums {
		nums[i+1] = v
	}

	for _, n := range numbers {
		delete(nums, n)
	}

	fiveNums := make([]int, 0)
	fiveMnums := make(map[int]int, 0)
	for _, v := range nums {
		fiveNums = append(fiveNums, v)
	}
	fmt.Println(fiveNums)

	for i, v := range nums {
		fiveMnums[i] = v
	}
	fmt.Println(fiveMnums)
	sort.Ints(fiveNums)

	selectNums := make([]int, 0)

	for k, v := range fiveMnums {
		if v == fiveNums[0] {
			fmt.Println("选择第1个号码是 %v", k)
			selectNums = append(selectNums, k)
			continue
		}
		if v == fiveNums[1] {
			fmt.Println("选择第2个号码是 %v", k)
			selectNums = append(selectNums, k)
			continue
		}
		if v == fiveNums[4] {
			fmt.Println("选择第3个号码是 %v", k)
			selectNums = append(selectNums, k)
			continue
		}

	}
	sort.Ints(selectNums)
	fmt.Println("本期新增三个号码 %v", selectNums)
	return selectNums
}


func GetDateData() {
	dbconn := db.GetDB()
	dd, err := dbconn.GetPutHistoryData()
	if err != nil {
		logs.Debug("获取当天历史数据失败 %v", err)
	}
        DateData=dd
        logs.Debug("历史数据是 %v",DateData)
}


func RestoreLotterResult(lps []model.LotterPlan) {
	LotterPlans = lps
}

