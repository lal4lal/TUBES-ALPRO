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

type dataPasien struct {
	infoPasien [nmax]pasien
	n          int
}
type dataDokter [nmax]dokter

func menu(patient dataPasien) {
	var option int
	fmt.Println("*=============Selamat Datang=============*")
	fmt.Println("|          1. Sebagai pasien             |")
	fmt.Println("|          2. Sebagai dokter             |")
	fmt.Println("|          0. keluar                     |")
	fmt.Println("*========================================*")
	fmt.Print("Masukkan pilihan anda: ")
	fmt.Scan(&option)
	for option < 1 || option > 2 {
		fmt.Println("Pilihan yang anda masukkan salah, Silahkan masukkan pilihan and kembali")
		fmt.Print("Masukkan pilihan anda: ")
		fmt.Scan(&option)
	}
	if option == 1 {
		menuPasien(patient)
	} else if option == 2 {
		login_dokter()
	}
}

func menuguest() {
	fmt.Println("guest")
}

func signUp(patient *dataPasien) {

	fmt.Println("*==================Sign Up===================*")
	fmt.Println("|   Silahkan masukkan data yang dibutuhkan   |")
	fmt.Println("*============================================*")
	fmt.Print("Username: ")
	fmt.Scan(&patient.infoPasien[patient.n].nama)
	fmt.Print("Umur: ")
	fmt.Scan(&patient.infoPasien[patient.n].umur)
	fmt.Print("Password: ")
	fmt.Scan(&patient.infoPasien[patient.n].password)
	patient.n++
	fmt.Printf("---Anda akan diarahkan kembali menuju login--- \n")
	login_pasien(*patient)
}

func menuPasien(patient dataPasien) {
	var option int
	fmt.Println("*==================Login==================*")
	fmt.Println("|    1. Sudah mendaftar sebagai pasien    |")
	fmt.Println("|    2. Belum terdaftar sebagai pasien    |")
	fmt.Println("|    3. Masuk sebagai tamu                |")
	fmt.Println("*=========================================*")
	fmt.Print("Masukkan pilihan anda: ")
	fmt.Scan(&option)
	for option < 1 || option > 3 {
		fmt.Println("Pilihan yang anda masukkan salah, Silahkan masukkan pilihan and kembali")
		fmt.Println("Masukkan pilihan anda: ")
		fmt.Scan(&option)
	}
	if option == 2 {
		signUp(&patient)
	} else if option == 3 {
		menuguest()
	} else if option == 1 {
		login_pasien(patient)
	}
}

func login_pasien(patient dataPasien) {
	var nama, pass string
	var success bool = false
	var i int
	fmt.Println("*==================Login==================*")
	fmt.Println("Input data anda :")
	fmt.Print("username :")
	fmt.Scan(&nama)
	fmt.Print("Password :")
	fmt.Scan(&pass)
	for !success {
		for i = 0; i < patient.n; i++ {
			if (patient.infoPasien[i].nama == nama || patient.infoPasien[i].nama == "dosen") && patient.infoPasien[i].password == pass {
				success = true
			}
		}
		if !success {
			fmt.Printf("Mohon maaf! username atau password yang anda masukkan salah \n")
			fmt.Println("Silahkan coba lagi")
			fmt.Print("username :")
			fmt.Scan(&nama)
			fmt.Print("Password :")
			fmt.Scan(&pass)
		}
	}
	homePasien()
}

func homePasien() {
	fmt.Println("sukses")
}

func login_dokter() {

}

func main() {
	var pasien dataPasien
	menu(pasien)

}
