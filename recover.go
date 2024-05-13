package main

import "fmt"

func stopApp() {
	fmt.Println("berhentikan aplikasi")
	pesan := recover()
	fmt.Println("terjadi error", pesan)
}

func jalanApp(err bool) {
	defer stopApp()
	if err {
		panic("ERROR COY APLIKASI KEBAKARAN")
	}
}

func main() {
	jalanApp(true)
	fmt.Println("cepatkan benerin")
}
