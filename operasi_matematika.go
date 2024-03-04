package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var c = a + b
	var d = a - b
	var e = a * b
	var f = a / b
	var g = a % b
	fmt.Println("hasil tambah", c)
	fmt.Println("hasil kurang", d)
	fmt.Println("hasil kali", e)
	fmt.Println("hasil bagi", f)
	fmt.Println("hasil sisa bagi", g)

	//	augmented assigment +=, -= , *=, /=, %=
	var nilai = 10
	nilai += 10 // nilai = nilai + 10
	fmt.Println(nilai)
	nilai /= 5 // nilai = nilai / 5
	fmt.Println(nilai)

	// Unary Operator ++, --, !,+,-
	var angka = 1
	angka++
	fmt.Println(angka)
	angka++
	fmt.Println(angka)
	angka--
	fmt.Println(angka)
	angka--
	fmt.Println(angka)
	fmt.Println(-angka)
	var boolean = true
	fmt.Println(!boolean)

}
