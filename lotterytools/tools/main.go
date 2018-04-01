package main

import (
	//"bufio"
	"fmt"
	//	"io"
	"io/ioutil"
	//	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("numbers.list")
	check(err)

	newlist := make(map[int][]string, 0)
	rewardlist := strings.Split(string(dat), "\n")

	for _, line := range rewardlist {
		slist := strings.Split(line, "|")
		list := make([]string, 0)
		list = append(list, slist[2])
		list = append(list, slist[3])
		list = append(list, slist[4])
		list = append(list, slist[5])
		list = append(list, slist[6])
		list = append(list, slist[7])
		list = append(list, slist[8])
		list = append(list, slist[9])
		list = append(list, slist[10])
		list = append(list, slist[11])
		c, _ := strconv.Atoi(slist[1])
		newlist[c] = list

	}
	var e bool
	newpushlist := make(map[int][]string, 0)
	for k, _ := range newlist {
		k_1 := k - 1
		_, ok := newlist[k_1]
		if ok {
			newpushlist[k] = make([]string, 0)
			for index := 0; index < 9; index++ {
				e = exist(newpushlist[k], newlist[k_1][index])
				if !e {
					newpushlist[k] = append(newpushlist[k], newlist[k_1][index])
				}

				e = exist(newpushlist[k], newlist[k][index+1])
				if !e {
					newpushlist[k] = append(newpushlist[k], newlist[k][index+1])
				}
			}
		}
	}

	fmt.Println(newpushlist)

	for period, vv := range newpushlist {
		fmt.Printf("period %v push %v ,misstime [%v]\n", period, vv[0:5], getRewardTime(newlist, period, vv[0:5]))
	}

}

func exist(l []string, n string) bool {
	for _, i := range l {
		if i == n {
			return true
		}
	}
	return false
}

func getRewardTime(newlist map[int][]string, period int, curs []string) int {
	least := 1000
	for p, v := range newlist {
		if p > period {
			e := exist(curs, v[0])
			if e {
				miss := p - period
				if miss < least {
					least = miss
				}
			}
		}
	}
	return least
}
