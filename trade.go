package ticker

import (
	"fmt"
	"time"
	"strings"

	//"github.com/aglyzov/ws-machine"
)

type TradeType byte

const (
	UNK_TRADE TradeType = 0
	BID       TradeType = 1
	SELL      TradeType = 1
	ASK       TradeType = 2
	BUY       TradeType = 2
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
	case BID: return "BID"
	case ASK: return "ASK"
	}
	return "UNK"
}

func ParseTradeType(typ string) (TradeType, error) {
	switch strings.ToUpper(typ) {
	case "ASK", "BUY":  return ASK, nil
	case "BID", "SELL": return BID, nil
	}
	return UNK_TRADE, fmt.Errorf("unknown trade type %#v", typ)
}
