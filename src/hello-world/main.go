package main

import "fmt"

func main() {
	result, err := fmt.Printf("Hola, %v!\n", "Maksimko")
	if (err == nil) {
		fmt.Println("result:", result)
	}
}
