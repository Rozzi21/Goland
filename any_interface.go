package main

import "fmt"

func Waduh() interface{} {
	return 1
}

func main() {
	var apasaja any = Waduh()
	fmt.Println(apasaja)
}
