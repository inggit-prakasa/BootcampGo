package main

import (
	"fmt"
	"github.com/rs/xid"
	"time"
)

type M map[string]interface {}

type Karcis struct {
	id string
	plat string
	time time.Time
	tipe string
}

func main() {
	var index int
	//var myMapSlice []M
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
			id,time,tipe,plat := generateKarcis()
			kas := Karcis{id,plat,time,tipe}

			kendaraan = append(kendaraan,kas)
			for i:=0; i<len(kendaraan); i++ {
				fmt.Println(kendaraan[i].id)
				fmt.Println(kendaraan[i].time.Weekday())
				fmt.Println(kendaraan[i].plat)
				fmt.Println(kendaraan[i].tipe)
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
			for i:=0; i<len(kendaraan); i++ {
				if(idAfter== kendaraan[i].id && tipeKendaraan==kendaraan[i].tipe && platNomor==kendaraan[i].plat ) {
					waktu := now.Sub(kendaraan[i].time).Seconds()
					fmt.Println(waktu)
					if kendaraan[i].tipe == "Motor" {
						if waktu > 1 {
							fmt.Print("Bayar parkir sebanyak : ")
							fmt.Println("Rp.",int(waktu*3000+5000))
						} else {
							fmt.Print("Bayar parkir sebanyak")
							fmt.Println("Rp.",5000)
						}
					} else if kendaraan[i].tipe == "Mobil" {
						if waktu > 1 {
							fmt.Print("Bayar parkir sebanyak : ")
							fmt.Println("Rp.",int(waktu*2000+3000))
						} else {
							fmt.Print("Bayar parkir sebanyak")
							fmt.Println("Rp.",3000)
						}
					}
					kendaraan = append(kendaraan[:i], kendaraan[i+1:]...)
				}

			}

		default:
			flag = false
			break;
		}
	}
}

func parkirMasuk() {

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
