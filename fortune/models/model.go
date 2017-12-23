package model

type StockData struct {
	Data *Candle `json:"data"`
}
type Candle struct {
	Infos map[string][]interface{} `json:"candle"`
}
type DayInfo struct {
	Code            string
	Date            string
	OpenPx          float64
	ClosePx         float64
	HighPx          float64
	LowPx           float64
	BusinessAmount  float64
	BusinessBalance float64
}
