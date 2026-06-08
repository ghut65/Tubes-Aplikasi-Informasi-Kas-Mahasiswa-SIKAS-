package main

import "fmt"

type Mahasiswa struct {
	NIM        string
	Nama       string
	Tunggakan  int
	TotalBayar int
}

type Pembayaran struct {
	NIM     string
	Nominal int
	Tanggal string
}

var kasKelas []Mahasiswa
var riwayatPembayaran []Pembayaran

const TargetIuran = 100000

func main() {
	kasKelas = []Mahasiswa{}
	riwayatPembayaran = []Pembayaran{}

	for {
		fmt.Println("\n-- Aplikasi Informasi Kas Mahasiswa (SIKAS) --")
		fmt.Println("1. Tambah Mahasiswa")
		fmt.Println("2. Lihat Data Mahasiswa")
		fmt.Println("3. Ubah Data Mahasiswa")
		fmt.Println("4. Hapus Data Mahasiswa")
		fmt.Println("5. Catat Pembayaran Iuran")
		fmt.Println("6. Cari Mahasiswa (Belum Lunas)")
		fmt.Println("7. Urutkan Data Mahasiswa")
		fmt.Println("8. Laporan Statistik & Riwayat")
		fmt.Println("9. Keluar")

		var pilih int
		fmt.Print("Pilih: ")
		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			tambah()
		case 2:
			lihat()
		case 3:
			ubah()
		case 4:
			hapus()
		case 5:
			catatPembayaran()
		case 6:
			cari()
		case 7:
			urutkan()
		case 8:
			laporan()
		case 9:
			fmt.Println("Anda telah keluar dari aplikasi SIKAS.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func tambah() {
	var nim string
	var nama string

	fmt.Print("NIM: ")
	fmt.Scan(&nim)
	fmt.Print("Nama Mahasiswa (Tanpa spasi): ")
	fmt.Scan(&nama)

	tunggakan := TargetIuran
	totalBayar := 0

	kasKelas = append(kasKelas, Mahasiswa{nim, nama, tunggakan, totalBayar})
	fmt.Println("Data mahasiswa berhasil ditambahkan.")
}

func lihat() {
	if len(kasKelas) == 0 {
		fmt.Println("Data kosong.")
		return
	}

	fmt.Println("\nDaftar Mahasiswa:")
	for i, m := range kasKelas {
		if m.Nama == "" {
			m.Nama = "Tidak Diketahui"
		}

		fmt.Printf("%d. NIM: %s | Nama: %s | Total Bayar: Rp%d | Tunggakan: Rp%d\n",
			i+1, m.NIM, m.Nama, m.TotalBayar, m.Tunggakan)
	}
}

func ubah() {
	if len(kasKelas) == 0 {
		fmt.Println("Data kosong.")
		return
	}

	var nim string
	fmt.Print("Masukkan NIM mahasiswa yang ingin diubah: ")
	fmt.Scan(&nim)

	found := -1
	for i := 0; i < len(kasKelas); i++ {
		if kasKelas[i].NIM == nim {
			found = i
			break
		}
	}

	if found == -1 {
		fmt.Println("Mahasiswa tidak ditemukan.")
		return
	}

	var namaBaru string
	fmt.Print("Nama baru (Tanpa spasi): ")
	fmt.Scan(&namaBaru)

	kasKelas[found].Nama = namaBaru
	fmt.Println("Data mahasiswa berhasil diperbarui.")
}

func hapus() {
	var nim string
	fmt.Print("Masukkan NIM mahasiswa yang ingin dihapus: ")
	fmt.Scan(&nim)

	for i := 0; i < len(kasKelas); i++ {
		if kasKelas[i].NIM == nim {
			kasKelas = append(kasKelas[:i], kasKelas[i+1:]...)
			fmt.Println("Mahasiswa berhasil dihapus.")
			return
		}
	}
	fmt.Println("Mahasiswa tidak ditemukan.")
}

func catatPembayaran() {
	if len(kasKelas) == 0 {
		fmt.Println("Data kosong.")
		return
	}

	var nim string
	var nominal int
	var tanggal string

	fmt.Print("Masukkan NIM mahasiswa: ")
	fmt.Scan(&nim)

	found := -1
	for i := 0; i < len(kasKelas); i++ {
		if kasKelas[i].NIM == nim {
			found = i
			break
		}
	}

	if found == -1 {
		fmt.Println("Mahasiswa tidak ditemukan.")
		return
	}

	if kasKelas[found].Tunggakan == 0 {
		fmt.Println("Mahasiswa ini sudah LUNAS.")
		return
	}

	fmt.Print("Masukkan tanggal pembayaran (Contoh: 15-05-2026): ")
	fmt.Scan(&tanggal)
	fmt.Print("Masukkan nominal pembayaran: Rp")
	fmt.Scan(&nominal)

	kasKelas[found].TotalBayar += nominal
	kasKelas[found].Tunggakan = TargetIuran - kasKelas[found].TotalBayar

	if kasKelas[found].Tunggakan < 0 {
		kasKelas[found].Tunggakan = 0
	}

	riwayatPembayaran = append(riwayatPembayaran, Pembayaran{nim, nominal, tanggal})
	fmt.Println("Pembayaran berhasil dicatat.")
}

func cari() {
	var metode int
	fmt.Println("Pilih metode pencarian Mahasiswa Belum Lunas:")
	fmt.Println("1. Sequential Search (Tampilkan semua yang belum bayar)")
	fmt.Println("2. Binary Search (Cek status lunas spesifik berdasarkan NIM)")
	fmt.Print("Pilihan: ")
	fmt.Scan(&metode)

	switch metode {
	case 1:
		seqSearch()
	case 2:
		binSearch()
	default:
		fmt.Println("Metode tidak valid.")
	}
}

func seqSearch() {
	fmt.Println("\n--- Hasil Sequential Search (Daftar Tunggakan) ---")
	found := false
	j := 0
	for j < len(kasKelas) {
		if kasKelas[j].Tunggakan > 0 {
			fmt.Printf("- %s (NIM: %s) menunggak: Rp%d\n", kasKelas[j].Nama, kasKelas[j].NIM, kasKelas[j].Tunggakan)
			found = true
		}
		j++
	}

	if !found {
		fmt.Println("Semua mahasiswa sudah lunas.")
	}
}

func binSearch() {
	var nim string
	fmt.Print("Masukkan NIM yang ingin dicek: ")
	fmt.Scan(&nim)

	sortNIMSebelumCari()

	found := -1
	kr := 0
	kn := len(kasKelas) - 1
	var med int

	for kr <= kn && found == -1 {
		med = (kr + kn) / 2
		if nim > kasKelas[med].NIM {
			kr = med + 1
		} else if nim < kasKelas[med].NIM {
			kn = med - 1
		} else {
			found = med
		}
	}

	if found != -1 {
		m := kasKelas[found]
		fmt.Printf("Ditemukan: %s | Tunggakan: Rp%d\n", m.Nama, m.Tunggakan)
		if m.Tunggakan > 0 {
			fmt.Println("Status: BELUM LUNAS (Memiliki Tunggakan)")
		} else {
			fmt.Println("Status: LUNAS")
		}
	} else {
		fmt.Println("Data mahasiswa tidak ditemukan.")
	}
}

func sortNIMSebelumCari() {
	n := len(kasKelas)
	for i := 0; i < n-1; i++ {
		idx_min := i
		for j := i + 1; j < n; j++ {
			if kasKelas[j].NIM < kasKelas[idx_min].NIM {
				idx_min = j
			}
		}
		temp := kasKelas[idx_min]
		kasKelas[idx_min] = kasKelas[i]
		kasKelas[i] = temp
	}
}
func urutkan() {
	if len(kasKelas) == 0 {
		fmt.Println("Data kosong.")
		return
	}

	fmt.Println("Pilih metode pengurutan:")
	fmt.Println("1. Selection Sort (Berdasarkan Nama - Ascending)")
	fmt.Println("2. Insertion Sort (Berdasarkan Tunggakan - Descending)")
	fmt.Print("Pilihan: ")

	var pilihan int
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		selectionSortNama()
		fmt.Println("Data diurutkan berdasarkan Nama (Ascending)")
	case 2:
		insertionSortTunggakan()
		fmt.Println("Data diurutkan berdasarkan Tunggakan (Descending)")
	default:
		fmt.Println("Pilihan tidak valid")
		return
	}

	lihat()
}

func selectionSortNama() {
	n := len(kasKelas)
	for i := 0; i < n-1; i++ {
		idx_min := i
		for j := i + 1; j < n; j++ {
			if kasKelas[j].Nama < kasKelas[idx_min].Nama {
				idx_min = j
			}
		}

		temp := kasKelas[idx_min]
		kasKelas[idx_min] = kasKelas[i]
		kasKelas[i] = temp
	}
}

func insertionSortTunggakan() {
	n := len(kasKelas)
	for i := 1; i < n; i++ {
		temp := kasKelas[i]
		j := i
		for j > 0 && kasKelas[j-1].Tunggakan < temp.Tunggakan {
			kasKelas[j] = kasKelas[j-1]
			j--
		}
		kasKelas[j] = temp
	}
}

func laporan() {
	if len(kasKelas) == 0 {
		fmt.Println("Tidak ada data untuk dilaporkan.")
		return
	}

	var totalSaldo int
	var countLunas, countMenunggak int

	for i := 0; i < len(kasKelas); i++ {
		totalSaldo += kasKelas[i].TotalBayar
		if kasKelas[i].Tunggakan == 0 {
			countLunas++
		} else {
			countMenunggak++
		}
	}

	fmt.Println("\n--- Laporan Statistik Kas Mahasiswa ---")
	fmt.Printf("Total Mahasiswa        : %d orang\n", len(kasKelas))
	fmt.Printf("Total Saldo Kas        : Rp%d\n", totalSaldo)
	fmt.Println("Status Pembayaran      :")
	fmt.Printf(" - Sudah Melunasi Iuran: %d mahasiswa\n", countLunas)
	fmt.Printf(" - Belum Melunasi Iuran: %d mahasiswa\n", countMenunggak)

	fmt.Println("\n--- Riwayat Transaksi Kas ---")
	if len(riwayatPembayaran) == 0 {
		fmt.Println("Belum ada transaksi pembayaran.")
	} else {
		for i, r := range riwayatPembayaran {
			fmt.Printf("%d. NIM: %s | Tanggal: %s | Nominal: Rp%d\n", i+1, r.NIM, r.Tanggal, r.Nominal)
		}
	}
}
