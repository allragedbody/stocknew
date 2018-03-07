package main

import (
	"fmt"
	"stocknew/lottery/model"
)

var FirstPeriods = 669754
var firstNumber int

func main() {
	fmt.Println("这是一个彩票游戏。")

	fmt.Println("当前信息。")

	n1 := &model.PKTen{}
	n1.Periods = 669754
	n1.FirstInfos = make([]*model.FirstInfo, 10)
	n1.FirstInfos[0] = &model.FirstInfo{Number: 1, Miss: 20, MAXMiss: 20, AvarageMiss: 20}
	n1.FirstInfos[1] = &model.FirstInfo{Number: 2, Miss: 3, MAXMiss: 3, AvarageMiss: 3}
	n1.FirstInfos[2] = &model.FirstInfo{Number: 3, Miss: 1, MAXMiss: 1, AvarageMiss: 1}
	n1.FirstInfos[3] = &model.FirstInfo{Number: 4, Miss: 5, MAXMiss: 5, AvarageMiss: 5}
	n1.FirstInfos[4] = &model.FirstInfo{Number: 5, Miss: 2, MAXMiss: 2, AvarageMiss: 2}
	n1.FirstInfos[5] = &model.FirstInfo{Number: 6, Miss: 20, MAXMiss: 20, AvarageMiss: 20}
	n1.FirstInfos[6] = &model.FirstInfo{Number: 7, Miss: 4, MAXMiss: 4, AvarageMiss: 4}
	n1.FirstInfos[7] = &model.FirstInfo{Number: 8, Miss: 0, MAXMiss: 0, AvarageMiss: 0}
	n1.FirstInfos[8] = &model.FirstInfo{Number: 9, Miss: 20, MAXMiss: 20, AvarageMiss: 20}
	n1.FirstInfos[9] = &model.FirstInfo{Number: 10, Miss: 8, MAXMiss: 8, AvarageMiss: 8}
	fmt.Printf("当前期次 %v：\n", n1.Periods)
	for _, nu := range n1.FirstInfos {
		fmt.Printf("号码 %v 遗漏次数 %v 最大遗漏次数 %v 平均遗漏次数%v。\n", nu.Number, nu.Miss, nu.MAXMiss, nu.AvarageMiss)
	}

	firstNumber = 6
	runLotteryFirst(firstNumber, n1)
	firstNumber = 10
	runLotteryFirst(firstNumber, n1)
	firstNumber = 6
	runLotteryFirst(firstNumber, n1)
	firstNumber = 4
	runLotteryFirst(firstNumber, n1)
	firstNumber = 5
	runLotteryFirst(firstNumber, n1)
	firstNumber = 4
	runLotteryFirst(firstNumber, n1)
	firstNumber = 6
	runLotteryFirst(firstNumber, n1)
	firstNumber = 9
	runLotteryFirst(firstNumber, n1)
	firstNumber = 1
	runLotteryFirst(firstNumber, n1)

	fmt.Println("进入下注环节。")
	fmt.Println("下注方案选择。")
	fmt.Println("下注。")
	fmt.Println("奖金结算。")
}

func runLotteryFirst(firstNumber int, n1 *model.PKTen) {
	fmt.Println("开奖。")
	fmt.Printf("本次开出第一名的数字是 %v。\n", firstNumber)
	fmt.Println("存储上一期开奖信息。")
	fmt.Println("当前信息修正。")

	rewardNumbers := make([]int, 0)
	rewardNumbers = append(rewardNumbers, firstNumber)
	n1.ChangeInfos(FirstPeriods, rewardNumbers, n1)
	fmt.Printf("当前期次 %v：\n", n1.Periods)
	for _, nu := range n1.FirstInfos {
		fmt.Printf("号码 %v 遗漏次数 %v 最大遗漏次数 %v 平均遗漏次数%v。\n", nu.Number, nu.Miss, nu.MAXMiss, nu.AvarageMiss)
	}
}
