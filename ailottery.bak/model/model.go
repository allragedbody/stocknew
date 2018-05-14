package model

type TrainingSet struct {
	KT *KNNTrainingSet
}

type CalculationSet struct {
	CS *KNNCalculationSet
}

type KNNTrainingSet struct {
	Size    int
	KNNList [][]int
}

type KNNCalculationSet struct {
	Size    int
	KNNList [][]int
}
