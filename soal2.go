package main

import (
	"fmt"
)

func getRangking(daftar_skor []int, skor_didapat []int) []int {

	var daftar_rangking = []int{}

	score_con := 1
	var last_rangking = 0
	for i, _ := range daftar_skor {

		if daftar_skor[i] < last_rangking {
			score_con++
		}
		last_rangking = daftar_skor[i]
		daftar_rangking = append(daftar_rangking, score_con)

	}
	// fmt.Println(daftar_rangking)
	// ===============================================

	// fmt.Println(skor_didapat[len(skor_didapat)-1])
	var result = []int{}
	score_con = 0

	for i, _ := range skor_didapat {
		var temp = 0
		for j, _ := range daftar_skor {
			if skor_didapat[i] > daftar_skor[0] {
				temp = 1
			}

			if skor_didapat[i] < daftar_skor[len(daftar_skor)-1] {
				// fmt.Println(j, ">>>", skor_didapat[i], "---", daftar_skor[len(daftar_skor)-1])

				temp = daftar_rangking[len(daftar_rangking)-1] + 1

			}

			if j < len(daftar_rangking)-2 {
				if skor_didapat[i] < daftar_skor[j] && skor_didapat[i] > daftar_skor[j+1] {
					temp = daftar_rangking[j+1]
				}

			}

			if skor_didapat[i] == daftar_skor[j] {
				temp = daftar_rangking[j]
			}

		}
		if temp != 0 {

			result = append(result, temp)
		}

	}

	return result
}

func main() {

	// var jumlah_pemain = 7
	// var jumlah_permainan = 4

	fmt.Println("case 1")
	var daftar_skor = []int{100, 80, 80, 70}
	var skor_didapat = []int{60, 70, 100}

	var result = getRangking(daftar_skor, skor_didapat)
	fmt.Println(result)

	fmt.Println("case 2")
	skor_didapat = []int{5, 25, 50, 120}
	daftar_skor = []int{100, 100, 50, 40, 40, 20, 10}
	result = getRangking(daftar_skor, skor_didapat)
	fmt.Println(result)
	// ============================================

}
