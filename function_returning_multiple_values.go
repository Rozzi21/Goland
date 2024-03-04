package main

import "fmt"

func main() {
	judul, penulis := komik()
	fmt.Println(judul, penulis)

	//	menghiraukan return

	title, _ := komik()
	fmt.Println(title)
}

func komik() (string, string) {
	return "naruto", "mashashi kishimoto"
}
