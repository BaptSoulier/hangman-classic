package main

func lettre_correct(fullWord []string, guess string) bool {
	for _, char := range fullWord {
		if char == guess {
			return true
		}
	}
	return false
}
