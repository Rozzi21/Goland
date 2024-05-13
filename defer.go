package main

import "fmt"

func logging() {
	fmt.Println("ini adalah loging terpanggil")
}

func runApps() {

	defer logging()
	fmt.Println("Run App")
}

func main() {
	runApps()
}
