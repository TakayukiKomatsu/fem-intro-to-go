package main

import "fmt"

func average(numbers ...float64) float64 {

	var sum float64

	arrayLen := len(numbers)

	for _, value := range numbers {
		sum += value
	}

	return sum / float64(arrayLen)

}

func main() {

	fmt.Println(average(10, 5, 7))

}
