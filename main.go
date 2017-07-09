package main

import (
	"os"
	"fmt"
	"time"
	"strings"
	"strconv"
	"encoding/json"

	"github.com/toorop/go-pusher"
	"github.com/alexflint/go-arg"
	"github.com/logrusorgru/aurora"
	"github.com/aglyzov/log15"
)

const BTCE_PUSHER_APP_KEY = "c354d4d129ee0faa5c92"

type TradeType byte
const (
	TRADE_UNK  TradeType = 0
	TRADE_BID  TradeType = 1 << iota
	TRADE_ASK
)
type Trade struct {
	Time	time.Time
	Pair	string
	Type	TradeType
	Amount	float64
	Price	float64
}

var Log = log15.New("pkg", "btce_ticker")

type Args struct {
	Color   bool    `arg:"-c,help:enable ANSI colors"`
	Debug   bool    `arg:"-d,help:enable DEBUG logging"`
}


func InitLoggers(args *Args) {
	var LogHandler = log15.StreamHandler(os.Stderr, log15.TerminalFormat())

	if ! args.Debug {
		LogHandler = log15.LvlFilterHandler(log15.LvlInfo, LogHandler)
	}

	Log.SetHandler(LogHandler)
}


func (typ TradeType) String() string {
	switch typ {
	case TRADE_BID: return "BID"
	case TRADE_ASK: return "ASK"
	}
	return "UNK"
}

func ParseTradeType(typ string) (TradeType, error) {
	switch strings.ToUpper(typ) {
	case "ASK", "BUY":  return TRADE_ASK, nil
	case "BID", "SELL": return TRADE_BID, nil
	}
	return TRADE_UNK, fmt.Errorf("unknown trade type %#v", typ)
}

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

func main() {
	var args Args
	arg.MustParse(&args)

	InitLoggers(&args)

	Log.Info("[i] program started")

	var au = aurora.NewAurora(args.Color)

	var p_conn, err = pusher.NewClient(BTCE_PUSHER_APP_KEY)
	if err != nil {
		Log.Fatal("could not connect to the Pusher service", "err", err)
	}
	defer p_conn.Close()
	Log.Debug("connected to the Pusher service")

	if err = p_conn.Subscribe("btc_usd.trades"); err != nil {
		Log.Fatal("could not subscribe to the btc_usd.trades channel", "err", err)
	}
	if err = p_conn.Subscribe("ltc_usd.trades"); err != nil {
		Log.Fatal("could not subscribe to the ltc_usd.trades channel", "err", err)
	}
	Log.Debug("subscribed to channels")

	var event_ch, b_err = p_conn.Bind("trades")
	if b_err != nil {
		Log.Fatal("could not bind to the trades event", "err", b_err)
	}
	Log.Debug("bound to the trades events")

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

	Log.Info("[.] program exited")
}
