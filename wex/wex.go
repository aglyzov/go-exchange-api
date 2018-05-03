package wex

import (
	"time"
	"strings"
	"strconv"
	"encoding/json"

	//"github.com/toorop/go-pusher"
)

//const WEX_PUSHER_APP_KEY = "ee987526a24ba107824c"
//const WEX_PUSHER_CLUSTER = "eu"

/*
func DecodeWexTrades(event *pusher.Event) (trades []*Trade, err error) {
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

/* TODO
func NewWexTicker() (ticker Ticker, err error) {
	var p_conn, err = pusher.NewClient(WEX_PUSHER_APP_KEY)
	var log = Log.New("exch", "WEX")

	if err != nil {
		log.Fatal("could not connect to the Pusher service", "err", err)
	}
	defer p_conn.Close()
	log.Debug("connected to the Pusher service")

	if err = p_conn.Subscribe("btc_usd.trades"); err != nil {
		log.Fatal("could not subscribe to the btc_usd.trades channel", "err", err)
	}
	if err = p_conn.Subscribe("ltc_usd.trades"); err != nil {
		log.Fatal("could not subscribe to the ltc_usd.trades channel", "err", err)
	}
	log.Debug("subscribed to channels")

	var event_ch, b_err = p_conn.Bind("trades")
	if b_err != nil {
		log.Fatal("could not bind to the trades event", "err", b_err)
	}
	log.Debug("bound to the trades events")

	for ! p_conn.Stopped() {
		select {
		case event := <- event_ch:
			Log.Debug("received an event", "channel", event.Channel, "data", event.Data)
			var trades, err = DecodeBTCETrades(event)
			if err != nil {
				Log.Error("could not decode BTC-e trades", "err", err)
				continue
			}
			for _, trade := range trades {
				var time_str   = trade.Time.Format(time.StampMilli)
				var color = aurora.GreenFg
				if trade.Type == TRADE_BID {
					color = aurora.RedFg
				}
				if strings.HasSuffix(trade.Pair, "/USD") {
					var usd_value = trade.Amount * trade.Price
					if usd_value >= 1000 && usd_value < 10000 {
						color |= aurora.BoldFm
					} else if usd_value >= 10000 {
						color |= aurora.InverseFm
					}
				}
				var colorize = func(arg interface{}) aurora.Value {return au.Colorize(arg, color)}
				fmt.Printf("%v [%v] %s %7.2f %7.2f\n", time_str, trade.Pair, colorize(trade.Type), colorize(trade.Price), colorize(trade.Amount))
			}
		}
	}
}
*/
