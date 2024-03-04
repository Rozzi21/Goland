package main

import "fmt"

func main() {
	//array pada go harus di definisan dulu kapasitasnya analoginya seperti sebuah lemari
	var games [4]string
	games[0] = "Pubg"
	games[1] = "Genshin Impact"
	games[3] = "Apex Legend"

	fmt.Println(games[0])
	fmt.Println(games[1])
	fmt.Println(games[3])

	//array langsung
	var token = [10]int{
		1231,
		1123,
		12412,
		4124,
		14134,
		124,
		4124,
	}

	fmt.Println(token)
	fmt.Println(len(token))

	token[9] = 120

	fmt.Println(token)

	//[...] akan menghintung jumlah data array yang di dalamnya jadi harus langsung di deklarasikan isinya

	var nilai = [...]int{
		123,
		1421,
		124,
		4124,
	}

	fmt.Println(nilai)
}
