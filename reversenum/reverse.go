package main

import (
	"fmt"
)

func main() {
	var n int
	var ctr int
	fmt.Print("enter the number: ")
	fmt.Scanln(&n)
	dup := n
	rev := 0
	for n > 0 {
		rev = rev*10 + (n % 10)
		n = n / 10
		ctr = ctr + 1
	}
	fmt.Printf("number of digit:%d\n", ctr)
	fmt.Printf("the reversed number: %d\n", rev)
	if rev == dup {
		fmt.Println("palindrome")
	} else {
		fmt.Println("Not palindrome")
	}

}
