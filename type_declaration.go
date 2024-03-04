package main

import "fmt"

func main() {

	type NoMahasiswa string

	var ktmOzzi NoMahasiswa = "21.52.0023"

	var contoh string = "222222222"
	var contohKtm NoMahasiswa = NoMahasiswa(contoh)

	fmt.Println(ktmOzzi)
	fmt.Println(contohKtm)

}
