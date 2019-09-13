package main

import "sync"

type LazySafety struct{}

var lazySafety *LazySafety

var smu sync.Mutex

func GetLazySafety() *LazySafety {
	smu.Lock()
	defer smu.Unlock()

	if nil == lazySafety {
		lazySafety = &LazySafety{}
	}
	return lazySafety
}
