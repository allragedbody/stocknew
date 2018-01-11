package models

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type StockData struct {
	Data *Candle `json:"data"`
}
type Candle struct {
	Infos map[string][]interface{} `json:"candle"`
}
type DayInfo struct {
	Code            string  `json:"code"`
	Date            string  `json:"date"`
	OpenPx          float64 `json:"openpx"`
	ClosePx         float64 `json:"closepx"`
	HighPx          float64 `json:"highpx"`
	LowPx           float64 `json:"lowpx"`
	BusinessAmount  float64 `json:"businessamount"`
	BusinessBalance float64 `json:"businessbalance"`
}

type DayPrice struct {
	Date string  `json:"date"`
	Px   float64 `json:"px"`
}

type Response struct {
	HistoryDatas []*DayInfo `json:"historydatas"`
	MaxPoints    []DayPrice `json:"maxpoints"`
	MinPoints    []DayPrice `json:"minpoints"`
}

func (*Response) GetMaxPoints(ds []*DayInfo) []DayPrice {
	dp := make([]DayPrice, 0)
	dataSize := len(ds)
	onePointWeight := 30000
	threePointWeight := 50000
	roleSize := 200000
	result = make(map[string]data, 0)
	datas = make(map[int]data, 0)

	for i, d := range ds {
		result[d.Date] = data{d.Date, d.HighPx, 0}
		datas[i] = data{d.Date, d.HighPx, 0}
	}

	for i1 := 0; i1 < roleSize; i1++ {
		td := make(map[string]float64, 0)

		nums := generateRandomNumber(1, dataSize, 15)
		for _, n := range nums {
			td[datas[n].date] = datas[n].value
		}

		var keys []string
		keys = make([]string, 0)

		for k, _ := range td {
			keys = append(keys, k)
		}
		//	sort.Strings(keys)

		// To perform the opertion you want
		for index, _ := range keys {
			//	fmt.Println(index)
			if index%3 == 1 {
				if result[keys[index]].value >= result[keys[index-1]].value && result[keys[index]].value >= result[keys[index+1]].value {
					t := result[keys[index]].size
					t = t + 1
					d := data{keys[index], result[keys[index]].value, t}
					//	fmt.Println(d, keys[index])
					result[keys[index]] = d
				}
			}
		}
		//}
	}
	for index, _ := range datas {
		if index > 0 && index < dataSize-1 {
			if datas[index].value >= datas[index-1].value && datas[index].value >= datas[index+1].value {
				t := result[datas[index].date].size
				t = t + onePointWeight
				d := data{datas[index].date, result[datas[index].date].value, t}
				//	fmt.Println(d, keys[index])
				result[datas[index].date] = d
			}
		}
	}

	for index, _ := range datas {
		if index > 1 && index < dataSize-2 {
			if datas[index].value+datas[index-1].value+datas[index+1].value >= datas[index-2].value+datas[index-1].value+datas[index].value && datas[index].value+datas[index-1].value+datas[index+1].value >= datas[index].value+datas[index+1].value+datas[index+2].value {
				t := result[datas[index].date].size
				t = t + threePointWeight
				d := data{datas[index].date, result[datas[index].date].value, t}
				//	fmt.Println(d, keys[index])
				result[datas[index].date] = d
			}
		}
	}

	sizeM := make(map[int]data, 0)
	sizeInt := make([]int, 0)
	for k, _ := range result {
		sizeM[result[k].size] = result[k]
		sizeInt = append(sizeInt, result[k].size)
	}
	sort.Ints(sizeInt)

	fmt.Printf("%v%v\n", sizeM[sizeInt[len(sizeInt)-1]], sizeM[sizeInt[len(sizeInt)-2]])

	d1 := DayPrice{sizeM[sizeInt[len(sizeInt)-2]].date, sizeM[sizeInt[len(sizeInt)-2]].value}
	dp = append(dp, d1)
	d2 := DayPrice{sizeM[sizeInt[len(sizeInt)-1]].date, sizeM[sizeInt[len(sizeInt)-1]].value}
	dp = append(dp, d2)
	return dp

}
func (*Response) GetMinPoints(ds []*DayInfo) []DayPrice {

	dp := make([]DayPrice, 0)
	dataSize := len(ds)
	onePointWeight := 30000
	threePointWeight := 50000
	roleSize := 200000
	result = make(map[string]data, 0)
	datas = make(map[int]data, 0)

	for i, d := range ds {
		result[d.Date] = data{d.Date, d.LowPx, 0}
		datas[i] = data{d.Date, d.LowPx, 0}
		fmt.Printf("11111111 %v\n", datas[i])
	}

	for i1 := 0; i1 < roleSize; i1++ {
		td := make(map[string]float64, 0)

		nums := generateRandomNumber(1, dataSize, 15)
		for _, n := range nums {
			td[datas[n].date] = datas[n].value
		}

		var keys []string
		keys = make([]string, 0)

		for k, _ := range td {
			keys = append(keys, k)
		}
		//	sort.Strings(keys)

		// To perform the opertion you want
		for index, _ := range keys {
			//	fmt.Println(index)
			if index%3 == 1 {
				if result[keys[index]].value <= result[keys[index-1]].value && result[keys[index]].value <= result[keys[index+1]].value {
					t := result[keys[index]].size
					t = t + 1
					d := data{keys[index], result[keys[index]].value, t}
					//	fmt.Println(d, keys[index])
					result[keys[index]] = d
				}
			}
		}
		//}
	}
	for index, _ := range datas {
		if index > 0 && index < dataSize-1 {
			if datas[index].value <= datas[index-1].value && datas[index].value <= datas[index+1].value {
				t := result[datas[index].date].size
				t = t + onePointWeight
				d := data{datas[index].date, result[datas[index].date].value, t}
				//	fmt.Println(d, keys[index])
				result[datas[index].date] = d
			}
		}
	}

	for index, _ := range datas {
		if index > 1 && index < dataSize-2 {
			if datas[index].value+datas[index-1].value+datas[index+1].value <= datas[index-2].value+datas[index-1].value+datas[index].value && datas[index].value+datas[index-1].value+datas[index+1].value <= datas[index].value+datas[index+1].value+datas[index+2].value {
				t := result[datas[index].date].size
				t = t + threePointWeight
				d := data{datas[index].date, result[datas[index].date].value, t}
				//	fmt.Println(d, keys[index])
				result[datas[index].date] = d
			}
		}
	}

	sizeM := make(map[int]data, 0)
	sizeInt := make([]int, 0)
	for k, _ := range result {
		sizeM[result[k].size] = result[k]
		sizeInt = append(sizeInt, result[k].size)
	}
	sort.Ints(sizeInt)

	fmt.Printf("%v%v\n", sizeM[sizeInt[len(sizeInt)-1]], sizeM[sizeInt[len(sizeInt)-2]])

	d1 := DayPrice{sizeM[sizeInt[len(sizeInt)-2]].date, sizeM[sizeInt[len(sizeInt)-2]].value}

	dp = append(dp, d1)
	d2 := DayPrice{sizeM[sizeInt[len(sizeInt)-1]].date, sizeM[sizeInt[len(sizeInt)-1]].value}

	dp = append(dp, d2)
	return dp
}

type data struct {
	date  string
	value float64
	size  int
}

var datas map[int]data
var result map[string]data

//生成count个[start,end)结束的不重复的随机数
func generateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}
