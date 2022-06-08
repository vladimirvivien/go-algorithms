package sort

import "fmt"

func InsertionSort(data []int) {
	fmt.Printf("sorting: %v\n", data)
	for i := 1; i < len(data); i++ {
		for j := i; j >= 0; j-- {
			if j == 0 {
				break
			}
			if data[j-1] >= data[j] {
				temp := data[j]
				data[j] = data[j-1]
				data[j-1] = temp
			}
		}
	}

}
