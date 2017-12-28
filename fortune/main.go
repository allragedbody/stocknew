package main

import (
	"fmt"
	_ "fortune/routers"
	"stocknew/fortune/meta"
	//	"stocknew/fortune/nettools"
	"stocknew/fortune/process"
	//	"unsafe"
	//	"strings"
	"stocknew/fortune/db"
	"stocknew/fortune/routers"
	"time"

	"github.com/astaxie/beego"
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
	stockdata(codes)
	go reloadStockdata(codes)

	routers.Init()
	beego.SetStaticPath("/views", "views")
	beego.Run()

}

func stockdata(codes []string) error {
	data, err := process.GetHistoryData(codes)
	if err != nil {
		fmt.Println("获取股票历史数据失败。", err)
		return err
	}

	for _, line := range data {
		fmt.Printf("Data will be insert into db  %v\n", line)
	}
	err = process.RestoreData(data)
	if err != nil {
		fmt.Println("插入股票数据失败。", err)
		return err
	}
	fmt.Println("插入股票数据成功。")
	return nil
}

func reloadStockdata(codes []string) {
	for {
		time.Sleep(time.Second * 3600)
		data, err := process.GetHistoryData(codes)
		if err != nil {
			fmt.Println("获取股票历史数据失败。", err)
			continue
		}

		for _, line := range data {
			fmt.Printf("Data will be insert into db  %v\n", line)
		}
		err = process.RestoreData(data)
		if err != nil {
			fmt.Println("插入股票数据失败。", err)
			continue
		}
		fmt.Println("插入股票数据成功。")
	}
}
