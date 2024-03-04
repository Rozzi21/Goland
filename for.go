package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("perulangan ke", counter)
		counter++
	}

	fmt.Println("sudah selesai perulangannya coy")

	//	for menggunakan statement
	// init statement dan post statement

	for count := 1; count <= 10; count++ {
		fmt.Println("perulangan ke", count)
	}

	// for range
	names := []string{"eko", "rozzi", "jhonny"}
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
