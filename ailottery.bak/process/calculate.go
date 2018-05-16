package process

import (
	"os/exec"
	"strconv"
	"strings"

	"github.com/astaxie/beego/logs"
)

func (oridata *OriData) KNNCalculate(param string, size int) (string, error) {
	cmdstr := "cd /var/source/src/stocknew/ailottery/aidir; sh knn" + strconv.Itoa(size) + ".sh " + param
	logs.Debug("推测命令 %v", cmdstr)
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
func (oridata *OriData) KNNMissCalculate(param string, size int, lastlist []int) ([]int, error) {
	cmdstr := "cd /var/source/src/stocknew/ailottery/aidir; sh knnmiss" + strconv.Itoa(size) + ".sh " + param
	logs.Debug("miss数据推测 %v", cmdstr)
	cmd := exec.Command("/bin/sh", "-c", cmdstr)
	stdout, err := cmd.Output()
	if err != nil {
		return "", err
	}
	str := string(stdout)
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)

	list := make([]int, 0)
	switch str {
	case "1":
		nm := 0
		for index, v := range lastlist {
			if v < 10 {
				list = append(list, index+1)
				nm += 1
			}
		}
		if nm <= 5 {
			return list, nil
		}

	case "2":
		s := ""
		nb := 0
		for index, v := range lastlist {
			if v >= 10 {
				list = append(list, index+1)
				nb += 1
			}
		}
		if nb <= 5 {
			return list, nil
		}
	}
	zerolist := make([]int, 0)
	return zerolist, nil
}

