//
// Package bitfinex implements the Connector, Ticker and Trader interfaces
// for the Bitfinex websocket API v2
//
package bitfinex

import (
	"time"
	"sync"
	"strings"
	"strconv"
	"encoding/json"

	"github.com/aglyzov/ws-machine"

	"github.com/aglyzov/exchange-api"
)

/*
func DecodeBitfinexTrades(event *pusher.Event) (trades []*Trade, err error) {
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
*/


/*
TODO

func InitTickers(exch_sym_pairs []string) ([]*Ticker, error) {
	for _, pair := range exch_sym_pairs {
		// each pair is "<EXCHANGE>:<SYMBOL>"
		pair = strings.TrimSpace(pair)
		var ok = len(pair) > 0
		var parts []string

		if ok { 
			parts = strings.Split(pair, ":")
			ok = len(parts) == 2
		}
		if ok {
			switch strings.ToLower(parts[0]) {
			case "bitfinex", "bfx":  // TODO: move this into the bitfinex module
			}
		}
		if !ok {
			Log.Error("bad argument", "ticker", pair)
		}
	}
}
*/
