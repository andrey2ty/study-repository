package dockerpractice

import (
	"sync"
	"sync/atomic"
)

var Id atomic.Int64
var RMtx sync.RWMutex
var Mtx sync.Mutex
var PeopleM = make(map[int]People)

type People struct {
	ID       int
	FullName string
	Position string
}

func NewPeople(fullname string, position string) People {
	return People{
		ID:       int(Id.Add(1)),
		FullName: fullname,
		Position: position,
	}
}
