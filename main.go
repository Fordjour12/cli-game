package main

import (
	"game/pkg"
	"strconv"
	"strings"
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

							   	ffor i := 0; i < tries; i++ {
							   		fmt.Println("Enter a number")
							   		fmt.Scan(&input)
						input = strings.TrimSpace(input)

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



		gameval := rand.Intn(100)+1
							   	for i := 0; i < tries; i++ {
							   		fmt.Println("Enter a number")

							   		reader := bufio.NewReader(os.Stdin)
							   		input, _ := reader.ReadString('\n')
							   		input = strings.TrimSpace(input)

				strToInt,err := strconv.Atoi(input)
				if err != nil{
				fmt.Println("Invalid input")
				i--
				}


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
	*/
}
