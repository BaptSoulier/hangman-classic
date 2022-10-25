package main

func mot_correct(fullWord []string, guess string) bool {
	str := ""
	for _, char := range fullWord {
		str += char
	}
	return str == guess

}
