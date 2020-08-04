package main

import (
	"fmt"
	"math"
)

type Persegi struct {
	sisi float64
}

type Segitiga struct {
	alas float32
	tinggi float32
}

type Hitung struct {
	angka1 int
	angka2 int
}

func main() {
	var index,x,y,z int

	flag := true
	for flag {
		fmt.Println("Pilihan :")
		fmt.Println("1. Pertambahan")
		fmt.Println("2. Pengurangan")
		fmt.Println("3. Perkalian")
		fmt.Println("4. Pembagian")
		fmt.Println("5. Akar")
		fmt.Println("6. Pangkat")
		fmt.Println("7. Luas Persegi")
		fmt.Println("8. Luas Lingkaran")
		fmt.Println("9. Volume tabung")
		fmt.Println("10. Volume Balok")
		fmt.Println("11. Volume Prisma")
		fmt.Print("Pilihan : ")
		fmt.Scan(&index)
		switch index {
		case 1:
			fmt.Println("Hasil Pertambahan:",tambah(input()))
		case 2:
			fmt.Scan(&x)
			fmt.Scan(&y)
			kurangi := Hitung{x,y}
			fmt.Println("Hasil Pengurangan:", kurangi.kurang())
		case 3:
			fmt.Scan(&x)
			fmt.Scan(&y)
			dibagi := Hitung{x,y}
			fmt.Println("Hasil Bagi:",dibagi.bagi())
		case 4:
			fmt.Scan(&x)
			fmt.Scan(&y)
			dikali := Hitung{x,y}
			fmt.Println("Hasil Bagi:", dikali.kali())
		case 5:
			fmt.Scan(&x)
			fmt.Println("Hasil Akar :",akar(float64(x)))
		case 6:
			fmt.Scan(&x)
			fmt.Scan(&y)
			pangkat1 := Hitung{x,y}
			fmt.Println("Hasil Pangkat :", pangkat1.pangkat())
		case 7:
			fmt.Scan(&x)
			persegi := Persegi{float64(x)}
			fmt.Println("Hasil Luas Persegi :",persegi.luaspersegi())
		case 8:
			fmt.Scan(&x)
			fmt.Println("Hasil Luas lingkaran :",luasLingkaran(float64(x)))
		case 9:
			fmt.Scan(&x)
			fmt.Scan(&y)
			fmt.Println("Hasil volume tabung :",volumeTabung(float64(x), float64(y)))
		case 10:
			fmt.Scan(&x)
			fmt.Scan(&y)
			fmt.Scan(&z)
			fmt.Println("Hasil volume Balok :",volumeBalok(float64(x), float64(y), float64(z)))
		case 11:
			fmt.Scan(&x)
			fmt.Scan(&y)
			fmt.Scan(&z)
			fmt.Println("Hasil volume Prisma :", volumePrisma(float64(x), float64(y), float64(z)))
		default:
			flag = false
			break
		}

	}
}

func input() (int,int) {
	var x,y int
	fmt.Scan(&x)
	fmt.Scan(&y)
	return x,y
}

func change(original *int, value int) {
	*original = value
}

func tambah(x int, y int) int {
	return x + y;
}

func (h Hitung) kurang() int {
	return h.angka1 - h.angka2;
}

func (h Hitung) bagi() int {
	return h.angka1/h.angka2;
}

func (h Hitung) kali() int {
	return h.angka1*h.angka2;
}

func akar(x float64) float64 {
	return math.Sqrt(x);
}

func (x Hitung) pangkat() float64 {
	return math.Pow(float64(x.angka1), float64(x.angka2))
}

func (x Persegi) luaspersegi() float64 {
	return math.Pow(x.sisi,2)
}

func luasLingkaran(x float64) float64 {
	return x*x*math.Pi;
}

func volumeTabung(r float64, t float64) float64 {
	return math.Pi*r*r*t;
}

func volumeBalok(p float64, l float64, t float64) float64{
	return p*l*t;
}

func volumePrisma(a float64,t float64, tB float64) float64 {
	return ((a*t)/2)*tB;
}