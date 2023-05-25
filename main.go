/*TEST */

package main

import (
	"fmt"
)

const nmax int = 2048

type pasien struct {
	pasienID string
	nama     string
	umur     int
	konsul   [nmax]string
	password string
}

type dokter struct {
	dokterID       string
	nama           string
	bidangKeahlian string
	password       string
}

type konsultasi struct {
	konsulID   string
	pasienID   string
	pertanyaan string
	tag        tag
}

type tag struct {
	konsulID string
	tag      [5]string
}

type dataPasien [nmax]pasien
type dataDokter [nmax]dokter

func menu(patient dataPasien) {
	var option int
	fmt.Println("==============Selamat Datang==============")
	fmt.Println("1. login sebagai pasien")
	fmt.Println("2. login sebagai dokter")
	fmt.Println("0. keluar")
	fmt.Print("Masukkan pilihan anda: ")
	fmt.Scan(&option)
	for option < 1 || option > 2 {
		fmt.Println("Pilihan yang anda masukkan salah, Silahkan masukkan pilihan and kembali")
		fmt.Print("Masukkan pilihan anda: ")
		fmt.Scan(&option)
	}
	if option == 1 {
		login_pasien(patient)
	} else if option == 2 {
		login_dokter()
	}
}

func signUp(patient *dataPasien, n *int) {

	fmt.Println("==============Sign Up==============")
	fmt.Println("Silahkan masukkan data yang dibutuhkan")
	fmt.Print("Nama Lengkap: ")
	fmt.Scan(&patient[*n].nama)
	fmt.Print("Umur: ")
	fmt.Scan(&patient[*n].umur)
	fmt.Print("Password:")
	fmt.Scan(&patient[*n].password)
	*n++
	fmt.Println("---Anda akan diarahkan kembali menuju login--- \n")
	login_pasien(*patient)
}

func login_pasien(patient dataPasien) {
	var option int
	var n int
	fmt.Println("==============Login==============")
	fmt.Println("1. Sudah mendaftar sebagai pasien")
	fmt.Println("2. Belum terdaftar sebagai pasien")
	fmt.Println("3. Masuk sebagai tamu")
	fmt.Print("Masukkan pilihan anda: ")
	fmt.Scan(&option)
	for option < 1 || option > 3 {
		fmt.Println("Pilihan yang anda masukkan salah, Silahkan masukkan pilihan and kembali")
		fmt.Println("Masukkan pilihan anda: ")
		fmt.Scan(&option)
	}
	if option == 2 {
		signUp(&patient, &n)
	}

}

func login_dokter() {
	fmt.Println("coba aja ini mah")
}

func main() {

}
