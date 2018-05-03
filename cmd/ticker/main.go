package main

import (
	"os"
	//"fmt"
	"time"
	//"strings"
	"context"
	"reflect"

	"github.com/aglyzov/log15"
	"github.com/alexflint/go-arg"
	//"github.com/logrusorgru/aurora"

	bfx   "github.com/bitfinexcom/bitfinex-api-go/v2"
	bfxws "github.com/bitfinexcom/bitfinex-api-go/v2/websocket"
)

var Log = log15.New("pkg", "ticker")

type Args struct {
	Debug		bool     `arg:"-d" help:"enable DEBUG logging"`
	Color		bool     `arg:"-c" help:"enable ANSI colors"`
	Ticker		[]string `arg:"-t,separate" help:"enable ticker for <exchange>:<symbol>"`
	//BitfinexKey string   `arg"env:BFX_API_KEY" help:"set Bitfinex API key"`
}


func InitLoggers(args *Args) {
	var LogHandler = log15.StreamHandler(os.Stderr, log15.TerminalFormat())

	if ! args.Debug {
		LogHandler = log15.LvlFilterHandler(log15.LvlInfo, LogHandler)
	}

	Log.SetHandler(LogHandler)
}

func main() {
	var args Args
	arg.MustParse(&args)

	InitLoggers(&args)

	Log.Debug("[i] program started")

	//var au = aurora.NewAurora(args.Color)

	var bitfinex = SetupBitfinexAPI()

	for obj := range bitfinex.Listen() {
		switch obj.(type) {
		case error:
			Log.Warn("channel closed", "err", obj)
			break
		case *bfx.TradeExecution:
			var trade = obj.(*bfx.TradeExecution)
			Log.Info("RECV", "trade", trade)
		default:
			Log.Info("RECV", "type", reflect.TypeOf(obj), "obj", obj)
		}
	}

	Log.Debug("[.] program exited")
}


func SetupBitfinexAPI() *bfxws.Client {
	c := bfxws.NewClient()

	err := c.Connect()
	if err != nil {
		Log.Fatal("could not connect to websocket", "err", err)
	}
	c.SetReadTimeout(time.Second * 8)

	/*
	// subscribe to XRPUSD ticker
	ctx, cxl1 := context.WithTimeout(context.Background(), time.Second*4)
	defer cxl1()
	_, err = c.SubscribeTicker(ctx, bfx.TradingPrefix+bfx.XRPUSD)
	if err != nil {
		Log.Fatal("could not subscribe", "err", err)
	}
	*/

	// subscribe to XRPUSD trades
	ctx, cxl2 := context.WithTimeout(context.Background(), time.Second*4)
	defer cxl2()
	_, err = c.SubscribeTrades(ctx, bfx.TradingPrefix+bfx.XRPUSD)
	if err != nil {
		Log.Fatal("could not subscribe", "err", err)
	}

	return c
}
