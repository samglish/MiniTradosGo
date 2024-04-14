package main

import "fmt"
import "strings"

func Translate(phrase string) string {
    motsFR := []string{"bonjour", "salut"} // a completer
    motsFB := []string{"bjr", "slt"} //

    for i := 0; i < len(motsFR); i++ {
        phrase = strings.ReplaceAll(phrase, motsFR[i], motsFB[i])
    }

    return phrase
}

func main() {
text := "bonjour"
    fmt.Println(Translate(text))
}