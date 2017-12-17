package main

import (
	"fmt"
	"takistan/stock/meta"
	"takistan/stock/nettools"
	//	"unsafe"
	"strings"

	"github.com/axgle/mahonia"
	"github.com/robertkrimen/otto"
)

func main() {
	fmt.Println("这是一款股票分析模拟交易软件。作者 邓云飞。")
	fmt.Println("加载股票代码数据。")
	codes := meta.LoadMeta()
	fmt.Println("加载需要获取数据的股票。")
	fmt.Println("对股票进行分析。")
	c := nettools.CreateClient()
	url := "http://hq.sinajs.cn/list="
	codesstr := strings.Join(codes, ",")

	body, err := c.HttpDo(url + codesstr)
	if err != nil {
		fmt.Printf("Http 接口请求失败，错误: %v", err)
		return
	}

	str := string(body)

	str1 := ConvertToString(str, "gbk", "utf-8")
	//	fmt.Println(str1)

	o := otto.New()

	_, err = o.Run(str1)
	if err != nil {
		fmt.Printf("js运行失败，错误: %v", err)
		return
	}
	for _, code := range codes {
		ov, err := o.Get("hq_str_" + code)
		if err != nil {
			fmt.Printf("解析失败，错误: %v", err)
			//			return
		}
		ovstr, _ := ov.ToString()
		fmt.Printf("股票信息：%v\n", ovstr)
	}
}

//http://hq.sinajs.cn/list=sz000158,sh601766
func ConvertToString(src string, srcCode string, tagCode string) string {
	srcCoder := mahonia.NewDecoder(srcCode)
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder(tagCode)
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	result := string(cdata)
	return result
}
