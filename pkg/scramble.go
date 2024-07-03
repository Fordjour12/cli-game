package pkg

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var words = []string{"apple", "banana", "cherry", "date", "elderberry", "fig", "grape", "honeydew", "kiwi", "lemon", "mango", "nectarine", "orange", "pear", "quince", "raspberry", "strawberry", "tangerine", "ugli", "watermelon"}

func shuffle(s string) string {
	r := []rune(s)
	rand.Shuffle(len(r), func(i, j int) {
		// shuffling
		r[i], r[j] = r[j], r[i]
	})
	return string(r)
}

func Scramble() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		getWord := words[rand.Intn(len(words))]
		gottenWord := shuffle(getWord)

		fmt.Println("Your shuffled work is: ", gottenWord)

		// give Attemps
		noAttempts := 3
		correct := false
		var userInput string

		for attemp := 0; attemp < noAttempts; attemp++ {
			fmt.Print("Your Guess: ")
			fmt.Scan(&userInput)
			userInput = strings.TrimSpace(userInput)

			if userInput == getWord {
				fmt.Println("You got the correct answer")
				//correct = true
				break
			} else {
				fmt.Println("Incorrect Answer")
			}
		}

		if !correct {
			fmt.Println("sorry the correct answer is: ", getWord)
		}

		fmt.Println("New Word coming..... ")
	}
}
