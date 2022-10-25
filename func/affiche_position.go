package main

import("fmt"
		"os"
	)
	
func affiche_position(str string) {
	position, err := os.ReadFile(str)
	if err != nil {
		fmt.Printf("Fichier introuvable\n")
		return
	}
	fmt.Printf("%c\n", position)
}