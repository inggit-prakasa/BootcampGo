package main

import (
"encoding/json"
"fmt"
"github.com/rs/xid"
"net/http"
"time"
)

type karcis struct {
	Id string
	Time time.Time
}

type price struct {
	rp int
}

type karcis3 struct {
	Id string
	Plat string
	Tipe string
	Hasil int
}

var kendaraan = []karcis {}

func main() {
	http.HandleFunc("/masuk", generateID)
	http.HandleFunc("/keluar", bayarParkir)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func generateID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id,time := generateKarcis2()
	kas := karcis{id,time}

	kendaraan = append(kendaraan,kas)

	if r.Method == "POST" {
		var result, err = json.Marshal(kas)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

func bayarParkir(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		now := time.Now()
		var id = r.FormValue("id")
		var tipeKendaraan = r.FormValue("tipe")
		var platNomor = r.FormValue("plat")
		var result []byte
		var pengunjung karcis3
		var err error

		for _, each := range kendaraan {
			if each.Id == id {
				waktu := int(now.Sub(each.Time).Seconds())
				fmt.Println("Waktu parkir anda :",waktu,"Detik")
				if tipeKendaraan == "Mobil" {
					if waktu > 1 {
						hasilbayar := (waktu-1)*3000+5000
						fmt.Println("Id : ",id)
						fmt.Println("tipe Kendaraan : ",tipeKendaraan)
						fmt.Println("plat nomor : ",platNomor)
						fmt.Println("Bayar : ",hasilbayar)
						pengunjung = karcis3{id,platNomor,tipeKendaraan,hasilbayar}
						result,err = json.Marshal(pengunjung)
						w.Write(result)
						return
					} else {
						hasilbayar := 5000
						fmt.Println("Id : ",id)
						fmt.Println("tipe Kendaraan : ",tipeKendaraan)
						fmt.Println("plat nomor : ",platNomor)
						fmt.Println("Bayar : ",hasilbayar)
						result,err = json.Marshal(hasilbayar)
						pengunjung = karcis3{id,platNomor,tipeKendaraan,hasilbayar}
						result,err = json.Marshal(pengunjung)
						w.Write(result)
						return

					}
				} else if tipeKendaraan == "Motor" {
					if waktu > 1 {
						hasilbayar := (waktu-1)*2000+3000
						fmt.Println("Id : ",id)
						fmt.Println("tipe Kendaraan : ",tipeKendaraan)
						fmt.Println("plat nomor : ",platNomor)
						fmt.Println("Bayar : ",hasilbayar)
						result,err = json.Marshal(hasilbayar)
						pengunjung = karcis3{id,platNomor,tipeKendaraan,hasilbayar}
						result,err = json.Marshal(pengunjung)
						w.Write(result)
						return

					} else {
						hasilbayar := 3000
						fmt.Println("Id : ",id)
						fmt.Println("tipe Kendaraan : ",tipeKendaraan)
						fmt.Println("plat nomor : ",platNomor)
						fmt.Println("Bayar : ",hasilbayar)
						result,err = json.Marshal(hasilbayar)
						pengunjung = karcis3{id,platNomor,tipeKendaraan,hasilbayar}
						result,err = json.Marshal(pengunjung)
						w.Write(result)
						return
					}
				}
				//kendaraan = append(kendaraan[:i], kendaraan[i+1:]...)

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				return
			}
		}

		http.Error(w, "Tidak ada kendaraan", http.StatusBadRequest)
		return
	}
	http.Error(w, "", http.StatusBadRequest)
}

func generateKarcis2() (string,time.Time){
	id := xid.New().String()
	time := time.Now()

	return id,time
}