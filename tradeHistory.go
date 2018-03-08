package kucoin


type TradeHistory [][5]interface{}


type rawTradeHistory struct{
	Success bool               `json:"success"`
	Code    string             `json:"code"`
	Msg     string             `json:"msg"`
	Data    [][5]interface{}      `json:"data"`
}

