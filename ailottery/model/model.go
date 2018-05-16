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
type LotterPlan struct {
	CurrentPierod string `json:"currentPierod"`
	NumberList    []int  `json:"numberList"`
	PutTime       int    `json:"putTime"`
	RealPutTime   int    `json:"realPutTime"`
	Status        string `json:"status"`
	GetReward     bool   `json:"getReward"`
	CreateTime    string `json:"createTime"`
}
type PushData struct {
	Touser    string  `json:"touser"`
	Toparty   string  `json:"toparty"`
	Totag     string  `json:"totag"`
	Msgtype   string  `json:"msgtype"`
	Agentid   int     `json:"agentid"`
	GetReward bool    `json:"getReward"`
	Text      Content `json:"text"`
	Safe      int     `json:"safe"`
}
type Content struct {
	Content string `json:"content"`
}

type TokenResp struct {
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

