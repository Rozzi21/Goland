package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println("integer32", nilai32)
	fmt.Println("integer64", nilai64)
	fmt.Println("integer16", nilai16)
	fmt.Println("interger8", nilai8)

	//konversi string
	var name = "fakhruddin arrozi"
	var e uint8 = name[0]

	var eString = string(e)

	fmt.Println(e)
	fmt.Println(eString)

}
