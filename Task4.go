package main

import "fmt"

func main() {
	var arr [6]int = [6]int {17,18,5,4,6,1}
	max := 1
	for i := 0;i<len(arr);i++ {
		if i == len(arr) - 1 {
			arr[i] = -1
			break
		}
		for j := i+1; j<len(arr); j++ {
			if arr[j] >= max{
				max = arr[j]
			}
		}
		arr[i] = max
		max = 1
	}
	fmt.Println(arr)
}
