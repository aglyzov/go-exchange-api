//
// Package bitfinex implements the exchange-api interfaces for the Bitfinex websocket API v2
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

const API_URL = "wss://api.bitfinex.com/ws/2"

type (
	// API represents a connection to the Bitfinex websocket API and
	// implements the Connector, Ticker and Trader interfaces
	API struct {
		sync.Mutex
		key           string
		wsm			  *machine.Machine
		status        Status
		tradesChan    <-chan *ticker.Trade
		statusChan	  <-chan ticker.Status
		subscriptions []string
	}
)

var Log = .Log.New("e", "BFX")


func NewAPI() *API {
	return &API{
		// TODO: init wsm, channels, ...
	}
}
