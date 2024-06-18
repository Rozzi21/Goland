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

func Greet(value Asking) {
	fmt.Println("hi ", value.getName())
}

type Somebody struct {
	Name    string
	address string
}

func (somebody Somebody) getName() string {
	return somebody.Name + " do you live in " + somebody.address
}

type Character struct {
	Name string
	Role string
}

func (Character Character) getName() string {
	return Character.Name + " your role is " + Character.Role + "?"
}

func main() {
	ozi := Person{"rozzi"}

	ilham := Somebody{"ilham", "lawang"}
	SayAsk(ozi)
	Greet(ilham)
	ultra := Character{"ultraman", "hyper"}
	SayAsk(ultra)

}
