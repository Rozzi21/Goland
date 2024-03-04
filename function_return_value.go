package main

import "fmt"

func main() {
	hasil := tambah(10, 20)
	fmt.Println(hasil)
}

func tambah(angkaPertama int, angkaKedua int) int {
	return angkaKedua + angkaPertama
}
