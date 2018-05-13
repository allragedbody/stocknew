package process

import (
	"io/ioutil"
	"os/exec"

	"github.com/astaxie/beego/logs"
)

func (oridata *OriData) KNNCalculate(param string) (string, error) {
	logs.Debug("/usr/bin/python aidir/knn.py %v", param)
	cmd := exec.Command("/usr/bin/python aidir/knn.py ", param)
	// 获取输出对象，可以从该对象中读取输出结果
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	// 保证关闭输出流
	defer stdout.Close()
	// 运行命令
	if err := cmd.Start(); err != nil {
		return "", err
	}
	// 读取输出结果
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", err
	}
	result := string(opBytes)

	switch result {
	case "1":
		return "小", nil
	case "2":
		return "大", nil
	}

	return "未知", nil
}

