package main

import "fmt"

func main() {
	name := "rozigaming"

	if name == "rozi" {
		fmt.Println("halo bang rozi")
	} else if name != "" {
		fmt.Println("halo", name, "boleh kenalan?")
	} else {
		fmt.Println("anda siapa")
	}

	if length := len(name); length < 6 {
		fmt.Println("namanya enggak panjang cuy")
	} else {
		fmt.Println("namanya kepanjangan anjay")
	}

}
