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

type dataTopik struct {
	ntopik int
	topik  string
}

type dataPasien struct {
	infoPasien [nmax]pasien
	n          int
	nama, pass string
	guest      bool
}

type dataDokter struct {
	infoDokter [nmax]dokter
	n          int
	nama, pass string
	reply      bool
}
type dataKonsul struct {
	infoKonsul [nmax]konsultasi
	n          int
	ntopik     [nmax]dataTopik
	topik      [nmax]string
}

func menu(patient *dataPasien, konsul *dataKonsul, dokter *dataDokter, topik *dataTopik) {
	/*
		I.S Terdefinisi data patient, konsul, dokter, dan topik
		F.S Memunculkan pilihan menu dan user dapat memilih tujuan sesuai menu
	*/
	var option int
	fmt.Println("*========Selamat Datang di HolaDoc=======*")
	fmt.Println("|          Silahkan Pilih Role           |")
	fmt.Println("|          1. Sebagai Pasien             |")
	fmt.Println("|          2. Sebagai Dokter             |")
	fmt.Println("|          3. Sebagai Tamu               |")
	fmt.Println("|          0. Keluar Aplikasi            |")
	fmt.Println("*========================================*")
	fmt.Print("Masukkan pilihan anda: ")
	fmt.Scan(&option)
	for option != 1 && option != 2 && option != 3 && option != 0 {
		fmt.Println("Pilihan yang anda masukkan salah, Silahkan masukkan pilihan and kembali")
		fmt.Print("Masukkan pilihan anda: ")
		fmt.Scan(&option)
	}
	if option == 1 {
		menuPasien(patient, konsul, dokter, topik)
	} else if option == 2 {
		login_dokter(patient, konsul, dokter, topik)
	} else if option == 3 {
		menuguest(patient, konsul, dokter, topik)
	} else {
		fmt.Println("Keluar Aplikasi...")
	}
}
func menuguest(patient *dataPasien, konsul *dataKonsul, dokter *dataDokter, topik *dataTopik) {
	/*
		I.S Terdefinisi data patient, konsul, dokter, dan topik
		F.S Memunculkan pilihan menu untuk tamu dan user dapat memilih tujuan sesuai menu
	*/
	var option int
	fmt.Println("*================Welcome==================*")
	fmt.Println("|      1. Lihat Konsultasi Pasien         |")
	fmt.Println("|      2. Cari Konsultasi Pasien          |")
	fmt.Println("|      0. Ganti Role                      |")
	fmt.Println("*=========================================*")
	fmt.Print("Masukkan pilihan anda: ")
	fmt.Scan(&option)
	for option != 1 && option != 0 && option != 2 {
		fmt.Println("Pilihan yang anda masukkan salah, Silahkan masukkan pilihan and kembali")
		fmt.Println("Masukkan pilihan anda: ")
		fmt.Scan(&option)
	}
	if option == 0 {
		menu(patient, konsul, dokter, topik)
	} else if option == 1 {
		patient.guest = true
		tampilanKonsul(patient, konsul, dokter, topik)
	} else if option == 2 {
		patient.guest = true
		searchKonsultasiTag(patient, konsul, dokter, topik)
	}
}

func signUp(patient *dataPasien, konsul *dataKonsul, dokter *dataDokter, topik *dataTopik) {
	/*
		I.S Terdefinisi data patient, konsul, dokter, dan topik
		F.S Nama dan password yang dimasukkan untuk signup akan masuk ke ARRAY dataPasien
	*/
	fmt.Println("*==================Sign Up===================*")
	fmt.Println("|   Silahkan masukkan data yang dibutuhkan   |")
	fmt.Println("*============================================*")
	fmt.Print("Username: ")
	fmt.Scan(&patient.infoPasien[patient.n].nama)
	fmt.Print("Password: ")
	fmt.Scan(&patient.infoPasien[patient.n].password)
	patient.n++
	fmt.Println("---Anda akan diarahkan kembali menuju login---")
	login_pasien(patient, konsul, dokter, topik)
}

func menuPasien(patient *dataPasien, konsul *dataKonsul, dokter *dataDokter, topik *dataTopik) {
	/*
		I.S Terdefinisi data patient, konsul, dokter, dan topik
		F.S Memunculkan menu untuk pasien dan user dapat memilih tujuan sesuai menu
	*/
	var option int

	fmt.Println("*================Welcome==================*")
	fmt.Println("|    1. Sudah mendaftar sebagai pasien    |")
	fmt.Println("|    2. Belum terdaftar sebagai pasien    |")
	fmt.Println("|    0. Ganti Role                        |")
	fmt.Println("*=========================================*")
	fmt.Print("Masukkan pilihan anda: ")
	fmt.Scan(&option)
	for option != 0 && option != 1 && option != 2 {
		fmt.Println("Pilihan yang anda masukkan salah, Silahkan masukkan pilihan anda kembali")
		fmt.Println("Masukkan pilihan anda: ")
		fmt.Scan(&option)
	}
	if option == 2 {
		signUp(patient, konsul, dokter, topik)
	} else if option == 0 {
		menu(patient, konsul, dokter, topik)
	} else if option == 1 {
		login_pasien(patient, konsul, dokter, topik)
	}
}

func login_pasien(patient *dataPasien, konsul *dataKonsul, dokter *dataDokter, topik *dataTopik) {
	/*
		I.S Terdefinisi data patient, konsul, dokter, dan topik
		F.S Nama dan password yang dimasukkan untuk login akan terdefinisi sebagai patient.nama dan patient.pass dan     memunculkan menu homePasien jika berhasil login
	*/
	var success bool = false
	var idxPasien int

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
		menuPasien(patient, konsul, dokter, topik)
	}

	for !success {
		//Searching login Nama Pasien dan Password pasien yang telah dimasukkan di ARRAY infoPasien
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
				menuPasien(patient, konsul, dokter, topik)
			}
		}
	}
	homePasien(patient, konsul, dokter, topik)
}

func homePasien(patient *dataPasien, konsul *dataKonsul, dokter *dataDokter, topik *dataTopik) {
	/*
		I.S Terdefinisi data patient, konsul, dokter, dan topik
		F.S Memunculkan tampilan home pasien dan pilihan menu pasien dan user dapat memilih tujuan sesuai dengan pilihan menu pasien
	*/
	var option, idxPasien int

	fmt.Println("")
	//Searching index Pasien yang sedang login saat ini
	for idxPasien = 0; idxPasien <= patient.n; idxPasien++ {
		if patient.infoPasien[idxPasien].nama == patient.nama && patient.infoPasien[idxPasien].password == patient.pass {
			fmt.Println("")
			fmt.Println("Selamat Datang", patient.infoPasien[idxPasien].nama)
		}
	}
	//tampilan home Pasien
	fmt.Println("*=========================================*")
	fmt.Println("|        1. Konsultasi pada dokter!       |")
	fmt.Println("|        2. Tanggapi Konsultasi           |")
	fmt.Println("|        3. Kunjungi Forum                |")
	fmt.Println("|        4. Cari Konsultasi Pasien        |")
	fmt.Println("|        0. Keluar dari Akun              |")
	fmt.Println("*=========================================*")
	fmt.Print("Masukkan pilihan anda: ")
	fmt.Scan(&option)
	for option != 1 && option != 0 && option != 2 && option != 3 && option != 4 {
		fmt.Println("Pilihan yang anda masukkan salah, Silahkan masukkan pilihan anda kembali")
		fmt.Println("Masukkan pilihan anda: ")
		fmt.Scan(&option)
	}

	if option == 0 {
		menuPasien(patient, konsul, dokter, topik)
	} else if option == 1 {
		postKonsul_fromPasien(patient, konsul, dokter, topik)
	} else if option == 2 {
		replyKonsultasiPasien(patient, konsul, dokter, topik)
	} else if option == 3 {
		tampilanKonsul(patient, konsul, dokter, topik)
	} else if option == 4 {
		searchKonsultasiTag(patient, konsul, dokter, topik)
	}
}

func searchKonsultasiTag(patient *dataPasien, konsul *dataKonsul, dokter *dataDokter, topik *dataTopik) {
	/*
		I.S Terdefinisi data patient, konsul, dokter, dan topik
		F.S Memunculkan pilihan 5 tag/topik penyakit yang telah ditentukan dan user dapat memilih salah satu topik penyakit dan menampilkan seluruh konsultasi tentang topik tersebut
	*/
	var option int
	var i, j, n int

	fmt.Println("*=========================================*")
	fmt.Println("|  Pilih topik yang ada ingin cari :      |")
	fmt.Println("|  1. Penyakit Umum     4. Spesialis Gigi |")
	fmt.Println("|  2. Kesehatan Mental  5. Spesialis THT  |")
	fmt.Println("|  3. Kandungan         0. Kembali        |")
	fmt.Println("*=========================================*")
	fmt.Print("Masukan Pilihan anda: ")
	fmt.Scan(&option)
	for option != 1 && option != 0 && option != 2 && option != 3 && option != 4 && option != 5 {
		fmt.Println("Pilihan yang anda masukkan salah, Silahkan masukkan pilihan anda kembali")
		fmt.Println("Masukkan pilihan anda: ")
		fmt.Scan(&option)
	}
	fmt.Println("Berikut Topik yang anda cari : ")
	fmt.Println("")

	if option == 1 {
		//Searching Semua data konsultasi yang memiliki topik Penyakit Umum
		n = 0
		for i = 0; i < patient.n; i++ {
			for j = 0; j < konsul.infoKonsul[i].nPertanyaan; j++ {
				if konsul.infoKonsul[i].topik[j] == "Penyakit Umum" {
					fmt.Printf("%v.Dari Pasien: %v\n", n+1, patient.infoPasien[i].nama)
					n++
				}
				if konsul.infoKonsul[i].topik[j] == "Penyakit Umum" {
					fmt.Printf("  Topik : %v\n %v\n", konsul.infoKonsul[i].topik[j], konsul.infoKonsul[i].pertanyaan[j])
					fmt.Println("")
				}
			}

		}
	} else if option == 2 {
		//Searching Semua data konsultasi yang memiliki topik Kesehatan Mental
		n = 0
		for i = 0; i < patient.n; i++ {
			for j = 0; j < konsul.infoKonsul[i].nPertanyaan; j++ {
				if konsul.infoKonsul[i].topik[j] == "Kesehatan Mental" {
					fmt.Printf("%v.Dari Pasien: %v\n", n+1, patient.infoPasien[i].nama)
					n++
				}
				if konsul.infoKonsul[i].topik[j] == "Kesehatan Mental" {
					fmt.Printf("  Topik : %v\n %v\n", konsul.infoKonsul[i].topik[j], konsul.infoKonsul[i].pertanyaan[j])
					fmt.Println("")
				}
			}

		}
	} else if option == 3 {
		//Searching Semua data konsultasi yang memiliki topik Kandungan
		n = 0
		for i = 0; i < patient.n; i++ {
			for j = 0; j < konsul.infoKonsul[i].nPertanyaan; j++ {
				if konsul.infoKonsul[i].topik[j] == "Kandungan" {
					fmt.Printf("%v.Dari Pasien: %v\n", n+1, patient.infoPasien[i].nama)
					n++
				}
				if konsul.infoKonsul[i].topik[j] == "Kandungan" {
					fmt.Printf("  Topik : %v\n %v\n", konsul.infoKonsul[i].topik[j], konsul.infoKonsul[i].pertanyaan[j])
					fmt.Println("")
				}
			}
		}
	} else if option == 4 {
		//Searching Semua data konsultasi yang memiliki topik Spesialis Gigi
		for i = 0; i < patient.n; i++ {
			for j = 0; j < konsul.infoKonsul[i].nPertanyaan; j++ {
				if konsul.infoKonsul[i].topik[j] == "Spesialis Gigi" {
					fmt.Printf("%v.Dari Pasien: %v\n", n+1, patient.infoPasien[i].nama)
					n++
				}
				if konsul.infoKonsul[i].topik[j] == "Spesialis Gigi" {
					fmt.Printf("  Topik : %v\n %v\n", konsul.infoKonsul[i].topik[j], konsul.infoKonsul[i].pertanyaan[j])
					fmt.Println("")
				}
			}

		}
	} else if option == 5 {
		//Searching Semua data konsultasi yang memiliki topik Spesialis THT
		n = 0
		for i = 0; i < patient.n; i++ {
			for j = 0; j < konsul.infoKonsul[i].nPertanyaan; j++ {
				if konsul.infoKonsul[i].topik[j] == "Spesialis THT" {
					fmt.Printf("%v.Dari Pasien: %v\n", n+1, patient.infoPasien[i].nama)
					n++
				}
				if konsul.infoKonsul[i].topik[j] == "Spesialis THT" {
					fmt.Printf("  Topik : %v\n %v\n", konsul.infoKonsul[i].topik[j], konsul.infoKonsul[i].pertanyaan[j])
					fmt.Println("")
				}
			}

		}

	}
	//mengecek apakah yang login adalah tamu, jika iya maka user hanya dapat melihat konsultasi
	if patient.guest {
		fmt.Print("[Pada mode tamu, anda hanya bisa melihat pertanyaan.] Kembali? (0 untuk kembali) : ")
		fmt.Scan(&option)
		for option != 0 {
			fmt.Println("Pilihan yang anda masukkan salah, Silahkan masukkan pilihan anda kembali")
			fmt.Println("Masukkan pilihan anda: ")
			fmt.Scan(&option)
		}
		if option == 0 {
			patient.guest = false
			menuguest(patient, konsul, dokter, topik)
		}
	} else {
		fmt.Println("*===========================*")
		fmt.Println("|         0. Kembali        |")
		fmt.Println("*===========================*")
		fmt.Print("Masukkan Pilihan Anda: ")
		fmt.Scan(&option)
		for option != 0 {
			fmt.Println("Pilihan yang anda masukkan salah, Silahkan masukkan pilihan anda kembali")
			fmt.Println("Masukkan pilihan anda: ")
			fmt.Scan(&option)
		}
		if option == 0 {
			homePasien(patient, konsul, dokter, topik)
		}
	}
}

func tampilanKonsul(patient *dataPasien, konsul *dataKonsul, dokter *dataDokter, topik *dataTopik) {
	/*
		I.S Terdefinisi data patient, konsul, dokter, dan topik
		F.S Menampilkan Daftar konsultasi Pasien dan mengembalikan ke menu pasien atau menu dokter atau menu guest sesuai kondisi login saat ini
	*/
	var i, j, option int

	fmt.Printf("*=========================================*\n")
	fmt.Printf("|  Berikut Daftar Konsultasi pasien :     |\n")
	fmt.Printf("*=========================================*\n")
	fmt.Printf("|  Jumlah Konsultasi Kesehatan Mental: %v  |\n", konsul.ntopik[0].ntopik)
	fmt.Printf("|  Jumlah Konsultasi Penyakit Umum: %v     |\n", konsul.ntopik[1].ntopik)
	fmt.Printf("|  Jumlah Konsultasi Kandungan: %v         |\n", konsul.ntopik[2].ntopik)
	fmt.Printf("|  Jumlah Konsultasi Masalah Gigi: %v      |\n", konsul.ntopik[3].ntopik)
	fmt.Printf("|  Jumlah Konsultasi Masalah THT: %v       |\n", konsul.ntopik[4].ntopik)
	fmt.Printf("*=========================================*\n\n")
	for i = 0; i < patient.n; i++ {
		//PROSES MENAMPILKAN PERTANYAAN
		fmt.Printf("%v. Dari Pasien: %v\n", i+1, patient.infoPasien[i].nama)
		for j = 0; j < konsul.infoKonsul[i].nPertanyaan; j++ {
			fmt.Printf("   Topik : %v\n", konsul.infoKonsul[i].topik[j])
			fmt.Printf("  %v\n", konsul.infoKonsul[i].pertanyaan[j])
		}
		fmt.Println("   Tanggapan : ")
		//PROSES MENAMPILKAN TANGGAPAN
		if konsul.infoKonsul[i].ntanggapan == 0 {
			fmt.Printf("   [BELUM ADA TANGGAPAN]\n\n")
		} else {
			for j = 0; j < konsul.infoKonsul[i].ntanggapan; j++ {
				fmt.Printf("	%v\n\n", konsul.infoKonsul[i].tanggapan[j])
			}
		}
	}
	//mengecek apakah yang sedang login adalah guest, jika iya maka user hanya dapat melihat daftar konsultasi
	if patient.guest {
		fmt.Print("[Pada mode tamu, anda hanya bisa melihat pertanyaan.] Kembali? (0 untuk kembali) : ")
		fmt.Scan(&option)
		for option != 0 {
			fmt.Println("Pilihan yang anda masukkan salah, Silahkan masukkan pilihan anda kembali")
			fmt.Println("Masukkan pilihan anda: ")
			fmt.Scan(&option)
		}
		if option == 0 {
			patient.guest = false
			menuguest(patient, konsul, dokter, topik)
		}
	} else if dokter.reply {
		//mengecek apakah yang login adalah dokter, jika iya makan user akan dikembalikan ke home dokter tanpa kembali ke menu pasien
		fmt.Print("Kembali? (0 untuk kembali) : ")
		fmt.Scan(&option)
		for option != 0 {
			fmt.Println("Pilihan yang anda masukkan salah, Silahkan masukkan pilihan anda kembali")
			fmt.Println("Masukkan pilihan anda: ")
			fmt.Scan(&option)
		}
		if option == 0 {
			dokter.reply = false
			homedokter(patient, konsul, dokter, topik)
		}
	} else {
		fmt.Println("*================================================*")
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
			homePasien(patient, konsul, dokter, topik)
		} else {
			menuPasien(patient, konsul, dokter, topik)
		}
	}
}

func postKonsul_fromPasien(patient *dataPasien, konsul *dataKonsul, dokter *dataDokter, topik *dataTopik) {
	/*
		I.S Terdefinisi data patient, konsul, dokter, dan topik
		F.S Menampilkan pilihan topik yang dapat dipilih oleh pasien yang akan menjadi tag konsultasi tersebut, pasien dapat memposting konsultasi dan meemilih untuk memposting lagi atau kembali ke home pasien
	*/
	var kalimat, kata string
	var option, idxPasien int
	var key bool = true
	var found bool = false
	fmt.Println("Selamat Datang, Silahkan konsultasi", patient.nama)

	for key {
		//mengecek apakah pasien yang saat ini sedang login sudah pernah membuat konsultasi
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
		//pemilihan topik/ tag yang akan dikonsultasikan pasien
		fmt.Println("*=========================================*")
		fmt.Println("|          Pilih Topik Konsultasi         |")
		fmt.Println("*=========================================*")
		fmt.Println("|           1. Penyakit Umum              |")
		fmt.Println("|           2. Kesehatan Mental           |")
		fmt.Println("|           3. Kandungan                  |")
		fmt.Println("|           4. Spesialis Gigi             |")
		fmt.Println("|           5. Spesialis THT              |")
		fmt.Println("*=========================================*")
		fmt.Print("Masukan Pilihan anda: ")
		fmt.Scan(&option)
		for option != 1 && option != 2 && option != 3 && option != 4 && option != 5 {
			fmt.Println("Pilihan yang anda masukkan salah, Silahkan masukkan pilihan anda kembali")
			fmt.Println("Masukkan pilihan anda: ")
			fmt.Scan(&option)
		}
		if option == 1 {
			konsul.infoKonsul[idxPasien].topik[konsul.infoKonsul[idxPasien].ntopik] = "Penyakit Umum"
			konsul.ntopik[1].ntopik++
		} else if option == 2 {
			konsul.infoKonsul[idxPasien].topik[konsul.infoKonsul[idxPasien].ntopik] = "Kesehatan Mental"
			konsul.ntopik[0].ntopik++
		} else if option == 3 {
			konsul.infoKonsul[idxPasien].topik[konsul.infoKonsul[idxPasien].ntopik] = "Kandungan"
			konsul.ntopik[2].ntopik++
		} else if option == 4 {
			konsul.infoKonsul[idxPasien].topik[konsul.infoKonsul[idxPasien].ntopik] = "Spesialis Gigi"
			konsul.ntopik[3].ntopik++
		} else if option == 5 {
			konsul.infoKonsul[idxPasien].topik[konsul.infoKonsul[idxPasien].ntopik] = "Spesialis THT"
			konsul.ntopik[4].ntopik++
		}
		//Proses Posting konsultasi pasien
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
		konsul.infoKonsul[idxPasien].nPertanyaan++
		fmt.Println("*================================================================================*")
		fmt.Println("|             Terima Kasih Atas Konsultasi Anda, Semoga Lekas Sembuh             |")
		fmt.Println("*================================================================================*")
		fmt.Println("|                            1. Posting Konsultasi lain                          |")
		fmt.Println("|                            0. kembali                                          |")
		fmt.Println("*================================================================================*")
		fmt.Print("Masukan Pilihan anda: ")
		fmt.Scan(&option)
		if option != 1 && option != 0 {
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
	homePasien(patient, konsul, dokter, topik)
}

func replyKonsultasiPasien(patient *dataPasien, konsul *dataKonsul, dokter *dataDokter, topik *dataTopik) {
	/*
		I.S Terdefinisi data patient, konsul, dokter, dan topik
		F.S Menampilkan Daftar konsultasi Pasien dan tanggapan jika tersedia, Pasien atau pun dokter dapat memilih konsultasi mana yang akan diberikan tanggapan, dan akan dikembalikan ke menu pasien adatau menu dokter sesuai kondisi login saat ini
	*/
	var idxPasien, i, j, option int
	var kalimat, kata, ceknama string
	var found, key bool

	fmt.Println("*================================================*")
	fmt.Println("|       Tanggapi Konsultasi Pasien lain!         |")
	fmt.Println("|        (pilih nama pasien konsultasi)          |")
	fmt.Println("*================================================*")

	//SEQUENSIAL SEARCH (MENAMPILKAN NAMA BESERTA PERTANYAAN)
	for i = 0; i < patient.n; i++ {
		//PROSES MENAMPILKAN PERTANYAAN
		fmt.Printf("%v. Dari Pasien: %v\n", i+1, patient.infoPasien[i].nama)
		for j = 0; j < konsul.infoKonsul[i].nPertanyaan; j++ {
			fmt.Printf("   Topik : %v\n", konsul.infoKonsul[i].topik[j])
			fmt.Printf("  %v\n", konsul.infoKonsul[i].pertanyaan[j])
		}
		fmt.Println("   Tanggapan : ")
		//PROSES MENAMPILKAN TANGGAPAN
		if konsul.infoKonsul[i].ntanggapan == 0 {
			fmt.Printf("   [BELUM ADA TANGGAPAN]\n\n")
		} else {
			for j = 0; j < konsul.infoKonsul[i].ntanggapan; j++ {
				fmt.Printf("   %v\n\n", konsul.infoKonsul[i].tanggapan[j])
			}
		}
	}
	fmt.Print("Pilih nama pasien : ")
	fmt.Scan(&ceknama)
	key = true
	for key {
		//SEQUENSIAL SEARCH (POSTING TANGGAPAN/REPLY)
		found = false
		idxPasien = 0
		for idxPasien < patient.n && !found {
			if patient.infoPasien[idxPasien].nama == ceknama {
				found = true
				idxPasien--
			}
			idxPasien++
		}
		if !found {
			konsul.n++
		}
		//proses memasukkan teks tanggapan untuk konsul pasien
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
		//mengecek apakah yang memberi tanggapan dokter, jika bukan berarti pasien yang memberi tanggapan
		if dokter.reply {
			kalimat = "Dari Dr. Jordy :" + kalimat
		} else {
			kalimat = "Dari " + patient.nama + ":" + kalimat
		}

		konsul.infoKonsul[idxPasien].tanggapan[konsul.infoKonsul[idxPasien].ntanggapan] = kalimat
		//mengecek apakah kondisi login sekarang adalah dokter, jika iya maka akan dikembalikan ke home dokter
		if dokter.reply {
			fmt.Print("Kembali? (0 untuk kembali) : ")
			fmt.Scan(&option)
			konsul.infoKonsul[idxPasien].ntanggapan++
			for option != 0 {
				fmt.Println("Pilihan yang anda masukkan salah, Silahkan masukkan pilihan anda kembali")
				fmt.Print("Masukkan pilihan anda: ")
				fmt.Scan(&option)
			}
			if option == 0 {
				key = false
			}
		} else {
			konsul.infoKonsul[idxPasien].ntanggapan++
			fmt.Println("*=======================================================*")
			fmt.Println("|           Terima Kasih Atas Tanggapan Anda            |")
			fmt.Println("*=======================================================*")
			fmt.Println("|           1. Posting tanggapan lain                   |")
			fmt.Println("|           0. kembali                                  |")
			fmt.Println("*=======================================================*")
			fmt.Print("Masukan Pilihan anda: ")
			fmt.Scan(&option)
			if option != 1 && option != 0 {
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

	}
	if dokter.reply {
		dokter.reply = false
		homedokter(patient, konsul, dokter, topik)

	} else {
		homePasien(patient, konsul, dokter, topik)
	}
}

func homedokter(patient *dataPasien, konsul *dataKonsul, dokter *dataDokter, topik *dataTopik) {
	/*
		I.S Terdefinisi data patient, konsul, dokter, dan topik
		F.S Menampilkan menu pilihan untuk dokter, user dapat memilih tujuan sesuai dengan menu yang tersedia
	*/
	var option int
	fmt.Println("")
	fmt.Println("Selamat Datang Dr.", dokter.infoDokter[0].nama)
	fmt.Println("*=========================================*")
	fmt.Println("|        1. Tampilkan Trend Topik         |")
	fmt.Println("|        2. Kunjungi Forum                |")
	fmt.Println("|        3. Tanggapi Konsultasi           |")
	fmt.Println("|        0. Ganti Role                    |")
	fmt.Println("*=========================================*")
	fmt.Print("Masukkan pilihan anda: ")
	fmt.Scan(&option)
	for option != 1 && option != 0 && option != 2 && option != 3 {
		fmt.Println("Pilihan yang anda masukkan salah, Silahkan masukkan pilihan anda kembali")
		fmt.Println("Masukkan pilihan anda: ")
		fmt.Scan(&option)
	}
	if option == 0 {
		menu(patient, konsul, dokter, topik)
	} else if option == 1 {
		sortingDokter(patient, konsul, dokter, topik)
	} else if option == 2 {
		dokter.reply = true
		tampilanKonsul(patient, konsul, dokter, topik)
	} else if option == 3 {
		dokter.reply = true
		replyKonsultasiPasien(patient, konsul, dokter, topik)
	}

}

func sortingDokter(patient *dataPasien, konsul *dataKonsul, dokter *dataDokter, topik *dataTopik) {
	/*
		I.S Terdefinisi data patient, konsul, dokter, dan topik
		F.S Menampilkan Daftar topik konsultasi Pasien yang sedang sering diperbincangkan, mengembalikan user ke menu dokter
	*/
	var i, j, n, idx_min int
	var x int
	var pil int
	var t dataTopik

	fmt.Println("*=========================================*")
	fmt.Println("|       BERIKUT TREND TOPIK PASIEN        |")
	fmt.Println("*=========================================*")

	i = 1
	//Pengurutan data menggunakan selection sort
	for i <= 5 {
		idx_min = i - 1
		j = i
		for j < 5 {
			if konsul.ntopik[idx_min].ntopik < konsul.ntopik[j].ntopik {
				idx_min = j
			}
			j = j + 1
		}
		t = konsul.ntopik[idx_min]
		konsul.ntopik[idx_min] = konsul.ntopik[i-1]
		konsul.ntopik[i-1] = t
		i = i + 1
	}
	//menampilkan topik konsultasi dari topik dengan jumlah konsultasi terbanyak
	for n < 5 {
		fmt.Println("")
		fmt.Printf("%v. Pertanyaan Mengenai %v ditanya sebanyak = %v", x+1, konsul.ntopik[n].topik, konsul.ntopik[n].ntopik)
		fmt.Println("")
		n++
		x++
	}
	fmt.Println("")

	fmt.Print("Kembali? (0 untuk kembali) : ")
	fmt.Scan(&pil)
	for pil != 0 {
		fmt.Println("Pilihan yang anda masukkan salah, Silahkan masukkan pilihan anda kembali")
		fmt.Print("Masukkan pilihan anda: ")
		fmt.Scan(&pil)
	}
	if pil == 0 {
		homedokter(patient, konsul, dokter, topik)
	}
}

func login_dokter(patient *dataPasien, konsul *dataKonsul, dokter *dataDokter, topik *dataTopik) {
	/*
		I.S Terdefinisi data patient, konsul, dokter, dan topik
		F.S Dialihkan ke homedokter jika dokter.nama dan dokter.pass yang dimasukkan oleh user benar atau terdapat dalam ARRAY infoDokter
	*/
	var success bool
	var idxdoctor int
	fmt.Println("")
	fmt.Println("*==================Login==================*")
	fmt.Println("Input data anda dibawah ini")
	fmt.Print("username :")
	fmt.Scan(&dokter.nama)
	fmt.Print("Password :")
	fmt.Scan(&dokter.pass)

	//mengecek apakah dokter.nama dan dokter.pass yang dimasukkan terdapat dalam infoDokter
	for !success {
		for idxdoctor = 0; idxdoctor <= dokter.n; idxdoctor++ {
			if dokter.infoDokter[idxdoctor].nama == dokter.nama && dokter.infoDokter[idxdoctor].password == dokter.pass {
				success = true
			}
		}
		//jika tidak maka akan terus mengambil input
		if !success {
			fmt.Printf("Mohon maaf! username atau password yang anda masukkan salah \n")
			fmt.Println("Silahkan coba lagi")
			fmt.Print("username :")
			fmt.Scan(&dokter.nama)
			fmt.Print("Password :")
			fmt.Scan(&dokter.pass)
		}
	}
	homedokter(patient, konsul, dokter, topik)
}

func dataset(patient *dataPasien, konsul *dataKonsul, dokter *dataDokter, topik *dataTopik) {
	/*
		I.S Terdefinisi data patient, konsul, dokter, dan topik
		F.S data patient, konsul, dokter, dan topik terisi sebagai data awal aplikasi
	*/

	//melakukan pengisian 5 data topik
	konsul.ntopik[0].topik = "Kesehatan Mental"
	konsul.ntopik[1].topik = "Penyakit Umum"
	konsul.ntopik[2].topik = "Kandungan"
	konsul.ntopik[3].topik = "Spesialis Gigi"
	konsul.ntopik[4].topik = "Spesialis THT"

	//melakukan pengisian 10 data akun pasien dan konsultasi pasien tersebut
	patient.infoPasien[0].nama = "Andito"
	patient.infoPasien[0].password = "dito"
	konsul.infoKonsul[0].topik[0] = "Kesehatan Mental"
	konsul.infoKonsul[0].pertanyaan[0] = " Akhir Akhir ini saya mengalami gangguan tidur dikarenakan trauma di masalalu."
	konsul.infoKonsul[0].ntopik = 1
	konsul.infoKonsul[0].nPertanyaan = 1

	patient.infoPasien[1].nama = "Hilal"
	patient.infoPasien[1].password = "hilal"
	konsul.infoKonsul[1].topik[0] = "Penyakit Umum"
	konsul.ntopik[1].ntopik = 1
	konsul.infoKonsul[1].pertanyaan[0] = " Dokter, saya mengalami demam tinggi, batuk kering, dan kesulitan bernapas, apakah saya mungkin terinfeksi COVID-19?"
	konsul.infoKonsul[1].ntopik = 1
	konsul.infoKonsul[1].nPertanyaan = 1

	patient.infoPasien[2].nama = "Andi"
	patient.infoPasien[2].password = "andi"
	konsul.infoKonsul[2].topik[0] = "Kesehatan Mental"
	konsul.ntopik[0].ntopik = 2
	konsul.infoKonsul[2].pertanyaan[0] = " Dokter, saya merasa sangat sulit berkonsentrasi dan seringkali merasa gelisah dan cemas tanpa alasan yang jelas."
	konsul.infoKonsul[2].ntopik = 1
	konsul.infoKonsul[2].nPertanyaan = 1

	patient.infoPasien[3].nama = "Ajeng"
	patient.infoPasien[3].password = "ajeng"
	konsul.infoKonsul[3].topik[0] = "Kandungan"
	konsul.ntopik[2].ntopik = 1
	konsul.infoKonsul[3].pertanyaan[0] = " Dokter, saya sedang hamil, kira-kira apa ya dok makanan/minuman yang harus saya hindari?"
	konsul.infoKonsul[3].ntopik = 1
	konsul.infoKonsul[3].nPertanyaan = 1

	patient.infoPasien[4].nama = "Icha"
	patient.infoPasien[4].password = "icha"
	konsul.infoKonsul[4].topik[0] = "Kandungan"
	konsul.ntopik[2].ntopik = 2
	konsul.infoKonsul[4].pertanyaan[0] = " Dokter, apakah mengkonsumsi suplemen vitamin tertentu aman bagi kandungan saya, mengingat saya sedang dalam masa kehamilan."
	konsul.infoKonsul[4].ntopik = 1
	konsul.infoKonsul[4].nPertanyaan = 1

	patient.infoPasien[5].nama = "Asep"
	patient.infoPasien[5].password = "asep"
	konsul.infoKonsul[5].topik[0] = "Spesialis Gigi"
	konsul.ntopik[3].ntopik = 1
	konsul.topik[4] = "Spesialis THT"
	konsul.infoKonsul[5].pertanyaan[0] = " Dokter, gigi belakang saya terasa sakit dan terkadang berdarah saat mengunyah, apakah ada masalah yang perlu diperiksa?"
	konsul.infoKonsul[5].ntopik = 1
	konsul.infoKonsul[5].nPertanyaan = 1

	patient.infoPasien[6].nama = "Agung"
	patient.infoPasien[6].password = "agung"
	konsul.infoKonsul[6].topik[0] = "Spesialis Gigi"
	konsul.ntopik[3].ntopik = 2
	konsul.infoKonsul[6].pertanyaan[0] = " Dokter, gigi depan saya ada perubahan warna yang sensitif, apakah ada perawatan yang dapat membantu memperbaikinya?"
	konsul.infoKonsul[6].ntopik = 1
	konsul.infoKonsul[6].nPertanyaan = 1

	patient.infoPasien[7].nama = "Joni"
	patient.infoPasien[7].password = "joni"
	konsul.infoKonsul[7].topik[0] = "Spesialis THT"
	konsul.ntopik[4].ntopik = 1
	konsul.infoKonsul[7].pertanyaan[0] = " Dokter, saya sering mengalami hidung tersumbat, pilek, dan batuk pada malam hari, apakah ada kaitannya dengan THT?"
	konsul.infoKonsul[7].ntopik = 1
	konsul.infoKonsul[7].nPertanyaan = 1

	patient.infoPasien[8].nama = "Doni"
	patient.infoPasien[8].password = "doni"
	konsul.infoKonsul[8].topik[0] = "Spesialis THT"
	konsul.ntopik[4].ntopik = 2
	konsul.infoKonsul[8].pertanyaan[0] = " Dokter, saya mengalami tinnitus atau denging di telinga yang konstan, apakah ada pengobatan yang efektif untuk kondisi ini?"
	konsul.infoKonsul[8].ntopik = 1
	konsul.infoKonsul[8].nPertanyaan = 1

	patient.infoPasien[9].nama = "Restu"
	patient.infoPasien[9].password = "restu"
	konsul.infoKonsul[9].topik[0] = "Spesialis THT"
	konsul.ntopik[4].ntopik = 3
	konsul.infoKonsul[9].pertanyaan[0] = " Dokter, saya mengalami kehilangan pendengaran diwaktu waktu tertentu, apakah ini perlu ditangani secepatnya di bagian THT?"
	konsul.infoKonsul[9].ntopik = 1
	konsul.infoKonsul[9].nPertanyaan = 1

	//pengisian 1 data dokter yang digunakan pada aplikasi ini
	dokter.infoDokter[0].nama = "Nurul"
	dokter.infoDokter[0].password = "Nurul"

	//jumlah data dokter
	dokter.n = 1
	//jumlah data konsul
	konsul.n = 10
	//jumlah data pasien
	patient.n = 10
}
func main() {
	/* Asumsi Dokter hanya 1 untuk aplikasi ini, tidak dapat menambah dokter jika tidak terdaftar
	username dokter = Nurul
	pass dokter = Nurul */
	var patient dataPasien
	var konsul dataKonsul
	var topik dataTopik
	var dokter dataDokter
	dataset(&patient, &konsul, &dokter, &topik)
	menu(&patient, &konsul, &dokter, &topik)
}
