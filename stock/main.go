package main

import (
	"fmt"
	"takistan/stock/meta"
	//	"takistan/stock/nettools"
	"takistan/stock/process"
	//	"unsafe"
	//	"strings"

	//	"github.com/robertkrimen/otto"
)

func main() {
	fmt.Println("这是一款股票分析模拟交易软件。作者 邓云飞。")
	fmt.Println("加载股票代码数据。")
	codes := meta.LoadMeta()
	fmt.Println("加载需要获取数据的股票。")
	fmt.Println("对股票进行分析。")

	process.Init()
	process.GetHistoryData(codes)

}
