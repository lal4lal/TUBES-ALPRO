package main

import "fmt"

func menu() {
	var option int
	fmt.Println("==============Selamat Datang==============")
	fmt.Println("1. login sebagai pasien")
	fmt.Println("2. login sebagai dokter")
	fmt.Println("3. login sebagai guest")
	fmt.Println("4. keluar")
	fmt.Print("Masukkan pilihan anda: ")
	fmt.Scan(&option)
	for option < 1 || option > 3 {
		fmt.Println("Pilihan yang anda masukkan salah, Silahkan masukkan pilihan and kembali")
		fmt.Print("Masukkan pilihan anda: ")
		fmt.Scan(&option)
	}
	if option == 1 {
		fmt.Println(1)
	} else if option == 2 {
		fmt.Println(2)
	} else if option == 3 {
		fmt.Println(3)
	}
}

func main() {
	menu()
}
