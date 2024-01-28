package handlers

import (
	"Hangman-Web/input"
	"Hangman-Web/lettre"
	"Hangman-Web/mag"
	"Hangman-Web/random"
	"fmt"
	"net/http"
	"text/template"
)

var data mag.GameStruct

const port = ":5500"

func HandleFunc() {
	http.HandleFunc("/", Home)

	http.HandleFunc("/difficulty", Difficulty)

	http.HandleFunc("/pseudo", Pseudo)

	http.HandleFunc("/play/loose", Loose)

	http.HandleFunc("/play/win", Win)

	http.HandleFunc("/play", Game)

	http.HandleFunc("/rules", Rules)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./src/css/"))))
	fmt.Println("http://localhost:5500 - Server started on port :5500")

	http.ListenAndServe(port, nil)
}

func Loose(w http.ResponseWriter, r *http.Request) {
	template2 := template.Must(template.ParseFiles("src/html/loose.html"))

	template2.Execute(w, data)
}

func Win(w http.ResponseWriter, r *http.Request) {

	renderTemplate(w, "src/html/win")

}

func Pseudo(w http.ResponseWriter, r *http.Request) {
	pseudo := r.FormValue("pseudo")
	// Mettre à jour la variable pseudo
	if len(pseudo) > 0 {
		data.Pseudo = pseudo
		http.Redirect(w, r, "/play", http.StatusSeeOther)
	}
	renderTemplate(w, "src/html/pseudo")
}

func Difficulty(w http.ResponseWriter, r *http.Request) {

	data1 := mag.GameStruct{ // la variable data prend en compte la struct de jeu

		Attempts:    10,    //nb d'essais
		NumberOfPos: 10,    //  numero  pos image pendu
		Win:         false, // Victoire ou défaite
		EndGame:     false, // Fin de la partie

	}
	difficulty := r.FormValue("difficulty")
	// Mettre à jour la variable de difficulté
	if len(difficulty) > 0 {
		data1.Difficulty = difficulty

		//Choisis un mot au hasard
		data2 := random.Random(data1)
		//Choisis les lettres à révéler
		data3 := lettre.Lettre(data2)

		data = data3
		http.Redirect(w, r, "/pseudo", http.StatusSeeOther)
	}
	renderTemplate(w, "src/html/difficulty")

}

func Game(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("src/html/play.html"))
	character := r.FormValue("character")
	if len(character) > 0 {
		letter_submit := []rune(character)
		data.LetterSubmit = letter_submit[0]
		data1 := data
		data = input.Input(data1)
	}
	if data.EndGame && data.Win {
		http.Redirect(w, r, "/play/win", http.StatusSeeOther)
	}
	if data.EndGame && !data.Win {
		http.Redirect(w, r, "/play/loose", http.StatusSeeOther)
	}
	template.Execute(w, data)

}
func HandleProcess(w http.ResponseWriter, r *http.Request) rune {
	// Récupérer la valeur du caractère depuis le formulaire
	character := r.FormValue("character")
	letter_submit := []rune(character)
	return letter_submit[0]
}

func Home(w http.ResponseWriter, r *http.Request) { //renvoies à la page d'accueil
	renderTemplate(w, "src/html/Menu_Principal")
}

func Rules(w http.ResponseWriter, r *http.Request) { //renvoies à la page d'accueil
	renderTemplate(w, "src/html/règles")
}

func renderTemplate(w http.ResponseWriter, tmpl string) { // Parser fichie
	t, err := template.ParseFiles("./" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}
