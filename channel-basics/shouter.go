package main

import (
	"fmt"
	"strings"
)

func shout(in chan string, out chan string) {
	s := <-in

	out <- fmt.Sprintf("HEY %s!!!", strings.ToUpper(s))

}
func main() {
	in := make(chan string)
	out := make(chan string)

	go shout(in, out)

	fmt.Println("Enter your name or Enter Q to Quit")

	for {
		var userInput string

		fmt.Print("-> ")
		fmt.Scanln(&userInput)
		if userInput == strings.ToLower("q") {
			break
		}
		in <- userInput

		response := <-out
		fmt.Println("Response: ", response)

	}
	fmt.Println("Closing channels")
	close(in)
	close(out)
}
