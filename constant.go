package main

import "fmt"

func main() {
	//constant adalah variable yang tidak berubah dan tidak ada error jika constant tidak dipakai berbeda dengan var
	const id = 11223344

	//error
	//id = 12312
	fmt.Println(id)

	const (
		anime  = "gintama"
		rating = 9
	)

	fmt.Println("anime =", anime, "rating =", rating)

}
