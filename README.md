# MiniTradosGo
mini traducteur avec Go

### traduire un text normal en langage sms

```go

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

```
<img src="Go.png" width="75%">

### vous pouvez utiliser cette approche dpour traduire des textes en langue local etc.

Disponible en d'autres langages
<hr>
<a href=""> mini traducteur java </a>
<a href="">mini traducteur python </a>
<a href="">mini traducteur kotlin </a>