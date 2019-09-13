package main

type Lazy struct{}

var lazy *Lazy

func GetLazy() *Lazy {

	if lazy == nil {
		lazy = &Lazy{}
	}
	return lazy
}
