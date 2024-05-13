package main

import "fmt"

type Anime struct {
	Name, Genre string
	Rating      int
}

func main() {
	//	membuat data struct menggunakan var

	var black Anime
	var red Anime
	black.Name = "Black clover"
	black.Genre = "fiksi"
	black.Rating = 9.0

	fmt.Println(black)
	fmt.Println(black.Rating)
	fmt.Println(red)

	naruto := Anime{
		Name:   "Naruto",
		Genre:  "fight",
		Rating: 9.0,
	}

	fmt.Println(naruto)

	naga := Anime{"nagabonar", "fiksi", 8.0}

	fmt.Println(naga)
}
