package main

import "fmt"

func main() {
	contoh := getGoodBye
	fmt.Println(contoh("jhoni"))
}

func getGoodBye(nama string) string {
	return "dada bro " + nama
}
