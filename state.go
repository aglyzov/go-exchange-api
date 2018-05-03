//
// Package ticker defines common interfaces, constants and utility functions
// for all exchange clients
//
package ticker

import "fmt"

type (
	Command byte
	State   byte
	Status	struct {
		State	State
		Error	error
	}
)

const (
	// states
	DISCONNECTED State = iota
	CONNECTING
	CONNECTED
	WAITING
)

func (s State) String() string {
	switch s {
	case DISCONNECTED:  return "DISCONNECTED"
	case CONNECTING:    return "CONNECTING"
	case CONNECTED:     return "CONNECTED"
	case WAITING:       return "WAITING"
	}
	return fmt.Sprintf("UNKNOWN STATUS %v", s)
}

