package main

import (
	"github.com/astaxie/beego/logs"
)

func main() {
	log := logs.NewLogger(10000) // 创建一个日志记录器，参数为缓冲区的大小
	// 设置配置文件
	jsonConfig := `{
        "filename" : "test.log", 
        "maxlines" : 1000,       
        "maxsize"  : 10240       
    }`
	log.SetLogger("file", jsonConfig) // 设置日志记录方式：本地文件记录
	log.SetLevel(logs.LevelDebug)     // 设置日志写入缓冲区的等级
	log.EnableFuncCallDepth(true)     // 输出log时能显示输出文件名和行号（非必须）

	log.Emergency("Emergency")
	log.Alert("Alert")
	log.Critical("Critical")
	log.Error("Error")
	log.Warning("Warning")
	log.Notice("Notice")
	log.Informational("Informational")
	log.Debug("Debug")

	//	log.Flush() // 将日志从缓冲区读出，写入到文件
	//	log.Close()
}
