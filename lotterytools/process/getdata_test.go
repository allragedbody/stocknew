package process

import (
	"fmt"
	"stocknew/lotterytools/db"
	"strconv"
	"testing"
)

func Test_NextNumberStatistics(t *testing.T) {
	dbconn := db.GetDB()
	ns := &numberStatistics{}
	ns.numbers = make(map[string]map[string]int, 0)

	for i := 1; i < 10; i++ {
		ns.numbers[strconv.Itoa(i)] = make(map[string]int, 0)
		for j := 1; j < 10; j++ {
			ns.numbers[strconv.Itoa(i)][strconv.Itoa(j)] = 0
		}
	}

	datas, err := dbconn.GetLotterDataPositive(0, 1000)
	if err != nil {
		return
	}
	for index, data := range datas {
		if index == 0 {
			ns.cur = data[1]
		} else {
			if nextHitMode(data[1], "number", "1") {
				ns.numbers[ns.cur]["1"] += 1
			}
			if nextHitMode(data[1], "number", "2") {
				ns.numbers[ns.cur]["2"] += 1
			}
			if nextHitMode(data[1], "number", "3") {
				ns.numbers[ns.cur]["3"] += 1
			}
			if nextHitMode(data[1], "number", "4") {
				ns.numbers[ns.cur]["4"] += 1
			}
			if nextHitMode(data[1], "number", "5") {
				ns.numbers[ns.cur]["5"] += 1
			}
			if nextHitMode(data[1], "number", "6") {
				ns.numbers[ns.cur]["6"] += 1
			}
			if nextHitMode(data[1], "number", "7") {
				ns.numbers[ns.cur]["7"] += 1
			}
			if nextHitMode(data[1], "number", "8") {
				ns.numbers[ns.cur]["8"] += 1
			}
			if nextHitMode(data[1], "number", "9") {
				ns.numbers[ns.cur]["9"] += 1
			}
			if nextHitMode(data[1], "number", "10") {
				ns.numbers[ns.cur]["10"] += 1
			}
			if nextHitMode(data[1], "oddeven", "odd") {
				ns.numbers[ns.cur]["odd"] += 1
			}
			if nextHitMode(data[1], "oddeven", "Even") {
				ns.numbers[ns.cur]["even"] += 1
			}
			ns.cur = data[1]
		}
	}
	fmt.Println(ns)
}

