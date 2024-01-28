package jose

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func Jose(attempts int) {
	fichier, err := os.ReadFile("jose/hangman.txt") //Avec la fonction "oS" on lit "hangman.txt", la variable "fichier" a pour valeur la lecture de "hangman.txt"
	if err != nil {                                 //Si on ne peut pas lire le fichier on retourne une erreur
		fmt.Println("Erreur lors de l'ouverture du fichier", err)
		return
	}
	skip := (8 * attempts)
	lecture := bufio.NewScanner(bytes.NewReader(fichier)) //On utilise la bibliothèque "bufio" pour lire le fichier ligne par ligne, la variable lecture prend en valeur cela
	nblignes := 0
	for i := 80; i > skip; i-- { //Grace à cette boucle on donne le nombre de ligne à sauter pour accéder au dessin souhaité
		if !lecture.Scan() { //Si on dépasse les limites du fichier on sort de la boucle
			break
		}
	}
	for nblignes < 8 && lecture.Scan() { //On demande a ne lire que 8 lignes gràce a une boucle for, chaque dessin étant de cette taille
		ligne := lecture.Text()
		fmt.Println(ligne) //On affiche la ligne actuellement lu
		nblignes++
	}
}
