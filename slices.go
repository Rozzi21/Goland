package main

import "fmt"

func main() {

	//hati hati ya bang dalam bikin array atau slice
	iniArray := [...]int{1, 2, 3}
	iniSlice := []int{1, 2, 3}

	fmt.Println(iniArray, iniSlice)

	nama := [...]string{"jhon", "leyry", "herry", "manroe", "cronjob", "xiao", "ali"}
	fmt.Println(nama)
	namaslice1 := nama[5:]
	namaslice2 := nama[:4]
	namaslice3 := nama[1:4]
	namaslice4 := nama[:]

	//deklarasi slice seperti biasa
	var slice5 []string = nama[4:6]
	fmt.Println(namaslice1)
	fmt.Println(namaslice2)
	fmt.Println(namaslice3)
	fmt.Println(namaslice4)
	fmt.Println(slice5)

	//cek jumlah array
	fmt.Println(cap(namaslice1))
	//menambahkan data ke slice
	fmt.Println(append(slice5, "larry"))

	hari := [...]string{"senin", "selasa", "rabu"}
	harislice := hari[1:]
	fmt.Println(harislice)
	harislice[0] = "selasa baru"
	harislice[1] = "rabu baru"
	fmt.Println(hari)

	harislice1 := (append(harislice, "kamis"))
	fmt.Println(harislice)
	fmt.Println(harislice1)
	fmt.Println(hari)
	harislice1[0] = "senin lama"
	fmt.Println(harislice1)

	//	membuat slice baru otomatis buat array

	sliceBaru := make([]string, 2, 5)
	sliceBaru[0] = "satu"
	sliceBaru[1] = "dua"
	//sliceBaru[2] = "tiga" akan terjadi error karena lengh nya kita set 2 harus menggunakan append

	fmt.Println(sliceBaru)
	fmt.Println(cap(sliceBaru))
	fmt.Println(len(sliceBaru))

	sliceBaru1 := append(sliceBaru, "tiga")
	fmt.Println(sliceBaru1)
	fmt.Println(cap(sliceBaru1))
	fmt.Println(len(sliceBaru1))
	sliceBaru1[0] = "1"
	fmt.Println(sliceBaru1)
	fmt.Println(sliceBaru)

	//	copy slice

	fromSlice := hari[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)
}
