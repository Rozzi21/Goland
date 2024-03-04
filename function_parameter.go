package main

import "fmt"

func main() {
	greeting("Fakhruddin", "Arrozi")
	greeting("Hafiz", "Nafi'uddin")
}

func greeting(firstName string, lastName string) {
	fmt.Println("halo", firstName, lastName)
}
