package pkg

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Hangman() {
	fmt.Println(`
		Welcome to Hangman!
		I have a word for you to guess.
		`)

	words := []string{"programming", "developer", "hangman", "gopher", "computer"}
	maxTries := 6
	tries := 0

	rand.New(rand.NewSource(time.Now().UnixNano()))

	word := words[rand.Intn(len(words))]

	letter := strings.Split(word, "")
	guessed := make([]string, len(letter))

	for i := range letter {
		guessed[i] = "__"
	}

	guessedWord := make(map[string]bool)

	reader := bufio.NewReader(os.Stdin)

	// game state
	printGameState := func() {
		fmt.Println("word: ", strings.Join(guessed, " "))
		fmt.Printf("tries: %d/%d", tries, maxTries)
		fmt.Println("Guessed Letter: ")
		for k := range guessedWord {
			fmt.Printf("%s", k)
		}
		fmt.Println()
	}

	// main game loop

	for tries < maxTries {
		printGameState()

		fmt.Print("Guess a letter: ")
		guess, _ := reader.ReadString('\n')

		guess = strings.TrimSpace(guess)

		// check if the letter is in the word
		if guessedWord[guess] {
			fmt.Printf("You already guessed %s\n", guess)
			continue
		}

		guessedWord[guess] = true

		correct := false
		for i, l := range letter {
			if l == guess {
				correct = true
				guessed[i] = l
			}
		}

		if !correct {
			tries++
			fmt.Println("Incorrect!")
		}

		if strings.Join(guessed, "") == word {
			fmt.Println("You win! the word was ", word)
			return
		}
	}
	fmt.Println("You lose!")
}
