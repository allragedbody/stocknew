package process

import (
	"os/exec"
	"strconv"
	"strings"

	"github.com/astaxie/beego/logs"
)

func (oridata *OriData) KNNCalculate(param string, size int) (string, error) {
	cmdstr := "cd /var/source/src/stocknew/ailottery/aidir; sh knn" + strconv.Itoa(size) + ".sh " + param
	logs.Debug("cmdstring %v", cmdstr)
	cmd := exec.Command("/bin/sh", "-c", cmdstr)
	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}
	str := string(stdout)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	switch str {
	case "1":
		return "1 2 3 4 5", nil
	case "2":
		return "6 7 8 9 10", nil
	}

	return "未知", nil
}

