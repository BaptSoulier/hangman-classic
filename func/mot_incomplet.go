package main

func mot_incomplet(fullWord []string) bool {
	for _, char := range fullWord {
		if char == "_" {
			return true
		}
	}
	return false
}
