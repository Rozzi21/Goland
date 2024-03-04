package main

import "fmt"

func main() {
	var name string

	name = "Fakhruddin Arrozi"
	fmt.Println(name)

	name = "fakhruddin ozzi"
	fmt.Println(name)

	//	tipe data otomatis

	var umur = 20
	fmt.Println(umur)

	// simplify menyerdehanakan tanpa menggunakan var
	hobi := "coding"
	fmt.Println(hobi)

	// multiple variable
	var (
		bahasa = "go"
		game   = "apex legend"
	)

	fmt.Println(bahasa)
	fmt.Println(game)
}

//note saat mendeklarasikan variable semuanya harus digunakan
//kalau tidak akan terjadi error
