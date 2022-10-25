package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var dictionnaire = []string{
	//"./words.txt",
	"banc",
	"bureau",
	"cabinet",
	"carreau",
	"chaise",
	"classe",
	"cle",
	"coin",
	"couloir",
	"dossier",
	"eau",
	"ecole",
	"ecriture",
	"entree",
	"escalier",
}
var position = []string{
	// "./Hangman-Golang/position/hangman0",
	"./position/hangman0",
	"./position/hangman1",
	"./position/hangman2",
	"./position/hangman3",
	"./position/hangman4",
	"./position/hangman5",
	"./position/hangman6",
	"./position/hangman7",
	"./position/hangman8",
	"./position/hangman9",
}

func main() {
	var inputReader = bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(len(dictionnaire))
	fullWord := strings.Split(dictionnaire[random], "")
	gameWord := cree_mot(fullWord, "_")
	guess := ""
	i := 0
	fmt.Printf("%v\n", gameWord)
	affiche_position(position[0])

	for mot_incomplet(gameWord) {
		fmt.Printf("Trouve les lettres\n")
		guess, _ = inputReader.ReadString('\n')
		guess = strings.TrimSpace(guess)

		if len(guess) > 1 {
			if mot_correct(fullWord, guess) {
				fmt.Printf("Le mot est correct. BRAVO !\n")
				break
			} else {
				i++
				fmt.Printf("Hangman position:\n ")
				affiche_position(position[i])
				if i == len(position)-1 {
					break
				} else {
					continue
				}
			}

		}
		gameWord = change_caractere(gameWord, fullWord, guess)

		if lettre_correct(gameWord, guess) {
			fmt.Printf("%v\n", gameWord)
			affiche_position(position[i])
		} else {
			i++
			fmt.Printf("%v\n", gameWord)
			fmt.Printf("Hangman position \n")
			affiche_position(position[i])
			if i == len(position)-1 {
				break
			}
		}
	}
	if i == len(position)-1 {
		fmt.Printf("TU AS PERDU\n")
	} else {
		fmt.Printf("TU AS GAGNE\n")
	}

}
func affiche_position(str string) {
	position, err := os.ReadFile(str)
	if err != nil {
		fmt.Printf("Fichier introuvable\n")
		return
	}
	fmt.Printf("%c\n", position)
}
func cree_mot(fullWord []string, _ string) []string {
	var word []string
	for i := 0; i < len(fullWord); i++ {
		word = append(word, "_")
	}
	return word
}
func change_caractere(gameWord []string, fullWord []string, guess string) []string {
	for i, char := range fullWord {
		if char == guess {
			gameWord[i] = guess
		}
	}
	return gameWord
}
func mot_incomplet(fullWord []string) bool {
	for _, char := range fullWord {
		if char == "_" {
			return true
		}
	}
	return false
}
func lettre_correct(fullWord []string, guess string) bool {
	for _, char := range fullWord {
		if char == guess {
			return true
		}
	}
	return false
}
func mot_correct(fullWord []string, guess string) bool {
	str := ""
	for _, char := range fullWord {
		str += char
	}
	return str == guess

}
