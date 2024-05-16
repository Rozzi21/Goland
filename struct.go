package main

import "fmt"

type Anime struct {
	Name, Genre string
	Rating      int
}

func (anime Anime) sayReview(review string) {
	fmt.Println("my review about", anime.Name, "its very", review)
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

	naga.sayReview("good")
	naruto.sayReview("it is mastrepiece")

	//	struct Method

}
