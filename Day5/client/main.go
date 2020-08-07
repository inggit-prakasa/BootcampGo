package main

import (
	"context"
	"fmt"
	parkir "github.com/inggit_prakasa/BootcampGo/Day5/parkir"
	"google.golang.org/grpc"
	"log"
)

const (
	address     = "localhost:50051"
)

func main() {
	var index int

	flag := true
	for flag {
		Smenu := "------- Sistem Parkir --------- \n" +
			"1. Parkir Masuk \n" +
			"2. Parkir Keluar \n" +
			"3. Keluar\n" +
			"Pilihan : "
		fmt.Print(Smenu)
		fmt.Scan(&index)
		switch index {
		case 1:
			kendaraanMasuk()
		case 2:
			kendaraanKeluar()
		case 3:
			fmt.Println("Terima kasih telah parkir disini :)")
			flag = false
			break
		default:
			fmt.Println("Pilihan anda salah")
		}
	}
}

func kendaraanKeluar() {
	var tipeKendaraan,idAfter,platNo string
	fmt.Print("ID : ")
	fmt.Scan(&idAfter)
	fmt.Print("Tipe Kendaraan : ")
	fmt.Scan(&tipeKendaraan)
	fmt.Print("Plat Nomor : ")
	fmt.Scan(&platNo)


	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	car := parkir.NewParkirClient(conn)

	r, err := car.KeluarParkir(context.Background(), &parkir.Kendaraan{Id: idAfter,Tipe: tipeKendaraan,Plat: platNo})
	if err != nil {
		log.Fatalf("Tidak bisa parkir : %v", err)
	}

	if r.Id == "0" {
		fmt.Println("Tidak ada kendaraan")
	} else {
		fmt.Println("-------------")
		fmt.Printf("Id : %s\n",r.Id)
		fmt.Printf("Tipe : %s\n",r.Tipe)
		fmt.Printf("Plat : %s\n",r.Plat)
		fmt.Printf("Jam Masuk : %s\n",r.TMasuk)
		fmt.Printf("Jam Keluar : %s\n",r.TKeluar)
		fmt.Printf("Bayar : Rp.%d,00\n",r.Price)
	}


}

func kendaraanMasuk() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	car := parkir.NewParkirClient(conn)

	r, err := car.MasukParkir(context.Background(), &parkir.Empty{})
	if err != nil {
		log.Fatalf("Tidak bisa parkir : %v", err)
	}

	fmt.Println("------------------------")
	fmt.Printf("Id karcis : %s \n", r.Id)
	fmt.Printf("Waktu kedatangan : %s \n", r.Time)
}
