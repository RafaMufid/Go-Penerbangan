package main

import (
	"fmt"
)

const NMAX int = 100

type penerbangan struct {
	idPenerbangan        string
	maskapai             string
	asal                 string
	tujuan               string
	harga                int
	tanggalKeberangkatan struct {
		tanggal int
		bulan   int
		tahun   int
	}
}

type tabPenerbangan [NMAX]penerbangan

func main() {
	var flight tabPenerbangan
	var jumlahFlight int
	var pilihan int
	var batas int

	fmt.Println("Masukkan jumlah penerbangan:")
	//batas input penerbangan
	fmt.Scan(&batas)
	//jumlah flight otomatis di assign 0
	tambahPenerbangan(&flight, &jumlahFlight, batas)
	//melakukan sorting agar data terurut
	insertionSort(&flight, batas)

	//akan terus menjalankan menu utama selama pilihan tidak 7
	for pilihan != 7 {
		fmt.Println("===============================================")
		fmt.Println("============== Selamat Datang! ================")
		fmt.Println("===============================================")
		fmt.Println("Pilih Menu : ")
		fmt.Println("1. Lihat Data Penerbangan.")
		fmt.Println("2. Tambah Data Penerbangan.")
		fmt.Println("3. Cari Jadwal Penerbangan.")
		fmt.Println("4. Pesan Penerbangan.")
		fmt.Println("5. Edit Data Penerbangan.")
		fmt.Println("6. Hapus Data Penerbangan.")
		fmt.Println("7. Selesai.")
		fmt.Println("===============================================")

		fmt.Scan(&pilihan)
		if pilihan == 1 {
			cetakPenerbangan(flight, jumlahFlight)
		} else if pilihan == 2 {
			fmt.Println("Masukkan jumlah penerbangan:")
			fmt.Scan(&batas)
			tambahPenerbangan(&flight, &jumlahFlight, batas)
			insertionSort(&flight, jumlahFlight)
		} else if pilihan == 3 {
			cariJadwal(flight, jumlahFlight)
		} else if pilihan == 4 {
			pesanPenerbangan(flight, jumlahFlight)
		} else if pilihan == 5 {
			fmt.Println("Edit Data Penerbangan")
			editPenerbangan(&flight, jumlahFlight)
			insertionSort(&flight, jumlahFlight)
		} else if pilihan == 6 {
			fmt.Println("Hapus Data Penerbangan")
			hapusPenerbangan(&flight, &jumlahFlight)
		} else if pilihan == 7 {
			fmt.Println("Terima kasih telah menggunakan program ini!")
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tambahPenerbangan(A *tabPenerbangan, n *int, lim int) {
	var i int
	var j int
	var nPenerbangan int
	var id string
	var idTaken bool
	var slash string

	//untuk mengupdate batas loop jika ingin melakukan input data kembali
	nPenerbangan = *n + lim

	//jika jumlah penerbangan melebihi kapasitas NMAX, maka jumlah penerbangan akan di set menjadi NMAX
	if *n > NMAX {
		*n = NMAX
	}

	//i adalah jumlah penerbangan yang tersedia; nPenerbangan adalah batas penerbangan yang akan di input
	for i = *n; i < nPenerbangan; i++ {
		//idTaken menjadi true di awal tiap iterasi
		idTaken = true
		//menjalankan loop selama idTaken true
		for idTaken {
			fmt.Println("Masukkan ID Penerbangan:")
			fmt.Scan(&id)
			//idTaken menjadi false dengan asumsi id selanjutnya unik
			idTaken = false
			//kondisi cek apakah id sudah ada di array tersimpan dan sudah di input sebelumnya
			for j = 0; j < *n || j < i; j++ { //kondisi saat id yang diinput sudah ada di dalam array
				if A[j].idPenerbangan == id {
					fmt.Println("Id Penerbangan sudah terdaftar. Silahkan masukkan Id Penerbangan lain")
					//kembali menjalankan loop sampai id unik
					idTaken = true
				}
			}
		}
		//assign inputan id ke dalam array
		A[i].idPenerbangan = id
		fmt.Println("Masukkan data penerbangan:")
		fmt.Scan(&A[i].maskapai, &A[i].asal, &A[i].tujuan, &A[i].harga, &A[i].tanggalKeberangkatan.tanggal, &slash, &A[i].tanggalKeberangkatan.bulan, &slash, &A[i].tanggalKeberangkatan.tahun)
	}
	//assign n dengan n penerbangan untuk mengupdate panjang array
	*n = nPenerbangan
	fmt.Println("Data berhasil ditambahkan.")
	fmt.Println(".")
	//kembali ke menu utama
}

func cetakPenerbangan(A tabPenerbangan, n int) {
	var i int

	if n == 0 {
		fmt.Println("Belum ada penerbangan yang tersedia.")
		fmt.Println(".")
	} else {
		fmt.Println("Daftar Penerbangan:")
		fmt.Println("----------------------------------------------------------------------------------------------------------------")
		fmt.Printf("%-16s %-18s %-12s %-12s %-12s %-12s %-12s %-12s\n", "ID Penerbangan", "Maskapai", "Asal", "Tujuan", "Harga", "Tanggal", "Bulan", "Tahun")
		fmt.Println("----------------------------------------------------------------------------------------------------------------")
		for i = 0; i < n; i++ {
			fmt.Printf("%-16s %-18s %-12s %-12s %-12d %-12d %-12d %-12d\n", A[i].idPenerbangan, A[i].maskapai, A[i].asal, A[i].tujuan, A[i].harga, A[i].tanggalKeberangkatan.tanggal, A[i].tanggalKeberangkatan.bulan, A[i].tanggalKeberangkatan.tahun)
		}
		fmt.Println(".")
		//kembali ke menu utama
	}
}

func cariJadwal(A tabPenerbangan, n int) {
	fmt.Println("===============================================")
	fmt.Println("Pencarian Penerbangan")
	fmt.Println("===============================================")
	fmt.Println("1. Cari Berdasarkan Tujuan.")
	fmt.Println("2. Cari Berdasarkan Asal.")
	fmt.Println("3. Kembali ke Menu Utama.")

	var pilihan int
	var asal string
	var tujuan string

	fmt.Scan(&pilihan)

	if pilihan == 1 {
		fmt.Print("Masukkan Tujuan: ")
		fmt.Scan(&tujuan)
		cariTujuan(A, n, tujuan)
	} else if pilihan == 2 {
		fmt.Print("Masukkan Asal: ")
		fmt.Scan(&asal)
		cariAsal(A, n, asal)
	} else if pilihan > 3 || pilihan < 1 {
		fmt.Println("Pilihan tidak valid.")
	}
}

func cariAsal(A tabPenerbangan, n int, asal string) {
	var i int
	var found bool

	//asumsi sebelum melakukan pencarian
	found = false
	fmt.Printf("Daftar Penerbangan %s:\n", asal)
	fmt.Println("----------------------------------------------------------------------------------------------------------------")
	fmt.Printf("%-16s %-18s %-12s %-12s %-12s %-12s %-12s %-12s\n", "ID Penerbangan", "Maskapai", "Asal", "Tujuan", "Harga", "Tanggal", "Bulan", "Tahun")
	fmt.Println("----------------------------------------------------------------------------------------------------------------")
	//melakukan pencarian hingga array terakhir
	for i = 0; i < n; i++ {
		if A[i].asal == asal { //kondisi asal ditemukan dalam array
			fmt.Printf("%-16s %-18s %-12s %-12s %-12d %-12d %-12d %-12d\n", A[i].idPenerbangan, A[i].maskapai, A[i].asal, A[i].tujuan, A[i].harga, A[i].tanggalKeberangkatan.tanggal, A[i].tanggalKeberangkatan.bulan, A[i].tanggalKeberangkatan.tahun)
			//pencarian sesuai; penerbangan ditemukan; mengupdate nilai found menjadi true
			found = true
		}
	}
	fmt.Println(".")

	if !found {
		//kondisi setelah dilakukan pencarian tidak ada asal yang sesuai
		fmt.Println("Penerbangan tidak ditemukan.")
		fmt.Println(".")
	}
}

func cariTujuan(A tabPenerbangan, n int, tujuan string) {
	var i int
	var found bool

	found = false
	//asumsi sebelum melakukan pencarian
	fmt.Printf("Daftar Penerbangan %s:\n", tujuan)
	fmt.Println("----------------------------------------------------------------------------------------------------------------")
	fmt.Printf("%-16s %-18s %-12s %-12s %-12s %-12s %-12s %-12s\n", "ID Penerbangan", "Maskapai", "Asal", "Tujuan", "Harga", "Tanggal", "Bulan", "Tahun")
	fmt.Println("----------------------------------------------------------------------------------------------------------------")
	for i = 0; i < n; i++ {
		if A[i].tujuan == tujuan {
			fmt.Printf("%-16s %-18s %-12s %-12s %-12d %-12d %-12d %-12d\n", A[i].idPenerbangan, A[i].maskapai, A[i].asal, A[i].tujuan, A[i].harga, A[i].tanggalKeberangkatan.tanggal, A[i].tanggalKeberangkatan.bulan, A[i].tanggalKeberangkatan.tahun)
			//pencarian selesai; penerbangan ditemukan; mengupdate nilai found menjadi true
			found = true
		}
	}
	fmt.Println(".")

	//kondisi setelah dilakukan pencarian tidak ada tujuan yang sesuai
	if !found {
		fmt.Println("Penerbangan tidak ditemukan.")
		fmt.Println(".")
	}
}

func pesanPenerbangan(A tabPenerbangan, n int) {
	var i int
	var tanggal, bulan, tahun int
	var asal, tujuan, IdP string
	var index int
	var found bool
	var slash string //variabel slash untuk format input tanggal keberangkatan

	fmt.Print("Masukkan asal: ")
	fmt.Scan(&asal)
	fmt.Print("Masukkan tujuan: ")
	fmt.Scan(&tujuan)
	fmt.Print("Masukkan tanggal keberangkatan (format: dd / mm / yyyy): ")
	fmt.Scan(&tanggal, &slash, &bulan, &slash, &tahun)

	//memvalidasi input tanggal keberangkatan
	if !validasiTanggal(tanggal, bulan, tahun) {
		fmt.Println("Format tanggal tidak valid. Gunakan format dd / mm / yyyy.")
		//menjalankan kembali procedure pesan penerbangan
		pesanPenerbangan(A, n)
	}

	//asumsi asal, tujuan, dan tanggal keberangkatan belum ditemukan
	found = false

	fmt.Println("Daftar Penerbangan:")
	fmt.Println("----------------------------------------------------------------------------------------------------------------")
	fmt.Printf("%-16s %-18s %-12s %-12s %-12s %-12s %-12s %-12s\n", "ID Penerbangan", "Maskapai", "Asal", "Tujuan", "Harga", "Tanggal", "Bulan", "Tahun")
	fmt.Println("----------------------------------------------------------------------------------------------------------------")
	for i = 0; i < n; i++ {
		if A[i].asal == asal && A[i].tujuan == tujuan && A[i].tanggalKeberangkatan.tanggal == tanggal && A[i].tanggalKeberangkatan.bulan == bulan && A[i].tanggalKeberangkatan.tahun == tahun {
			fmt.Printf("%-16s %-18s %-12s %-12s %-12d %-12d %-12d %-12d\n", A[i].idPenerbangan, A[i].maskapai, A[i].asal, A[i].tujuan, A[i].harga, A[i].tanggalKeberangkatan.tanggal, A[i].tanggalKeberangkatan.bulan, A[i].tanggalKeberangkatan.tahun)
			fmt.Println(".")
			//pencarian selesai; asal, tujuan dan tanggal keberangkatan ditemukan
			found = true
		}
	}

	//kondisi setelah melakukan pengecekan tetap tidak ditemukan
	if !found {
		fmt.Println("Penerbangan tidak tersedia.")
		fmt.Println(".")
	} else {
		fmt.Print("Masukkan id penerbangan yang ingin dipesan: ")
		fmt.Scan(&IdP)

		//melakukan binary search untuk index penerbangan yang ingin dipesan
		index = binarySearch(A, n, IdP)

		if index != -1 {
			fmt.Printf("Anda telah memilih penerbangan %s dengan maskapai %s dari %s ke %s pada %d/%d/%d.\n", A[index].idPenerbangan, A[index].maskapai, A[index].asal, A[index].tujuan, A[index].tanggalKeberangkatan.tanggal, A[index].tanggalKeberangkatan.bulan, A[index].tanggalKeberangkatan.tahun)
			fmt.Printf("Total biaya: Rp%d\n", A[index].harga)
			fmt.Println(".")
			fmt.Println("--Pemesanan berhasil!--")
			fmt.Println(".")
		} else {
			fmt.Println("Penerbangan dengan ID tersebut tidak ditemukan.")
			fmt.Println(".")
		}
	}
}

func validasiTanggal(tanggal, bulan, tahun int) bool {
	// Lakukan validasi tanggal, bulan, dan tahun
	// Contoh sederhana: tanggal harus berada dalam rentang 1-31, bulan 1-12, tahun > 0
	var valid bool
	valid = true
	if tanggal < 1 || tanggal > 31 || bulan < 1 || bulan > 12 || tahun <= 0 {
		valid = false
	}
	return valid
}

func editPenerbangan(A *tabPenerbangan, n int) {
	var pilihan int
	var IdP string
	var index int
	var idBaru string
	var idTaken bool
	var slash string
	var i int

	fmt.Println("Masukkan ID Penerbangan yang ingin diedit: ")
	//input id penerbangan yang ingin diedit
	fmt.Scan(&IdP)

	//melakukan binary search untuk mencari index penerbangan yang ingin di edit
	index = binarySearch(*A, n, IdP)

	if index != -1 {
		fmt.Println("===============================================")
		fmt.Println("Pilih data yang ingin di edit:")
		fmt.Println("===============================================")
		fmt.Println("1. Edit Id Penerbangan")
		fmt.Println("2. Edit maskapai.")
		fmt.Println("3. Edit Asal.")
		fmt.Println("4. Edit Tujuan.")
		fmt.Println("5. Edit Harga.")
		fmt.Println("6. Edit Tanggal Keberangkatan.")
		fmt.Println("7. Edit Seluruh Data.")

		fmt.Scan(&pilihan)
		if pilihan == 1 {
			//idTaken menjadi true di awal tiap iterasi
			idTaken = true
			for idTaken {
				fmt.Println("Masukkan ID Penerbangan Baru:")
				//input id penerbangan baru
				fmt.Scan(&idBaru)
				//idTaken menjadi false dengan asumsi id selanjutnya unik
				idTaken = false
				for i = 0; i < index; i++ {
					if A[i].idPenerbangan == idBaru {
						fmt.Println("Id Penerbangan sudah terdaftar. Silahkan masukkan Id Penerbangan lain")
						//kembali menjalankan loop sampai id unik
						idTaken = true
					}
				}
			}
			//assign inputan id penerbangan yang baru ke dalam array
			A[index].idPenerbangan = idBaru
		} else if pilihan == 2 {
			fmt.Println("Masukkan maskapai:")
			fmt.Scan(&A[index].maskapai)
		} else if pilihan == 3 {
			fmt.Println("Masukkan asal:")
			fmt.Scan(&A[index].asal)
		} else if pilihan == 4 {
			fmt.Println("Masukkan tujuan:")
			fmt.Scan(&A[index].tujuan)
		} else if pilihan == 5 {
			fmt.Println("Masukkan harga tiket:")
			fmt.Scan(&A[index].harga)
		} else if pilihan == 6 {
			fmt.Println("Masukkan tanggal keberangkatan:")
			fmt.Scan(&A[index].tanggalKeberangkatan.tanggal, &slash, &A[index].tanggalKeberangkatan.bulan, &slash, &A[index].tanggalKeberangkatan.tahun)
		} else if pilihan == 7 {
			//idTaken menjadi true di awal tiap iterasi
			idTaken = true
			for idTaken {
				fmt.Println("Masukkan ID Penerbangan:")
				fmt.Scan(&idBaru)
				//idTaken menjadi false dengan asumsi id selanjutnya unik
				idTaken = false
				for i = 0; i < index; i++ {
					if A[i].idPenerbangan == idBaru {
						fmt.Println("Id Penerbangan sudah terdaftar. Silahkan masukkan Id Penerbangan lain")
						//kembali menjalankan loop sampai id unik
						idTaken = true
					}
				}
			}
			//assign idBaru ke dalam array
			A[index].idPenerbangan = idBaru
			fmt.Println("Masukkan data penerbangan:")
			fmt.Scan(&A[index].maskapai, &A[index].asal, &A[index].tujuan, &A[index].harga, &A[index].tanggalKeberangkatan.tanggal, &slash, &A[index].tanggalKeberangkatan.bulan, &slash, &A[index].tanggalKeberangkatan.tahun)
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
		fmt.Println("Data berhasil diedit.")
		fmt.Println(".")
	} else {
		fmt.Println("Penerbangan dengan ID tersebut tidak ditemukan.")
		fmt.Println(".")
	}
}

func hapusPenerbangan(A *tabPenerbangan, n *int) {
	var IdP string
	var index int
	var i int

	//base case saat data penerbangan 0
	if *n == 0 {
		fmt.Println("Belum ada penerbangan yang tersedia.")
		fmt.Println(".")
	} else {
		//print penerbangan yang tersedia
		cetakPenerbangan(*A, *n)
		fmt.Println("Masukkan ID Penerbangan yang ingin dihapus: ")
		fmt.Scan(&IdP)

		//melakukan binary search untuk mencari index yang ingin dihapus
		index = binarySearch(*A, *n, IdP)
		if index != -1 {
			for i = index; i < *n-1; i++ {
				//mengupdate nilai index i dengan index berikutnya i+1
				A[i] = A[i+1]
			}
			//mengurangi panjang array setelah dihapus
			*n--
			fmt.Println("Data penerbangan berhasil dihapus.")
			fmt.Println(".")
		} else {
			//kondisi setelah pencarian id penerbangan tidak ditemukan
			fmt.Println("Penerbangan dengan ID tersebut tidak ditemukan.")
			fmt.Println(".")
		}
	}
}

func binarySearch(A tabPenerbangan, n int, IdP string) int {
	//melakukan pencarian untuk id penerbangan yang diinginkan dengan algoritma binary search
	var left int
	var right int
	var mid int

	left = 0      //memulai dari index paling kiri
	right = n - 1 //batas index paling kanan

	for left <= right {
		//assign nilai tengah sebagai index yang akan dikembalikan
		mid = (left + right) / 2
		if A[mid].idPenerbangan == IdP {
			return mid
		} else if A[mid].idPenerbangan < IdP {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func insertionSort(A *tabPenerbangan, n int) {
	//mengurutkan jadwal penerbangan berdasarkan tanggal, bulan, dan tahun
	var i, pass int
	var temp penerbangan

	for pass = 1; pass < n; pass++ {
		temp = A[pass]
		i = pass - 1
		for i >= 0 && compareTanggal(A[i].tanggalKeberangkatan, temp.tanggalKeberangkatan) > 0 {
			A[i+1] = A[i]
			i--
		}
		A[i+1] = temp
	}
}

func compareTanggal(t1, t2 struct{ tanggal, bulan, tahun int }) int {
	//melakukan komparasi nilai yang akan di sorting
	//mengembalikan nilai -1 jika index sebelumnya lebih kecil dari index yang di cek
	var sortValue int

	//melakukan pengecekan tahun
	if t1.tahun < t2.tahun {
		sortValue = -1
	} else if t1.tahun > t2.tahun {
		sortValue = 1
	} else {
		//jika kedua tahun sama maka dilakukan pengecekan bulan
		if t1.bulan < t2.bulan {
			sortValue = -1
		} else if t1.bulan > t2.bulan {
			sortValue = 1
		} else {
			//jika kedua bulan sama maka dilakukan pengecekan tanggal
			if t1.tanggal < t2.tanggal {
				sortValue = -1
			} else if t1.tanggal > t2.tanggal {
				sortValue = 1
			} else {
				//jika kedua tanggal sama mengembalikan 0
				sortValue = 0
			}
		}
	}
	return sortValue
}
