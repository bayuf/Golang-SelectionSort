package main

import "fmt"

func main() {

	datas := []int{5, 4}

	fmt.Println("Data Belum Urut \n", datas)

	for i := 0; i <= len(datas)-2; i++ {
		minIndex := i

		for j := i + 1; j <= len(datas)-1; j++ {
			if datas[j] < datas[minIndex] {
				minIndex = j
			}
		}

		temp := datas[i]
		datas[i] = datas[minIndex]
		datas[minIndex] = temp
	}

	fmt.Println("Data Urut \n", datas)
}
