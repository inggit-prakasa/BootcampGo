package main

import (
	"fmt"
	"github.com/rs/xid"
	"time"
)

type Karcis struct {
	id string
	time time.Time
}

func main() {
	var index int
	var kendaraan []Karcis

	flag := true
	for flag {
		Smenu := "------- Sistem Parkir --------- \n" +
			"1. Parkir Masuk \n" +
			"2. Parkir Keluar \n" +
			"Pilihan : "
		fmt.Print(Smenu)
		fmt.Scan(&index)
		switch index {
		case 1:
			id,time := generateKarcis2()
			kas := Karcis{id,time}

			kendaraan = append(kendaraan,kas)

			for i:=0; i<len(kendaraan); i++ {
				fmt.Println(kendaraan[i].id)
			}

		case 2:
			var (
				idAfter string
				platNomor string
				tipeKendaraan string
			)
			now := time.Now()
			fmt.Print("Tipe Kendaraan : ")
			fmt.Scan(&tipeKendaraan)
			fmt.Print("Plat Nomor : ")
			fmt.Scan(&platNomor)
			fmt.Print("ID Parkir : ")
			fmt.Scan(&idAfter)

			if len(kendaraan)<1 {
				fmt.Println("Belum ada kendaraan yang masuk")
			}

			counter := false
			for i:=0; i<len(kendaraan); i++ {
				if(idAfter== kendaraan[i].id ) {
					waktu := int(now.Sub(kendaraan[i].time).Seconds())
					fmt.Println("Waktu parkir anda :",waktu,"Detik")
					if tipeKendaraan == "Mobil" {
						if waktu > 1 {
							fmt.Print("Bayar parkir sebanyak : ")
							fmt.Println("Rp.",int((waktu-1)*3000+5000))
						} else {
							fmt.Print("Bayar parkir sebanyak")
							fmt.Println("Rp.",5000)
						}
					} else if tipeKendaraan == "Motor" {
						if waktu > 1 {
							fmt.Print("Bayar parkir sebanyak : ")
							fmt.Println("Rp.",int((waktu-1)*2000+3000))
						} else {
							fmt.Print("Bayar parkir sebanyak : ")
							fmt.Println("Rp.",3000)
						}
					}
					kendaraan = append(kendaraan[:i], kendaraan[i+1:]...)
				} else {
					counter = true;
				}
			}

			if counter {
				fmt.Println("ID Parkir Salah")
			}

		default:
			flag = false
			break;
		}
	}
}

func generateKarcis() (string,time.Time,string,string){
	id := xid.New().String()
	time := time.Now()
	var (
		tipe,plat string
	)
	fmt.Print("Tipe : ")
	fmt.Scanf("%s", &tipe)
	fmt.Print("Plat Nomor : ")
	fmt.Scanf("%s", &plat)

	return id,time,tipe,plat
}

func generateKarcis2() (string,time.Time){
	id := xid.New().String()
	time := time.Now()

	return id,time
}
