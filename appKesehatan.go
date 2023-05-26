package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const NMAX int = 1000

type User struct {
	id                       int
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
	tag          string
	konten       string
	tabTanggapan [NMAX]Tanggapan
	tanggapanLen int
}

type Forum struct {
	tabPertanyaan [NMAX]Pertanyaan
	pertanyaanLen int
}

func guestMenu(users UserType, forums Forum) {
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
				dokterMenu(users, userData, forums)
			} else {
				pasienMenu(users, userData, forums)
			}
		} else if opsi == 3 {
			session := "guest"
			lihatForum(users, forums, session)
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
		fmt.Print("Apakah Anda seorang pasien? (y/n): ")
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
		if strings.ToLower(isDokter) == "n" {
			n = users.dokterLen

			users.Dokter[n].id = n
			users.Dokter[n].nama = nama
			users.Dokter[n].username = username
			users.Dokter[n].password = password
			users.dokterLen++
		} else {
			n = users.pasienLen

			users.Pasien[n].id = n
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
				result.id = users.Dokter[i].id
				found++
			}
			if (users.Pasien[i].username == username) && (users.Pasien[i].password == password) {
				result.isDokter = false
				result.id = users.Pasien[i].id
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

func lihatForum(users UserType, forums Forum, session string) {
	fmt.Println("\n=== Forum Konsultasi ===")

	for j := 0; j < forums.pertanyaanLen; j++ {
		pertanyaan := forums.tabPertanyaan[j]
		author := pertanyaan.author.id
		fmt.Printf("\nID: %d\t", pertanyaan.id)
		fmt.Printf("Oleh: %s\t", users.Pasien[author].nama)
		fmt.Printf("Tag: %s\n", pertanyaan.tag)
		fmt.Printf("Pertanyaan: %s\n", pertanyaan.konten)
		fmt.Printf("Tanggapan: %d\n", pertanyaan.tanggapanLen)
		for k := 0; k < pertanyaan.tanggapanLen; k++ {
			tanggapan := pertanyaan.tabTanggapan[k]
			if tanggapan.author.isDokter {
				fmt.Printf("- %s (dokter): %s\n", users.Dokter[tanggapan.author.id].nama, tanggapan.konten)
			} else {
				fmt.Printf("- %s (pasien): %s\n", users.Pasien[tanggapan.author.id].nama, tanggapan.konten)
			}
		}
	}

	fmt.Println("\n=== Menu ===")

	if session == "guest" {
		for {
			var opsi int
			fmt.Println("0. Kembali")
			fmt.Print("\nPilihan Anda: ")
			fmt.Scan(&opsi)

			if opsi == 0 {
				guestMenu(users, forums)
			} else {
				fmt.Println("Pilihan tidak valid.")
			}
		}
	} else if session == "pasien" {
		fmt.Println("1. Ajukan Pertanyaan")
		fmt.Println("2. Jawab Pertanyaan")
		fmt.Println("0. Kembali")
	} else if session == "dokter" {
		//
	}
}

func cariTag(users UserType, forums Forum) {
	var tag string
	fmt.Print("\nMasukkan tag yang ingin dicari: ")
	fmt.Scan(&tag)

	fmt.Println("\n=== Hasil Pencarian ===")
	found := false

	for j := 0; j < forums.pertanyaanLen; j++ {
		pertanyaan := forums.tabPertanyaan[j]
		if pertanyaan.tag == tag {
			if !found {
				found = true
			}
			author := pertanyaan.author.id
			fmt.Printf("\nID: %d\t", pertanyaan.id)
			fmt.Printf("Penanya: %s\n", users.Pasien[author].nama)
			fmt.Printf("Pertanyaan: %s\n", pertanyaan.konten)
			fmt.Printf("Tanggapan: %d\n", pertanyaan.tanggapanLen)
			for k := 0; k < pertanyaan.tanggapanLen; k++ {
				tanggapan := pertanyaan.tabTanggapan[k]
				if tanggapan.author.isDokter {
					fmt.Printf("- %s (dokter): %s\n", users.Dokter[tanggapan.author.id].nama, tanggapan.konten)
				} else {
					fmt.Printf("- %s (pasien): %s\n", users.Pasien[tanggapan.author.id].nama, tanggapan.konten)
				}
			}
		}
	}

	if !found {
		fmt.Println("Tidak ditemukan pertanyaan dengan tag tersebut.")
	}
}

func lihatTagAtas(users UserType, forums Forum, data UserData) {
	tags := make(map[string]int)

	for i := 0; i < forums.pertanyaanLen; i++ {
		pertanyaan := forums.tabPertanyaan[i]

		tag := pertanyaan.tag
		if tag != "" {
			tags[tag]++
		}
	}

	fmt.Println("\n=== Tag Populer ===")
	fmt.Println("Tag\t\tJumlah Pertanyaan")

	for tag, count := range tags {
		fmt.Printf("%s\t\t%d\n", tag, count)
	}

	fmt.Println("\n=== Menu ===")
	for {
		var opsi int
		fmt.Println("0. Kembali")
		fmt.Print("\nPilihan Anda: ")
		fmt.Scan(&opsi)

		if opsi == 0 {
			dokterMenu(users, data, forums)
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func postPertanyaan(users UserType, forums *Forum, data UserData) {
	var pertanyaan string
	var tags string
	var submit bool
	var input string

	reader := bufio.NewReader(os.Stdin)

	for !submit {
		fmt.Print("Masukkan pertanyaan Anda: ")
		pertanyaan, _ = reader.ReadString('\n')
		pertanyaan = strings.TrimSpace(pertanyaan)

		fmt.Println("Pilih tag:")
		tagsOpsi := [10]string{"kanker", "diabetes", "stroke", "jantung", "hiv", "asma", "tbc", "flu", "depresi", "lainnya"}
		for i := 0; i < 10; i++ {
			fmt.Printf("#%d: %s  |  ", i+1, tagsOpsi[i])
			if (i+1)%5 == 0 {
				fmt.Println()
			}
		}
		valid := false
		var opsi int
		for !valid {
			fmt.Print("\nPilih tag (1-10): ")
			fmt.Scanln(&opsi)
			if opsi < 1 || opsi > 10 {
				fmt.Println("Pilihan tag tidak valid")
			} else {
				valid = true
			}
		}
		tags = tagsOpsi[opsi-1]

		fmt.Print("Submit pertanyaan? (y/n):")
		fmt.Scan(&input)
		submit = (input == "y")
	}

	author := data.id
	id := forums.pertanyaanLen

	forums.tabPertanyaan[id].author.id = author
	forums.tabPertanyaan[id].id = id
	forums.tabPertanyaan[id].tag = tags
	forums.tabPertanyaan[id].konten = pertanyaan
	forums.pertanyaanLen++

	fmt.Println("Pertanyaan berhasil diposting!")
	lihatForum(users, *forums, "pasien")
}

func postJawaban(users *UserType, forums *Forum, data UserData, idPertanyaan int) {
	var jawaban string

	fmt.Print("Masukkan jawaban Anda: ")
	fmt.Scan(&jawaban)

	id := forums.tabPertanyaan[idPertanyaan].tanggapanLen
	author := data.id

	forums.tabPertanyaan[idPertanyaan].tabTanggapan[id].author.id = author
	forums.tabPertanyaan[idPertanyaan].tabTanggapan[id].konten = jawaban
	forums.tabPertanyaan[idPertanyaan].tanggapanLen++

	fmt.Println("Jawaban berhasil diposting!")
}

func pasienMenu(users UserType, data UserData, forums Forum) {
	var id int = data.id

	opsiMenu := func() {
		fmt.Println("\n=== Aplikasi Konsultasi Kesehatan ===")
		fmt.Printf("Halo, %s (pasien)\n", users.Pasien[id].nama)
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
			postPertanyaan(users, &forums, data)
		} else if opsi == 2 {
			session := "pasien"
			lihatForum(users, forums, session)
		} else if opsi == 00 {
			fmt.Println("Terima kasih! Sampai jumpa lagi :)")
			guestMenu(users, forums)
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func dokterMenu(users UserType, data UserData, forums Forum) {
	var id int = data.id

	opsiMenu := func() {
		fmt.Println("\n=== Aplikasi Konsultasi Kesehatan ===")
		fmt.Printf("Halo, %s (dokter)\n", users.Dokter[id].nama)
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
			lihatTagAtas(users, forums, data)
		} else if opsi == 2 {
			session := "dokter"
			lihatForum(users, forums, session)
		} else if opsi == 00 {
			fmt.Println("Terima kasih! Sampai jumpa lagi :)")
			guestMenu(users, forums)
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

func dummy(users *UserType, forums *Forum) {
	users.Pasien[0] = User{
		id:       0,
		nama:     "Jon",
		username: "jon123",
		password: "123",
	}
	users.pasienLen++

	users.Pasien[1] = User{
		id:       1,
		nama:     "Stefi",
		username: "stef1",
		password: "123",
	}
	users.pasienLen++

	users.Dokter[0] = User{
		id:       0,
		nama:     "Bob",
		username: "bob123",
		password: "123",
	}
	users.dokterLen++

	forums.tabPertanyaan[0] = Pertanyaan{
		author: UserData{
			id:       0,
			isDokter: false,
		},
		id:     0,
		tag:    "flu",
		konten: "Berapa lama indra penciuman hilang saat mengalami flu?",
	}
	forums.pertanyaanLen++

	forums.tabPertanyaan[1] = Pertanyaan{
		author: UserData{
			id:       1,
			isDokter: false,
		},
		id:     1,
		tag:    "kanker",
		konten: "What are the treatment options for lung cancer?",
	}
	forums.pertanyaanLen++

	forums.tabPertanyaan[2] = Pertanyaan{
		author: UserData{
			id:       0,
			isDokter: false,
		},
		id:     2,
		tag:    "flu",
		konten: "How can I manage my blood sugar levels effectively?",
	}
	forums.pertanyaanLen++

	forums.tabPertanyaan[0].tabTanggapan[0] = Tanggapan{
		author: UserData{
			id:       1,
			isDokter: false,
		},
		konten: "You're welcome! If you have any more questions, feel free to ask.",
	}
	forums.tabPertanyaan[0].tanggapanLen++

	forums.tabPertanyaan[1].tabTanggapan[0] = Tanggapan{
		author: UserData{
			id:       0,
			isDokter: true,
		},
		konten: "Thank you for your question. The treatment options for lung cancer include surgery, chemotherapy, radiation therapy, targeted therapy, and immunotherapy.",
	}
	forums.tabPertanyaan[1].tanggapanLen++
}

func main() {
	var users UserType
	var forums Forum

	dummy(&users, &forums)

	guestMenu(users, forums)
}
