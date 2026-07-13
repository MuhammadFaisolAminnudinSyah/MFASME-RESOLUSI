package main

import "fmt"

const NMAX = 40

type atributMahasiswa struct {
	nama            string
	NIM             int64
	statusKehadiran string
	sumMeet         int
	persentase      float64
}
type pertemuanKelas struct {
	dataMahasiswa [NMAX]atributMahasiswa
}
type tabMahasiswa [NMAX]atributMahasiswa

type pertemuan [NMAX]pertemuanKelas

func sortSemuaPertemuanSelectionSort(S *pertemuan, jumlahMahasiswa int, meet int) {
	var i, pass, idx, j int
	var tempMhs atributMahasiswa

	for i = 1; i <= meet; i++ {
		for pass = 1; pass < jumlahMahasiswa; pass++ {
			idx = pass - 1
			for j = pass; j < jumlahMahasiswa; j++ {
				if (*S)[i].dataMahasiswa[idx].nama > (*S)[i].dataMahasiswa[j].nama {
					idx = j
				}
			}
			tempMhs = (*S)[i].dataMahasiswa[idx]
			(*S)[i].dataMahasiswa[idx] = (*S)[i].dataMahasiswa[pass-1]
			(*S)[i].dataMahasiswa[pass-1] = tempMhs
		}
	}
}

func sortSemuaPertemuanInsertionSort(S *pertemuan, jumlahMahasiswa int, meet int) {
	var i, pass, j int
	var tempMhs atributMahasiswa

	for i = 1; i <= meet; i++ {
		for pass = 1; pass < jumlahMahasiswa; pass++ {
			j = pass
			tempMhs = (*S)[i].dataMahasiswa[pass]
			for j > 0 && (*S)[i].dataMahasiswa[j-1].nama > tempMhs.nama {
				(*S)[i].dataMahasiswa[j] = (*S)[i].dataMahasiswa[j-1]
				j--
			}
		}
		(*S)[i].dataMahasiswa[j] = tempMhs
	}
}

func tampilSemuaPertemuan(S *pertemuan, Pertemuan tabMahasiswa, kelas string, jumlahMahasiswa int, meet int) {
	var i, j, k, l int

	for i = 1; i <= meet; i++ {
		for l = 0; l < 100; l++ {
			fmt.Print("-")
		}
		fmt.Println("")
		fmt.Print("Pertemuan ke-", Pertemuan[i].sumMeet)
		fmt.Printf("%-20s%s\n", "", kelas)
		for k = 0; k < 100; k++ {
			fmt.Print("-")
		}
		fmt.Println("")
		fmt.Printf("%-20s %-20s %-20s\n", "Nama", "NIM", "Status Kehadiran")
		fmt.Println("")
		for j = 0; j < jumlahMahasiswa; j++ {
			fmt.Printf("%-20s %-20d %-20s\n", (*S)[i].dataMahasiswa[j].nama, (*S)[i].dataMahasiswa[j].NIM, (*S)[i].dataMahasiswa[j].statusKehadiran)
		}
	}
	for k = 0; k < 100; k++ {
		fmt.Print("-")
	}
	fmt.Println("")
}

func input_namaMahasiswa(T *tabMahasiswa, Pertemuan *tabMahasiswa, S *pertemuan, kelas string, jumlahMahasiswa int, meet int) {
	var i, j int
	var k, l, m int

	fmt.Println("")
	for l = 0; l < 100; l++ {
		fmt.Print("-")
	}
	fmt.Println("\n=== INPUT DATA INDUK MAHASISWA ===")

	// TAHAP 1: Input Nama dan NIM di awal saja
	for j = 0; j < jumlahMahasiswa; j++ {
		fmt.Printf("Data Mahasiswa ke-%d\n", j+1)
		fmt.Print("Masukan nama Mahasiswa: ")
		fmt.Scan(&T[j].nama)
		fmt.Print("Masukan NIM Mahasiswa: ")
		fmt.Scan(&T[j].NIM)
		fmt.Println("")
	}

	// TAHAP 2: Looping setiap pertemuan untuk input absensi
	for i = 1; i <= meet; i++ {
		for l = 0; l < 100; l++ {
			fmt.Print("-")
		}

		fmt.Println("")
		fmt.Println("")
		fmt.Print("Pertemuan ke-")
		fmt.Scan(&Pertemuan[i].sumMeet)

		fmt.Println("--- Isi Status Kehadiran ---")
		for j = 0; j < jumlahMahasiswa; j++ {
			(*S)[i].dataMahasiswa[j].nama = T[j].nama
			(*S)[i].dataMahasiswa[j].NIM = T[j].NIM

			fmt.Printf("Status Kehadiran %s (NIM: %d): ", T[j].nama, T[j].NIM)
			fmt.Scan(&(*S)[i].dataMahasiswa[j].statusKehadiran)
		}

		for m = 1; m <= 100; m++ {
			fmt.Print("-")
		}
		fmt.Println("")
		fmt.Print("Pertemuan ke- ", Pertemuan[i].sumMeet)
		fmt.Printf("%-20s%s\n", "", kelas)

		for k = 0; k < 100; k++ {
			fmt.Print("-")
		}

		fmt.Println("")
		fmt.Printf("%-20s %-20s %-20s\n", "Nama", "NIM", "Status Kehadiran")
		fmt.Println("")

		for j = 0; j < jumlahMahasiswa; j++ {
			fmt.Printf("%-20s %-20d %-20s\n", (*S)[i].dataMahasiswa[j].nama, (*S)[i].dataMahasiswa[j].NIM, (*S)[i].dataMahasiswa[j].statusKehadiran)
		}
	}
	for m = 1; m <= 100; m++ {
		fmt.Print("-")
	}
	fmt.Println("")
}

func addData(T *tabMahasiswa, S *pertemuan, Pertemuan tabMahasiswa, kelas string, jumlahMahasiswa int, add int, meet int) {
	var batasbawah, batasatas int
	var i, j, m int
	var jenisPencarian string
	var hasilBinary, hasilSequential int
	var target int64

	batasbawah = jumlahMahasiswa
	batasatas = jumlahMahasiswa + add - 1

	fmt.Println("")
	for m = 1; m <= 100; m++ {
		fmt.Print("-")
	}
	fmt.Println("")

	for j = batasbawah; j <= batasatas; j++ {
		fmt.Print("Masukan nama Mahasiswa: ")
		fmt.Scan(&T[j].nama)
		fmt.Print("Masukan NIM Mahasiswa: ")
		fmt.Scan(&T[j].NIM)
		fmt.Println("")
	}
	for i = 1; i <= meet; i++ {
		for m = 1; m <= 100; m++ {
			fmt.Print("-")
		}
		fmt.Println("")
		fmt.Print("Status Kehadiran untuk Pertemuan ke-", Pertemuan[i].sumMeet)
		fmt.Println("")
		for j = batasbawah; j <= batasatas; j++ {
			fmt.Print("Status Kehadiran ", T[j].nama, ": ")
			fmt.Scan(&(*S)[i].dataMahasiswa[j].statusKehadiran)
			(*S)[i].dataMahasiswa[j].nama = T[j].nama
			(*S)[i].dataMahasiswa[j].NIM = T[j].NIM
			fmt.Println("")
		}
	}
	fmt.Print("Jenis Pencarian: ")
	fmt.Scan(&jenisPencarian)

	if jenisPencarian == "Binary" {
		sortSemuaPertemuanSelectionSort(S, batasatas+1, meet)
		tampilSemuaPertemuan(S, Pertemuan, kelas, batasatas+1, meet)
		fmt.Print("Mencari absen mahasiswa: ")
		fmt.Scan(&target)
		hasilBinary = binarySearch(S, batasatas+1, meet, target)
		if hasilBinary != -1 {
			fmt.Print("Siswa dengan NIM (", target, ") dengan nama ", (*S)[1].dataMahasiswa[hasilBinary].nama, " memiliki urutan absen nomor : ", hasilBinary+1)
		} else {
			fmt.Print("NIM TIDAK DITEMUKAN!")
		}
	} else if jenisPencarian == "Sequential" {
		sortSemuaPertemuanInsertionSort(S, batasatas+1, meet)
		tampilSemuaPertemuan(S, Pertemuan, kelas, batasatas+1, meet)
		fmt.Print("Mencari absen mahasiswa: ")
		fmt.Scan(&target)
		hasilSequential = SequentialSearch(S, batasatas+1, meet, target)
		if hasilSequential != -1 {
			fmt.Print("Siswa dengan NIM (", target, ") dengan nama ", (*S)[1].dataMahasiswa[hasilSequential].nama, " memiliki urutan absen nomor : ", hasilSequential+1)
		} else {
			fmt.Print("NIM TIDAK DITEMUKAN!")
		}
	}
}

func changeData(T *tabMahasiswa, dataBaru int, S *pertemuan, Pertemuan tabMahasiswa, change string, dataMahasiswaBaru tabMahasiswa, kelas string, jumlahMahasiswa int, meet int) {
	var i, j int
	var changetarget int64
	var k, l int
	var kondisi string
	var hasilSequential, hasilBinary int
	var jenisPencarian string
	var target int64

	if change == "komprehensif" {
		for i = 0; i < dataBaru; i++ {
			fmt.Print("Masukan NIM mahasiswa yang ingin diubah : ")
			fmt.Scan(&changetarget)
			fmt.Print("APA YANG INGIN DIUBAH: ")
			fmt.Scan(&kondisi)
			fmt.Println("")
			fmt.Println("MASUKKAN DATA MAHASISWA BARU")
			for k = 1; k <= 100; k++ {
				fmt.Print("-")
			}
			fmt.Println("")

			if kondisi == "All" {
				fmt.Print("Masukan nama Mahasiswa: ")
				fmt.Scan(&dataMahasiswaBaru[i].nama)
				fmt.Print("Masukan NIM Mahasiswa: ")
				fmt.Scan(&dataMahasiswaBaru[i].NIM)
				fmt.Print("Status Kehadiran: ")
				fmt.Scan(&dataMahasiswaBaru[i].statusKehadiran)
				fmt.Println("")
			} else if kondisi == "NAMA&&NIM" {
				fmt.Print("Masukan nama Mahasiswa: ")
				fmt.Scan(&dataMahasiswaBaru[i].nama)
				fmt.Print("Masukan NIM Mahasiswa: ")
				fmt.Scan(&dataMahasiswaBaru[i].NIM)
				fmt.Println("")
			} else if kondisi == "NAMA&&SK" {
				fmt.Print("Masukan nama Mahasiswa: ")
				fmt.Scan(&dataMahasiswaBaru[i].nama)
				fmt.Print("Status Kehadiran: ")
				fmt.Scan(&dataMahasiswaBaru[i].statusKehadiran)
				fmt.Println("")
			} else if kondisi == "NIM&&SK" {
				fmt.Print("Masukan NIM Mahasiswa: ")
				fmt.Scan(&dataMahasiswaBaru[i].NIM)
				fmt.Print("Status Kehadiran: ")
				fmt.Scan(&dataMahasiswaBaru[i].statusKehadiran)
				fmt.Println("")
			} else if kondisi == "NAMA" {
				fmt.Print("Masukan nama Mahasiswa: ")
				fmt.Scan(&dataMahasiswaBaru[i].nama)
				fmt.Println("")
			} else if kondisi == "NIM" {
				fmt.Print("Masukan NIM Mahasiswa: ")
				fmt.Scan(&dataMahasiswaBaru[i].NIM)
				fmt.Println("")
			} else if kondisi == "SK" {
				fmt.Print("Status Kehadiran: ")
				fmt.Scan(&dataMahasiswaBaru[i].statusKehadiran)
				fmt.Println("")
			}

			for j = 0; j < jumlahMahasiswa; j++ {
				if changetarget == T[j].NIM {
					if kondisi == "All" {
						T[j].nama = dataMahasiswaBaru[i].nama
						T[j].NIM = dataMahasiswaBaru[i].NIM
						T[j].statusKehadiran = dataMahasiswaBaru[i].statusKehadiran
					} else if kondisi == "NAMA&&NIM" {
						T[j].nama = dataMahasiswaBaru[i].nama
						T[j].NIM = dataMahasiswaBaru[i].NIM
					} else if kondisi == "NAMA&&SK" {
						T[j].nama = dataMahasiswaBaru[i].nama
						T[j].statusKehadiran = dataMahasiswaBaru[i].statusKehadiran
					} else if kondisi == "NIM&&SK" {
						T[j].NIM = dataMahasiswaBaru[i].NIM
						T[j].statusKehadiran = dataMahasiswaBaru[i].statusKehadiran
					} else if kondisi == "NAMA" {
						T[j].nama = dataMahasiswaBaru[i].nama
					} else if kondisi == "NIM" {
						T[j].NIM = dataMahasiswaBaru[i].NIM
					} else if kondisi == "SK" {
						T[j].statusKehadiran = dataMahasiswaBaru[i].statusKehadiran
					}
				}
			}

			for l = 1; l <= meet; l++ {
				for j = 0; j < jumlahMahasiswa; j++ {
					if changetarget == (*S)[l].dataMahasiswa[j].NIM {
						if kondisi == "All" {
							(*S)[l].dataMahasiswa[j].nama = dataMahasiswaBaru[i].nama
							(*S)[l].dataMahasiswa[j].NIM = dataMahasiswaBaru[i].NIM
							(*S)[l].dataMahasiswa[j].statusKehadiran = dataMahasiswaBaru[i].statusKehadiran
						} else if kondisi == "NAMA&&NIM" {
							(*S)[l].dataMahasiswa[j].nama = dataMahasiswaBaru[i].nama
							(*S)[l].dataMahasiswa[j].NIM = dataMahasiswaBaru[i].NIM
						} else if kondisi == "NAMA&&SK" {
							(*S)[l].dataMahasiswa[j].nama = dataMahasiswaBaru[i].nama
							(*S)[l].dataMahasiswa[j].statusKehadiran = dataMahasiswaBaru[i].statusKehadiran
						} else if kondisi == "NIM&&SK" {
							(*S)[l].dataMahasiswa[j].NIM = dataMahasiswaBaru[i].NIM
							(*S)[l].dataMahasiswa[j].statusKehadiran = dataMahasiswaBaru[i].statusKehadiran
						} else if kondisi == "NAMA" {
							(*S)[l].dataMahasiswa[j].nama = dataMahasiswaBaru[i].nama
						} else if kondisi == "NIM" {
							(*S)[l].dataMahasiswa[j].NIM = dataMahasiswaBaru[i].NIM
						} else if kondisi == "SK" {
							(*S)[l].dataMahasiswa[j].statusKehadiran = dataMahasiswaBaru[i].statusKehadiran
						}
					}
				}
			}
		}
		fmt.Print("Jenis Pencarian: ")
		fmt.Scan(&jenisPencarian)

		if jenisPencarian == "Binary" {
			sortSemuaPertemuanSelectionSort(S, jumlahMahasiswa, meet)
			tampilSemuaPertemuan(S, Pertemuan, kelas, jumlahMahasiswa, meet)
			fmt.Print("Mencari absen mahasiswa: ")
			fmt.Scan(&target)
			hasilBinary = binarySearch(S, jumlahMahasiswa, meet, target)
			if hasilBinary != -1 {
				fmt.Print("Siswa dengan NIM (", target, ") dengan nama ", (*S)[1].dataMahasiswa[hasilBinary].nama, " memiliki urutan absen nomor : ", hasilBinary+1)
			} else {
				fmt.Print("NIM TIDAK DITEMUKAN!")
			}
		} else if jenisPencarian == "Sequential" {
			sortSemuaPertemuanInsertionSort(S, jumlahMahasiswa, meet)
			tampilSemuaPertemuan(S, Pertemuan, kelas, jumlahMahasiswa, meet)
			fmt.Print("Mencari absen mahasiswa: ")
			fmt.Scan(&target)
			hasilSequential = SequentialSearch(S, jumlahMahasiswa, meet, target)
			if hasilSequential != -1 {
				fmt.Print("Siswa dengan NIM (", target, ") dengan nama ", (*S)[1].dataMahasiswa[hasilSequential].nama, " memiliki urutan absen nomor : ", hasilSequential+1)
			} else {
				fmt.Print("NIM TIDAK DITEMUKAN!")
			}
		}

	} else if change == "parsial" {
		// PERBAIKAN LOGIKA: Loop berdasarkan mahasiswa terlebih dahulu
		for p := 0; p < dataBaru; p++ {
			fmt.Print("Masukan NIM mahasiswa yang ingin diubah : ")
			fmt.Scan(&changetarget)
			fmt.Print("APA YANG INGIN DIUBAH (contoh: SK): ")
			fmt.Scan(&kondisi)
			fmt.Println("")

			if kondisi == "SK" {
				fmt.Println("=== MASUKKAN DATA MAHASISWA BARU ===")
				fmt.Println("Petunjuk: Ketik tanda '-' jika tidak ingin mengubah data pada pertemuan tertentu.")
				for k := 1; k <= 50; k++ {
					fmt.Print("-")
				}
				fmt.Println("")

				// Lalu loop ke setiap pertemuan, tanyakan statusnya
				for i = 1; i <= meet; i++ {
					var statusBaru string
					fmt.Printf("Status Kehadiran untuk Pertemuan ke-%d: ", Pertemuan[i].sumMeet)
					fmt.Scan(&statusBaru)

					// Validasi: Hanya timpa data jika input bukan "-"
					if statusBaru != "-" {
						for j = 0; j < jumlahMahasiswa; j++ {
							if changetarget == (*S)[i].dataMahasiswa[j].NIM {
								(*S)[i].dataMahasiswa[j].statusKehadiran = statusBaru
								break // Keluar loop lebih cepat jika NIM sudah ketemu
							}
						}
					}
				}
			} else {
				fmt.Println("Kondisi tidak valid untuk mode parsial saat ini.")
			}
		}

		fmt.Print("Jenis Pencarian: ")
		fmt.Scan(&jenisPencarian)

		if jenisPencarian == "Binary" {
			sortSemuaPertemuanSelectionSort(S, jumlahMahasiswa, meet)
			tampilSemuaPertemuan(S, Pertemuan, kelas, jumlahMahasiswa, meet)
			fmt.Print("Mencari absen mahasiswa: ")
			fmt.Scan(&target)
			hasilBinary = binarySearch(S, jumlahMahasiswa, meet, target)
			if hasilBinary != -1 {
				fmt.Print("Siswa dengan NIM (", target, ") dengan nama ", (*S)[1].dataMahasiswa[hasilBinary].nama, " memiliki urutan absen nomor : ", hasilBinary+1)
			} else {
				fmt.Print("NIM TIDAK DITEMUKAN!")
			}
		} else if jenisPencarian == "Sequential" {
			sortSemuaPertemuanInsertionSort(S, jumlahMahasiswa, meet)
			tampilSemuaPertemuan(S, Pertemuan, kelas, jumlahMahasiswa, meet)
			fmt.Print("Mencari absen mahasiswa: ")
			fmt.Scan(&target)
			hasilSequential = SequentialSearch(S, jumlahMahasiswa, meet, target)
			if hasilSequential != -1 {
				fmt.Print("Siswa dengan NIM (", target, ") dengan nama ", (*S)[1].dataMahasiswa[hasilSequential].nama, " memiliki urutan absen nomor : ", hasilSequential+1)
			} else {
				fmt.Print("NIM TIDAK DITEMUKAN!")
			}
		}
	}
}

func removeData(T *tabMahasiswa, S *pertemuan, Pertemuan tabMahasiswa, kelas string, jumlahMahasiswa *int, meet int) {
	var i, remove, j, k, m int
	var targetRemove int64
	var hasilSequential, hasilBinary int
	var jenisPencarian string
	var target int64

	fmt.Print("Berapa data yang akan dihapus: ")
	fmt.Scan(&remove)
	fmt.Println("")

	for i = 0; i < remove; i++ {
		fmt.Print("Masukkan NIM mahasiswa yang ingin dihapus: ")
		fmt.Scan(&targetRemove)

		for j = 0; j < *jumlahMahasiswa; j++ {
			if targetRemove == T[j].NIM {
				for k = j; k < *jumlahMahasiswa-1; k++ {
					T[k].nama = T[k+1].nama
					T[k].NIM = T[k+1].NIM
					T[k].statusKehadiran = T[k+1].statusKehadiran
				}
				break
			}
		}
		for m = 1; m <= meet; m++ {
			for j = 0; j < *jumlahMahasiswa; j++ {
				if targetRemove == (*S)[m].dataMahasiswa[j].NIM {
					for k = j; k < *jumlahMahasiswa-1; k++ {
						(*S)[m].dataMahasiswa[k].nama = (*S)[m].dataMahasiswa[k+1].nama
						(*S)[m].dataMahasiswa[k].NIM = (*S)[m].dataMahasiswa[k+1].NIM
						(*S)[m].dataMahasiswa[k].statusKehadiran = (*S)[m].dataMahasiswa[k+1].statusKehadiran
					}
					break
				}
			}
		}
		*jumlahMahasiswa--
	}
	fmt.Print("Jenis Pencarian: ")
	fmt.Scan(&jenisPencarian)

	if jenisPencarian == "Binary" {
		sortSemuaPertemuanSelectionSort(S, *jumlahMahasiswa, meet)
		tampilSemuaPertemuan(S, Pertemuan, kelas, *jumlahMahasiswa, meet)
		fmt.Print("Mencari absen mahasiswa: ")
		fmt.Scan(&target)
		hasilBinary = binarySearch(S, *jumlahMahasiswa, meet, target)
		if hasilBinary != -1 {
			fmt.Print("Siswa dengan NIM (", target, ") dengan nama ", (*S)[1].dataMahasiswa[hasilBinary].nama, " memiliki urutan absen nomor : ", hasilBinary+1)
		} else {
			fmt.Print("NIM TIDAK DITEMUKAN!")
		}
	} else if jenisPencarian == "Sequential" {
		sortSemuaPertemuanInsertionSort(S, *jumlahMahasiswa, meet)
		tampilSemuaPertemuan(S, Pertemuan, kelas, *jumlahMahasiswa, meet)
		fmt.Print("Mencari absen mahasiswa: ")
		fmt.Scan(&target)
		hasilSequential = SequentialSearch(S, *jumlahMahasiswa, meet, target)
		if hasilSequential != -1 {
			fmt.Print("Siswa dengan NIM (", target, ") dengan nama ", (*S)[1].dataMahasiswa[hasilSequential].nama, " memiliki urutan absen nomor : ", hasilSequential+1)
		} else {
			fmt.Print("NIM TIDAK DITEMUKAN!")
		}
	}
}

func binarySearch(S *pertemuan, jumlahMahasiswa int, meet int, target int64) int {
	var left, right, mid int
	var found int

	found = -1
	left = 0
	right = jumlahMahasiswa - 1

	for left <= right && found == -1 {
		mid = (left + right) / 2
		if (*S)[1].dataMahasiswa[mid].NIM == target {
			found = mid
		} else if (*S)[1].dataMahasiswa[mid].NIM < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return found
}

func SequentialSearch(S *pertemuan, jumlahMahasiswa int, meet int, target int64) int {
	var i, ketemu int
	i = 0
	ketemu = -1

	for i < jumlahMahasiswa && ketemu == -1 {
		if (*S)[1].dataMahasiswa[i].NIM == target {
			ketemu = i
		}
		i = i + 1
	}
	return ketemu
}

func persentaseKehadiran(S pertemuan, Persentase tabMahasiswa, kelas string, jumlahMahasiswa int, meet int) {
	var i, j, k, l int
	var sumHadir int

	for j = 0; j < jumlahMahasiswa; j++ {
		sumHadir = 0
		for i = 1; i <= meet; i++ {
			if S[i].dataMahasiswa[j].statusKehadiran == "HADIR" {
				sumHadir++
			}
		}
		Persentase[j].persentase = (float64(sumHadir) / float64(meet)) * 100
	}

	fmt.Println("")
	for k = 0; k < 100; k++ {
		fmt.Print("-")
	}

	fmt.Println("")
	fmt.Printf("%-20s%s%s\n", "", "PERSENTASE KEHADIRAN ", kelas)

	for l = 0; l < 100; l++ {
		fmt.Print("-")
	}
	fmt.Println("")

	fmt.Printf("%-20s%s\n", "NAMA", "KEHADIRAN(%)")
	fmt.Println("")

	for j = 0; j < jumlahMahasiswa; j++ {
		fmt.Printf("%-20s %.2f\n", (S)[1].dataMahasiswa[j].nama, Persentase[j].persentase)
	}

}

func main() {
	var T, dataMahasiswa, Pertemuan, Persentase tabMahasiswa
	var S pertemuan
	var kelas, kondisi, change string
	var jenisPencarian string
	var hasilSequential, hasilBinary int
	var target int64
	var jumlahMahasiswa, add, meet, dataBaru, m int

	fmt.Print("Masukan Kelas: ")
	fmt.Scan(&kelas)
	fmt.Print("Masukan Jumlah Mahasiswa: ")
	fmt.Scan(&jumlahMahasiswa)
	fmt.Print("Banyak Pertemuan: ")
	fmt.Scan(&meet)
	input_namaMahasiswa(&T, &Pertemuan, &S, kelas, jumlahMahasiswa, meet)
	fmt.Println("")

	fmt.Print("SILAHKAN BILA INGIN MENAMBAHKAN, MENGUBAH, ATAUPUN MENGHAPUS DATA: ")
	fmt.Scan(&kondisi)

	if kondisi == "ADD" {
		fmt.Print("Berapa total mahasiswa yang ingin ditambahkan: ")
		fmt.Scan(&add)
		addData(&T, &S, Pertemuan, kelas, jumlahMahasiswa, add, meet)
		jumlahMahasiswa = jumlahMahasiswa + add
		persentaseKehadiran(S, Persentase, kelas, jumlahMahasiswa, meet)

	} else if kondisi == "CHANGE" {
		fmt.Print("Ada berapa mahasiswa yang ingin diubah : ")
		fmt.Scan(&dataBaru)
		fmt.Println("")
		for m = 1; m <= 100; m++ {
			fmt.Print("-")
		}
		fmt.Println("")
		fmt.Print("Data tersebut diubah secara: ")
		fmt.Scan(&change)
		fmt.Println("")
		if change == "komprehensif" {
			changeData(&T, dataBaru, &S, Pertemuan, change, dataMahasiswa, kelas, jumlahMahasiswa, meet)
			persentaseKehadiran(S, Persentase, kelas, jumlahMahasiswa, meet)

		} else if change == "parsial" {
			changeData(&T, dataBaru, &S, Pertemuan, change, dataMahasiswa, kelas, jumlahMahasiswa, meet)
			persentaseKehadiran(S, Persentase, kelas, jumlahMahasiswa, meet)
		}

	} else if kondisi == "REMOVE" {
		removeData(&T, &S, Pertemuan, kelas, &jumlahMahasiswa, meet)
		persentaseKehadiran(S, Persentase, kelas, jumlahMahasiswa, meet)
	} else if kondisi == "NOTINGS" {
		fmt.Print("Jenis Pencarian: ")
		fmt.Scan(&jenisPencarian)
		if jenisPencarian == "Binary" {
			sortSemuaPertemuanSelectionSort(&S, jumlahMahasiswa, meet)
			tampilSemuaPertemuan(&S, Pertemuan, kelas, jumlahMahasiswa, meet)
			fmt.Print("Mencari absen mahasiswa: ")
			fmt.Scan(&target)
			hasilBinary = binarySearch(&S, jumlahMahasiswa, meet, target)
			if hasilBinary != -1 {
				fmt.Print("Siswa dengan NIM (", target, ") dengan nama ", S[1].dataMahasiswa[hasilBinary].nama, " memiliki urutan absen nomor : ", hasilBinary+1)
			} else {
				fmt.Print("NIM TIDAK DITEMUKAN!")
			}
		} else if jenisPencarian == "Sequential" {
			sortSemuaPertemuanInsertionSort(&S, jumlahMahasiswa, meet)
			tampilSemuaPertemuan(&S, Pertemuan, kelas, jumlahMahasiswa, meet)
			fmt.Print("Mencari absen mahasiswa: ")
			fmt.Scan(&target)
			hasilSequential = SequentialSearch(&S, jumlahMahasiswa, meet, target)
			if hasilSequential != -1 {
				fmt.Print("Siswa dengan NIM (", target, ") dengan nama ", S[1].dataMahasiswa[hasilSequential].nama, " memiliki urutan absen nomor : ", hasilSequential+1)
			} else {
				fmt.Print("NIM TIDAK DITEMUKAN!")
			}
		}
		persentaseKehadiran(S, Persentase, kelas, jumlahMahasiswa, meet)
	}
}
