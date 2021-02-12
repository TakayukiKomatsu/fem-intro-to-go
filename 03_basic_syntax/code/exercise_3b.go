package main

import "fmt"

func main() {
	var name, hometown string

	var time int
	var weather bool

	fmt.Println("Type your name")
	fmt.Scan(&name)

	fmt.Println("Type your hometown name")
	fmt.Scan(&hometown)

	fmt.Println("Type the time lived in your hometown")
	fmt.Scan(&time)

	fmt.Println("is the weather nice? (true/false)")
	fmt.Scan(&weather)

	fmt.Printf("Hi! My name is %s. I have lived in %s for %d years. They say the weather is amazing, which is %t\n", name, hometown, time, weather)

}
