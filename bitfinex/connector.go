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

func (_ *API) GetName() string {
	return "Bitfinex websocket API v2"
}

func (a *API) Connect(key string) (<-chan api.Status, error) {
	a.Lock()
	defer a.Unlock()

	switch a.status.State {
	case api.DISCONNECTED:  // TODO
	case api.CONNECTING:  // TODO
	case api.CONNECTED:  // TODO
	case api.WAITING:  // TODO
	}

	return a.statusChan, nil
}

func (a *API) Disconnect() error {
	a.Lock()
	defer a.Unlock()

	/*
	switch a.status.State {
	case api.DISCONNECTED:  // TODO
	case api.CONNECTING:  // TODO
	case api.CONNECTED:  // TODO
	case api.WAITING:  // TODO
	}
	*/

	return nil
}

