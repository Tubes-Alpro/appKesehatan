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

type Pertanyaan struct {
	author User
	id     int
	tag    [5]string //max 5 tag
	konten string
}

type Tanggapan struct {
	author User
	konten string
}

type Forum struct {
	tabPertanyaan [NMAX]Pertanyaan
}

func guestMenu(users UserType) {
	opsiMenu := func() {
		fmt.Println("\n=== Aplikasi Konsultasi Kesehatan ===")
		fmt.Println("1. Daftar")
		fmt.Println("2. Masuk")
		fmt.Println("3. Keluar")
		fmt.Println("4. Lihat Forum")
		fmt.Println("00. debug")
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
			loginUser(users)
		} else if opsi == 3 {
			fmt.Println("Terima kasih! Sampai jumpa lagi :)")
			return
		} else if opsi == 4 {
			// tampilkan forum
			// jumlah balasan
			// lihat isi pertanyaan berdasarkan id
			// tampilkan tipe penjawab (dokter/pasien)
		} else if opsi == 00 {
			debugUser(users)
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
			inputUser()
		}
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

func loginUser(users UserType) {
	var username, password string
	var n int = maxLen(users)
	var found int = 0
	var isDokter bool

	inputUser := func() {
		fmt.Print("Masukkan username: ")
		fmt.Scan(&username)
		fmt.Print("Masukkan password: ")
		fmt.Scan(&password)
	}

	inputUser()

	i := 0
	for i < n && found > 0 {
		if users.Dokter[i].username == username && users.Dokter[i].password == password {
			isDokter = true
			found++
		}
		if users.Pasien[i].username == username && users.Pasien[i].password == password {
			isDokter = false
			found++
		}
		if found == 0 {
			fmt.Println("Username atau password tidak valid")
			inputUser()
		}
	}
	if isDokter == true {
		fmt.Println("login dokter")
	} else {
		fmt.Println("login pasien")
	}
}

// func lihatForum() {

// }

// func cariTag() {

// }

// func lihatTagAtas() {

// }

// func postPertanyaan() {

// }

// func postJawaban() {

// }

// func pasienMenu() {

// }

// func dokterMenu() {

// }

func debugUser(users UserType) {
	fmt.Println("Dokter list")
	for i := 0; i < users.dokterLen; i++ {
		fmt.Printf("Nama: %s \tUsername: %s \tPass: %s\n", users.Dokter[i].nama, users.Dokter[i].username, users.Dokter[i].password)
	}
	fmt.Println(users.dokterLen)

	fmt.Println("Pasien list")
	for j := 0; j < users.pasienLen; j++ {
		fmt.Printf("Nama: %s \tUsername: %s \tPass: %s\n", users.Pasien[j].nama, users.Pasien[j].username, users.Pasien[j].password)
	}
	fmt.Println(users.pasienLen)
}

func main() {
	var users UserType
	guestMenu(users)
}
