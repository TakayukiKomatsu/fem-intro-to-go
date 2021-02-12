package main

import "fmt"

func main() {
	var sentence = "Hello this is a sentence.\n"

	for index, letter := range sentence {
		if index%2 == 1 {
			fmt.Println("Index:", index, "Letter:", string(letter))
		}
	}
}
