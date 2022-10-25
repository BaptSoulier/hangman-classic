package main

func cree_mot(fullWord []string, guess string) []string {
	var word []string
	for i := 0; i < len(fullWord); i++ {
		word = append(word, "_")
	}
	return word
}
