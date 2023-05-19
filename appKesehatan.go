package main

const NMAX int = 1000

type User struct {
	nama, username, password string
}

var pasien [NMAX]User
var dokter [NMAX]User

type Pertanyaan struct {
	id      int
	tag     string
	kontent string
}

type Tanggapan struct {
	konten  string
	penulis User
}

// menu
func menu() {

}

// fungsi registrasi pasien

// fungsi login

// fungsi registrasi dokter

// fungsi login

// pasien posting pertanyaan

//

func main() {

}
