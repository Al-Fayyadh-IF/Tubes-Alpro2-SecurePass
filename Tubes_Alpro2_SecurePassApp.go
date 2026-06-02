package main

import (
	"fmt"
	"os"
	"os/exec"
)

const NMAX = 100

type akun struct {
	namaLayanan, email, password, namaPengguna, tglUpdate string
	idInput                                               int
}

var data [NMAX]akun
var n, nextID int

func clearScreen() {
	var cmd *exec.Cmd
	cmd = exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func menu() {
	fmt.Println("============ SecurePass ============")
	fmt.Println("1. Tambah Akun")
	fmt.Println("2. Ubah Akun")
	fmt.Println("3. Hapus Akun")
	fmt.Println("4. Cari Akun")
	fmt.Println("5. Urutkan Berdasarkan Nama")
	fmt.Println("6. Urutkan Berdasarkan Waktu Input")
	fmt.Println("7. Statistik")
	fmt.Println("0. Keluar")
	fmt.Println("====================================")
}

func tambahAkun() {
	var akunBaru akun
	clearScreen()
	if n >= NMAX {
		fmt.Println("Data akun sudah penuh, tidak bisa tambah lagi")
		menu()
	}
	fmt.Println()
	fmt.Println("--- Tambah Akun Baru ---")
	fmt.Print("Nama layanan               : ")
	fmt.Scan(&akunBaru.namaLayanan)
	fmt.Print("Email                      : ")
	fmt.Scan(&akunBaru.email)
	fmt.Print("Password                   : ")
	fmt.Scan(&akunBaru.password)
	fmt.Print("Nama pengguna              : ")
	fmt.Scan(&akunBaru.namaPengguna)
	fmt.Print("Tanggal update (dd-mm-yyyy): ")
	fmt.Scan(&akunBaru.tglUpdate)
	akunBaru.idInput = nextID
	data[n] = akunBaru
	nextID++
	n++
	fmt.Println("== Akun berhasil ditambahkan ==")
}

func ubahAkun() {
	var nL, email, emailBaru, pwBaru, npBaru, tglBaru string
	var idx, jumlah int
	clearScreen()
	if n == 0 {
		fmt.Println("Belum ada data akun")
		clearScreen()
		menu()
	}
	fmt.Print("Nama layanan yang ingin diubah: ")
	fmt.Scan(&nL)
	jumlah = banyakAkunSatuLayanan(nL)
	if jumlah == 0 {
		fmt.Println("Akun tidak ditemukan")
		clearScreen()
		menu()
	} else if jumlah == 1 {
		idx = sequentialSearchL(nL)
		fmt.Println()
		fmt.Println("Data lama:")
		tampilkanAkun(idx)
		fmt.Println()
		fmt.Println("Masukkan data baru:")
		fmt.Print("Email baru                 : ")
		fmt.Scan(&emailBaru)
		fmt.Print("Password baru              : ")
		fmt.Scan(&pwBaru)
		fmt.Print("Nama pengguna baru         : ")
		fmt.Scan(&npBaru)
		fmt.Print("Tanggal update (dd-mm-yyyy): ")
		fmt.Scan(&tglBaru)
		data[idx].email = emailBaru
		data[idx].password = pwBaru
		data[idx].namaPengguna = npBaru
		data[idx].tglUpdate = tglBaru
		fmt.Println("== Akun berhasil diubah ==")
	} else if jumlah > 1 {
		idx = -1
		for idx == -1 {
			fmt.Print("Email akun yang ingin diubah: ")
			fmt.Scan(&email)
			idx = sequentialSearchE(email)
			if idx == -1 {
				fmt.Println("Email tidak ditemukan, coba lagi.")
			}
		}
		fmt.Println()
		fmt.Println("Data lama:")
		tampilkanAkun(idx)
		fmt.Println()
		fmt.Println("Masukkan data baru:")
		fmt.Print("Email baru                 : ")
		fmt.Scan(&emailBaru)
		fmt.Print("Password baru              : ")
		fmt.Scan(&pwBaru)
		fmt.Print("Nama pengguna baru         : ")
		fmt.Scan(&npBaru)
		fmt.Print("Tanggal update (dd-mm-yyyy): ")
		fmt.Scan(&tglBaru)
		data[idx].email = emailBaru
		data[idx].password = pwBaru
		data[idx].namaPengguna = npBaru
		data[idx].tglUpdate = tglBaru
		fmt.Println("== Akun berhasil diubah ==")
	}
}

func hapusAkun() {
	var nL, email string
	var i, idx, jumlah int
	clearScreen()
	if n == 0 {
		fmt.Println("Belum ada data akun")
		clearScreen()
		menu()
	}
	fmt.Print("Nama layanan yang ingin dihapus: ")
	fmt.Scan(&nL)
	jumlah = banyakAkunSatuLayanan(nL)
	if jumlah == 0 {
		fmt.Println("Akun tidak ditemukan")
		clearScreen()
		menu()
	} else if jumlah == 1 {
		idx = sequentialSearchL(nL)
		for i = idx; i < n-1; i++ {
			data[i] = data[i+1]
		}
		n -= 1
		fmt.Println("== Akun berhasil dihapus ==")
	} else if jumlah > 1 {
		idx = -1
		for idx == -1 {
			fmt.Print("Email akun yang ingin diubah: ")
			fmt.Scan(&email)
			idx = sequentialSearchE(email)
			if idx == -1 {
				fmt.Println("Email tidak ditemukan, coba lagi.")
			}
		}
		for i = idx; i < n-1; i++ {
			data[i] = data[i+1]
		}
		n -= 1
		fmt.Println("== Akun berhasil dihapus ==")
	}
}

func cekPassword(pass string) string {
	var i int
	var c byte
	var huruf, angka, simbol bool
	for i = 0; i < len(pass); i++ {
		c = pass[i]
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			huruf = true
		} else if c >= '0' && c <= '9' {
			angka = true
		} else {
			simbol = true
		}
	}
	if len(pass) >= 8 && huruf && angka && simbol {
		return "Kuat"
	} else if len(pass) >= 6 && huruf && angka {
		return "Sedang"
	} else {
		return "Lemah"
	}
}

func tampilkanAkun(i int) {
	fmt.Printf("Nama Layanan  : %s\n", data[i].namaLayanan)
	fmt.Printf("Email         : %s\n", data[i].email)
	fmt.Printf("Password      : %s [%s]\n", data[i].password, cekPassword(data[i].password))
	fmt.Printf("Nama Pengguna : %s\n", data[i].namaPengguna)
	fmt.Printf("Update        : %s\n", data[i].tglUpdate)
}

func tampilkanSemuaAkun() {
	var i int
	clearScreen()
	if n == 0 {
		fmt.Println("Belum ada data akun yang tersimpan")
		clearScreen()
		menu()
	} else {
		fmt.Println()
		fmt.Println("--- Daftar Akun ---")
		for i = 0; i < n; i++ {
			fmt.Printf("Akun ke-%d\n", i+1)
			tampilkanAkun(i)
			fmt.Println()
		}
	}
}

func banyakAkunSatuLayanan(cari string) int {
	var i, jumlah int
	for i = 0; i < n; i++ {
		if data[i].namaLayanan == cari {
			jumlah++
		}
	}
	return jumlah
}

func sequentialSearchL(cari string) int {
	var i int
	for i = 0; i < n; i++ {
		if data[i].namaLayanan == cari {
			return i
		}
	}
	return -1
}

func sequentialSearchE(cari string) int {
	var i int
	for i = 0; i < n; i++ {
		if data[i].email == cari {
			return i
		}
	}
	return -1
}

func binarySearch(cari string) int {
	var kiri, kanan, tengah, found int
	kiri = 0
	kanan = n - 1
	found = -1
	for kiri <= kanan && found == -1 {
		tengah = (kiri + kanan) / 2
		if cari > data[tengah].namaLayanan {
			kiri = tengah + 1
		} else if cari < data[tengah].namaLayanan {
			kanan = tengah - 1
		} else {
			found = tengah
		}
	}
	return found
}

func binarySearchAw(cari string, posisi int) int {
	var kiri, kanan, tengah, found int
	kiri = 0
	kanan = posisi
	found = posisi
	for kiri <= kanan {
		tengah = (kiri + kanan) / 2
		if data[tengah].namaLayanan == cari {
			found = tengah
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}
	return found
}

func binarySearchAk(cari string, posisi int) int {
	var kiri, kanan, tengah, found int
	kiri = posisi
	kanan = n - 1
	found = posisi
	for kiri <= kanan {
		tengah = (kiri + kanan) / 2
		if data[tengah].namaLayanan == cari {
			found = tengah
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return found
}

func cariAkun() {
	var pilih int
	var jln bool
	clearScreen()
	jln = false
	for jln == false {
		fmt.Println("1. Sequential Search")
		fmt.Println("2. Binary Search")
		fmt.Println("====================================")
		fmt.Print("Pilih metode: ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			cariAkunSeq()
			jln = true
		} else if pilih == 2 {
			cariAkunBin()
			jln = true
		} else {
			clearScreen()
			fmt.Println("Pilihan tidak tersedia. Silakan pilih lagi.")
			fmt.Println()
		}
	}
}

func cariAkunSeq() {
	var nama string
	var i, jumlah int
	clearScreen()
	if n == 0 {
		fmt.Println("Belum ada data akun")
		clearScreen()
		menu()
	}
	clearScreen()
	fmt.Print("Nama layanan yang dicari: ")
	fmt.Scan(&nama)
	jumlah = banyakAkunSatuLayanan(nama)
	if jumlah == 0 {
		fmt.Println("Akun tidak ditemukan")
	} else {
		fmt.Println("== Akun ditemukan! ==")
		for i = 0; i < n; i++ {
			if data[i].namaLayanan == nama {
				tampilkanAkun(i)
				fmt.Println()
			}
		}
	}
}

func cariAkunBin() {
	var nama string
	var i, jumlah, pT, pAw, pAk int
	clearScreen()
	if n == 0 {
		fmt.Println("Belum ada data akun")
		clearScreen()
		menu()
	}
	clearScreen()
	selectionSort()
	fmt.Print("Nama layanan yang dicari: ")
	fmt.Scan(&nama)
	jumlah = banyakAkunSatuLayanan(nama)
	if jumlah == 0 {
		fmt.Println("Akun tidak ditemukan")
	} else {
		fmt.Println("== Akun ditemukan! ==")
		pT = binarySearch(nama)
		pAw = binarySearchAw(nama, pT)
		pAk = binarySearchAk(nama, pT)
		for i = pAw; i <= pAk; i++ {
			tampilkanAkun(i)
			fmt.Println()
		}
	}
}

func selectionSort() {
	var pass, idx, i int
	var temp akun
	pass = 1
	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if data[idx].namaLayanan > data[i].namaLayanan {
				idx = i
			}
			i = i + 1
		}
		temp = data[pass-1]
		data[pass-1] = data[idx]
		data[idx] = temp
		pass = pass + 1
	}
}

func insertionSort() {
	var pass, i int
	var temp akun
	pass = 1
	for pass <= n-1 {
		i = pass
		temp = data[pass]
		for i > 0 && temp.idInput < data[i-1].idInput {
			data[i] = data[i-1]
			i = i - 1
		}
		data[i] = temp
		pass = pass + 1
	}
}

func urutkanNama() {
	clearScreen()
	if n == 0 {
		fmt.Println("Belum ada data akun")
		clearScreen()
		menu()
	}
	selectionSort()
	fmt.Println("== Data berhasil diurutkan alfabetis ==")
	tampilkanSemuaAkun()
}

func urutkanInput() {
	clearScreen()
	if n == 0 {
		fmt.Println("Belum ada data akun")
		clearScreen()
		menu()
	}
	insertionSort()
	fmt.Println("== Data berhasil diurutkan berdasarkan waktu input ==")
	tampilkanSemuaAkun()
}

func statistik() {
	var i, lemah, sedang, kuat int
	var kategori string
	clearScreen()
	fmt.Println("--- Statistik Akun ---")
	fmt.Printf("Total akun tersimpan: %d\n", n)
	if n == 0 {
		clearScreen()
		menu()
	}
	for i = 0; i < n; i++ {
		kategori = cekPassword(data[i].password)
		if kategori == "Lemah" {
			lemah = lemah + 1
		} else if kategori == "Sedang" {
			sedang = sedang + 1
		} else {
			kuat = kuat + 1
		}
	}
	fmt.Println("Klasifikasi kekuatan password:")
	fmt.Printf("  Lemah  : %d akun\n", lemah)
	fmt.Printf("  Sedang : %d akun\n", sedang)
	fmt.Printf("  Kuat   : %d akun\n", kuat)
}

func isiDataDummy() {
	data[0] = akun{namaLayanan: "Roblox", email: "fayyadh@gmail.com", password: "Fayyadh123*", namaPengguna: "fayyadhgaming", tglUpdate: "26-05-2026", idInput: 0}
	data[1] = akun{namaLayanan: "Instagram", email: "fatih@gmail.com", password: "fatih123", namaPengguna: "fatih", tglUpdate: "27-05-2026", idInput: 1}
	data[2] = akun{namaLayanan: "Gmail", email: "rijal@gmail.com", password: "pass", namaPengguna: "rijalaslam", tglUpdate: "28-05-2026", idInput: 2}
	data[3] = akun{namaLayanan: "Discord", email: "rehan@gmail.com", password: "Reh@n2025", namaPengguna: "rehanwangsap", tglUpdate: "28-05-2026", idInput: 3}
	data[4] = akun{namaLayanan: "Spotify", email: "rafi@gmail.com", password: "musik", namaPengguna: "rafiSapu", tglUpdate: "29-05-2026", idInput: 4}
	data[5] = akun{namaLayanan: "Github", email: "ghifary@gmail.com", password: "Ghif4ry#", namaPengguna: "ghifaryyes", tglUpdate: "30-05-2026", idInput: 5}
	data[6] = akun{namaLayanan: "X", email: "bagas@gmail.com", password: "bagas123", namaPengguna: "bagasjimat", tglUpdate: "31-05-2026", idInput: 6}
	data[7] = akun{namaLayanan: "Tokopedia", email: "rizky@gmail.com", password: "R1zky@Top", namaPengguna: "rizkyplay", tglUpdate: "01-06-2026", idInput: 7}
	data[8] = akun{namaLayanan: "Shopee", email: "dimas@gmail.com", password: "dimas", namaPengguna: "dimasmmamang", tglUpdate: "02-06-2026", idInput: 8}
	data[9] = akun{namaLayanan: "Netflix", email: "arya@gmail.com", password: "Arya99!Nflx", namaPengguna: "aryaSMH", tglUpdate: "02-06-2026", idInput: 9}
	n = 10
	nextID = 10
}

func main() {
	var pilih int
	var jln bool
	jln = true
	isiDataDummy()
	fmt.Println("Selamat datang di SecurePass!")
	for jln == true {
		menu()
		fmt.Print("Pilih menu: ")
		fmt.Scan(&pilih)
		if pilih == 0 {
			fmt.Println("Terima kasih sudah menggunakan SecurePass. Sampai jumpa!")
			jln = false
		} else if pilih == 1 {
			tambahAkun()
		} else if pilih == 2 {
			ubahAkun()
		} else if pilih == 3 {
			hapusAkun()
		} else if pilih == 4 {
			cariAkun()
		} else if pilih == 5 {
			urutkanNama()
		} else if pilih == 6 {
			urutkanInput()
		} else if pilih == 7 {
			statistik()
		} else {
			fmt.Println("Pilihan tidak tersedia. Silakan pilih lagi.")
		}
	}
}
