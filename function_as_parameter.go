package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) {
	filteredname := filter(name)
	fmt.Println("hello", filteredname)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Eko", spamFilter)

	filter := spamFilter

	sayHelloWithFilter("anjing", filter)
}
