package main

import "fmt"

func main() {

	// for i := 0; i < 5; i++ {
	// 	fmt.Print("Angka", i)
	// }

	// for dengan keadaan
	// i := 0
	// for i < 3 {
	// 	fmt.Print("Kondisi", i)
	// 	i++
	// }

	//for true
	// j := 0
	// for{
	// 	fmt.Print("Tanpa Argumen", j)
	// 	j++
	// 	if j == 3 {
	// 		break
	// 	}
	// }

	// for i:=1; i <= 10; i++ {
	// 	if i % 2 == 1 {
	// 		continue
	// 	}
	// 	if i >8{
	// 		break
	// 	}
	// 	fmt.Println("Angka", i)
	// }

	// tanpa label, ini tidak menghentikan perulangan luar
	// for kelas := 1; kelas <= 3; kelas++ {
	// 	for siswa := 1; siswa <= 3; siswa++ {
	// 		if kelas == 2 && siswa == 2 {
	// 			fmt.Println("Siswa ketemu! Berhenti.")
	// 			break
	// 		}
	// 		fmt.Println("Kelas", kelas, "Siswa nomor", siswa)
	// 	}
	// }

	//dengan label= menghentikan perulangan luar
	// luar:	// ditaruh di luar for, penamaannya bebas
	// for kelas := 1; kelas <= 3; kelas++ {
	// 	for siswa := 1; siswa <= 3; siswa++ {
	// 		if kelas == 2 && siswa == 2 {
	// 			fmt.Println("Siswa ketemu! Berhenti.")
	// 			break luar //pelengkapnya diatruh disini
	// 		}
	// 		fmt.Println("Kelas", kelas, "Siswa nomor", siswa)
	// 	}
	// }

	// cara 1 buat array
	var names [4]string
	names[0] = "Nabila"
	names[1] = "Monica"
	names[2] = "Tarkim"
	names[3] = "KimMon"

	// cara 2
	names2 := [4]string{"Raysha", "Fauzia", "Andani"}

	// cara 3
	// var names3 = [...]string{"Andini", "Siti", "Nuryanti", "Yanto"}

	// kalo pake cara ini buat nampilinnya nanti langsung nampilin isi tanpa []
	fmt.Println(names[0], names[1], names[2], names[3])
	fmt.Println(names2[0], names2[1], names2[2], names2[3])
	// fmt.Println(names3[0], names3[1], names3[2], names3[3])

	// kalo pake cara ini buat nampilinnya nanti langsung nampilin isi dengan []
	fmt.Println(names)
	fmt.Println("jumlah :", len(names))

	// var fruits = [4]string{"Apple", "Banana", "Grape", "Melon"}
	// var numbers = [...]int{2,3, 2, 4}
	// var kosong [3]int // zero value 0

	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Println("Buah", fruits[i])
	// }

	// numbers := [2][3]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// }
	// fmt.Println(numbers)
	// fmt.Println(numbers[0][0], numbers[0][1], numbers[0][2])
	// fmt.Println("jumlah :", len(numbers))

	var fruits = [4]string{"Apple", "Banana", "Grape", "Melon"}

	// for i:= 0; i < len(fruits); i++ {
	// 	fmt.Println("Buah", fruits[i])
	// }

	for i, fruit := range fruits {
		fmt.Println("Buah", i, fruit)
	}

	// 	Kesimpulan
	// Kode	Hasil
	// for _, fruit := range fruits = Hanya nilai buah
	// for index, fruit := range fruits	= Index dan nilai buah
	// for index := range fruits = Hanya index


	for _, fruit := range fruits {
		fmt.Println("Buah", fruit)
	}
}
