package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type data struct {
	date  string
	value float64
	size  int
}

var datas map[int]data
var result map[string]data

func main() {
	dataSize := 60
	onePointWeight := 10000
	threePointWeight := 30000
	result = make(map[string]data, 0)
	datas = make(map[int]data, 0)

	result["20170926"] = data{"20170926", 53.3, 0}
	result["20170927"] = data{"20170927", 52.48, 0}
	result["20170928"] = data{"20170928", 53.15, 0}
	result["20170929"] = data{"20170929", 52.99, 0}
	result["20171009"] = data{"20171009", 54.5, 0}
	result["20171010"] = data{"20171010", 55.5, 0}
	result["20171011"] = data{"20171011", 56.32, 0}
	result["20171012"] = data{"20171012", 55.5, 0}
	result["20171013"] = data{"20171013", 54.72, 0}
	result["20171016"] = data{"20171016", 56.4, 0}
	result["20171017"] = data{"20171017", 52.9, 0}
	result["20171018"] = data{"20171018", 50.01, 0}
	result["20171019"] = data{"20171019", 49.58, 0}
	result["20171020"] = data{"20171020", 49.3, 0}
	result["20171023"] = data{"20171023", 50.3, 0}
	result["20171024"] = data{"20171024", 50.12, 0}
	result["20171025"] = data{"20171025", 50.1, 0}
	result["20171026"] = data{"20171026", 51, 0}
	result["20171027"] = data{"20171027", 50.42, 0}
	result["20171030"] = data{"20171030", 50, 0}
	result["20171031"] = data{"20171031", 52.23, 0}
	result["20171101"] = data{"20171101", 53.6, 0}
	result["20171102"] = data{"20171102", 53.25, 0}
	result["20171103"] = data{"20171103", 56.5, 0}
	result["20171106"] = data{"20171106", 57.11, 0}
	result["20171107"] = data{"20171107", 58.61, 0}
	result["20171108"] = data{"20171108", 60.15, 0}
	result["20171109"] = data{"20171109", 59.31, 0}
	result["20171110"] = data{"20171110", 59.95, 0}
	result["20171113"] = data{"20171113", 58, 0}
	result["20171114"] = data{"20171114", 65.5, 0}
	result["20171115"] = data{"20171115", 66, 0}
	result["20171116"] = data{"20171116", 65.19, 0}
	result["20171117"] = data{"20171117", 67.36, 0}
	result["20171120"] = data{"20171120", 68, 0}
	result["20171121"] = data{"20171121", 70.25, 0}
	result["20171122"] = data{"20171122", 70.7, 0}
	result["20171123"] = data{"20171123", 69.6, 0}
	result["20171124"] = data{"20171124", 66.8, 0}
	result["20171127"] = data{"20171127", 63.51, 0}
	result["20171128"] = data{"20171128", 62, 0}
	result["20171129"] = data{"20171129", 66.5, 0}
	result["20171130"] = data{"20171130", 65.06, 0}
	result["20171201"] = data{"20171201", 65.98, 0}
	result["20171204"] = data{"20171204", 68.95, 0}
	result["20171205"] = data{"20171205", 67, 0}
	result["20171206"] = data{"20171206", 62.93, 0}
	result["20171207"] = data{"20171207", 62.85, 0}
	result["20171208"] = data{"20171208", 61.69, 0}
	result["20171211"] = data{"20171211", 60.45, 0}
	result["20171212"] = data{"20171212", 62.4, 0}
	result["20171213"] = data{"20171213", 61.97, 0}
	result["20171214"] = data{"20171214", 62, 0}
	result["20171215"] = data{"20171215", 61.97, 0}
	result["20171218"] = data{"20171218", 62.52, 0}
	result["20171219"] = data{"20171219", 58.74, 0}
	result["20171220"] = data{"20171220", 59.91, 0}
	result["20171221"] = data{"20171221", 61, 0}
	result["20171222"] = data{"20171222", 61.78, 0}
	result["20171225"] = data{"20171225", 61.5, 0}

	datas[0] = data{"20170926", 53.3, 0}
	datas[1] = data{"20170927", 52.48, 0}
	datas[2] = data{"20170928", 53.15, 0}
	datas[3] = data{"20170929", 52.99, 0}
	datas[4] = data{"20171009", 54.5, 0}
	datas[5] = data{"20171010", 55.5, 0}
	datas[6] = data{"20171011", 56.32, 0}
	datas[7] = data{"20171012", 55.5, 0}
	datas[8] = data{"20171013", 54.72, 0}
	datas[9] = data{"20171016", 56.4, 0}
	datas[10] = data{"20171017", 52.9, 0}
	datas[11] = data{"20171018", 50.01, 0}
	datas[12] = data{"20171019", 49.58, 0}
	datas[13] = data{"20171020", 49.3, 0}
	datas[14] = data{"20171023", 50.3, 0}
	datas[15] = data{"20171024", 50.12, 0}
	datas[16] = data{"20171025", 50.1, 0}
	datas[17] = data{"20171026", 51, 0}
	datas[18] = data{"20171027", 50.42, 0}
	datas[19] = data{"20171030", 50, 0}
	datas[20] = data{"20171031", 52.23, 0}
	datas[21] = data{"20171101", 53.6, 0}
	datas[22] = data{"20171102", 53.25, 0}
	datas[23] = data{"20171103", 56.5, 0}
	datas[24] = data{"20171106", 57.11, 0}
	datas[25] = data{"20171107", 58.61, 0}
	datas[26] = data{"20171108", 60.15, 0}
	datas[27] = data{"20171109", 59.31, 0}
	datas[28] = data{"20171110", 59.95, 0}
	datas[29] = data{"20171113", 58, 0}
	datas[30] = data{"20171114", 65.5, 0}
	datas[31] = data{"20171115", 66, 0}
	datas[32] = data{"20171116", 65.19, 0}
	datas[33] = data{"20171117", 67.36, 0}
	datas[34] = data{"20171120", 68, 0}
	datas[35] = data{"20171121", 70.25, 0}
	datas[36] = data{"20171122", 70.7, 0}
	datas[37] = data{"20171123", 69.6, 0}
	datas[38] = data{"20171124", 66.8, 0}
	datas[39] = data{"20171127", 63.51, 0}
	datas[40] = data{"20171128", 62, 0}
	datas[41] = data{"20171129", 66.5, 0}
	datas[42] = data{"20171130", 65.06, 0}
	datas[43] = data{"20171201", 65.98, 0}
	datas[44] = data{"20171204", 68.95, 0}
	datas[45] = data{"20171205", 67, 0}
	datas[46] = data{"20171206", 62.93, 0}
	datas[47] = data{"20171207", 62.85, 0}
	datas[48] = data{"20171208", 61.69, 0}
	datas[49] = data{"20171211", 60.45, 0}
	datas[50] = data{"20171212", 62.4, 0}
	datas[51] = data{"20171213", 61.97, 0}
	datas[52] = data{"20171214", 62, 0}
	datas[53] = data{"20171215", 61.97, 0}
	datas[54] = data{"20171218", 62.52, 0}
	datas[55] = data{"20171219", 58.74, 0}
	datas[56] = data{"20171220", 59.91, 0}
	datas[57] = data{"20171221", 61, 0}
	datas[58] = data{"20171222", 61.78, 0}
	datas[59] = data{"20171225", 61.5, 0}
	for i1 := 0; i1 < 100000; i1++ {
		td := make(map[string]float64, 0)

		nums := generateRandomNumber(1, 60, 15)
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
			if datas[index].value > datas[index-1].value && datas[index].value > datas[index+1].value {
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
			if datas[index].value+datas[index-1].value+datas[index+1].value > datas[index-2].value+datas[index-1].value+datas[index].value && datas[index].value+datas[index-1].value+datas[index+1].value > datas[index].value+datas[index+1].value+datas[index+2].value {
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
}

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
