package models

type StockData struct {
	Data *Candle `json:"data"`
}
type Candle struct {
	Infos map[string][]interface{} `json:"candle"`
}
type DayInfo struct {
	Code            string  `json:"code"`
	Date            string  `json:"date"`
	OpenPx          float64 `json:"openpx"`
	ClosePx         float64 `json:"closepx"`
	HighPx          float64 `json:"highpx"`
	LowPx           float64 `json:"lowpx"`
	BusinessAmount  float64 `json:"businessamount"`
	BusinessBalance float64 `json:"businessbalance"`
}
