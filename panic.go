package main

import "fmt"

func endApps() {
	fmt.Println("aplikasi di hentikan")
}

func runApp(error bool) {
	defer endApps()

	if error {
		panic("Error ada kesalahan")
	}
}

func main() {
	runApp(true)
}
