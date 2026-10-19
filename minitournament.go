package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	const jumlahPertandingan = 3

	var arrayVenueStadion [jumlahPertandingan]string
	var klasementPoint = make(map[string]int)
	var totalGol int
	var jumlahSeri int
	var riwayatPertandingan []Riwayat
	var reader = bufio.NewReader(os.Stdin)

	fmt.Println("========================================")
	fmt.Println("     INPUT DATA VENUE STADION (ARRAY)")
	fmt.Println("========================================")

	// Input nama stadion untuk setiap pertandingan
	for i := 1; i <= jumlahPertandingan; i++ {
		arrayVenueStadion[i-1] = inputString(fmt.Sprintf("Masukkan nama stadion untuk laga ke-%d : ", i), reader)
	}

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("     SISTEM TURNAMEN MINI (REFACTOR)")
	fmt.Println("========================================")

	// Input data pertandingan
	for i := 1; i <= jumlahPertandingan; i++ {

		fmt.Println()
		fmt.Printf("=== Pertandingan Ke-%d di Stadion %s ===\n", i, arrayVenueStadion[i-1])

		homeTeam := inputString("Nama Tim Home : ", reader)
		awayTeam := inputString("Nama Tim Away : ", reader)

		homeGoal := inputInt("Gol "+homeTeam+" : ", reader)
		awayGoal := inputInt("Gol "+awayTeam+" : ", reader)

		// Menentukan pemenang, jumlah seri dan update klasemen
		pemenang := countPoints(homeTeam, awayTeam, homeGoal, awayGoal, klasementPoint, &jumlahSeri)

		// Update total gol
		totalGol += homeGoal + awayGoal

		// Menyimpan riwayat pertandingan
		riwayatPertandingan = append(riwayatPertandingan, Riwayat{
			Stadion:  arrayVenueStadion[i-1],
			HomeTeam: homeTeam,
			AwayTeam: awayTeam,
			HomeGoal: homeGoal,
			AwayGoal: awayGoal,
			Pemenang: pemenang,
		})
	}

	// Menghitung rata-rata gol
	rataRataGol := float64(totalGol) / float64(jumlahPertandingan)

	fmt.Println()
	fmt.Println("===========================================")
	fmt.Println("     RINGKASAN TURNAMEN & KLASEMEN AKHIR")
	fmt.Println("===========================================")

	// Menampilkan daftar venue stadion dari array
	fmt.Println("Daftar Venue Stadion (Dari Array):")
	for i, venue := range arrayVenueStadion {
		fmt.Printf("- Laga %d bertempat di %s\n", i+1, venue)
	}
	fmt.Println()

	// Menampilkan riwayat pertandingan dari slice
	fmt.Println("Riwayat Pertandingan:")
	for i, r := range riwayatPertandingan {
		fmt.Printf("%d. [%s] %s %d - %d %s --> Pemenang: %s\n",
			i+1, r.Stadion, r.HomeTeam, r.HomeGoal, r.AwayGoal, r.AwayTeam, r.Pemenang)
	}
	fmt.Println()

	// Menghitung panjang nama tim terpanjang untuk format output di klasemen akhir
	var maxLengthTeamName int = 0
	for team := range klasementPoint {
		if len(team) > maxLengthTeamName {
			maxLengthTeamName = len(team)
		}
	}

	// Menampilkan klasemen akhir
	fmt.Println("Klasemen Point Akhir:")
	for team, point := range klasementPoint {
		fmt.Printf("- %-*s : %d Poin\n", maxLengthTeamName, team, point)
	}

	// Menampilkan ringkasan statistik turnamen
	fmt.Println("-----------------------------------")
	fmt.Println("Jumlah Pertandingan :", jumlahPertandingan)
	fmt.Println("Total Gol           :", totalGol)
	fmt.Println("Jumlah Seri         :", jumlahSeri)
	fmt.Printf("Rata-rata Gol       : %.2f\n", rataRataGol)
}

// Fungsi untuk input string dengan validasi
func inputString(prompt string, reader *bufio.Reader) string {
	for {
		fmt.Print(prompt)
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("ERROR:", err)
			continue
		}
		text = strings.TrimSpace(text)
		if text == "" {
			fmt.Println("ERROR: Input tidak boleh kosong!")
			continue
		}
		return text
	}
}

// Fungsi untuk input integer dengan validasi
func inputInt(prompt string, reader *bufio.Reader) int {
	for {
		fmt.Print(prompt)
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("ERROR:", err)
			continue
		}

		text = strings.TrimSpace(text)
		if text == "" {
			fmt.Println("ERROR: Input tidak boleh kosong!")
			continue
		}

		value, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("ERROR: Gol harus berupa angka yang valid!")
			continue
		}

		if value < 0 {
			fmt.Println("ERROR: Gol tidak boleh negatif!")
			continue
		}

		return value
	}
}

// Fungsi untuk menghitung poin, jumlah seri dan menentukan pemenang
func countPoints(homeTeam, awayTeam string, homeGoal, awayGoal int, klasement map[string]int, jumlahSeri *int) string {
	if homeGoal > awayGoal {
		klasement[homeTeam] += 3
		klasement[awayTeam] += 0
		return homeTeam
	} else if awayGoal > homeGoal {
		klasement[awayTeam] += 3
		klasement[homeTeam] += 0
		return awayTeam
	} else {
		klasement[homeTeam]++
		klasement[awayTeam]++
		*jumlahSeri++
		return "Seri"
	}
}

// Struct untuk menyimpan riwayat pertandingan
type Riwayat struct {
	Stadion  string
	HomeTeam string
	AwayTeam string
	HomeGoal int
	AwayGoal int
	Pemenang string
}
