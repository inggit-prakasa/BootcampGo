package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

var baseURL = "http://localhost:8080"

type karcis2 struct {
	Id string
	Time time.Time
}

type karcis4 struct {
	Id string
	Plat string
	Tipe string
	Hasil int
}

func main() {
	var index int

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
			kend,err := kendaraan1()
			if err != nil {
				fmt.Println("Error!", err.Error())
				return
			}

			fmt.Println("Id Anda : ",kend.Id)
		case 2:
			var tipeKendaraan,idAfter,platNo string
			fmt.Print("ID : ")
			fmt.Scan(&idAfter)
			fmt.Print("Tipe Kendaraan : ")
			fmt.Scan(&tipeKendaraan)
			fmt.Print("Plat Nomor : ")
			fmt.Scan(&platNo)
			gms,err := membayar(idAfter,tipeKendaraan,platNo)
			if err != nil {
				fmt.Println("Error!", err.Error())
				return
			}
			fmt.Println("Anda Membayar sebanyak : ",gms.Hasil)
		default:
			flag = false
			break;
		}
	}
}

func kendaraan1() (karcis2, error) {
	var err error
	var client = &http.Client{}
	var data karcis2

	request, err := http.NewRequest("POST", baseURL+"/masuk", nil)
	if err != nil {
		return data, err
	}

	response, err := client.Do(request)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)

	if err != nil {
		return data, err
	}

	return data,nil
}

func membayar(ID string,tipe string, plat string) (karcis4, error) {
	var err error
	var client = &http.Client{}
	var pesan karcis4

	var param = url.Values{}
	param.Set("id", ID)
	param.Set("tipe", tipe)
	param.Set("plat", plat)
	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", baseURL+"/keluar", payload)

	if err != nil {
		return pesan, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		return pesan, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&pesan)

	if err != nil {
		return pesan, err
	}
	return pesan, nil
}