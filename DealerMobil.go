package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Mobil struct {
	ID    string
	Merk  string
	Tahun int
	Harga int
	tJual int
}

type Pabrikan struct {
	ID     string
	Nama   string
	Mobil  [2023]Mobil
	jMobil int
}

type DataMobil struct {
	Pabrikans [2023]Pabrikan
	jPabrik   int
}

func main() {
	var arr DataMobil
	var idp, idm string
	var running bool = true

	//data dummy
	arr.Pabrikans[0].ID = "A02"
	arr.Pabrikans[0].Nama = "PT_Astra"
	arr.jPabrik++

	arr.Pabrikans[0].Mobil[0].ID = "A1"
	arr.Pabrikans[0].Mobil[0].Merk = "ayla"
	arr.Pabrikans[0].Mobil[0].Tahun = 2018
	arr.Pabrikans[0].Mobil[0].Harga = 1000
	arr.Pabrikans[0].Mobil[0].tJual = 5
	arr.Pabrikans[0].jMobil++

	arr.Pabrikans[0].Mobil[1].ID = "A2"
	arr.Pabrikans[0].Mobil[1].Merk = "Luxio"
	arr.Pabrikans[0].Mobil[1].Tahun = 2020
	arr.Pabrikans[0].Mobil[1].Harga = 1000
	arr.Pabrikans[0].Mobil[1].tJual = 2
	arr.Pabrikans[0].jMobil++

	arr.Pabrikans[0].Mobil[2].ID = "A3"
	arr.Pabrikans[0].Mobil[2].Merk = "Sirion"
	arr.Pabrikans[0].Mobil[2].Tahun = 2023
	arr.Pabrikans[0].Mobil[2].Harga = 1000
	arr.Pabrikans[0].Mobil[2].tJual = 7
	arr.Pabrikans[0].jMobil++

	arr.Pabrikans[1].ID = "A03"
	arr.Pabrikans[1].Nama = "PT_Toyota"
	arr.jPabrik++
	for running {

		var pilih string
		menu(&pilih)

		switch pilih {
		case "1":
			fmt.Print("Masukan ID Pabrik : ")
			fmt.Scan(&idp)
			addMobil(&arr, idp)
		case "2":
			addPabrik(&arr)
		case "3":
			fmt.Print("Masukan ID Pabrik : ")
			fmt.Scan(&idp)
			fmt.Print("Masukan ID Mobil Yang Ingin Dihapus : ")
			fmt.Scan(&idm)
			hapusDataMobil(&arr, idm, idp)
		case "4":
			fmt.Print("Masukan ID Pabrik : ")
			fmt.Scan(&idp)
			hapusDataPabrik(&arr, idp)
		case "5":
			fmt.Print("Masukan ID Pabrik : ")
			fmt.Scan(&idp)
			fmt.Print("Masukan ID Mobil Yang Ingin Diubah : ")
			fmt.Scan(&idm)
			ubahDataMobil(&arr, idm, idp)
		case "6":
			fmt.Print("Masukan ID Pabrik : ")
			fmt.Scan(&idp)
			ubahDataPabrik(&arr, idp)
		case "7":
			fmt.Print("Masukan ID Pabrik : ")
			fmt.Scan(&idp)
			DaftarMobilByPabrik(&arr, idp)
		case "8":
			fmt.Print("Masukan ID Mobil : ")
			fmt.Scan(&idm)
			SearchDataMobil(&arr, idm)
		case "9":
			sortByValue(&arr)
			for i := 0; i < arr.jPabrik; i++ {

				fmt.Println("Pabrik :", arr.Pabrikans[i].Nama, "jumlah Mobil :", arr.Pabrikans[i].jMobil)
			}
		case "10":
			sortCarByYear(&arr)
			cetakMobil(arr)
		case "11":
			sortByPenjualan(&arr)
			for i := 0; i < arr.jPabrik; i++ {
				fmt.Println("ID Pabrik :", arr.Pabrikans[i].ID, "Pabrik :", arr.Pabrikans[i].Nama)
				for j := 0; j < arr.Pabrikans[i].jMobil && i < 3; j++ {
					fmt.Println("ID Mobil :", arr.Pabrikans[i].Mobil[j].ID, "Merk :", arr.Pabrikans[i].Mobil[j].Merk, "Tahun Keluar :", arr.Pabrikans[i].Mobil[j].Tahun, "Harga : $", arr.Pabrikans[i].Mobil[j].Harga, "Total terjual :", arr.Pabrikans[i].Mobil[j].tJual)
				}
			}
		case "12":
			cetakMobil(arr)
		case "0":
			fmt.Print("exit")
			running = false
		default:
			fmt.Println("Masukkan tidak valid!")
		}

		fmt.Scanln()

		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
func menu(pilih *string) {
	//menampilkan pilihan menu
	fmt.Println("SELAMAT DATANG \nSILAHKAN PILIH OPTION BERIKUT\n1. Tambah Data Mobil\n2. Tambah Data Pabrik\n3. Hapus Data Mobil\n4. Hapus Data Pabrik\n5. Ubah data Mobil\n6. Ubah Data Pabrik\n7. Cari Daftar Mobil\n8. Cari Data Mobil\n9. Ranking Pabrik\n10. Sort Car By Year\n11. Sort Car by Sale\n12. Print All Data\n0. exit (selesai)")
	fmt.Print("Masukkan pilihan : ")
	fmt.Scan(&*pilih)
}
func searchMobil(t *DataMobil, idmobil string, idpabrik string) int {
	//sequential search
	//mengembalikan index mobil apabila id mobil sama dengan parameter idpabrik
	var idx, temp int
	var found bool = false
	temp = searchPabrik(t, idpabrik)
	for i := 0; i < t.Pabrikans[temp].jMobil && !found; i++ {
		if t.Pabrikans[temp].Mobil[i].ID == idmobil {
			idx = i
			found = true
		}
	}
	if found {
		return idx
	} else {
		return -1
	}
}
func searchPabrik(t *DataMobil, idpabrik string) int {
	//sequential search
	//mengembalikan index pabrik apabila id pabrik sama dengan parameter idpabrik
	var idx int
	var found bool = false
	for i := 0; i < t.jPabrik && !found; i++ {
		if t.Pabrikans[i].ID == idpabrik {
			idx = i
			found = true
		}
	}
	if found {
		return idx
	} else {
		return -1
	}
}
func addMobil(t *DataMobil, idpabrik string) {
	/*IS. data siap pada piranti masukan
	FS. data telah ditambahkan kedalam array data mobil berdasarkan id Pabrik dan jmobil bertambah*/
	var idx int
	var idm string
	idx = searchPabrik(t, idpabrik)
	if idx < 0 {
		fmt.Println("ID Pabrik Tidak Tersedia")
	} else {
		fmt.Print("Masukkan ID Mobil , Merk ,Tahun Produksi , Harga dan Total Terjual :")
		fmt.Scan(&idm, &t.Pabrikans[idx].Mobil[t.Pabrikans[idx].jMobil].Merk, &t.Pabrikans[idx].Mobil[t.Pabrikans[idx].jMobil].Tahun, &t.Pabrikans[idx].Mobil[t.Pabrikans[idx].jMobil].Harga, &t.Pabrikans[idx].Mobil[t.Pabrikans[idx].jMobil].tJual)
		if searchMobil(t, idm, idpabrik) < 0 {
			t.Pabrikans[idx].Mobil[t.Pabrikans[idx].jMobil].ID = idm
			t.Pabrikans[idx].jMobil++
		} else {
			fmt.Println("ID Mobil Tersedia")
		}

	}
}
func addPabrik(t *DataMobil) {
	/*IS. data siap pada piranti masukan
	FS. data telah ditambahkan kedalam array data pabrik dan jpabrik bertambah*/
	var idp string
	fmt.Scan(&idp, &t.Pabrikans[t.jPabrik].Nama)
	if searchPabrik(t, idp) < 0 {
		t.Pabrikans[t.jPabrik].ID = idp
		t.jPabrik++
	} else {
		fmt.Println("ID Pabrik Tersedia")
	}
}
func ubahDataMobil(t *DataMobil, idmobil string, idpabrik string) {
	/*IS. arr t data mobil,idmobil dan idpabrik diketahui
	FS. data sesuai id pabrik dan id mobil telah diubah*/
	var temp, idx int
	temp = searchPabrik(t, idpabrik)
	idx = searchMobil(t, idmobil, idpabrik)
	if idx < 0 {
		fmt.Println("ID Mobil Tidak Tersedia")
	} else {
		fmt.Print("Masukan ID Mobil , Merk Mobil , Tahun , Harga dan Total Terjual :")
		fmt.Scan(&t.Pabrikans[temp].Mobil[idx].ID, &t.Pabrikans[temp].Mobil[idx].Merk, &t.Pabrikans[temp].Mobil[idx].Tahun, &t.Pabrikans[temp].Mobil[idx].Harga, &t.Pabrikans[temp].Mobil[idx].tJual)
	}
}
func ubahDataPabrik(t *DataMobil, idpabrik string) {
	/*IS. arr t data mobil dan idpabrik diketahui
	FS. data sesuai id pabrik telah diubah*/
	var idx int
	idx = searchPabrik(t, idpabrik)
	if idx < 0 {
		fmt.Println("ID Pabrik Tidak Tersedia")
	} else {
		fmt.Print("Masuka ID Pabrik dan Nama Pabrik Yang Ingin Diubah :")
		fmt.Scan(&t.Pabrikans[idx].ID, &t.Pabrikans[idx].Nama)
	}
}
func hapusDataMobil(t *DataMobil, idmobil string, idpabrik string) {
	/*IS. t array data mobil ,id mobil,idpabrik telah terdefinisi
	FS. Data mobil sesuai idpabrik telah dihapus*/
	temp := searchPabrik(t, idpabrik)
	if temp < 0 {
		fmt.Println("ID Pabrik Tidak Tersedia")
		return
	}
	idx := searchMobil(t, idmobil, idpabrik)
	if idx < 0 {
		fmt.Println("ID Mobil Tidak Tersedia")
		return
	}
	mobil := t.Pabrikans[temp].Mobil
	for j := idx; j < t.Pabrikans[temp].jMobil-1; j++ {
		mobil[j] = mobil[j+1]
	}
	t.Pabrikans[temp].jMobil--
}
func hapusSemuaDataMobil(t *DataMobil, idx int) {
	/* IS: arr t diketahui
	   FS: semua data mobil telah dihapus */
	for j := 0; j < t.Pabrikans[idx].jMobil; j++ {
		t.Pabrikans[idx].Mobil[j] = Mobil{}
	}

	t.Pabrikans[idx].jMobil = 0
}
func hapusDataPabrik(t *DataMobil, idpabrik string) {
	/* IS: arr t dan id pabrik diketahui
	   FS: data pabrik sesuai id pabrik telah dihapus */

	idx := searchPabrik(t, idpabrik)
	hapusSemuaDataMobil(t, idx)
	if idx < 0 {
		fmt.Println("ID Pabrik Tidak Tersedia")
		return
	}
	for i := idx; i < t.jPabrik-1; i++ {
		t.Pabrikans[i] = t.Pabrikans[i+1]
	}
	t.Pabrikans[t.jPabrik-1] = Pabrikan{}
	t.jPabrik--
}
func DaftarMobilByPabrik(t *DataMobil, idpabrik string) {
	//menampilan daftar mobil berdasarkan pabik tertentu
	temp := searchPabrik(t, idpabrik)
	if temp > 0 {
		fmt.Println("Pabrik :", t.Pabrikans[temp].Nama)
		for i := 0; i < t.Pabrikans[temp].jMobil; i++ {
			fmt.Println("ID Mobil :", t.Pabrikans[temp].Mobil[i].ID, "Merk :", t.Pabrikans[temp].Mobil[i].Merk, "Tahun Keluar :", t.Pabrikans[temp].Mobil[i].Tahun, "Harga : $", t.Pabrikans[temp].Mobil[i].Harga, "Total Terjual :", t.Pabrikans[temp].Mobil[i].tJual)
		}
	} else {
		fmt.Println("ID Pabrik Tidak Tersedia")
	}

}
func SearchDataMobil(t *DataMobil, idmobil string) {
	//sequential search
	//menampilkan data mobil yang dicari
	var idx, temp int
	var found bool = false
	for i := 0; i < t.jPabrik && !found; i++ {
		for j := 0; j < t.Pabrikans[i].jMobil && !found; j++ {
			if t.Pabrikans[i].Mobil[j].ID == idmobil {
				found = true
				temp = i
				idx = j
			}
		}
	}
	if found {
		fmt.Println("Pabrik :", t.Pabrikans[temp].Nama, "ID Mobil :", t.Pabrikans[temp].Mobil[idx].ID, "Merk :", t.Pabrikans[temp].Mobil[idx].Merk, "Tahun Keluar :", t.Pabrikans[temp].Mobil[idx].Tahun, "Harga : $", t.Pabrikans[temp].Mobil[idx].Harga, "Total Terjual :", t.Pabrikans[temp].Mobil[idx].tJual)
	} else {
		fmt.Println("Data Mobil Tidak Ditemukan")
	}
}
func sortByValue(t *DataMobil) {
	//selection sort
	//mengurutkan daftar mobil berdasarkan jumlah mobil yang tersedia di pabrik
	for i := 0; i < t.jPabrik-1; i++ {
		minIndex := i
		for j := i + 1; j < t.jPabrik; j++ {
			if t.Pabrikans[j].jMobil > t.Pabrikans[minIndex].jMobil {
				minIndex = j
			}
		}

		// Menukar pabrik dengan jumlah mobil terkecil dengan pabrik pertama pada sisa array
		t.Pabrikans[i], t.Pabrikans[minIndex] = t.Pabrikans[minIndex], t.Pabrikans[i]
	}
}

func sortCarByYear(t *DataMobil) {
	//selection sort
	//mengurutkan daftar mobil berdasarkan tahun keluar
	for y := 0; y < t.jPabrik; y++ {
		for i := 0; i < t.Pabrikans[y].jMobil-1; i++ {
			minIndex := i
			for j := i + 1; j < t.Pabrikans[y].jMobil; j++ {
				if t.Pabrikans[y].Mobil[j].Tahun > t.Pabrikans[y].Mobil[minIndex].Tahun {
					minIndex = j
				}
			}

			// Menukar elemen terkecil dengan elemen pertama pada sisa array yang belum diurutkan
			t.Pabrikans[y].Mobil[i], t.Pabrikans[y].Mobil[minIndex] = t.Pabrikans[y].Mobil[minIndex], t.Pabrikans[y].Mobil[i]
		}

	}
}
func sortByPenjualan(t *DataMobil) {
	// selection sort
	// mengurutkan daftar mobil berdasarkan total penjualan
	for i := 0; i < t.jPabrik; i++ {
		for j := 0; j < t.Pabrikans[i].jMobil-1; j++ {
			maxIndex := j
			for k := j + 1; k < t.Pabrikans[i].jMobil; k++ {
				if t.Pabrikans[i].Mobil[k].tJual > t.Pabrikans[i].Mobil[maxIndex].tJual {
					maxIndex = k
				}
			}

			// Menukar elemen terbesar dengan elemen pertama pada sisa array yang belum diurutkan
			t.Pabrikans[i].Mobil[j], t.Pabrikans[i].Mobil[maxIndex] = t.Pabrikans[i].Mobil[maxIndex], t.Pabrikans[i].Mobil[j]
		}
	}
}
func cetakMobil(t DataMobil) {
	//menampilkan seluruh data pabrik dan mobil
	for i := 0; i < t.jPabrik; i++ {
		fmt.Println("ID Pabrik :", t.Pabrikans[i].ID, "Pabrik :", t.Pabrikans[i].Nama)
		for j := 0; j < t.Pabrikans[i].jMobil; j++ {
			fmt.Println("ID Mobil :", t.Pabrikans[i].Mobil[j].ID, "Merk :", t.Pabrikans[i].Mobil[j].Merk, "Tahun Keluar :", t.Pabrikans[i].Mobil[j].Tahun, "Harga : $", t.Pabrikans[i].Mobil[j].Harga, "Total terjual :", t.Pabrikans[i].Mobil[j].tJual)
		}
	}
}
