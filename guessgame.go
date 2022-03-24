// guess challenges players to guess a random number.

package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between ! and 100.")
	fmt.Println("can you guess it?")

	reader := bufio.NewReader(os.Stdin)

	sucess := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")

		fmt.Print("Make a guess.")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else if guess > target {
			fmt.Println("OOps. your guess was High")
		} else {
			fmt.Println("Good job! You guessed it.")
			break
		}
	}
	if !sucess {
		fmt.Println("Sorry, you didn't guess my number. It was:", target)
	}

}
