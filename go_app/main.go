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
	if input == "Positive" || input == "positive" {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		randomNumber := r1.Intn(6) + 1
		fmt.Println("Your number is: ", randomNumber)
		fmt.Println()
		fmt.Println("Done.")
	} else if input == "Negative" || input == "negative" {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		randomNumber := r1.Intn((6)+1) * -1
		fmt.Println("Your number is:", randomNumber)
		fmt.Println()
		fmt.Println("Done.")
	} else {
		fmt.Println("Please enter a valid input")
	}
}
