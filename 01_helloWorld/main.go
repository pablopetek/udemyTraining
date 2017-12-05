package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	for i := 0; i < 2000; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
