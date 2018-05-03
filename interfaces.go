//
// Package ticker defines common interfaces, constants and utility functions
// for all exchange clients
//
package ticker

type (
	Connector interface {
		GetURL()              string
		GetName()             string
		GetStatus()           Status
		Connect(key string)   (<-chan Status, error)
		Disconnect()          error
	}
	Ticker interface {
		Subscribe([]string)   (<-chan *Trade, error)
		Unsubscribe([]string) error
	}
	Trader interface {
		// TODO: orders, etc.
	}
	Exchanger interface {
		Connector
		Ticker
		Trader
	}
)

