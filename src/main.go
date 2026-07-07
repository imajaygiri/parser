package main

import (
	// "fmt"
	"github.com/imajaygiri/parser/src/lexer"
	// "os"
)

func main() {
	// bytes, err := os.ReadFile("example/00.lang")
	// if err != nil {
	// 	fmt.Printf("Error: %v\n", err)
	// }
	// fmt.Printf("source: %s\n", string(bytes))

	tokens := lexer.Tokenize("3 + 5 * 2 / 2.5")

	for _, token := range tokens {
		token.Debug()
		println()
	}
}
