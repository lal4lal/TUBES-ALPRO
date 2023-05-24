package main

import "fmt"

const nmax int = 2048

type pasien struct {
	nama   string
	umur   int
	konsul [nmax]string
}

type dokter struct {
}

type dataPasien [nmax]pasien
type dataDokter [nmax]dokter

func menu() {
	var option int
	fmt.Println("==============Selamat Datang==============")
	fmt.Println("1. login sebagai pasien")
	fmt.Println("2. login sebagai guest")
	fmt.Println("3. keluar")
	fmt.Print("Masukkan pilihan anda: ")
	fmt.Scan(option)
	for option < 1 || option > 3 {
		fmt.Println("Pilihan yang anda masukkan salah, Silahkan masukkan pilihan and kembali")
		fmt.Print("Masukkan pilihan anda: ")
		fmt.Scan(option)
	}
	if option == 1 {
		login_pasien()
	} else if option == 2 {
		login_guest()
	}
}

func signUp(patient *dataPasien, n *int) {

	fmt.Println("==============Sign Up==============")
	fmt.Println("Silahkan masukkan data yang dibutuhkan")
	fmt.Print("Nama Lengkap: ")
	fmt.Scan(&patient[*n].nama)
	fmt.Print("Umur: ")
	fmt.Scan(&patient[*n].umur)
}

func login_pasien() {

}

func login_guest() {

}

func main() {

}
