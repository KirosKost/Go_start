// Пакет strings содержит тип Replacer, который ищет под-
// строку в строке и заменяет каждое вхождение этой подстроки
// в другой строке. Следующий код заменяет каждый символ #
// в строке буквой o:

package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "G# r#cks!"
	rep := "G# is Str#ng lang"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	repl := replacer.Replace(rep)
	fmt.Println(fixed)
	fmt.Println(repl)

}
