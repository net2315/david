package mag

type GameStruct struct { // structure des paramètres de jeu
	WordToFind      string
	WordState       string
	LetterSubmit    rune
	Attempts        int
	NumberOfPos     int
	Win             bool
	EndGame         bool
	Difficulty      string
	Pseudo          string
	ListLetterUsed  []rune
	ListLetterUsed2 string
}
