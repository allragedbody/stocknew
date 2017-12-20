package model

type StockData struct {
	Data *Candle `json:"data"`
}
type Candle struct {
	Infos map[string][]interface{} `json:"candle"`
}
