package main

import "fmt"

func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func PrimeGenerator(n int, ch chan int) {
	defer close(ch)
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			ch <- i
		}
	}
}

func PrintPrimes(ch chan int, done chan bool) {
	for v := range ch {
		fmt.Printf("Printed Prime: %v\n", v)
	}
	done <- true

}
func main() {
	var num int
	fmt.Print("Enter the Number:")
	fmt.Scanln(&num)
	ch := make(chan int)
	done := make(chan bool)
	go PrimeGenerator(num, ch)
	go PrintPrimes(ch, done)
	<-done
}
