package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

type Subject interface {
	Attach(Observer) //注册观察者
	Detach(Observer) //释放观察者
	Notify()         //通知所有注册的观察者
}

type ConcreteSubject struct {
	observers *list.List
	value     int
}

func (s *ConcreteSubject) Attach(observe Observer) { //注册观察者
	s.observers.PushBack(observe)
}

func (s *ConcreteSubject) Detach(observer Observer) { //释放观察者
	for ob := s.observers.Front(); ob != nil; ob = ob.Next() {
		if ob.Value.(*Observer) == &observer {
			s.observers.Remove(ob)
			break
		}
	}
}

func (s *ConcreteSubject) Notify() { //通知所有观察者
	for ob := s.observers.Front(); ob != nil; ob = ob.Next() {
		ob.Value.(Observer).Update(s)
	}
}

func (s *ConcreteSubject) setValue(value int) {
	s.value = value
	s.Notify()
}

func (s *ConcreteSubject) getValue() int {
	return s.value
}

type Observer interface {
	Update(Subject) //观察者进行更新状态
}

func NewConcreteSubject() *ConcreteSubject {
	s := new(ConcreteSubject)
	s.observers = list.New()
	return s
}

type HexaObserver struct{}

func (h *HexaObserver) Update(subject Subject) {
	fmt.Println("Hex String: " + strings.ToUpper(strconv.FormatInt(int64(subject.(*ConcreteSubject).getValue()), 16)))
}

type OctalObserver struct{}

func (o *OctalObserver) Update(subject Subject) {
	fmt.Println("Octal String: " + strconv.FormatInt(int64(subject.(*ConcreteSubject).getValue()), 8))
}

type BinaryObserver struct{}

func (b *BinaryObserver) Update(subject Subject) {
	fmt.Println("Octal String: " + strconv.FormatInt(int64(subject.(*ConcreteSubject).getValue()), 2))
}

func main() {
	subject := NewConcreteSubject()

	binaryObserver := new(BinaryObserver)
	octalObserver := new(OctalObserver)
	hexaObserver := new(HexaObserver)

	subject.Attach(hexaObserver)
	subject.Attach(binaryObserver)
	subject.Attach(octalObserver)

	subject.setValue(15)
	fmt.Println("----------------")
	subject.setValue(10)
	// fmt.Println(strconv.FormatInt(10, 2))
}
