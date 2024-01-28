package random

import (
	"Hangman-Web/mag"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func Random(g mag.GameStruct) mag.GameStruct {

	// Chemin d'accès du fichier qui contient la liste de mot
	var file string
	if g.Difficulty == "hard" {
		file = "words3"
	}
	if g.Difficulty == "medium" {
		file = "words2"
	}
	if g.Difficulty == "easy" {
		file = "words"
	}
	path := "./" + file + ".txt"

	// Lis le fichier
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
	} else {
		// Choisis un mot aléatoire de la liste
		mots := strings.Fields(string(content))

		indice := rand.Intn(len(mots))

		motrandom := mots[indice]

		g.WordToFind = string(motrandom)

	}

	return g
}
