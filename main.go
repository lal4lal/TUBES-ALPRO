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

func menu() {
	var patient dataPasien
	var option int

	fmt.Println("==============Selamat Datang==============")
	fmt.Println("1. Sebagai pasien")
	fmt.Println("2. Sebagai dokter")
	fmt.Println("0. keluar")
	fmt.Print("Masukkan pilihan anda: ")
	fmt.Scan(&option)
	for option != 1 && option != 2 && option != 0 {
		fmt.Println("Pilihan yang anda masukkan salah, Silahkan masukkan pilihan and kembali")
		fmt.Print("Masukkan pilihan anda: ")
		fmt.Scan(&option)
	}
	if option == 1 {
		menu_pasien(patient)
	} else if option == 2 {
		login_dokter()
	} else if option == 0 {
		fmt.Println("Keluar dari aplikasi...")
	}
}

func menu_pasien(patient dataPasien) {
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
		signUpPasien(&patient, &n)
		fmt.Println(patient[n].nama)
	} else if option == 1 {
		/* biar mudahin ibunya, kita assign aja value dititik ini seakan udah daftar */
		patient[n].nama = "dosen"
		patient[n].umur = 30
		patient[n].password = "alpro"
		login_pasien(patient, n)
	} else if option == 3 {
		menu_tamu()
	}
}

func mainmenu_pasien() {
	fmt.Println("yeyy")
}

func menu_tamu() {
	fmt.Println("ini menu tamu")
}

func signUpPasien(patient *dataPasien, n *int) {

	fmt.Println("==============Sign Up==============")
	fmt.Println("Silahkan masukkan data yang dibutuhkan")
	fmt.Print("Nama Lengkap: ")
	fmt.Scan(&patient[*n].nama)
	fmt.Print("Umur: ")
	fmt.Scan(&patient[*n].umur)
	fmt.Print("Password:")
	fmt.Scan(&patient[*n].password)
	*n++
	// fmt.Println("---Anda akan diarahkan kembali menuju login--- \n")
	// if login_pasien(*patient, *n) {
	// 	mainmenu_pasien()
	// } else {
	// 	fmt.Println("data tidak ditemukan")
	// 	login_pasien(*patient, *n)
	// }
}

func login_pasien(patient dataPasien, n int) bool {
	var nama, pass string
	fmt.Println("--- Login ---")
	fmt.Println("Input data anda :")
	fmt.Print("username :")
	fmt.Scan(&nama)
	fmt.Print("Password :")
	fmt.Scan(&pass)
	if patient[n].nama == nama && patient[n].password == pass || patient[n].nama == "dosen" {
		return true
	}
	return false
}

func login_dokter() {

}

func main() {
	menu()

}
