package main

import (
	"time"
	"strings"
	"strconv"
	"encoding/json"

	"github.com/toorop/go-pusher"
)

func DecodeBTCETrades(event *pusher.Event) (trades []*Trade, err error) {
	type TradeArr [3]string
	var trades_arr = make([]*TradeArr, 0, 1)

	if err = json.Unmarshal([]byte(event.Data), & trades_arr); err != nil {
		return
	}

	for _, trade_arr := range trades_arr {
		var pair  = strings.ToUpper(strings.Split(event.Channel, ".")[0])
		pair = strings.Replace(pair, "_", "/", 1)
		var typ, t_err = ParseTradeType(trade_arr[0]) 
		if err != nil {return trades, t_err}
		var amount, a_err = strconv.ParseFloat(trade_arr[2], 64)
		if err != nil {return trades, a_err}
		var price, p_err = strconv.ParseFloat(trade_arr[1], 64)
		if err != nil {return trades, p_err}
		var trade = & Trade {
			Time	: time.Now(),
			Pair	: pair,
			Type	: typ,
			Amount	: amount,
			Price   : price,
		}
		trades = append(trades, trade)
	}

	return
}
