// Created by: Dominic M.
// Created on: April 2023
//
// This program generates a random number between 1 and 6 or -6 and -1.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var input string

	// input
	fmt.Println("This program generates a random number between 1 and 6 or -6 and -1.")
	fmt.Println()
	fmt.Print("What number do you want? Positive or Negative: ")
	fmt.Println()
	fmt.Scanln(&input)

	// process
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	randomNumber := r1.Intn(6) + 1
	if input == "Positive" || input == "positive" {
		fmt.Println("Your number is: ", randomNumber)
		fmt.Println()
		fmt.Println("\nDone.")
	} else if input == "Negative" || input == "negative" {
		randomNumberFormatted := 0 - randomNumber
		fmt.Println("Your number is:", randomNumberFormatted)
		fmt.Println()
		fmt.Println("\nDone.")
	} else {
		fmt.Println()
		fmt.Println("Please enter a valid input")
		fmt.Println()
		fmt.Println("\nDone.")
	}
}
