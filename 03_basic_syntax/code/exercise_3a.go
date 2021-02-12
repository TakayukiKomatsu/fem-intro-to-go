package main

import (
	"fmt"
	"reflect"
)

func main() {
	var sentence = "Hello this is a sentence.\n"

	fmt.Printf(sentence)
	fmt.Println(reflect.TypeOf(sentence))

}
