package main

type Hungry struct{}

var hungry *Hungry = &Hungry{}

func GetHungry() *Hungry {
	return hungry
}
