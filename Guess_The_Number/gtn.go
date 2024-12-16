package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	target := rand.Intn(100) + 1
	attempts := 0
	var guess int

	fmt.Println("Welcome to our Guess the Number Game")
	fmt.Println("I am thinking of a number between 1 and 100. can you guess it?")

	for {
		fmt.Println("Enter your guess: ")
		fmt.Scanln(&guess)
		attempts++

		if guess < attempts {
			fmt.Println("Too low! Try again!")
		} else if guess > target {
			fmt.Println("Too high! Try again!")
		} else {
			fmt.Printf("Congratulations! You guessed the number %d with %d attempts.\n", target, attempts)
			break
		}
	}
}
