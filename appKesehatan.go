package main

import (
	"fmt"
	"strings"
)

const NMAX int = 1000

type User struct {
	nama, username, password string
}

type UserType struct {
	Pasien               [NMAX]User
	Dokter               [NMAX]User
	pasienLen, dokterLen int
}

type UserData struct {
	isDokter bool
	id       int
}

type Tanggapan struct {
	author UserData
	konten string
}

type Pertanyaan struct {
	author       UserData
	id           int
	tag          [5]string //max 5 tag
	konten       string
	tabTanggapan [NMAX]Tanggapan
	tanggapanLen int
}

type Forum struct {
	tabPertanyaan [NMAX]Pertanyaan
	pertanyaanLen int
}

func guestMenu(users UserType) {
	opsiMenu := func() {
		fmt.Println("\n=== Aplikasi Konsultasi Kesehatan ===")
		fmt.Println("1. Daftar")
		fmt.Println("2. Masuk")
		fmt.Println("3. Lihat Forum")
		fmt.Println("00. Keluar")
	}

	opsiMenu()

	for {
		var opsi int
		fmt.Print("\nPilihan Anda: ")
		fmt.Scan(&opsi)

		if opsi == 1 {
			registerUser(&users)
			opsiMenu()
		} else if opsi == 2 {
			userData := loginUser(users)

			if userData.isDokter {
				dokterMenu(users, userData)
			} else {
				pasienMenu(users, userData)
			}
		} else if opsi == 3 {
			lihatForum()
		} else if opsi == 00 {
			fmt.Println("Terima kasih! Sampai jumpa lagi :)")
			return
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func maxLen(users UserType) int {
	if users.dokterLen > users.pasienLen {
		return users.dokterLen
	} else {
		return users.pasienLen
	}
}

func registerUser(users *UserType) {
	var nama, username, password string
	var isDokter string
	var n int = maxLen(*users)
	var hasUsername bool = false

	inputUser := func() {
		fmt.Print("Masukkan nama: ")
		fmt.Scan(&nama)
		fmt.Print("Masukkan username: ")
		fmt.Scan(&username)
		fmt.Print("Masukkan password: ")
		fmt.Scan(&password)
		fmt.Print("Apakah Anda seorang dokter? (y/n): ")
		fmt.Scan(&isDokter)
	}

	inputUser()

	i := 0
	for i < n {
		if users.Dokter[i].username == username || users.Pasien[i].username == username {
			fmt.Printf("\nUsername %s telah terdaftar. Ulangi proses pendaftaran!\n", username)
			hasUsername = true
			i = 0
			inputUser()
		}
		i++
	}

	if !hasUsername {
		if strings.ToLower(isDokter) == "y" {
			n = users.dokterLen

			users.Dokter[n].nama = nama
			users.Dokter[n].username = username
			users.Dokter[n].password = password
			users.dokterLen++
		} else {
			n = users.pasienLen

			users.Pasien[n].nama = nama
			users.Pasien[n].username = username
			users.Pasien[n].password = password
			users.pasienLen++
		}
		fmt.Println("\nPendaftaran berhasil!")
	}
}

func loginUser(users UserType) UserData {
	var username, password string
	var n int = maxLen(users)
	var found int = 0
	var result UserData

	inputUser := func() {
		fmt.Print("Masukkan username: ")
		fmt.Scan(&username)
		fmt.Print("Masukkan password: ")
		fmt.Scan(&password)
	}

	inputUser()

	for found == 0 {
		for i := 0; i < n && found == 0; i++ {
			if (users.Dokter[i].username == username) && (users.Dokter[i].password == password) {
				result.isDokter = true
				result.id = i
				found++
			}
			if (users.Pasien[i].username == username) && (users.Pasien[i].password == password) {
				result.isDokter = false
				result.id = i
				found++
			}
		}

		if found == 0 {
			fmt.Println("Username atau password tidak valid")
			inputUser()
		}
	}

	return result
}

func lihatForum() {
	// jumlah balasan
	// lihat isi pertanyaan berdasarkan id
	// tampilkan tipe penjawab (dokter/pasien)
}

// func cariTag() {

// }

// func lihatTagAtas() {

// }

// func postPertanyaan() {

// }

// func postJawaban() {

// }

func pasienMenu(users UserType, data UserData) {
	var id int = data.id

	opsiMenu := func() {
		fmt.Println("\n=== Aplikasi Konsultasi Kesehatan ===")
		fmt.Printf("Halo, %s\n", users.Pasien[id].nama)
		fmt.Println("1. Ajukan Pertanyaan")
		fmt.Println("2. Lihat Forum")
		fmt.Println("00. Keluar")
	}

	opsiMenu()

	for {
		var opsi int
		fmt.Print("\nPilihan Anda: ")
		fmt.Scan(&opsi)

		if opsi == 1 {
			// postPertanyaan()
		} else if opsi == 2 {
			// lihatForum()
		} else if opsi == 00 {
			fmt.Println("Terima kasih! Sampai jumpa lagi :)")
			guestMenu(users)
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func dokterMenu(users UserType, data UserData) {
	var id int = data.id

	opsiMenu := func() {
		fmt.Println("\n=== Aplikasi Konsultasi Kesehatan ===")
		fmt.Printf("Halo, %s\n", users.Dokter[id].nama)
		// fmt.Printf("Notifikasi: %d pertanyaan belum dijawab\n", )
		fmt.Println("1. Lihat Topik Populer")
		fmt.Println("2. Lihat Forum")
		fmt.Println("00. Keluar")
	}

	opsiMenu()

	for {
		var opsi int
		fmt.Print("\nPilihan Anda: ")
		fmt.Scan(&opsi)

		if opsi == 1 {
			// lihatTagAtas()
		} else if opsi == 2 {
			// lihatForum()
		} else if opsi == 00 {
			fmt.Println("Terima kasih! Sampai jumpa lagi :)")
			guestMenu(users)
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// func debugUser(users UserType) {
// 	fmt.Println("Dokter list")
// 	for i := 0; i < users.dokterLen; i++ {
// 		fmt.Printf("Nama: %s \tUsername: %s \tPass: %s\n", users.Dokter[i].nama, users.Dokter[i].username, users.Dokter[i].password)
// 	}
// 	fmt.Println(users.dokterLen)

// 	fmt.Println("Pasien list")
// 	for j := 0; j < users.pasienLen; j++ {
// 		fmt.Printf("Nama: %s \tUsername: %s \tPass: %s\n", users.Pasien[j].nama, users.Pasien[j].username, users.Pasien[j].password)
// 	}
// 	fmt.Println(users.pasienLen)
// }

func main() {
	var users UserType
	guestMenu(users)
}
