package main

import (
	"fmt"
	"os"
	"os/exec"
	"sort"
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
	arr.Pabrikans[0].Mobil[1].Merk = "luxio"
	arr.Pabrikans[0].Mobil[1].Tahun = 2020
	arr.Pabrikans[0].Mobil[1].Harga = 1000
	arr.Pabrikans[0].Mobil[1].tJual = 2
	arr.Pabrikans[0].jMobil++

	arr.Pabrikans[0].Mobil[2].ID = "A3"
	arr.Pabrikans[0].Mobil[2].Merk = "sirion"
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
			sortByMerk(&arr)
			cetakMobil(arr)
		case "13":
			displayTop3Penjualan(&arr)
		case "14":
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
	fmt.Println("SELAMAT DATANG \nSILAHKAN PILIH OPTION BERIKUT\n1. Tambah Data Mobil\n2. Tambah Data Pabrik\n3. Hapus Data Mobil\n4. Hapus Data Pabrik\n5. Ubah data Mobil\n6. Ubah Data Pabrik\n7. Cari Daftar Mobil\n8. Cari Data Mobil\n9. Mengurutkan Nama Pabrik Berdasarkan Jumlah Mobil Di Setiap Pabrik\n10. Mengurutkan Mobil Berdasarkan Tahun Produksi Pada Setiap Pabrik\n11. Mengurutkan Mobil Berdasarkan Total Penjualan Pada Setiap Pabrik\n12. Mengurutkan Berdasarkan Merk Pada Setiap Pabrik\n13. Menampilkan 3 Daftar Mobil Dan Pabrikan Dengan Jumlah Penjualan Tertinggi\n14. Print All Data\n0. exit (selesai)")
	fmt.Print("Masukkan pilihan : ")
	fmt.Scan(&*pilih)
}

func searchMobil(t *DataMobil, idmobil string, idpabrik string) int {
	// Mengembalikan indeks mobil jika ID mobil sama dengan parameter idpabrik

	// Mencari indeks pabrik menggunakan fungsi searchPabrik
	temp := searchPabrik(t, idpabrik)

	// Pastikan data mobil dalam setiap pabrik telah diurutkan berdasarkan ID mobil
	sort.Slice(t.Pabrikans[temp].Mobil[:t.Pabrikans[temp].jMobil], func(i, j int) bool {
		return t.Pabrikans[temp].Mobil[i].ID < t.Pabrikans[temp].Mobil[j].ID
	})

	// Melakukan binary search pada data mobil yang telah diurutkan dan mengembalikan indeks mobil jika ditemukan
	idx := binarySearch(t.Pabrikans[temp].Mobil[:t.Pabrikans[temp].jMobil], idmobil)
	return idx
}


func binarySearch(mobil []Mobil, idmobil string) int {
	low, high := 0, len(mobil)-1

	// Melakukan binary search selama low tidak melebihi high
	for low <= high {
		mid := (low + high) / 2

		// Jika ID mobil pada tengah sesuai dengan idmobil, kembalikan indeks tengah
		if mobil[mid].ID == idmobil {
			return mid
		} else if mobil[mid].ID < idmobil {
			// Jika ID mobil pada tengah lebih kecil dari idmobil, ubah low menjadi mid+1 untuk mencari di setengah kanan
			low = mid + 1
		} else {
			// Jika ID mobil pada tengah lebih besar dari idmobil, ubah high menjadi mid-1 untuk mencari di setengah kiri
			high = mid - 1
		}
	}
	// Jika tidak ditemukan, kembalikan nilai -1
	return -1
}
func searchPabrik(t *DataMobil, idpabrik string) int {
	// Sequential search
	// Mengembalikan indeks pabrik jika ID pabrik sama dengan parameter idpabrik

	var idx int
	var found bool = false

	// Melakukan sequential search pada setiap pabrik
	for i := 0; i < t.jPabrik && !found; i++ {
		// Memeriksa apakah ID pabrik pada indeks i sama dengan idpabrik
		if t.Pabrikans[i].ID == idpabrik {
			idx = i
			found = true
		}
	}

	// Jika ditemukan, mengembalikan indeks pabrik
	if found {
		return idx
	} else {
		// Jika tidak ditemukan, mengembalikan nilai -1 sebagai indikasi tidak ditemukannya pabrik dengan ID yang dicari
		return -1
	}
}

func addMobil(t *DataMobil, idpabrik string) {
	/* IS. data siap pada piranti masukan
	   FS. data telah ditambahkan kedalam array data mobil berdasarkan id Pabrik dan jmobil bertambah */

	var idx int
	var idm string

	// Mencari indeks pabrik berdasarkan idpabrik
	idx = searchPabrik(t, idpabrik)

	// Jika indeks pabrik tidak ditemukan
	if idx < 0 {
		fmt.Println("ID Pabrik Tidak Tersedia")
	} else {
		fmt.Print("Masukkan ID Mobil, Merk, Tahun Produksi, Harga, dan Total Terjual:")
		// Meminta input ID Mobil, Merk, Tahun Produksi, Harga, dan Total Terjual
		fmt.Scan(&idm, &t.Pabrikans[idx].Mobil[t.Pabrikans[idx].jMobil].Merk, &t.Pabrikans[idx].Mobil[t.Pabrikans[idx].jMobil].Tahun, &t.Pabrikans[idx].Mobil[t.Pabrikans[idx].jMobil].Harga, &t.Pabrikans[idx].Mobil[t.Pabrikans[idx].jMobil].tJual)

		// Memeriksa apakah ID Mobil sudah ada dalam pabrik tersebut
		if searchMobil(t, idm, idpabrik) < 0 {
			// Jika ID Mobil belum ada, tambahkan data mobil ke dalam pabrik
			t.Pabrikans[idx].Mobil[t.Pabrikans[idx].jMobil].ID = idm
			t.Pabrikans[idx].jMobil++
		} else {
			fmt.Println("ID Mobil Tersedia")
		}
	}
}

func addPabrik(t *DataMobil) {
	/* IS. data siap pada piranti masukan
	   FS. data telah ditambahkan kedalam array data pabrik dan jpabrik bertambah */

	var idp string

	// Meminta input ID Pabrik dan Nama Pabrik
	fmt.Print("Masukkan ID Pabrik, Nama Pabrik:")
	fmt.Scan(&idp, &t.Pabrikans[t.jPabrik].Nama)

	// Memeriksa apakah ID Pabrik sudah ada dalam data pabrik
	if searchPabrik(t, idp) < 0 {
		// Jika ID Pabrik belum ada, tambahkan data pabrik ke dalam array
		t.Pabrikans[t.jPabrik].ID = idp
		t.jPabrik++
	} else {
		fmt.Println("ID Pabrik Tersedia")
	}
}

func ubahDataMobil(t *DataMobil, idmobil string, idpabrik string) {
	// Fungsi untuk mengubah data mobil berdasarkan ID mobil dan ID pabrik
	// Menerima parameter tipe pointer *DataMobil, idmobil, dan idpabrik
	var temp, idx int
	temp = searchPabrik(t, idpabrik) // Mencari indeks pabrik berdasarkan ID pabrik
	idx = searchMobil(t, idmobil, idpabrik) // Mencari indeks mobil berdasarkan ID mobil dan ID pabrik
	if idx < 0 {
		fmt.Println("ID Mobil Tidak Tersedia") // Jika indeks mobil tidak ditemukan, cetak pesan bahwa ID mobil tidak tersedia
	} else {
		fmt.Print("Masukan ID Mobil, Merk Mobil, Tahun, Harga, dan Total Terjual:")
		// Meminta pengguna memasukkan ID Mobil, Merk Mobil, Tahun, Harga, dan Total Terjual
		fmt.Scan(&t.Pabrikans[temp].Mobil[idx].ID, &t.Pabrikans[temp].Mobil[idx].Merk, &t.Pabrikans[temp].Mobil[idx].Tahun, &t.Pabrikans[temp].Mobil[idx].Harga, &t.Pabrikans[temp].Mobil[idx].tJual)
		// Mengubah data mobil dengan data yang dimasukkan pengguna
	}
}

func ubahDataPabrik(t *DataMobil, idpabrik string) {
	// Fungsi untuk mengubah data pabrik berdasarkan ID pabrik
	// Menerima parameter tipe pointer *DataMobil dan idpabrik
	var idx int
	idx = searchPabrik(t, idpabrik) // Mencari indeks pabrik berdasarkan ID pabrik
	if idx < 0 {
		fmt.Println("ID Pabrik Tidak Tersedia") // Jika indeks pabrik tidak ditemukan, cetak pesan bahwa ID pabrik tidak tersedia
	} else {
		fmt.Print("Masuka ID Pabrik dan Nama Pabrik:") // Meminta pengguna memasukkan ID Pabrik dan Nama Pabrik baru
		fmt.Scan(&t.Pabrikans[idx].ID, &t.Pabrikans[idx].Nama) // Mengubah data pabrik dengan data yang dimasukkan pengguna
	}
}

func hapusDataMobil(t *DataMobil, idmobil string, idpabrik string) {
	// Fungsi untuk menghapus data mobil berdasarkan ID mobil dan ID pabrik
	// Menerima parameter tipe pointer *DataMobil, idmobil, dan idpabrik
	temp := searchPabrik(t, idpabrik) // Mencari indeks pabrik berdasarkan ID pabrik
	if temp < 0 {
		fmt.Println("ID Pabrik Tidak Tersedia") // Jika indeks pabrik tidak ditemukan, cetak pesan bahwa ID pabrik tidak tersedia
		return
	}
	idx := searchMobil(t, idmobil, idpabrik) // Mencari indeks mobil berdasarkan ID mobil dan ID pabrik
	if idx < 0 {
		fmt.Println("ID Mobil Tidak Tersedia") // Jika indeks mobil tidak ditemukan, cetak pesan bahwa ID mobil tidak tersedia
		return
	}
	mobil := t.Pabrikans[temp].Mobil // Menyimpan array mobil dari pabrik yang ditemukan
	for j := idx; j < t.Pabrikans[temp].jMobil-1; j++ {
		mobil[j] = mobil[j+1] // Menggeser data mobil ke kiri untuk mengisi celah data yang dihapus
	}
	t.Pabrikans[temp].jMobil-- // Mengurangi jumlah mobil pada pabrik
}

func hapusSemuaDataMobil(t *DataMobil, idx int) {
	// Fungsi untuk menghapus semua data mobil pada indeks pabrik tertentu
	// Menerima parameter tipe pointer *DataMobil dan idx (indeks pabrik)
	for j := 0; j < t.Pabrikans[idx].jMobil; j++ {
		t.Pabrikans[idx].Mobil[j] = Mobil{} // Menghapus data mobil pada indeks j dengan mengganti nilainya dengan objek Mobil kosong
	}

	t.Pabrikans[idx].jMobil = 0 // Mengatur jumlah mobil pada pabrik menjadi 0
}

func hapusDataPabrik(t *DataMobil, idpabrik string) {
	// Fungsi untuk menghapus data pabrik berdasarkan ID pabrik
	// Menerima parameter tipe pointer *DataMobil dan idpabrik
	idx := searchPabrik(t, idpabrik) // Mencari indeks pabrik berdasarkan ID pabrik
	hapusSemuaDataMobil(t, idx) // Menghapus semua data mobil pada pabrik dengan indeks tersebut
	if idx < 0 {
		fmt.Println("ID Pabrik Tidak Tersedia") // Jika indeks pabrik tidak ditemukan, cetak pesan bahwa ID pabrik tidak tersedia
		return
	}
	for i := idx; i < t.jPabrik-1; i++ {
		t.Pabrikans[i] = t.Pabrikans[i+1] // Menggeser data pabrik ke kiri untuk mengisi celah data yang dihapus
	}
	t.Pabrikans[t.jPabrik-1] = Pabrikan{} // Mengganti data pabrik terakhir dengan objek Pabrikan kosong
	t.jPabrik-- // Mengurangi jumlah pabrik
}

func DaftarMobilByPabrik(t *DataMobil, idpabrik string) {
	// Fungsi untuk menampilkan daftar mobil berdasarkan pabrik tertentu
	// Menerima parameter tipe pointer *DataMobil dan idpabrik
	temp := searchPabrik(t, idpabrik) // Mencari indeks pabrik berdasarkan ID pabrik
	if temp < 0 {
		fmt.Println("ID Pabrik Tidak Tersedia") // Jika indeks pabrik tidak ditemukan, cetak pesan bahwa ID pabrik tidak tersedia
	} else {
		fmt.Println("Pabrik :", t.Pabrikans[temp].Nama) // Menampilkan nama pabrik yang ditemukan
		for i := 0; i < t.Pabrikans[temp].jMobil; i++ {
			// Menampilkan informasi mobil pada pabrik tersebut
			fmt.Println("ID Mobil :", t.Pabrikans[temp].Mobil[i].ID, "Merk :", t.Pabrikans[temp].Mobil[i].Merk, "Tahun Keluar :", t.Pabrikans[temp].Mobil[i].Tahun, "Harga : $", t.Pabrikans[temp].Mobil[i].Harga, "Total Terjual :", t.Pabrikans[temp].Mobil[i].tJual)
		}
	}
}

func SearchDataMobil(t *DataMobil, idmobil string) {
	// Fungsi untuk mencari data mobil berdasarkan ID mobil
	// Menerima parameter tipe pointer *DataMobil dan idmobil
	var idx, temp int
	var found bool = false
	for i := 0; i < t.jPabrik && !found; i++ {
		// Melakukan iterasi melalui pabrik-pabrik yang ada
		for j := 0; j < t.Pabrikans[i].jMobil && !found; j++ {
			// Melakukan iterasi melalui mobil-mobil di setiap pabrik
			if t.Pabrikans[i].Mobil[j].ID == idmobil {
				// Jika ditemukan ID mobil yang sesuai, set found menjadi true dan menyimpan indeks pabrik dan mobil
				found = true
				temp = i
				idx = j
			}
		}
	}
	if found {
		// Jika data mobil ditemukan, menampilkan informasi mobil
		fmt.Println("Pabrik :", t.Pabrikans[temp].Nama, "ID Mobil :", t.Pabrikans[temp].Mobil[idx].ID, "Merk :", t.Pabrikans[temp].Mobil[idx].Merk, "Tahun Keluar :", t.Pabrikans[temp].Mobil[idx].Tahun, "Harga : $", t.Pabrikans[temp].Mobil[idx].Harga, "Total Terjual :", t.Pabrikans[temp].Mobil[idx].tJual)
	} else {
		// Jika data mobil tidak ditemukan, cetak pesan bahwa data mobil tidak ditemukan
		fmt.Println("Data Mobil Tidak Ditemukan")
	}
}

func sortByValue(t *DataMobil) {
	// Fungsi untuk mengurutkan daftar mobil berdasarkan jumlah mobil yang tersedia di pabrik
	// Menerima parameter tipe pointer *DataMobil
	// Menggunakan algoritma selection sort
	for i := 0; i < t.jPabrik-1; i++ {
		// Iterasi melalui pabrik-pabrik
		minIndex := i
		for j := i + 1; j < t.jPabrik; j++ {
			// Iterasi melalui pabrik-pabrik yang tersisa untuk mencari jumlah mobil terkecil
			if t.Pabrikans[j].jMobil > t.Pabrikans[minIndex].jMobil {
				// Jika jumlah mobil pada pabrik saat ini lebih besar dari jumlah mobil pada pabrik dengan indeks minIndex, perbarui minIndex
				minIndex = j
			}
		}

		// Menukar posisi pabrik dengan jumlah mobil terkecil dengan pabrik pertama pada sisa array
		t.Pabrikans[i], t.Pabrikans[minIndex] = t.Pabrikans[minIndex], t.Pabrikans[i]
	}
}

func sortCarByYear(t *DataMobil) {
	// Fungsi untuk mengurutkan daftar mobil berdasarkan tahun keluar
	// Menerima parameter tipe pointer *DataMobil
	// Menggunakan algoritma selection sort
	for y := 0; y < t.jPabrik; y++ {
		// Iterasi melalui pabrik-pabrik
		for i := 0; i < t.Pabrikans[y].jMobil-1; i++ {
			// Iterasi melalui mobil-mobil pada pabrik saat ini
			minIndex := i
			for j := i + 1; j < t.Pabrikans[y].jMobil; j++ {
				// Iterasi melalui mobil-mobil yang belum diurutkan pada pabrik saat ini untuk mencari tahun keluar terkecil
				if t.Pabrikans[y].Mobil[j].Tahun > t.Pabrikans[y].Mobil[minIndex].Tahun {
					// Jika tahun keluar mobil pada posisi saat ini lebih besar dari tahun keluar mobil pada posisi minIndex, perbarui minIndex
					minIndex = j
				}
			}

			// Menukar posisi mobil dengan tahun keluar terkecil dengan mobil pertama pada sisa array yang belum diurutkan
			t.Pabrikans[y].Mobil[i], t.Pabrikans[y].Mobil[minIndex] = t.Pabrikans[y].Mobil[minIndex], t.Pabrikans[y].Mobil[i]
		}
	}
}

func sortByPenjualan(t *DataMobil) {
	// Fungsi untuk mengurutkan daftar mobil berdasarkan total penjualan
	// Menerima parameter tipe pointer *DataMobil
	// Menggunakan algoritma insertion sort
	for i := 0; i < t.jPabrik; i++ {
		// Iterasi melalui pabrik-pabrik
		for j := 1; j < t.Pabrikans[i].jMobil; j++ {
			// Iterasi melalui mobil-mobil pada pabrik saat ini, dimulai dari indeks 1
			key := t.Pabrikans[i].Mobil[j]
			k := j - 1

			// Pindahkan elemen-elemen yang lebih besar dari kunci ke posisi setelahnya
			for k >= 0 && t.Pabrikans[i].Mobil[k].tJual < key.tJual {
				t.Pabrikans[i].Mobil[k+1] = t.Pabrikans[i].Mobil[k]
				k--
			}

			t.Pabrikans[i].Mobil[k+1] = key
		}
	}
}

func sortByMerk(t *DataMobil) {
	// Fungsi untuk mengurutkan daftar mobil berdasarkan Merk (nama)
	// Menerima parameter tipe pointer *DataMobil
	for i := 0; i < t.jPabrik; i++ {
		// Iterasi melalui pabrik-pabrik
		// Pastikan hanya mengurutkan jumlah mobil yang valid
		jMobil := t.Pabrikans[i].jMobil
		mobils := t.Pabrikans[i].Mobil[:jMobil]

		// Insertion Sort berdasarkan Merk (nama)
		for j := 1; j < jMobil; j++ {
			// Iterasi melalui mobil-mobil pada pabrik saat ini, dimulai dari indeks 1
			key := mobils[j]
			k := j - 1
			for k >= 0 && mobils[k].Merk > key.Merk {
				// Pindahkan elemen-elemen yang lebih besar dari kunci ke posisi setelahnya
				mobils[k+1] = mobils[k]
				k--
			}
			mobils[k+1] = key
		}
	}
}

func displayTop3Penjualan(t *DataMobil) {
	// Fungsi untuk menampilkan 3 daftar mobil dan pabrikan dengan penjualan tertinggi
	// Menerima parameter tipe pointer *DataMobil

	type Penjualan struct {
		NamaPabrikan string
		NamaMobil    string
		Penjualan    int
	}

	var penjualan []Penjualan

	for i := 0; i < t.jPabrik; i++ {
		// Iterasi melalui pabrik-pabrik
		for j := 0; j < t.Pabrikans[i].jMobil; j++ {
			// Iterasi melalui mobil-mobil pada pabrik saat ini
			p := Penjualan{
				NamaPabrikan: t.Pabrikans[i].Nama,
				NamaMobil:    t.Pabrikans[i].Mobil[j].Merk,
				Penjualan:    t.Pabrikans[i].Mobil[j].tJual,
			}
			penjualan = append(penjualan, p)
		}
	}

	// Mengurutkan daftar penjualan berdasarkan jumlah penjualan secara menurun
	sort.Slice(penjualan, func(i, j int) bool {
		return penjualan[i].Penjualan > penjualan[j].Penjualan
	})

	// Menampilkan 3 daftar mobil dan pabrikan dengan penjualan tertinggi
	fmt.Println("Daftar 3 Mobil dan Pabrikan dengan Jumlah Penjualan Tertinggi:")

	for i := 0; i < 3 && i < len(penjualan); i++ {
		// Iterasi melalui 3 data pertama atau seluruh data jika kurang dari 3
		fmt.Printf("%d. Pabrikan: %s, Mobil: %s, Jumlah Penjualan: %d\n", i+1, penjualan[i].NamaPabrikan, penjualan[i].NamaMobil, penjualan[i].Penjualan)
	}
}

func cetakMobil(t DataMobil) {
	// Fungsi untuk mencetak seluruh data pabrik dan mobil
	// Menerima parameter tipe DataMobil (tidak pointer)

	for i := 0; i < t.jPabrik; i++ {
		// Iterasi melalui pabrik-pabrik
		fmt.Println("ID Pabrik :", t.Pabrikans[i].ID, "Pabrik :", t.Pabrikans[i].Nama)
		for j := 0; j < t.Pabrikans[i].jMobil; j++ {
			// Iterasi melalui mobil-mobil pada pabrik saat ini
			fmt.Println("ID Mobil :", t.Pabrikans[i].Mobil[j].ID, "Merk :", t.Pabrikans[i].Mobil[j].Merk, "Tahun Keluar :", t.Pabrikans[i].Mobil[j].Tahun, "Harga : $", t.Pabrikans[i].Mobil[j].Harga, "Total terjual :", t.Pabrikans[i].Mobil[j].tJual)
		}
	}
}

