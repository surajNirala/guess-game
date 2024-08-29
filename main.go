package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	target := r.Intn(100) + 1
	fmt.Println(target)
	Describe()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter your guess:")
		//Read user input

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred while reading input. Please try again.", err)
			continue
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil || guess < 0 || guess > 100 {
			fmt.Println("Invalid input. Please enter a number between 1 to 100.")
			continue
		}
		if guess < target {
			fmt.Println("Too Low!")
		} else if guess > target {
			fmt.Println("To High!")
		} else {
			fmt.Println("Congratulations! you guessed the correct number.")
			break
		}
	}
	fmt.Println("Thanks for playing the Number Guessing Game!")
}

func Describe() {
	fmt.Println("Welcome to the Number Guessing Game!!!")
	fmt.Println("I have selected a random number between 1 to 100")
	fmt.Println("Can you guess whta it is?")
}
