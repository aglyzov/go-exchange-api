package main

import (
	"fmt"
	"time"
	"strings"

	//"github.com/aglyzov/ws-machine"
)

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
