package main

import "fmt"

var data = [...]int{5, 1, 6, 4, 5, 2, 3, 7, 8, 2, 3, 5, 11}

func main() {

	length := len(data)
	currentMax := 0
	maxWithItem := 0

	for length > 0 {

		if data[length-1]*data[length-1] > currentMax {
			currentMax = data[length-1] * data[length-1]
		}

		maxWithItem = findMax(length-1, len(data)-1)
		if findMax(length-1, len(data)-1) > currentMax {
			currentMax = maxWithItem
		}
		length--
		fmt.Println(currentMax)
	}
}

func findMax(i int, j int) (res int) {
	min := data[i]
	sum := 0
	res = 0
	for index := i; index <= j; index++ {
		if data[index] < min {
			min = data[index]
		}
		sum += data[index]
		if min*sum > res {
			res = min * sum
			//fmt.Printf("%d %d\n",min,sum)
		}
	}
	return res
}
