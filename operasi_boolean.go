package main

import "fmt"

func main() {
	// && and, || or, ! kebalikan

	var nilaiAkhir = 90
	var absen = 90

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsen bool = absen > 80

	var lulus bool = lulusNilaiAkhir && lulusAbsen

	fmt.Println(lulus)
	fmt.Println(true && !false && true || false)
}
