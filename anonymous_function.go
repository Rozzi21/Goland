package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("halo", name)
	}
}

func main() {
	kata := func(name string) bool {
		{
			return name == "asu"
		}
	}
	registerUser("jhon", kata)
}
