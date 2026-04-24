package gobyexample

import "fmt"

const (
	Value1 = 3.14
	Value2 = 3.15
)

type ServerState int

const (
	StateIdle ServerState = iota // iota is 0 here
	// when the operation is omitted (type = value), it gets copied to the next lines
	StateConnected // ServerState = iota (iota is 1 here)
	StateError     // ServerState = iota (iota is 2 here)
	StateRetrying  // ServerState = iota (iota is 3 here)
)

var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func Enums() {
	fmt.Println(StateIdle)
	fmt.Println(StateConnected)
	fmt.Println(StateError)
	fmt.Println(StateRetrying)
}
