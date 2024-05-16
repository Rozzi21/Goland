package main

import (
	"fmt"
)

type Asking interface {
	getName() string
}

func SayAsk(value Asking) {
	fmt.Println("hello", value.getName())
}

type Person struct {
	Name string
}

func (person Person) getName() string {
	return person.Name + ""
}

//try

func About(value Asking) {
	fmt.Println("hi ", value.getName())
}

type Somebody struct {
	Name   string
	addres string
}

func (somebody Somebody) getName() string {
	return somebody.Name + " do you live in " + somebody.addres
}

func main() {
	ozi := Person{"rozzi"}

	ilham := Somebody{"ilham", "lawang"}
	SayAsk(ozi)
	About(ilham)
}
