package main

import (
	"game/pkg"
)

const (
	tries   = 3
	gameVal = 10
)

func main() {

	//pkg.Hangman()

	pkg.Scramble()

	/*
	   var input int

	   	for i := 0; i < tries; i++ {
	   		fmt.Println("Enter a number")
	   		fmt.Scan(&input)

	   		if input == gameVal {
	   			fmt.Println("You won")
	   		}
	   		if input > gameVal {
	   			fmt.Println("Too high")
	   		}
	   		if input < gameVal {
	   			fmt.Println("Too low")
	   		}
	   	}

	   	for i := 0; i < tries; i++ {
	   		fmt.Println("Enter a number")

	   		reader := bufio.NewReader(os.Stdin)
	   		input, _ := reader.ReadString('\n')
	   		input = strings.TrimSpace(input)

	   		fmt.Println(reflect.TypeOf(input))

	   		if input == string(gameVal) {
	   			fmt.Println("You won")
	   		}
	   		if input > gameVal {
	   			fmt.Println("Too high")
	   		}
	   		if input < gameVal {
	   			fmt.Println("Too low")
	   		}
	   	}
	*/
}
