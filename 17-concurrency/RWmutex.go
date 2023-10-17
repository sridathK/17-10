package main

import "sync"

type cabs struct {
	driver int
	rw     sync.RWMutex
}
