package main

import "fmt"

func main() {

	//bikin map mode 1
	var person map[string]string = map[string]string{}
	person["name"] = "ozzi"
	person["address"] = "malang"
	fmt.Println(person)

	//	bikin map mode 2 lebih cepat cuy

	persons := map[string]string{
		"name":    "mihdan",
		"address": "lamongan",
	}

	fmt.Println(persons)
	fmt.Println(persons["name"])

	anime := make(map[string]string)

	anime["title"] = "gintama"
	anime["rate"] = "9,7"
	anime["salah"] = "lah kok iso"

	fmt.Println(anime)
	delete(anime, "salah")
	fmt.Println(anime)

	fmt.Println(len(anime))
	fmt.Println("rating dari", anime["title"], "adalah", anime["rate"])

}
