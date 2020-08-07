package main

import (
	"fmt"
	"time"
)

const (
	layoutISO = "2006-01-02"
)

func main() {

	// Pemanggilan awal dari tahun dan tanggal
	date := "2020-01-01"
	time, _ := time.Parse(layoutISO, date)
	year := time.Year()

	// Inisiasi dari bulan dan tanggal
	month := []string{"", "Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"}

	days := []int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	// Penentuan hari kosong pada tahun pertama
	kosong := int(time.Weekday())

	for m := 1; m <= 12; m++ {

		// pada tahun kabisat
		if (year%4 == 0) && m == 2 {
			days[m] = 29
		}

		// Print awal pada bulan
		fmt.Println("====== " + month[m] + " ======")
		fmt.Println("  S   S   R   K   J   S   M")

		// Penentuan hari kosong
		kosong = (days[m-1] + kosong) % 7

		// tanggal terakhir sebelum
		for i := 1; i <= kosong; i++ {
			fmt.Print(" -- ")
		}

		// jumlah tanggal hari setelah hari terakhir
		kosong2 := (42 - (kosong + days[m])) % 7

		// Print
		for i := 1; i <= days[m]; i++ {

			fmt.Printf(" %2d ", i)
			if i == days[m] {
				for j := 0; j < kosong2; j++ {
					fmt.Print(" -- ")
				}
			}
			if ((i+kosong)%7 == 0) || (i == days[m]) {
				fmt.Println("")
			}
		}

		fmt.Println("")

	}

}
