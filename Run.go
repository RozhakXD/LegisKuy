package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Fungsi Membersihkan Terminal
func Bersihkan_Terminal() {
	var clear_cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		clear_cmd = exec.Command("cmd", "/c", "cls")
	default:
		clear_cmd = exec.Command("clear")
	}
	clear_cmd.Stdout = os.Stdout
	clear_cmd.Run()
}

// Deklarasikan Struct Calon
type Pemilih struct {
	nama         string
	sudahMemilih bool
}

// Deklarasikan Struct Pemilih
type Calon struct {
	nama   string
	partai string
	suara  int
}

const MAX_PEMILIH = 100
const MAX_CALON = 100

var calon_array [MAX_CALON]Calon
var pemilih_array [MAX_PEMILIH]Pemilih

// Deklarasikan Struct DataPemilu
type DataPemilu struct {
	jumlah_calon   int
	jumlah_pemilih int
}

// Fungsi Menambahkan Calon
func Tambah_Calon(nama, partai string, data *DataPemilu) {
	if data.jumlah_calon < MAX_CALON {
		calon_array[data.jumlah_calon] = Calon{nama: nama, partai: partai, suara: 0}
		data.jumlah_calon++
	} else {
		fmt.Println(Panel("[!] Jumlah Calon Sudah Mencapai Batas Maksimum.", "1", 1, 2))
	}
}

// Fungsi Mengubah Data Calon
func Ubah_Calon(index int, nama_baru, partai_baru string, data DataPemilu) {
	if index >= 0 && index < data.jumlah_calon {
		calon_array[index].nama = nama_baru
		calon_array[index].partai = partai_baru
	} else {
		fmt.Println(Panel("[!] Index Tidak Valid.", "1", 1, 2))
	}
}

// Fungsi Menghapus Calon
func Hapus_Calon(index int, data *DataPemilu) {
	if index >= 0 && index < data.jumlah_calon {
		for i := index; i < data.jumlah_calon-1; i++ {
			calon_array[i] = calon_array[i+1]
		}
		data.jumlah_calon--
	} else {
		fmt.Println(Panel("[!] Index Tidak Valid.", "1", 1, 2))
	}
}

// Fungsi Tambah Data Pemilih
func Tambah_Pemilih(nama string, data *DataPemilu) {
	if data.jumlah_pemilih < MAX_PEMILIH {
		pemilih_array[data.jumlah_pemilih] = Pemilih{nama: nama, sudahMemilih: false}
		data.jumlah_pemilih++
	} else {
		fmt.Println(Panel("[!] Jumlah Pemilih Sudah Mencapai Batas Maksimum.", "1", 1, 2))
	}
}

// Fungsi Menampilkan Data Calon
func Tampilkan_Calon(data DataPemilu) {
	if data.jumlah_calon == 0 {
		fmt.Println(Panel("[!] Tidak Ada Calon Yang Terdaftar.", "1", 1, 2))
		return
	}

	var tabel_data string
	tabel_data = "\nDaftar Calon Terdaftar:\n"
	for i := 0; i < data.jumlah_calon; i++ {
		tabel_data += fmt.Sprintf("[%d]. %s - %s\n", i, calon_array[i].nama, calon_array[i].partai)
	}

	fmt.Println(Panel_Tabel(tabel_data, "2", 1, 2))
}

// Fungsi Menampilkan Data Pemilih
func Tampilkan_Pemilih(data DataPemilu) {
	if data.jumlah_pemilih == 0 {
		fmt.Println(Panel("[!] Tidak Ada Pemilih yang Terdaftar.", "1", 1, 2))
		return
	}

	var tabel_data string
	tabel_data = "\nDaftar Pemilih Terdaftar:\n"
	for i := 0; i < data.jumlah_pemilih; i++ {
		tabel_data += fmt.Sprintf("[%d]. %s\n", i, pemilih_array[i].nama)
	}

	fmt.Println(Panel_Tabel(tabel_data, "2", 1, 2))
}

// Fitur Manajemen Data
func Manajemen_Data(data *DataPemilu) {
	for {
		var pilihan_kamu int

		fmt.Println(Panel("[1]. Tambah Calon Legislatif\n[2]. Edit Data Calon Legislatif\n[3]. Hapus Calon Legislatif\n[4]. Tambah Data Pemilih\n[5]. Lihat Daftar Calon Legislatif\n[6]. Lihat Daftar Pemilih\n[7]. Kembali ke Fitur Utama", "7", 0, 2))

		fmt.Print("[?] Pilihan Kamu: ")
		fmt.Scanln(&pilihan_kamu)

		switch pilihan_kamu {
		case 1:
			var nama, partai string

			fmt.Print("[?] Nama Calon: ")
			reader := bufio.NewReader(os.Stdin)
			nama, _ = reader.ReadString('\n')
			nama = strings.TrimSpace(nama)

			fmt.Print("[?] Partai: ")
			partai, _ = reader.ReadString('\n')
			partai = strings.TrimSpace(partai)

			Tambah_Calon(nama, partai, data)

			fmt.Println(Panel("[*] Calon Berhasil Ditambahkan.", "2", 1, 2))
			Mengulang_Program()
		case 2:
			var index int
			var nama_baru, partai_baru string

			if data.jumlah_calon == 0 {
				fmt.Println(Panel("[!] Tidak Ada Calon Yang Terdaftar.", "1", 1, 2))
				Mengulang_Program()
				continue
			}
			Tampilkan_Calon(*data)

			fmt.Print("[?] Index Calon yang Ingin Diubah: ")
			fmt.Scanln(&index)

			fmt.Print("[?] Nama Baru: ")
			reader := bufio.NewReader(os.Stdin)
			nama_baru, _ = reader.ReadString('\n')
			nama_baru = strings.TrimSpace(nama_baru)

			fmt.Print("[?] Partai Baru: ")
			partai_baru, _ = reader.ReadString('\n')
			partai_baru = strings.TrimSpace(partai_baru)
			Ubah_Calon(index, nama_baru, partai_baru, *data)

			fmt.Println(Panel("[*] Data Calon Berhasil Diubah.", "2", 1, 2))
			Mengulang_Program()
		case 3:
			var index int

			if data.jumlah_calon == 0 {
				fmt.Println(Panel("[!] Tidak Ada Calon Yang Terdaftar.", "1", 1, 2))
				Mengulang_Program()
				continue
			}
			Tampilkan_Calon(*data)

			fmt.Print("[?] Index Calon yang Ingin Dihapus: ")
			fmt.Scanln(&index)
			Hapus_Calon(index, data)

			fmt.Println(Panel("[*] Data Calon Berhasil Dihapus.", "2", 1, 2))
			Mengulang_Program()
		case 4:
			var nama string

			fmt.Print("[?] Nama Pemilih: ")
			reader := bufio.NewReader(os.Stdin)
			nama, _ = reader.ReadString('\n')
			nama = strings.TrimSpace(nama)

			Tambah_Pemilih(nama, data)

			fmt.Println(Panel("[*] Pemilih Berhasil Ditambahkan.", "2", 1, 2))
			Mengulang_Program()
		case 5:
			if data.jumlah_calon == 0 {
				fmt.Println(Panel("[!] Tidak Ada Calon Yang Terdaftar.", "1", 1, 2))
				Mengulang_Program()
				continue
			}

			Tampilkan_Calon(*data)
			Mengulang_Program()
		case 6:
			if data.jumlah_pemilih == 0 {
				fmt.Println(Panel("[!] Tidak Ada Pemilih Yang Terdaftar.", "1", 1, 2))
				Mengulang_Program()
				continue
			}

			Tampilkan_Pemilih(*data)
			Mengulang_Program()
		case 7:
			return
		default:
			fmt.Println(Panel("[!] Pilihan Yang Anda Masukkan Tidak Valid. Silakan Coba Lagi.", "1", 1, 2))
			Mengulang_Program()
		}
	}
}

// Fungsi Memilih Calon Legislatif
func Pilih_Calon(index_pemilih, index_calon int, data DataPemilu, pemilu_berlangsung bool) {
	if !pemilu_berlangsung {
		fmt.Println(Panel("[!] Pemilu Belum Dimulai atau Sudah Selesai.", "1", 1, 2))
		return
	}

	if index_pemilih < 0 || index_pemilih >= data.jumlah_pemilih {
		fmt.Println(Panel("[!] Index Pemilih Tidak Valid.", "1", 1, 2))
		return
	}

	if index_calon < 0 || index_calon >= data.jumlah_calon {
		fmt.Println(Panel("[!] Index Calon Tidak Valid.", "1", 1, 2))
		return
	}

	if !pemilih_array[index_pemilih].sudahMemilih {
		pemilih_array[index_pemilih].sudahMemilih = true
		calon_array[index_calon].suara++
		fmt.Println(Panel("[*] Suara Berhasil Diberikan.", "2", 1, 2))
	} else {
		fmt.Println(Panel("[!] Pemilih Sudah Memberikan Suara.", "1", 1, 2))
		return
	}
}

// Fungsi Lihat Hasil Sementara Dan Akhir
func Lihat_Hasil_Sementara_Akhir(data DataPemilu, message string) {
	if data.jumlah_calon == 0 {
		fmt.Println(Panel("[!] Tidak Ada Calon Yang Terdaftar.", "1", 1, 2))
		Mengulang_Program()
		return
	}
	var tabel_data string
	tabel_data = fmt.Sprintf("\nDaftar Hasil %s Pemilu:\n", message)
	for i := 0; i < data.jumlah_calon; i++ {
		tabel_data += fmt.Sprintf("[%d]. %s - %s: %d Suara\n", i, calon_array[i].nama, calon_array[i].partai, calon_array[i].suara)
	}
	fmt.Println(Panel_Tabel(tabel_data, "2", 1, 2))
}

// Fungsi Memproses Pemilu
func Proses_Pemilu(berlangsung bool, data DataPemilu) {
	for {
		var pilihan_kamu int

		// Pilih Calon Legislatif = Hanya Untuk Pemilih
		// Lihat Hasil Sementara = Hanya Untuk Petugas KPU
		fmt.Println(Panel("[1]. Pilih Calon Legislatif\n[2]. Lihat Hasil Sementara\n[3]. Lihat Hasil Akhir\n[4]. Kembali ke Fitur Utama", "7", 0, 2))

		fmt.Print("[?] Pilihan Kamu: ")
		fmt.Scanln(&pilihan_kamu)

		switch pilihan_kamu {
		case 1:
			if !berlangsung {
				fmt.Println(Panel("[!] Pemilu Belum Dimulai atau Sudah Selesai.", "1", 1, 2))
				Mengulang_Program()
				return
			}
			if data.jumlah_calon == 0 {
				fmt.Println(Panel("[!] Tidak Ada Calon Yang Terdaftar.", "1", 1, 2))
				Mengulang_Program()
				continue
			}
			if data.jumlah_pemilih == 0 {
				fmt.Println(Panel("[!] Tidak Ada Pemilih Yang Terdaftar.", "1", 1, 2))
				Mengulang_Program()
				continue
			}

			var index_pemilih, index_calon int

			Tampilkan_Pemilih(data)
			Tampilkan_Calon(data)

			fmt.Print("[?] Index Pemilih: ")
			fmt.Scanln(&index_pemilih)

			fmt.Print("[?] Index Calon: ")
			fmt.Scanln(&index_calon)

			Pilih_Calon(index_pemilih, index_calon, data, berlangsung)
			Mengulang_Program()
		case 2:
			if !berlangsung {
				fmt.Println(Panel("[!] Pemilu Belum Dimulai atau Sudah Selesai.", "1", 1, 2))
				Mengulang_Program()
				return
			}
			Lihat_Hasil_Sementara_Akhir(data, "Sementara")
			Mengulang_Program()
		case 3:
			if berlangsung {
				fmt.Println(Panel("[!] Pemilu Belum Selesai.", "1", 1, 2))
				Mengulang_Program()
				continue
			}
			Lihat_Hasil_Sementara_Akhir(data, "Akhir")
			Mengulang_Program()
		case 4:
			return
		default:
			fmt.Println(Panel("[!] Pilihan Yang Anda Masukkan Tidak Valid. Silakan Coba Lagi.", "1", 1, 2))
			Mengulang_Program()
		}
	}
}

// Fungsi Sequential Search Berdasarkan Nama
func Sequential_Search_By_Name(nilai string, data DataPemilu) {
	var ditemukan bool
	var tabel_data string
	tabel_data = "\n"
	for i := 0; i < data.jumlah_calon; i++ {
		if strings.Contains(strings.ToLower(calon_array[i].nama), strings.ToLower(nilai)) {
			tabel_data += fmt.Sprintf("[%d]. %s - %s\n", i, calon_array[i].nama, calon_array[i].partai)
			ditemukan = true
		}
	}

	if !ditemukan {
		fmt.Println(Panel("[!] Tidak Ada Calon Yang Ditemukan.", "1", 1, 2))
	} else {
		fmt.Println(Panel_Tabel(tabel_data, "2", 1, 2))
	}
}

// Fungsi Binary Search Berdasarkan Partai
func Binary_Search_By_Party(nilai string, data DataPemilu) {
	Insertion_Sort_By_Party("ascending", data)

	left, right := 0, data.jumlah_calon-1
	found := false
	var tabel_data string
	tabel_data = "\n"

	for left <= right {
		mid := (left + right) / 2
		if strings.EqualFold(calon_array[mid].partai, nilai) {
			tabel_data += fmt.Sprintf("[%d]. %s - %s\n", mid, calon_array[mid].nama, calon_array[mid].partai)
			found = true
			break
		} else if strings.ToLower(calon_array[mid].partai) < strings.ToLower(nilai) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if !found {
		fmt.Println(Panel("[!] Tidak Ada Calon Yang Ditemukan.", "1", 1, 2))
	} else {
		fmt.Println(Panel_Tabel(tabel_data, "2", 1, 2))
	}
}

// Fungsi Sequential Search Berdasarkan Nama Pemilih
func Sequential_Search_By_Pemilih(nilai string, data DataPemilu) {
	var tabel_data string
	var ditemukan bool
	tabel_data = "\n"
	for i := 0; i < data.jumlah_pemilih; i++ {
		if strings.Contains(strings.ToLower(pemilih_array[i].nama), strings.ToLower(nilai)) {
			tabel_data += fmt.Sprintf("[%d]. %s\n", i, pemilih_array[i].nama)
			ditemukan = true
		}
	}

	if !ditemukan {
		fmt.Println(Panel("[!] Tidak Ada Pemilih Yang Ditemukan.", "1", 1, 2))
	} else {
		fmt.Println(Panel_Tabel(tabel_data, "2", 1, 2))
	}
}

// Fungsi Pencarian Calon Berdasarkan Nama, Partai, Dan Pemilih
func Pencarian_Calon_Dan_Pemilih(kriteria, nilai string, data DataPemilu) {
	if kriteria == "pemilih" {
		Sequential_Search_By_Pemilih(nilai, data)
	} else if kriteria == "partai" {
		Binary_Search_By_Party(nilai, data) // Implementasi Binary Search
	} else {
		Sequential_Search_By_Name(nilai, data)
	}
}

// Fungsi Pencarian Data
func Pencarian_Data(data DataPemilu) {
	for {
		var pilihan_kamu int

		fmt.Println(Panel("[1]. Cari Calon Berdasarkan Nama\n[2]. Cari Calon Berdasarkan Partai\n[3]. Cari Pemilih Berdasarkan Nama\n[4]. Kembali ke Menu Utama", "7", 0, 2))

		fmt.Print("[?] Pilihan Kamu: ")
		fmt.Scanln(&pilihan_kamu)

		switch pilihan_kamu {
		case 1:
			if data.jumlah_calon == 0 {
				fmt.Println(Panel("[!] Tidak Ada Calon Yang Terdaftar.", "1", 1, 2))
				Mengulang_Program()
				return
			}
			var nama string

			fmt.Print("[?] Nama Calon: ")
			reader := bufio.NewReader(os.Stdin)
			nama, _ = reader.ReadString('\n')
			nama = strings.TrimSpace(nama)

			Pencarian_Calon_Dan_Pemilih("nama", nama, data)
			Mengulang_Program()
		case 2:
			if data.jumlah_calon == 0 {
				fmt.Println(Panel("[!] Tidak Ada Calon Yang Terdaftar.", "1", 1, 2))
				Mengulang_Program()
				return
			}
			var partai string

			fmt.Print("[?] Partai Calon: ")
			reader := bufio.NewReader(os.Stdin)
			partai, _ = reader.ReadString('\n')
			partai = strings.TrimSpace(partai)

			Pencarian_Calon_Dan_Pemilih("partai", partai, data)
			Mengulang_Program()
		case 3:
			if data.jumlah_pemilih == 0 {
				fmt.Println(Panel("[!] Tidak Ada Pemilih Yang Terdaftar.", "1", 1, 2))
				Mengulang_Program()
				return
			}
			var nama string

			fmt.Print("[?] Nama Pemilih: ")
			reader := bufio.NewReader(os.Stdin)
			nama, _ = reader.ReadString('\n')
			nama = strings.TrimSpace(nama)

			Pencarian_Calon_Dan_Pemilih("pemilih", nama, data)
			Mengulang_Program()
		case 4:
			return
		default:
			fmt.Println(Panel("[!] Pilihan Yang Anda Masukkan Tidak Valid. Silakan Coba Lagi.", "1", 1, 2))
			Mengulang_Program()
			continue
		}
	}
}

// Fungsi Mengurutkan Data Berdasarkan Jumlah Suara
func Insertion_Sort_By_Votes(urutan string, data DataPemilu) {
	for i := 1; i < data.jumlah_calon; i++ {
		key := calon_array[i]
		j := i - 1
		if urutan == "ascending" {
			for j >= 0 && calon_array[j].suara > key.suara {
				calon_array[j+1] = calon_array[j]
				j--
			}
		} else {
			for j >= 0 && calon_array[j].suara < key.suara {
				calon_array[j+1] = calon_array[j]
				j--
			}
		}
		calon_array[j+1] = key
	}
}

// Fungsi Mengurutkan Data Berdasarkan Nama
func Selection_Sort_By_Name(urutan string, data DataPemilu) {
	for i := 0; i < data.jumlah_calon-1; i++ {
		for j := i + 1; j < data.jumlah_calon; j++ {
			if urutan == "ascending" {
				if calon_array[i].nama > calon_array[j].nama {
					calon_array[i], calon_array[j] = calon_array[j], calon_array[i]
				}
			} else {
				if calon_array[i].nama < calon_array[j].nama {
					calon_array[i], calon_array[j] = calon_array[j], calon_array[i]
				}
			}
		}
	}
}

// Fungsi Mengurutkan Data Berdasarkan Partai
func Insertion_Sort_By_Party(urutan string, data DataPemilu) {
	for i := 1; i < data.jumlah_calon; i++ {
		key := calon_array[i]
		j := i - 1
		if urutan == "ascending" {
			for j >= 0 && calon_array[j].partai > key.partai {
				calon_array[j+1] = calon_array[j]
				j--
			}
		} else {
			for j >= 0 && calon_array[j].partai < key.partai {
				calon_array[j+1] = calon_array[j]
				j--
			}
		}
		calon_array[j+1] = key
	}
}

// Fungsi Mengurutkan Data Berdasarkan Jumlah Suara, Nama Partai, Dan Nama Calon
func Mengurutkan_Data_Pemilu(kriteria string, urutan string, data DataPemilu) {
	if kriteria == "suara" {
		Insertion_Sort_By_Votes(urutan, data)
	} else if kriteria == "nama" {
		Selection_Sort_By_Name(urutan, data) // Implementasi Selection Sort
	} else if kriteria == "partai" {
		Insertion_Sort_By_Party(urutan, data)
	}
}

// Fungsi Pengurutan Data
func Mengurutkan_Data(data DataPemilu) {
	for {
		var pilihan_kamu int

		fmt.Println(Panel("[1]. Berdasarkan Jumlah Suara\n[2]. Berdasarkan Nama Partai\n[3]. Berdasarkan Nama Calon\n[4]. Kembali ke Menu Utama", "7", 0, 2))

		fmt.Print("[?] Pilihan Kamu: ")
		fmt.Scanln(&pilihan_kamu)

		switch pilihan_kamu {
		case 1:
			if data.jumlah_calon == 0 {
				fmt.Println(Panel("[!] Tidak Ada Calon Yang Terdaftar.", "1", 1, 2))
				Mengulang_Program()
				return
			}
			var urutan string
			fmt.Print("[?] Urutkan (ascending/descending): ")
			fmt.Scanln(&urutan)

			Mengurutkan_Data_Pemilu("suara", urutan, data)
			Tampilkan_Calon(data)
			Mengulang_Program()
		case 2:
			if data.jumlah_calon == 0 {
				fmt.Println(Panel("[!] Tidak Ada Calon Yang Terdaftar.", "1", 1, 2))
				Mengulang_Program()
				return
			}

			var urutan string
			fmt.Print("[?] Urutkan (ascending/descending): ")
			fmt.Scanln(&urutan)

			Mengurutkan_Data_Pemilu("partai", urutan, data)
			Tampilkan_Calon(data)
			Mengulang_Program()
		case 3:
			if data.jumlah_calon == 0 {
				fmt.Println(Panel("[!] Tidak Ada Calon Yang Terdaftar.", "1", 1, 2))
				Mengulang_Program()
				return
			}

			var urutan string
			fmt.Print("[?] Urutkan (ascending/descending): ")
			fmt.Scanln(&urutan)

			Mengurutkan_Data_Pemilu("nama", urutan, data)
			Tampilkan_Calon(data)
			Mengulang_Program()
		case 4:
			return
		default:
			fmt.Println(Panel("[!] Pilihan Yang Anda Masukkan Tidak Valid. Silakan Coba Lagi.", "1", 1, 2))
			Mengulang_Program()
		}
	}
}

// Fungsi Tampilkan Berdasarkan Threshold
func Tampilkan_Berdasarkan_Threshold(threshold int, data DataPemilu) {
	if data.jumlah_calon == 0 {
		fmt.Println(Panel("[!] Tidak Ada Calon Yang Terdaftar.", "1", 1, 2))
		return
	}
	var tabel_data string
	tabel_data = "\n"
	for i := 0; i < data.jumlah_calon; i++ {
		if calon_array[i].suara >= threshold {
			tabel_data += fmt.Sprintf("[%d]. %s - %s: %d Suara\n", i, calon_array[i].nama, calon_array[i].partai, calon_array[i].suara)
		}
	}
	if tabel_data == "\n" {
		fmt.Println(Panel("[!] Tidak Ada Calon Yang Memenuhi Ambang Batas.", "1", 1, 2))
		return
	}
	fmt.Println(Panel_Tabel(tabel_data, "2", 1, 2))
}

// Fungsi Panel
func Panel(text string, color string, padding1 int, padding2 int) string {
	return lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(padding1, padding2).
		BorderForeground(lipgloss.Color("63")).
		Width(45).
		Foreground(lipgloss.Color(color)).
		Render(text)
}

// Fungsi Panel Tabel
func Panel_Tabel(text string, color string, padding1 int, padding2 int) string {
	return lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Padding(padding1, padding2).
		BorderForeground(lipgloss.Color("63")).
		Width(45).
		Foreground(lipgloss.Color(color)).
		Render(text)
}

// Fungsi Utama / Fitur Utama
func main() {
	var data DataPemilu
	var berlangsung bool

	for {
		Bersihkan_Terminal()
		var pilihan_kamu int
		var status string

		fmt.Println(Panel("  _    ____ ____ _ ____ _  _ _  _ _   _ \n  |    |___ | __ | [__  |_/  |  |  \\_/  \n  |___ |___ |__] | ___] | \\_ |__|   |   ", "33", 1, 2))
		if berlangsung {
			status = "Berlangsung"
		} else {
			status = "Selesai"
		}
		fmt.Println("[*] Status Pemilu:", status)
		// Pengaturan Waktu Pemilu = Hanya Untuk Petugas KPU
		// Manajemen Data = Hanya Untuk Petugas KPU
		// Ambang Batas Calon = Hanya Untuk Petugas KPU
		fmt.Println(Panel("[1]. Manajemen Data\n[2]. Proses Pemilu\n[3]. Temukan Data\n[4]. Urutkan Data\n[5]. Ambang Batas Calon\n[6]. Pengaturan Waktu Pemilu\n[7]. Keluar", "7", 0, 2))

		fmt.Print("[?] Pilihan Kamu: ")
		fmt.Scanln(&pilihan_kamu)
		switch pilihan_kamu {
		case 1:
			Manajemen_Data(&data)
		case 2:
			Proses_Pemilu(berlangsung, data)
		case 3:
			Pencarian_Data(data)
		case 4:
			Mengurutkan_Data(data)
		case 5:
			var threshold int
			fmt.Print("[?] Masukkan Ambang Batas Suara: ")
			fmt.Scanln(&threshold)
			Tampilkan_Berdasarkan_Threshold(threshold, data)
			Mengulang_Program()
		case 6:
			var status string
			fmt.Print("[?] Pemilu Berlangsung? (y/n): ")
			fmt.Scanln(&status)
			if status == "y" {
				berlangsung = true
				fmt.Println(Panel("[*] Pemilu Dimulai. Pemilih Dapat Memberikan Suara.", "2", 1, 2))
			} else {
				berlangsung = false
				fmt.Println(Panel("[*] Pemilu Selesai. Pemilih Tidak Dapat Memberikan Suara.", "1", 1, 2))
			}
			Mengulang_Program()
		case 7:
			fmt.Println(Panel("[*] Terima Kasih Telah Menggunakan Aplikasi LegisKuy.", "3", 1, 2))
			os.Exit(0)
		default:
			fmt.Println(Panel("[!] Pilihan Yang Anda Masukkan Tidak Valid. Silakan Coba Lagi.", "1", 1, 2))
			Mengulang_Program()
		}
	}
}

// Fungsi Mengulang Program
func Mengulang_Program() {
	var ulangi string
	fmt.Print("[?] Ulangi Program? (y/n): ")
	fmt.Scanln(&ulangi)
	if ulangi != "y" {
		os.Exit(0)
	}
}
