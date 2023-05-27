package main

import (
	"bufio"
	"fmt"
	"os"
)

const nmax int = 2048

type pasien struct {
	pasienID string
	nama     string
	umur     int
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
type dataDokter struct {
	infoDokter [nmax]dokter
	n          int
}
type dataKonsul struct {
	infoKonsul [nmax]konsultasi
	n          int
}

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
		menuPasien(&patient)
	} else if option == 2 {
		login_dokter()
	}
}
func menuguest(patient dataPasien) {
	var option int
	fmt.Println("*================Welcome==================*")
	fmt.Println("|      1. Lihat Konsultasi Pasien         |")
	fmt.Println("|      2. Cari Konsultasi Pasien          |")
	fmt.Println("|      0. Kembali ke menu                 |")
	fmt.Println("*=========================================*")
	fmt.Print("Masukkan pilihan anda: ")
	fmt.Scan(&option)
	for option != 1 && option != 0 && option != 2 {
		fmt.Println("Pilihan yang anda masukkan salah, Silahkan masukkan pilihan and kembali")
		fmt.Println("Masukkan pilihan anda: ")
		fmt.Scan(&option)
	}
	if option == 0 {
		menuPasien(&patient)
	} else if option == 1 {
		sortingKonsultasiTag()
	} else if option == 2 {
		searchKonsultasiTag()
	}

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
	login_pasien(patient)
}

func menuPasien(patient *dataPasien) {
	var option int
	fmt.Println("*================Welcome==================*")
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
		signUp(patient)
	} else if option == 3 {
		menuguest(*patient)
	} else if option == 1 {
		login_pasien(patient)
	}
}

func login_pasien(patient *dataPasien) {
	var nama, pass string
	var success bool = false
	var idxPasien int
	fmt.Println("*==================Login==================*")
	fmt.Println("Input data anda dibawah ini")
	fmt.Print("username :")
	fmt.Scan(&nama)
	fmt.Print("Password :")
	fmt.Scan(&pass)
	for !success {
		for idxPasien = 0; idxPasien < patient.n; idxPasien++ {
			if patient.infoPasien[idxPasien].nama == nama && patient.infoPasien[idxPasien].password == pass {
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
	homePasien(patient, idxPasien)
}

func homePasien(patient *dataPasien, idxPasien int) {
	var i, option int
	var konsul dataKonsul
	fmt.Println("")
	fmt.Println("Selamat Datang", patient.infoPasien[i].nama)
	fmt.Println("-----------------------------------")
	fmt.Println(" 1. Konsultasi pada dokter! ")
	fmt.Println(" 2. Tanggapi Konsultasi     ")
	fmt.Println(" 0. Keluar dari Akun        ")
	fmt.Print("Masukkan pilihan anda: ")
	fmt.Scan(&option)
	for option != 1 && option != 0 && option != 2 {
		fmt.Println("Pilihan yang anda masukkan salah, Silahkan masukkan pilihan anda kembali")
		fmt.Println("Masukkan pilihan anda: ")
		fmt.Scan(&option)
	}
	if option == 0 {
		menuPasien(patient)
	} else if option == 1 {
		addKonsultasiPasien(patient, &konsul, idxPasien)
	} else if option == 2 {
		replyKonsultasiPasien()
	}
}

func searchKonsultasiTag() {
	fmt.Println("searchKonsultasiTag")
}

func sortingKonsultasiTag() {
	fmt.Println("sorting")
}

func addKonsultasiPasien(patient *dataPasien, konsul *dataKonsul, idxPasien int) {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		fmt.Printf("Silahkan masukkan masalah kesehatan anda: ")
		konsul.infoKonsul[idxPasien].pertanyaan = scanner.Text()
		fmt.Printf("Input was: %q\n", konsul.infoKonsul[idxPasien].pertanyaan)
	}
}

func replyKonsultasiPasien() {

}

func login_dokter() {

}

func main() {
	var pasien dataPasien
	menu(pasien)

}
