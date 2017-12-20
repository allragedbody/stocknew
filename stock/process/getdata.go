package process

import (
	"encoding/json"
	"fmt"
	"strconv"
	"takistan/stock/model"
	"takistan/stock/nettools"

	"github.com/axgle/mahonia"
)

var c *nettools.HttpClient

func Init() {
	c = nettools.CreateClient()
}

func GetHistoryData(codes []string) {
	url := "https://stock.api51.cn/stock_chart/"

	for _, code := range codes {
		meta := &nettools.StockMeter{CandlePeriod: "1", DataCount: "2", ProdCode: code, AppCode: ""}

		body, err := c.HttpDoGet(url, meta)
		if err != nil {
			fmt.Printf("Http 接口请求失败，错误: %v", err)
			return
		}

		sd := &model.StockData{}
		err = json.Unmarshal([]byte(body), &sd)
		if err != nil {
			fmt.Printf("解析数据失败，错误: %v", err)
			return
		}

		if v, ok := sd.Data.Infos[code]; ok {
			for _, ws := range v {
				oneday, _ := ws.([]interface{})
				o0, _ := oneday[0].(float64)

				sdate := strconv.FormatFloat(o0, 'E', -1, 64)
				fmt.Printf("%ss\n", sdate)
				//				for _, one := range oneday {
				//					switch one.(type) {
				//					case string:
				//						ws1, _ := one.(string)
				//						fmt.Printf("%v1\n", string(ws1))
				//					case int:
				//						ws1, _ := one.(int)
				//						fmt.Printf("%v2\n", int64(ws1))
				//					case float32:
				//						ws1, _ := one.(float32)
				//						fmt.Printf("%v3\n", float64(ws1))
				//					case float64:
				//						ws1, _ := one.(float32)
				//						fmt.Printf("%v4\n", float64(ws1))
				//					default:
				//						ws1, _ := one.(string)
				//						fmt.Printf("%v5\n", string(ws1))
				//					}
				//				}

			}
		}
		//		str := string(body)
		//		strencode := convertToString(str, "gbk", "utf-8")
		//		fmt.Printf("%v\n", strencode)
	}
}

//http://hq.sinajs.cn/list=sz000158,sh601766
func convertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
