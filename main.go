package main

import (
	"fmt"
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
	idxkonsul  int
}

type tag struct {
	konsulID string
	tag      [5]string
}

type dataPasien struct {
	infoPasien [nmax]pasien
	n          int
	nama, pass string
}

type dataDokter struct {
	infoDokter [nmax]dokter
	n          int
	nama, pass string
}
type dataKonsul struct {
	infoKonsul [nmax]konsultasi
	n          int
}

func menu(patient dataPasien) {
	var option int
	var konsul dataKonsul
	var doctor dataDokter
	fmt.Println("*=============Selamat Datang=============*")
	fmt.Println("|          Silahkan Pilih Role           |")
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
		menuPasien(&patient, &konsul)
	} else if option == 2 {
		login_dokter(&doctor, &konsul)
	}
}
func menuguest(patient *dataPasien, konsul *dataKonsul) {
	var option int
	// var konsul dataKonsul
	// var idx konsultasi
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
		menuPasien(&*patient, &*konsul)
	} else if option == 1 {
		// postKonsul_fromPasien(&patient, &konsul)
		sortingKonsultasiTag(&*patient, &*konsul)
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

func menuPasien(patient *dataPasien, konsul *dataKonsul) {
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
		menuguest(&*patient, &*konsul)
	} else if option == 1 {
		login_pasien(patient)
	}
}

func login_pasien(patient *dataPasien) {
	var success bool = false
	var konsul dataKonsul
	var idxPasien int
	fmt.Println("*==================Login==================*")
	fmt.Println("Input data anda dibawah ini")
	fmt.Print("username :")
	fmt.Scan(&patient.nama)
	fmt.Print("Password :")
	fmt.Scan(&patient.pass)

	for !success {
		for idxPasien = 0; idxPasien <= patient.n; idxPasien++ {
			if patient.infoPasien[idxPasien].nama == patient.nama && patient.infoPasien[idxPasien].password == patient.pass {
				success = true
			}
		}

		if !success {
			fmt.Printf("Mohon maaf! username atau password yang anda masukkan salah \n")
			fmt.Println("Silahkan coba lagi")
			fmt.Print("username :")
			fmt.Scan(&patient.nama)
			fmt.Print("Password :")
			fmt.Scan(&patient.pass)
		}
	}
	homePasien(&*patient, &konsul)
}

func homePasien(patient *dataPasien, konsul *dataKonsul) {
	var option, idxPasien int

	fmt.Println("")
	for idxPasien = 0; idxPasien <= patient.n; idxPasien++ {
		if patient.infoPasien[idxPasien].nama == patient.nama && patient.infoPasien[idxPasien].password == patient.pass {
			fmt.Println("Selamat Datang", patient.infoPasien[idxPasien].nama)
		}
	}
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
		menuPasien(&*patient, &*konsul)
	} else if option == 1 {
		postKonsul_fromPasien(&*patient, &*konsul)
	} else if option == 2 {
		replyKonsultasiPasien()
	}
}

func searchKonsultasiTag() {
	fmt.Println("searchKonsultasiTag")
}

func sortingKonsultasiTag(patient *dataPasien, konsul *dataKonsul) {
	var i, option int
	fmt.Println("")
	fmt.Println("Berikut Ini adalah daftar konsultasi pasien :")
	fmt.Println("---------------------------------------------")
	for i < 5 {
		fmt.Println("")
		fmt.Println(i+1, "Dari : ", patient.infoPasien[i].nama)
		fmt.Println(konsul.infoKonsul[i].pertanyaan)
		i++
	}
	fmt.Println("Anda hanya bisa melihat dalam mode tamu, Daftar akun?")
	fmt.Println("=====================================================")
	fmt.Println("1. Daftar akun                                       ")
	fmt.Println("2. Kembali ke menu tamu                              ")
	fmt.Println("0. Keluar mode tamu                                  ")
	fmt.Print("Masukan Pilihan anda: ")
	fmt.Scan(&option)
	if option != 1 && option != 0 && option != 2 {
		fmt.Println("Pilihan yang anda masukkan salah, Silahkan masukkan pilihan anda kembali")
		fmt.Println("Masukkan pilihan anda: ")
		fmt.Scan(&option)
	}
	if option == 1 {
		signUp(patient)
	} else if option == 2 {
		menuguest(patient, konsul)
	} else {
		menuPasien(patient, konsul)
	}

	// menuPasien(patient,konsul)

}

func postKonsul_fromPasien(patient *dataPasien, konsul *dataKonsul) {
	var kalimat, kata string
	var option, idx int
	var key bool = true

	for key != false {
		fmt.Println("")
		fmt.Println("Silahkan masukkan masalah kesehatan anda: ")
		fmt.Println("-----------------------------------------------------------------")
		fmt.Println("petunjuk : klik enter lalu ketik 'post' apabila ingin memposting ")
		fmt.Print("Apa yang ingin anda konsultasikan? ")
		kalimat = ""
		kata = ""
		for kata != "post" {
			kalimat += kata + " "
			fmt.Scan(&kata)
		}
		konsul.infoKonsul[idx].pertanyaan = kalimat
		fmt.Println("Anda berhasil memposting! ")
		fmt.Println("1. Posting Konsultasi lain")
		fmt.Println("2. Lihat postingan anda   ")
		fmt.Println("0. kembali  ")
		fmt.Print("Masukan Pilihan anda: ")
		fmt.Scan(&option)
		if option != 1 && option != 0 && option != 2 {
			fmt.Println("Pilihan yang anda masukkan salah, Silahkan masukkan pilihan anda kembali")
			fmt.Println("Masukkan pilihan anda: ")
			fmt.Scan(&option)
		}
		if option == 1 {
			key = true
		} else if option == 2 {
			sortingKonsultasiTag(patient, konsul)
		} else {
			key = false
		}
		idx++
	}
	// fmt.Println(konsul.infoKonsul[0].pertanyaan)
	homePasien(&*patient, &*konsul)

}

func replyKonsultasiPasien() {

}

func homedokter(doctor dataDokter, konsul dataKonsul) {
	var option int
	var patient dataPasien
	fmt.Println("")
	fmt.Println("Selamat Datang Dr.", doctor.infoDokter[0].nama)
	fmt.Println("-----------------------------------")
	fmt.Println(" 1. Tampilakan topik ")
	fmt.Println(" 2. Tanggapi Konsultasi     ")
	fmt.Println(" 0. Keluar dari mode dokter        ")
	fmt.Print("Masukkan pilihan anda: ")
	fmt.Scan(&option)
	for option != 1 && option != 0 && option != 2 {
		fmt.Println("Pilihan yang anda masukkan salah, Silahkan masukkan pilihan anda kembali")
		fmt.Println("Masukkan pilihan anda: ")
		fmt.Scan(&option)
	}

	if option == 0 {
		login_dokter(&doctor, &konsul)
	} else if option == 1 {
		sortingKonsultasiTag(&patient, &konsul)
	} else if option == 2 {
		replyKonsultasiPasien()
	}

}

func login_dokter(doctor *dataDokter, konsul *dataKonsul) {
	doctor.infoDokter[0].nama = "jordy"
	doctor.infoDokter[0].password = "jordy"
	var success bool
	var idxdoctor int
	fmt.Println("")
	fmt.Println("*==================Login==================*")
	fmt.Println("Input data anda dibawah ini")
	fmt.Print("username :")
	fmt.Scan(&doctor.nama)
	fmt.Print("Password :")
	fmt.Scan(&doctor.pass)

	for !success {
		for idxdoctor = 0; idxdoctor <= doctor.n; idxdoctor++ {
			if doctor.infoDokter[idxdoctor].nama == doctor.nama && doctor.infoDokter[idxdoctor].password == doctor.pass {
				success = true
			}
		}

		if !success {
			fmt.Printf("Mohon maaf! username atau password yang anda masukkan salah \n")
			fmt.Println("Silahkan coba lagi")
			fmt.Print("username :")
			fmt.Scan(&doctor.nama)
			fmt.Print("Password :")
			fmt.Scan(&doctor.pass)
		}
	}
	homedokter(*doctor, *konsul)
}

func main() {
	/* asumsi sementara kalo dokter ini udah tetap, karena di soal juga gadisebutin kalo dokter bisa daftar
	jadi fokusnya kita sekarang nyelsein function inti aja biar cepet
	username doctor = jordy
	pass = jordy */
	var patient dataPasien
	menu(patient)
}
