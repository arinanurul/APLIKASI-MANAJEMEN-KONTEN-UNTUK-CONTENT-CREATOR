package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Content struct {
	ID         int
	Title      string
	Platform   string
	Category   string
	Date       string
	Engagement int
}

var contents []Content
var reader = bufio.NewReader(os.Stdin)
var idCounter = 1

// ================= CRUD =================

func tambahContent() {
	var title, platform, category, date string
	var engagement int

	fmt.Print("Judul konten: ")
	title, _ = reader.ReadString('\n')

	fmt.Print("Platform: ")
	platform, _ = reader.ReadString('\n')

	fmt.Print("Kategori: ")
	category, _ = reader.ReadString('\n')

	fmt.Print("Tanggal (YYYY-MM-DD): ")
	date, _ = reader.ReadString('\n')

	fmt.Print("Engagement: ")
	fmt.Scanln(&engagement)

	content := Content{
		ID:         idCounter,
		Title:      strings.TrimSpace(title),
		Platform:   strings.TrimSpace(platform),
		Category:   strings.TrimSpace(category),
		Date:       strings.TrimSpace(date),
		Engagement: engagement,
	}

	contents = append(contents, content)
	idCounter++

	fmt.Println("✅ Konten berhasil ditambahkan!")
}

func tampilkanContent() {
	if len(contents) == 0 {
		fmt.Println("Belum ada konten.")
		return
	}

	for _, c := range contents {
		fmt.Printf("[%d] %s | %s | %s | %s | Engagement: %d\n",
			c.ID, c.Title, c.Platform, c.Category, c.Date, c.Engagement)
	}
}

func hapusContent() {
	var id int
	fmt.Print("Masukkan ID yang ingin dihapus: ")
	fmt.Scanln(&id)

	for i, c := range contents {
		if c.ID == id {
			contents = append(contents[:i], contents[i+1:]...)
			fmt.Println("🗑️ Konten dihapus.")
			return
		}
	}
	fmt.Println("ID tidak ditemukan.")
}

func updateContent() {
	var id int
	fmt.Print("Masukkan ID yang ingin diupdate: ")
	fmt.Scanln(&id)

	for i := range contents {
		if contents[i].ID == id {
			fmt.Print("Judul baru: ")
			title, _ := reader.ReadString('\n')
			contents[i].Title = strings.TrimSpace(title)

			fmt.Print("Engagement baru: ")
			fmt.Scanln(&contents[i].Engagement)

			fmt.Println("✏️ Konten berhasil diupdate!")
			return
		}
	}
	fmt.Println("ID tidak ditemukan.")
}

// ================= SEARCH =================

// Sequential Search
func sequentialSearch(keyword string) {
	found := false
	for _, c := range contents {
		if strings.Contains(strings.ToLower(c.Title), strings.ToLower(keyword)) {
			fmt.Println("Ditemukan:", c.Title)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ditemukan.")
	}
}

// Binary Search (harus sorted dulu berdasarkan Title)
func binarySearch(keyword string) {
	left, right := 0, len(contents)-1

	for left <= right {
		mid := (left + right) / 2
		if contents[mid].Title == keyword {
			fmt.Println("Ditemukan:", contents[mid].Title)
			return
		} else if contents[mid].Title < keyword {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	fmt.Println("Tidak ditemukan.")
}

// ================= SORTING =================

// Selection Sort (by engagement)
func selectionSort() {
	n := len(contents)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if contents[j].Engagement > contents[min].Engagement {
				min = j
			}
		}
		contents[i], contents[min] = contents[min], contents[i]
	}
	fmt.Println("Sorted by engagement (descending).")
}

// Insertion Sort (by date)
func insertionSort() {
	for i := 1; i < len(contents); i++ {
		key := contents[i]
		j := i - 1

		for j >= 0 && contents[j].Date > key.Date {
			contents[j+1] = contents[j]
			j--
		}
		contents[j+1] = key
	}
	fmt.Println("Sorted by date.")
}

// ================= ANALISIS =================

func topEngagement() {
	if len(contents) == 0 {
		fmt.Println("Tidak ada data.")
		return
	}

	max := contents[0]
	for _, c := range contents {
		if c.Engagement > max.Engagement {
			max = c
		}
	}

	fmt.Println("🔥 Konten dengan engagement tertinggi:")
	fmt.Println(max.Title, "|", max.Engagement)
}

// ================= MENU =================

func menu() {
	for {
		fmt.Println("\n=== Aplikasi Manajemen Konten ===")
		fmt.Println("1. Tambah Konten")
		fmt.Println("2. Lihat Konten")
		fmt.Println("3. Update Konten")
		fmt.Println("4. Hapus Konten")
		fmt.Println("5. Search (Sequential)")
		fmt.Println("6. Search (Binary)")
		fmt.Println("7. Sort Engagement (Selection)")
		fmt.Println("8. Sort Date (Insertion)")
		fmt.Println("9. Top Engagement")
		fmt.Println("0. Keluar")

		fmt.Print("Pilih: ")
		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tambahContent()
		case 2:
			tampilkanContent()
		case 3:
			updateContent()
		case 4:
			hapusContent()
		case 5:
			fmt.Print("Keyword: ")
			key, _ := reader.ReadString('\n')
			sequentialSearch(strings.TrimSpace(key))
		case 6:
			fmt.Print("Keyword: ")
			key, _ := reader.ReadString('\n')
			binarySearch(strings.TrimSpace(key))
		case 7:
			selectionSort()
		case 8:
			insertionSort()
		case 9:
			topEngagement()
		case 0:
			fmt.Println("Keluar...")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func main() {
	menu()
}