package main

import "sync"

type DoubleCheckLock struct{}

var dcl *DoubleCheckLock
var mu sync.Mutex

func GetDoubleCheckLock() *DoubleCheckLock {
	if nil == dcl {
		mu.Lock()
		defer mu.Unlock()

		if nil == dcl {
			dcl = &DoubleCheckLock{}
		}
	}

	return dcl
}
