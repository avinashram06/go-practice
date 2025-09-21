package main

import "fmt"

func main() {
	var i, j int
	fmt.Print("enter the first number:")
	fmt.Scanln(&i)
	fmt.Print("\nenter the second number:")
	fmt.Scanln(&j)
	for i > 0 && j > 0 {
		if i > j {
			i = i % j
		} else {
			j = j % i
		}
	}
	if i == 0 {
		fmt.Printf("GCD: %d", j)
	} else {
		fmt.Printf("GCD: %d", i)
	}

}
