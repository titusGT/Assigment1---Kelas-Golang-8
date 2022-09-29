package main

import (
	databaseclass "Assignment1/DatabaseClass"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var ID1 databaseclass.Student
	ID1.Nama = "Nama : Titus\n"
	ID1.ID = 1
	ID1.Alamat = "Alamat : Demangan\n"
	ID1.Pekerjaan = "Pekerjaan : Tukang\n"
	ID1.AlasanKelasGolang = "Alasan Kelas Go : Suka"

	var ID2 databaseclass.Student
	ID2.Nama = "Nama : Gigih\n"
	ID2.Alamat = "Alamat : Regency\n"
	ID2.Pekerjaan = "Pekerjaan : Kuli\n"
	ID2.AlasanKelasGolang = "Alasan Kelas Go : Sangat suka"
	ID2.ID = 2

	var ID3 databaseclass.Student
	ID3.Nama = "Nama : Trionggo\n"
	ID3.Alamat = "Alamat : Demangan Regency\n"
	ID3.Pekerjaan = "Pekerjaan : Tukang Kuli\n"
	ID3.AlasanKelasGolang = "Alasan Kelas Go : Sangat suka sekali"
	ID3.ID = 3

	Num := os.Args
	IDD, err := strconv.ParseUint(Num[1], 10, 32)

	if err == nil {
		fmt.Println("Selamat datang di kelas golang 8")
	}
	if IDD == 1 {
		fmt.Println(ID1)
	}
	if IDD == 2 {
		fmt.Println(ID2)
	}
	if IDD == 3 {
		fmt.Println(ID3)
	}
	if IDD > 3 {
		fmt.Println("student dengan id", IDD, "tidak ada pada database")
	} else {
		fmt.Println(err.Error())
	}

}

func init() {

}
