package main

import (
	"fmt"
)

const nmax int = 2048

type pasien struct {
	nama     string
	password string
}

type dokter struct {
	nama     string
	password string
}

type konsultasi struct {
	pertanyaan  [nmax]string
	nPertanyaan int
	tanggapan   [nmax]string
	ntanggapan  int
	topik       [nmax]string
	ntopik      int
}

type topik struct {
	nKesehatanMental  int
	nKebugaranJasmani int
	nPenyakitUmum     int
}

type dataPasien struct {
	infoPasien [nmax]pasien
	n          int
	nama, pass string
	tanggapan  string
	replier    int
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
	var dokter dataDokter
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
		tampilanKonsul(&*patient, &*konsul, &dokter)
	} else if option == 2 {
		searchKonsultasiTag()
	}

}

func signUp(patient *dataPasien, konsul *dataKonsul) {
	fmt.Println("*==================Sign Up===================*")
	fmt.Println("|   Silahkan masukkan data yang dibutuhkan   |")
	fmt.Println("*============================================*")
	fmt.Print("Username: ")
	fmt.Scan(&patient.infoPasien[patient.n].nama)
	fmt.Print("Password: ")
	fmt.Scan(&patient.infoPasien[patient.n].password)
	patient.n++
	fmt.Println("---Anda akan diarahkan kembali menuju login---")
	login_pasien(patient, konsul)
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
		fmt.Println("Pilihan yang anda masukkan salah, Silahkan masukkan pilihan anda kembali")
		fmt.Println("Masukkan pilihan anda: ")
		fmt.Scan(&option)
	}
	if option == 2 {
		signUp(patient, konsul)
	} else if option == 3 {
		menuguest(&*patient, &*konsul)
	} else if option == 1 {
		login_pasien(patient, konsul)
	}
}

func login_pasien(patient *dataPasien, konsul *dataKonsul) {
	var success bool = false
	var idxPasien int
	patient.infoPasien[0].nama = "Andito"
	patient.infoPasien[0].password = "dito"
	konsul.infoKonsul[0].topik[0] = "Kesehatan Mental"
	konsul.infoKonsul[0].topik[1] = "Kesehatan Mental"
	patient.infoPasien[1].nama = "Hilal"
	patient.infoPasien[1].password = "hilal"
	konsul.infoKonsul[1].topik[0] = "Kesehatan Mental"
	konsul.infoKonsul[1].topik[1] = "Penyakit Umum"
	patient.infoPasien[2].nama = "Andi"
	patient.infoPasien[2].password = "andi"
	konsul.infoKonsul[1].topik[0] = "Kesehatan Mental"
	konsul.infoKonsul[1].topik[1] = "Penyakit Umum"
	patient.infoPasien[3].nama = "Pragos"
	patient.infoPasien[3].password = "pragos"
	konsul.infoKonsul[1].topik[0] = "Kebugaran Jasmani"
	konsul.infoKonsul[1].topik[1] = "Penyakit Umum"
	patient.infoPasien[4].nama = "Icha"
	patient.infoPasien[4].password = "icha"
	konsul.infoKonsul[1].topik[0] = "Kebugaran Jasmani"
	konsul.infoKonsul[1].topik[1] = "Penyakit Umum"

	fmt.Println("*==================Login==================*")
	fmt.Println("|       Input data anda dibawah ini       |")
	fmt.Println("*=========================================*")
	fmt.Println("(!) ketik 0 apabila belum memiliki akun (!)")
	fmt.Println("*=========================================*")
	fmt.Print("username :")
	fmt.Scan(&patient.nama)
	fmt.Print("Password :")
	fmt.Scan(&patient.pass)

	if patient.nama == "0" && patient.pass == "0" {
		menuPasien(patient, &*konsul)
	}

	for !success {
		for idxPasien = 0; idxPasien <= patient.n; idxPasien++ {
			if patient.infoPasien[idxPasien].nama == patient.nama && patient.infoPasien[idxPasien].password == patient.pass {
				success = true
			}
		}

		if !success {
			fmt.Printf("Mohon maaf! username atau password yang anda masukkan salah \n")
			fmt.Println("Silahkan coba lagi")
			fmt.Println("(ketik 0 apabila belum memiliki akun)")
			fmt.Print("username :")
			fmt.Scan(&patient.nama)
			fmt.Print("Password :")
			fmt.Scan(&patient.pass)

			if patient.nama == "0" && patient.pass == "0" {
				menuPasien(patient, &*konsul)
			}
		}
	}
	homePasien(&*patient, &*konsul)
}

func homePasien(patient *dataPasien, konsul *dataKonsul) {
	var dokter dataDokter
	var option, idxPasien int

	fmt.Println("")
	for idxPasien = 0; idxPasien <= patient.n; idxPasien++ {
		if patient.infoPasien[idxPasien].nama == patient.nama && patient.infoPasien[idxPasien].password == patient.pass {
			fmt.Println("Selamat Datang", patient.infoPasien[idxPasien].nama)
		}
	}
	fmt.Println("*================================*")
	fmt.Println("|   1. Konsultasi pada dokter!   |")
	fmt.Println("|   2. Tanggapi Konsultasi       |")
	fmt.Println("|   0. Keluar dari Akun          |")
	fmt.Println("*================================*")
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
		replyKonsultasiPasien(patient, konsul, &dokter)
	}
}

func searchKonsultasiTag() {
	fmt.Println("searchKonsultasiTag")
}

func tampilanKonsul(patient *dataPasien, konsul *dataKonsul, dokter *dataDokter) {
	var i, j, option int
	var topic topik
	konsul.infoKonsul[0].topik[0] = "Kesehatan Mental"
	konsul.infoKonsul[0].topik[1] = "Kesehatan Mental"
	konsul.infoKonsul[1].topik[0] = "Kesehatan Mental"
	konsul.infoKonsul[1].topik[1] = "Penyakit Umum"
	konsul.infoKonsul[2].topik[0] = "Kesehatan Mental"
	konsul.infoKonsul[2].topik[1] = "Penyakit Umum"
	konsul.infoKonsul[3].topik[0] = "Kebugaran Jasmani"
	konsul.infoKonsul[3].topik[1] = "Penyakit Umum"
	konsul.infoKonsul[4].topik[0] = "Kebugaran Jasmani"
	konsul.infoKonsul[4].topik[1] = "Penyakit Umum"
	fmt.Println("*================================================*")
	fmt.Println("|  Berikut Ini adalah daftar konsultasi pasien:  |")
	fmt.Println("*================================================*")

	for i = 0; i < patient.n; i++ {

		/* -- PROSES MENAMPILKAN PERTANYAAN -- */
		fmt.Printf("%v.) Nama Pasien: %v\n", i+1, patient.infoPasien[i].nama)
		for j = 0; j < konsul.infoKonsul[i].nPertanyaan; j++ {
			fmt.Printf("  %v. Topik : %v\n %v\n", j+1, konsul.infoKonsul[i].topik[j], konsul.infoKonsul[i].pertanyaan[j])
			hitungTopikFav(&*patient, &*konsul, &*dokter, &topic, i)
		}
		fmt.Println("Tanggapan : ")
		/* -- PROSES MENAMPILKAN TANGGAPAN -- */
		if konsul.infoKonsul[i].ntanggapan == 0 {
			fmt.Println("  [BELUM ADA TANGGAPAN]")
		} else {
			for j = 0; j < konsul.infoKonsul[i].ntanggapan; j++ {
				fmt.Printf("   %v\n", konsul.infoKonsul[i].tanggapan[j])
			}
		}

		fmt.Println("*================================================*")

	}

	fmt.Println("|           1. Kembali ke menu pasien            |")
	fmt.Println("|           0. Keluar akun                       |")
	fmt.Println("*================================================*")
	fmt.Print("Masukan Pilihan anda: ")
	fmt.Scan(&option)
	if option != 1 && option != 0 {
		fmt.Println("Pilihan yang anda masukkan salah, Silahkan masukkan pilihan anda kembali")
		fmt.Println("Masukkan pilihan anda: ")
		fmt.Scan(&option)
	}
	if option == 1 {
		homePasien(patient, konsul)
	} else {
		menuPasien(patient, konsul)
	}

}

func postKonsul_fromPasien(patient *dataPasien, konsul *dataKonsul) {
	var kalimat, kata string
	var option, idxPasien int
	var dokter dataDokter
	var key bool = true
	var found bool = false
	// var topic topik
	konsul.infoKonsul[0].topik[0] = "Kesehatan Mental"
	konsul.infoKonsul[0].topik[1] = "Kesehatan Mental"
	konsul.infoKonsul[1].topik[0] = "Kesehatan Mental"
	konsul.infoKonsul[1].topik[1] = "Penyakit Umum"
	konsul.infoKonsul[2].topik[0] = "Kesehatan Mental"
	konsul.infoKonsul[2].topik[1] = "Penyakit Umum"
	konsul.infoKonsul[3].topik[0] = "Kebugaran Jasmani"
	konsul.infoKonsul[3].topik[1] = "Penyakit Umum"
	konsul.infoKonsul[4].topik[0] = "Kebugaran Jasmani"
	konsul.infoKonsul[4].topik[1] = "Penyakit Umum"

	fmt.Println("Selamat Datang, Silahkan konsultasi", patient.nama)

	for key {
		found = false
		idxPasien = 0
		for idxPasien < patient.n && !found {
			if patient.infoPasien[idxPasien].nama == patient.nama && patient.infoPasien[idxPasien].password == patient.pass {
				found = true
				idxPasien--
			}
			idxPasien++
		}
		if !found {
			konsul.n++
		}
		fmt.Println("Pilih Topik Konsultasi :")
		fmt.Println("1. Penyakit Umum")
		fmt.Println("2. Kesehatan Mental")
		fmt.Println("3. Kebugaran Jasmani")
		fmt.Print("Masukan Pilihan anda: ")
		fmt.Scan(&option)
		for option != 1 && option != 2 && option != 3 {
			fmt.Println("Pilihan yang anda masukkan salah, Silahkan masukkan pilihan anda kembali")
			fmt.Println("Masukkan pilihan anda: ")
			fmt.Scan(&option)
		}
		if option == 1 {
			konsul.infoKonsul[idxPasien].topik[konsul.infoKonsul[idxPasien].ntopik] = "Penyakit Umum"
		} else if option == 2 {
			konsul.infoKonsul[idxPasien].topik[konsul.infoKonsul[idxPasien].ntopik] = "Kesehatan Mental"
		} else {
			konsul.infoKonsul[idxPasien].topik[konsul.infoKonsul[idxPasien].ntopik] = "Kebugaran Jasmani"
		}
		fmt.Println("*================================================================================*")
		fmt.Println("|    (!) petunjuk : klik enter lalu ketik 'post' apabila ingin memposting (!)    |")
		fmt.Println("*================================================================================*")
		fmt.Printf("Topik : %v\n", konsul.infoKonsul[idxPasien].topik[konsul.infoKonsul[idxPasien].ntopik])
		fmt.Print("Apa yang ingin anda konsultasikan? ")
		kalimat = ""
		kata = ""
		for kata != "post" {
			kalimat += kata + " "
			fmt.Scan(&kata)
		}
		konsul.infoKonsul[idxPasien].ntopik++
		konsul.infoKonsul[idxPasien].pertanyaan[konsul.infoKonsul[idxPasien].nPertanyaan] = kalimat
		fmt.Println("*================================================================================*")
		fmt.Println("|             Terima Kasih Atas Konsultasi Anda, Semoga Lekas Sembuh             |")
		fmt.Println("*================================================================================*")
		konsul.infoKonsul[idxPasien].nPertanyaan++
		fmt.Println("|                            1. Posting Konsultasi lain                          |")
		fmt.Println("|                            2. Lihat postingan anda                             |")
		fmt.Println("|                            0. kembali                                          |")
		fmt.Println("*================================================================================*")
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
			tampilanKonsul(&*patient, &*konsul, &dokter)
		} else {
			key = false
		}
	}
	homePasien(&*patient, &*konsul)
}

func hitungTopikFav(patient *dataPasien, konsul *dataKonsul, doctor *dataDokter, topic *topik, idxPasien int) {
	fmt.Println(idxPasien)
	if konsul.infoKonsul[idxPasien].topik[konsul.infoKonsul[idxPasien].ntopik] == "Kesehatan Mental" {
		topic.nKesehatanMental++
	} else if konsul.infoKonsul[idxPasien].topik[konsul.infoKonsul[idxPasien].ntopik] == "Kebugaran Jasmani" {
		topic.nKebugaranJasmani++
	} else {
		topic.nPenyakitUmum++
	}

	fmt.Println(topic.nPenyakitUmum)
	fmt.Println(topic.nKesehatanMental)
	fmt.Println(topic.nKebugaranJasmani)

}

func replyKonsultasiPasien(patient *dataPasien, konsul *dataKonsul, dokter *dataDokter) {
	var idxPasien, i, j, option int
	var kalimat, kata string
	var found, key bool

	fmt.Println("*================================================*")
	fmt.Println("|       Tanggapi Konsultasi Pasien lain!         |")
	fmt.Println("|        (pilih nama pasien konsultasi)          |")
	fmt.Println("*================================================*")

	/* -- SEQUENSIAL SEARCH (MENAMPILKAN NAMA BESERTA PERTANYAAN) -- */
	for i = 0; i < patient.n; i++ {

		/* -- PROSES MENAMPILKAN PERTANYAAN -- */
		fmt.Printf("%v.) Nama Pasien: %v\n", i+1, patient.infoPasien[i].nama)
		for j = 0; j < konsul.infoKonsul[i].nPertanyaan; j++ {
			fmt.Printf("  %v. %v\n", j+1, konsul.infoKonsul[i].pertanyaan[j])
		}
		fmt.Println("Tanggapan : ")
		/* -- PROSES MENAMPILKAN TANGGAPAN -- */
		if konsul.infoKonsul[i].ntanggapan == 0 {
			fmt.Println("  [BELUM ADA TANGGAPAN]")
		} else {
			// fmt.Println(patient.infoPasien[i].nama)
			for j = 0; j < konsul.infoKonsul[i].ntanggapan; j++ {
				fmt.Printf("  %v\n", konsul.infoKonsul[i].tanggapan[j])
			}
		}
		// fmt.Println(" Tanggapan :                                       ")
		fmt.Println("*=================================================*")
	}
	/* -- SEQUENSIAL SEARCH (MENAMPILKAN NAMA BESERTA PERTANYAAN) -- */

	/* -- SEQUENSIAL SEARCH (POSTING TANGGAPAN/REPLY) -- */
	fmt.Print("Pilih nama pasien : ")
	fmt.Scan(&patient.nama)
	key = true
	for key {
		found = false
		idxPasien = 0
		for idxPasien < patient.n && !found {
			if patient.infoPasien[idxPasien].nama == patient.nama {
				found = true
				idxPasien--
			}
			idxPasien++
		}
		if !found {
			konsul.n++
		}
		fmt.Println("*================================================================================*")
		fmt.Println("|    (!) petunjuk : klik enter lalu ketik 'post' apabila ingin memposting (!)    |")
		fmt.Println("*================================================================================*")
		fmt.Print("Apa yang ingin anda Tanggapi? ")
		kalimat = ""
		kata = ""
		for kata != "post" {
			kalimat += kata + " "
			fmt.Scan(&kata)
		}
		konsul.infoKonsul[idxPasien].tanggapan[konsul.infoKonsul[idxPasien].ntanggapan] = kalimat
		fmt.Println("*================================================================================*")
		fmt.Println("|             Terima Kasih Atas Konsultasi Anda, Semoga Lekas Sembuh             |")
		fmt.Println("*================================================================================*")
		konsul.infoKonsul[idxPasien].ntanggapan++
		fmt.Println("|                            1. Posting tanggapan lain                          |")
		fmt.Println("|                            0. kembali                                          |")
		fmt.Println("*================================================================================*")
		fmt.Print("Masukan Pilihan anda: ")
		fmt.Scan(&option)
		if option != 1 && option != 0 && option != 2 {
			fmt.Println("Pilihan yang anda masukkan salah, Silahkan masukkan pilihan anda kembali")
			fmt.Println("Masukkan pilihan anda: ")
			fmt.Scan(&option)
		}
		if option == 1 {
			key = true
		} else {
			key = false
		}
	}
	/* -- SEQUENSIAL SEARCH (POSTING TANGGAPAN/REPLY) -- */

	// fmt.Println(konsul.infoKonsul[0].ntanggapan)
	// fmt.Println(konsul.infoKonsul[1].ntanggapan)
	homePasien(patient, konsul)

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
		tampilanKonsul(&patient, &konsul, &doctor)
	} else if option == 2 {
		replyKonsultasiPasien(&patient, &konsul, &doctor)
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
	patient.n = 5
	menu(patient)
}
