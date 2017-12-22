package main

import (
	"fmt"
	"takistan/stock/meta"
	//	"takistan/stock/nettools"
	"takistan/stock/process"
	//	"unsafe"
	//	"strings"
	"takistan/stock/db"
	"time"
	//	"github.com/robertkrimen/otto"
)

func main() {
	fmt.Println("这是一款股票分析模拟交易软件。作者 邓云飞。")
	fmt.Println("加载股票代码数据。")
	codes := meta.LoadMeta()
	fmt.Println("加载需要获取数据的股票。")
	fmt.Println("对股票进行分析。")
	err := db.Init()
	if err != nil {
		return
	}
	process.Init()
	data, _ := process.GetHistoryData(codes)
	for _, line := range data {
		fmt.Printf("Data will be insert into db  %v\n", line)
	}
	err = process.RestoreData(data)
	if err != nil {
		fmt.Println("插入股票数据失败。", err)
		return
	}
	fmt.Println("插入股票数据成功。")

	for {
		time.Sleep(time.Second * 3600)
		data, _ := process.GetHistoryData(codes)
		for _, line := range data {
			fmt.Printf("Data will be insert into db  %v\n", line)
		}
		err = process.RestoreData(data)
		if err != nil {
			fmt.Println("插入股票数据失败。", err)
			return
		}
		fmt.Println("插入股票数据成功。")
	}
}
