package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var dictionary = []string{
	"kardelen",
	"sandalye",
	"sigorta",
	"kastamonu",
}
var states = []string{ //The adresses have to be changed
	"./states/hangman0",
	"./states/hangman1",
	"./states/hangman2",
	"./states/hangman3",
	"./states/hangman4",
	"./states/hangman5",
	"./states/hangman6",
	"./states/hangman7",
	"./states/hangman8",
	"./states/hangman9",
}

func main() {
	var inputReader = bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(len(dictionary))
	fullWord := strings.Split(dictionary[random], "")
	gameWord := createGameWord(fullWord, "_")
	guess := ""
	i := 0
	fmt.Printf("%v\n", gameWord)
	showStates(states[0])

	for isWordNotFinished(gameWord) {
		fmt.Printf("Guess the letter\n")
		guess, _ = inputReader.ReadString('\n')
		guess = strings.TrimSpace(guess)

		if len(guess) > 1 { // Use function
			if isWholeWordCorrect(fullWord, guess) {
				fmt.Printf("The whole word is correct!\n")
				break
			} else {
				i++
				fmt.Printf("Hangman states:\n ")
				showStates(states[i])
				if i == len(states)-1 {
					break
				} else {
					continue
				}
			}

		}
		gameWord = changeChars(gameWord, fullWord, guess)

		if isLetterCorrect(gameWord, guess) {
			fmt.Printf("%v\n", gameWord)
			showStates(states[i])
		} else {
			i++
			fmt.Printf("%v\n", gameWord)
			fmt.Printf("Hangman states \n")
			showStates(states[i])
			if i == len(states)-1 {
				break
			}
		}
	}
	if i == len(states)-1 {
		fmt.Printf("You lose!\n")
	} else {
		fmt.Printf("You won!\n")
	}

}
func showStates(str string) {
	state, err := os.ReadFile(str)
	if err != nil {
		fmt.Printf("File is not found\n")
		return
	}
	fmt.Printf("%c\n", state) //format ascii? to character/string
}
func createGameWord(fullWord []string, guess string) []string {
	var word []string
	for i := 0; i < len(fullWord); i++ {
		word = append(word, "_")
	}
	return word
}
func changeChars(gameWord []string, fullWord []string, guess string) []string {
	for i, char := range fullWord {
		if char == guess {
			gameWord[i] = guess
		}
	}
	return gameWord
}
func isWordNotFinished(fullWord []string) bool {
	for _, char := range fullWord {
		if char == "_" {
			return true
		}
	}
	return false
}
func isLetterCorrect(fullWord []string, guess string) bool {
	for _, char := range fullWord {
		if char == guess {
			return true
		}
	}
	return false
}
func isWholeWordCorrect(fullWord []string, guess string) bool {
	str := ""
	for _, char := range fullWord {
		str += char
	}
	return str == guess

}
