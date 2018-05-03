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

// TODO: implement the Ticker interface

