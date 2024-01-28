package lettre

import (
	"Hangman-Web/mag"
	"math/rand"
)

func Lettre(g mag.GameStruct) mag.GameStruct {
	lettrerandom_rune := []rune(g.WordToFind)
	var reveal []rune
	b := 0
	var position []int
	// Choisit le nombre de lettre à révéler
	n := len(lettrerandom_rune)/2 - 1
	// Choisit des lettres au hasard et fais en sorte de ne pas tomber sur la meme
	for i := 0; i < n; i++ {
		b = 0
		a := rand.Intn(len(lettrerandom_rune))
		for j := 0; j < len(reveal); j++ {
			if lettrerandom_rune[a] == reveal[j] {
				b = b + 1
			}
		}
		if b > 0 {
			i--
		} else {
			reveal = append(reveal, lettrerandom_rune[a])
			position = append(position, a)
		}
	}
	//Révèle les lettres en plusieurs exemplaires
	// si l'une de celle choisie au début est contenue plusieurs fois dans le mot
	var total []rune
	var count int
	for i := 0; i < len(reveal); i++ {
		for j := 0; j < len(lettrerandom_rune); j++ {
			if lettrerandom_rune[j] == reveal[i] {
				count = count + 1
			}
		}
		for k := 0; k < count; k++ {
			total = append(total, reveal[i])
		}

		count = 0
	}

	//Place les lettres qui sont bonnes aux bons endroits
	var motjeu []rune
	underscore := '_'
	space := ' '
	var count2 int
	for k := 0; k < len(lettrerandom_rune); k++ {
		count2 = 0
		for m := 0; m < len(total); m++ {
			if lettrerandom_rune[k] == total[m] {
				count2 = count2 + 1

			}
		}
		if count2 == 0 {
			motjeu = append(motjeu, underscore)
			motjeu = append(motjeu, space)
		} else {
			motjeu = append(motjeu, lettrerandom_rune[k])
			motjeu = append(motjeu, space)
		}
	}
	// Affiche le mot de depart

	g.WordState = string(motjeu)
	//input.Input(lettrerandom_rune, motjeu)
	return g
}
