package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) < 2 {
		fmt.Println("Please enter a number")
		return
	}

	guess, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid Input")
		return
	}

	if guess < 1 || guess > 100 {
		fmt.Println("Guess the Number between 1-100")
		return
	}

	secret := rand.Intn(100) + 1

	if secret > guess {
		fmt.Printf("Too Low : Secret was %d\n", secret)
	} else if secret < guess {
		fmt.Printf("Too High : Secret was %d\n", secret)
	} else {
		fmt.Println("Correct, You won!!!")
	}
}
