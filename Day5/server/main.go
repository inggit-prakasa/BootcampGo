
package main

import (
	"context"
	"github.com/rs/xid"
	"log"
	"net"
	"time"

	parkir "github.com/inggit_prakasa/BootcampGo/Day5/parkir"
	"google.golang.org/grpc"
)

type karcis struct {
	Id string
	Time time.Time
}

var kendaraan = []karcis {}

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	parkir.UnimplementedParkirServer
}

func (s *server) MasukParkir(ctx context.Context, in *parkir.Empty) (*parkir.IdParkir, error) {
	id := xid.New().String()
	time := time.Now()

	kendaraan = append(kendaraan,karcis{id,time})

	return &parkir.IdParkir{Id: id,Time: time.Format("2006-01-02 15:04:05")},nil
}

func (s *server) KeluarParkir(ctx context.Context, in *parkir.Kendaraan) (*parkir.BayarParkir, error) {

	var (
		Tmasuk time.Time
		harga,hargaAwal,hargaMaks int64
	)

	if len(kendaraan)<1 {
		return &parkir.BayarParkir{
			Id: "0",
			Price:   0,
			Plat:    "0",
			Tipe:    "0",
			TMasuk:  "0",
			TKeluar: "0",
		},nil
	}

	Tkeluar := time.Now()
	flag := true
	for i:=0; i<len(kendaraan); i++ {
		if (kendaraan[i].Id == in.Id) {
			flag = false
			Tmasuk = kendaraan[i].Time
			kendaraan = append(kendaraan[:i], kendaraan[i+1:]...)
			break
		}
	}

	if flag {
		return &parkir.BayarParkir{
			Id: "0",
			Price:   0,
			Plat:    "0",
			Tipe:    "0",
			TMasuk:  "0",
			TKeluar: "0",
		},nil
	}

	switch in.Tipe {
		case "MOBIL" :
			harga = 3000
			hargaAwal = 5000
			hargaMaks = 500000
		case "MOTOR" :
			harga = 2000
			hargaAwal = 3000
			hargaMaks = 500000
	}

	Tdurasi := int64(Tkeluar.Sub(Tmasuk).Seconds())
	hargaTotal := hargaAwal + (Tdurasi-1) * harga

	if hargaTotal > hargaMaks {
		return &parkir.BayarParkir{
			Price:   hargaMaks,
			Plat:    in.Plat,
			Tipe:    in.Tipe,
			TMasuk:  Tmasuk.Format("2006-01-02 15:04:05"),
			TKeluar: Tkeluar.Format("2006-01-02 15:04:05"),
		},nil
	}

	return &parkir.BayarParkir{
		Id: in.Id,
		Price:   hargaTotal,
		Plat:    in.Plat,
		Tipe:    in.Tipe,
		TMasuk:  Tmasuk.Format("2006-01-02 15:04:05"),
		TKeluar: Tkeluar.Format("2006-01-02 15:04:05"),
	},nil

}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	parkir.RegisterParkirServer(s,&server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
