package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(findFirstStringInBracket("Indonesia, disebut juga dengan Negara Kesatuan Republik Indonesia (NKRI)"))
}

func findFirstStringInBracket(str string) string {
	if len(str) == 0 {
		return ""
	}

	strRunes := []rune(str)
	indexFirstBracketFound := strings.Index(str, "(")
	if indexFirstBracketFound < 0 {
		return ""
	}

	wordsAfterFirstBracket := string(strRunes[indexFirstBracketFound:len(str)])
	indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
	if indexClosingBracketFound < 0 {
		return ""
	}

	wordsInsideBracketsRunes := []rune(wordsAfterFirstBracket)
	wordsInsideBrackets := string(wordsInsideBracketsRunes[1:indexClosingBracketFound])

	return wordsInsideBrackets
}
