package process

import (
	"encoding/json"
	"fmt"
	"strconv"
	//	"takistan/stock/model"
	"takistan/stock/nettools"

	//	"github.com/axgle/mahonia"
)

var c *nettools.HttpClient

func Init() {
	c = nettools.CreateClient()
}

type dayinfo struct {
	Code            string
	Date            string
	OpenPx          float64
	ClosePx         float64
	HighPx          float64
	LowPx           float64
	BusinessAmount  float64
	BusinessBalance float64
}

func GetHistoryData(codes []string) {
	url := "https://stock.api51.cn/stock_chart/"

	for _, code := range codes {
		meta := &nettools.StockMeter{CandlePeriod: "6", DataCount: "5", ProdCode: code, AppCode: "27319841797a486cb7b634b2dfef7ecb"}

		body, err := c.HttpDoGet(url, meta)
		if err != nil {
			fmt.Printf("Http 接口请求失败，错误: %v", err)
			return
		}

		var sdata map[string]map[string]map[string]interface{}
		//{"data":{"candle":{"002237.SZ":[[20171215,10.13,10.3,10.35,10.06,10775921,110202713],[20171218,10.28,10.26,10.37,10.15,8126516,83330556],[20171219,10.24,10.23,10.33,10.21,5519450,56621137],[20171220,10.2,10.1,10.26,10.08,4861455,49426968],[20171221,10.1,10.13,10.16,10.06,4047912,40928963]],"fields":["min_time","open_px","close_px","high_px","low_px","business_amount","business_balance"]}}}
		//		str := string(body)
		//		strencode := convertToString(str, "gbk", "utf-8")
		//		fmt.Printf("%v\n", strencode)

		//		fmt.Printf("%v\n", string(body))
		//var jsonData
		//		sd := &model.StockData{}
		err = json.Unmarshal([]byte(body), &sdata)
		if err != nil {
			fmt.Printf("解析数据失败，错误: %v", err)
			return
		}

		for k, v := range sdata["data"]["candle"] {
			if k == "fields" {
				continue
			}
			di := &dayinfo{}
			di.Code = k

			oneday, _ := v.([]interface{})
			for _, one := range oneday {
				//				fmt.Printf("%v %v\n", k, one)
				o, _ := one.([]interface{})
				v0, _ := o[0].(float64)
				date := strconv.FormatFloat(v0, 'G', 8, 64)
				openPx, _ := o[1].(float64)
				closePx, _ := o[2].(float64)
				highPx, _ := o[3].(float64)
				lowPx, _ := o[4].(float64)
				businessAmount, _ := o[5].(float64)
				businessBalance, _ := o[6].(float64)

				fmt.Printf("股票代码:(%v) 日期:(%v) 开盘价:(%v) 收盘价:(%v) 最高价:(%v) 最低价:(%v) businessAmount:(%v) businessBalance:(%v)\n", k, date, openPx, closePx, highPx, lowPx, businessAmount, businessBalance)
			}

		}

		//		str := string(body)
		//		strencode := convertToString(str, "gbk", "utf-8")
		//		fmt.Printf("%v\n", strencode)
	}
}

//http://hq.sinajs.cn/list=sz000158,sh601766
//func convertToString(src string, srcCode string, tagCode string) string {
//	srcCoder := mahonia.NewDecoder(srcCode)
//	srcResult := srcCoder.ConvertString(src)
//	tagCoder := mahonia.NewDecoder(tagCode)
//	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
//	result := string(cdata)
//	return result
//}
