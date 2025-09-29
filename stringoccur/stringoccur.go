package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func CharCounter(str string, alphasort bool) {
	charset := make(map[rune]int)
	order := []rune{}
	for _, ch := range str {
		if _, v := charset[ch]; !v {
			order = append(order, ch)
		}
		charset[ch]++
	}

	// sort keys alphabetically
	if alphasort {
		sort.Slice(order, func(i, j int) bool {
			return order[i] < order[j]
		})
	}
	result := ""
	for _, ch := range order {
		result += string(ch) + fmt.Sprint(charset[ch])
	}
	fmt.Println(result)

}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var str string
	var alphasort bool
	fmt.Print("Enter the string:")
	fmt.Scanln(&str)
	for {
		fmt.Print("Do you want to sort in alphabetical order? [y/N]: ")
		response, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("error reading input: %v\n", err)
			os.Exit(1)
		}
		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" || response == "yes" {
			alphasort = true
			break
		} else if response == "n" || response == "no" || response == "" { // Default to No if empty
			alphasort = false
			break
		} else {
			fmt.Println("Invalid input. Please enter 'y' or 'n'.")
		}
	}
	CharCounter(str, alphasort)

}
