package main

import "sync"

// sync.Once实现

type singleton struct{}

var ins *singleton
var once sync.Once

func GetIns() *singleton {
	once.Do(func() {
		ins = &singleton{}
	})
	return ins
}
func main() {}
