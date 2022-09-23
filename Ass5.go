package main

import (
	"fmt"
)

func printNumber(num chan int) {

	for i := 1; i <= 6; i++ {

		num <- i

	}

	close(num)

}

func printAlphabets(alpha chan string) {

	for i := 'A'; i <= 'F'; i++ {

		alpha <- string(i)

	}

	close(alpha)

}

func main() {

	num := make(chan int)

	alpha := make(chan string)

	go printNumber(num)

	go printAlphabets(alpha)

	for {

		v, ok := <-num

		if !ok {

			return

		}

		fmt.Print(v)

		v1, ok1 := <-alpha

		if !ok1 {

			return

		}

		fmt.Print(v1)

	}

}
