package ascii

import (
	"bufio"
	"os"
	"strings"
)

func Ascii(word string, typee string) string {
	var Filename string

	if typee == "standard" {
		Filename = "files/standard.txt"
	} else if typee == "shadow" {
		Filename = "files/shadow.txt"
	} else if typee == "thinkertoy" {
		Filename = "files/thinkertoy.txt"
	}

	file, err := os.Open(Filename)
	if err != nil {
		return ""
	}
	defer file.Close()

	SliceRune := []string{}
	AsciiMap := map[rune][]string{}
	count := 0
	espace := ' '
	Myscanner := bufio.NewScanner(file)

	for Myscanner.Scan() {
		text := Myscanner.Text()
		if text != "" {
			SliceRune = append(SliceRune, text)
			count++
		}
		if count == 8 {
			AsciiMap[espace] = SliceRune
			espace++
			SliceRune = []string{}
			count = 0
		}
	}

	Splitslice := strings.Split(word, "\n")

	var LastResult string
	if strings.Replace(word, "\n", "", -1) == "" {
		for i := 0; i < strings.Count(word, "\n"); i++ {
			LastResult += "\n"
		}
	}
	LastResult = PrintAscii(Splitslice, AsciiMap)

	return LastResult
}
