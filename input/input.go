package input

import (
	"Hangman-Web/mag"
)

func Input(g mag.GameStruct) mag.GameStruct {
	status2 := false
	//lettrerandom_rune []rune, motjeu []rune
	lettrerandom_rune := []rune(g.WordToFind)
	motjeu := []rune(g.WordState)
	attempts := g.Attempts
	//var lettre string
	var count3 int
	var count4 bool
	count4 = false

	lettre2 := []rune(string(g.LetterSubmit))

	// Analyse l'input du joueur

	// S'il écrit plus q'un caractère
	if len(lettre2) > 1 {
		count4 = true

		// Si l'input est un caractère mais qu'il n'est pas compris entre a et z
	} else if lettre2[0] > 'z' || lettre2[0] < 'a' {
		count4 = true

		// Si l'input est compris entre a et z
	} else {
		a := lettre2[0]
		for i := 0; i < len(lettrerandom_rune); i++ {
			// Vérifie si la lettre proposée est dans le mot
			if a == lettrerandom_rune[i] {
				motjeu[i*2] = a
				count4 = true
			}
		}

		for i := 0; i < len(g.ListLetterUsed); i++ {
			if g.ListLetterUsed[i] == g.LetterSubmit {
				status2 = true
			}
		}
		// Si la lettre n'a pas déja était utilisée
		if !status2 {
			g.ListLetterUsed = append(g.ListLetterUsed, g.LetterSubmit)
			g.ListLetterUsed = append(g.ListLetterUsed, ' ')
			g.ListLetterUsed2 = string(g.ListLetterUsed)
		}

	}
	if status2 {
		return g
	}

	g.WordState = string(motjeu)
	// Si la lettre n'est pas dans le mot on affiche le dessin du pendu
	if !count4 {
		attempts = attempts - 1
		g.Attempts = attempts
	}
	// Vérifie si le nombre d'essais n'est pas egal à 0
	if attempts == 0 {
		g.EndGame = true
	}

	//vérifier si le mot a été trouver
	count3 = 0
	for j := 0; j < len(motjeu); j++ {
		if motjeu[j] == '_' {
			count3 = count3 + 1
		}
	}
	// Vérifie si le mot a été trouvé
	if count3 == 0 {
		g.Win = true
		g.EndGame = true
	}
	return g

}
