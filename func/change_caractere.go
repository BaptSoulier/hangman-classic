package main

func change_caractere(gameWord []string, fullWord []string, guess string) []string {
	for i, char := range fullWord {
		if char == guess {
			gameWord[i] = guess
		}
	}
	return gameWord
}
