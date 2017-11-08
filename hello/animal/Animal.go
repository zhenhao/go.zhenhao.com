package animal

import (
	"reflect"
)

type None interface {
}

type Animal struct {
	name           string
	age            int32
	height, weight float32
	this           None
}

func (a *Animal) SetName(name string) {
	a.name = name
}

func (a *Animal) SetAge(age int32) {
	a.age = age
}

func (a *Animal) SetHeight(h float32) {
	a.height = h
}

func (a *Animal) SetWeight(w float32) {
	a.weight = w
}

func (a *Animal) Name() string {
	return a.name
}

func (a *Animal) Age() int32 {
	return a.age
}

func (a *Animal) Height() float32 {
	return a.height
}

func (a *Animal) Weight() float32 {
	return a.weight
}

func (a *Animal) Bark() string {
	return a.Sound()
}

func (a *Animal) setThis(this None) {
	a.this = this
}

func (a *Animal) Sound() string {
	if nil != a.this {
		t := reflect.ValueOf(a.this).MethodByName("Sound").Call([]reflect.Value{})
		return t[0].String()
	}
	return "I am an Animal named " + a.name
}
