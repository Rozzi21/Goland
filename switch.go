package main

import "fmt"

func main() {
	name := "ozi gaming"

	switch name {
	case "Ozi":
		fmt.Println("Halo ozi")
	case "Hafiz":
		fmt.Println("halo Hafiz")
	default:
		fmt.Println("nama tidak terdaftar")
	}

	//	switch short statement

	switch length := len(name); length > 6 {
	case true:
		fmt.Println("nama kepanjangan cuy")
	case false:
		fmt.Println("nama enggak panjang sih")

	}

	//switch tanpa kondisi
	length := len(name)

	switch {
	case length < 10:
		fmt.Println("nama cukup panjang juga ya")
	case length < 5:
		fmt.Println("nama kependekan ya")
	default:
		fmt.Println("nama kamu kok gg gitu")
	}
}
