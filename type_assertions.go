package main

import (
	"fmt"
)

func random() any {
	return "YES"
}

func newRandom() any {
	return 1
}

func main() {
	var result any = random()
	var newResult any = newRandom()

	resultString := result.(string)

	fmt.Println(resultString)

	////make panic
	//var resultInt int = result.(int)
	//
	//fmt.Println(resultInt)

	switch value := newResult.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	case float32:
		fmt.Println("float", value)
	default:
		fmt.Println("ini tipe data apa ya", value)
	}

}
