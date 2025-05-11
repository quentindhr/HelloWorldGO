package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	text := AskText()
	ChangeColor(text)

}

func AskText() string {
	fmt.Println("Entrez le texte :")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')

	if err != nil {
		return fmt.Sprint("Erreur de lecture : ", err)
	}

	return line
}

func ChangeColor(t string) {
	r := YesOrNo()
	bool := CorrectText(r)
	for !bool {
		r = YesOrNo()
		bool = CorrectText(r)
	}

	if strings.Contains(r, "o") {
		c := strings.Trim(WhichColor(), "\n")
		switch c {
		case "rouge":
			red := color.New(color.FgRed).SprintFunc()
			fmt.Println("Texte saisi : ", red(t))
		case "vert":
			green := color.New(color.FgGreen).SprintFunc()
			fmt.Println("Texte saisi : ", green(t))
		case "jaune":
			yellow := color.New(color.FgYellow).SprintFunc()
			fmt.Println("Texte saisi : ", yellow(t))
		case "bleu":
			blue := color.New(color.FgBlue).SprintFunc()
			fmt.Println("Texte saisi : ", blue(t))
		default:
			fmt.Println("Texte saisi : ", t)
			fmt.Println("Couleur non reconnue, affichage du texte sans couleur.")
			ChangeColor(t)
		}
	} else {
		fmt.Println("Texte saisi : ", t)
	}

}

func YesOrNo() string {
	fmt.Println("Voulez-vous changer la couleur ? (o/n)")

	reader := bufio.NewReader(os.Stdin)

	line, err := reader.ReadString('\n')

	if err != nil {
		return "Erreur"
	}

	return line

}

func CorrectText(text string) bool {
	if strings.Contains(text, "o") || strings.Contains(text, "n") {
		return true
	}
	return false
}

func WhichColor() string {
	fmt.Println("Quelle couleur ?")

	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')

	if err != nil {
		return "Erreur"
	}

	return line
}
