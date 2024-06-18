package main

import "fmt"

func Waduh() interface{} {
	//return 1
	//return true
	return "ini text"
}

func main() {
	var apasaja any = Waduh()
	fmt.Println(apasaja)
}
