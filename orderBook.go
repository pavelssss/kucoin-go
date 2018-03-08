package kucoin

type OrderBook struct {
	Sell	[][3]float64
	Buy		[][3]float64
}

type rawOrderBook struct{
	Success   bool       `json:"success, omitempty"`
	Code      string     `json:"code, omitempty"`
	Msg       string     `json:"msg, omitempty"`
	//Timestamp int64      `json:"timestamp, omitempty"`
	Data      struct{
		Sell	[][3]float64	`json:"SELL, omitempty"`
		Buy		[][3]float64	`json:"BUY, omitempty"`
	} `json:"data, omitempty"`

}