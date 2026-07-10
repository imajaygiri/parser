package main

import (
	"fmt"
	"os"

	"github.com/imajaygiri/parser/src/lexer"
	"github.com/imajaygiri/parser/src/parser"
	"github.com/sanity-io/litter"
)

func main() {
	bytes, err := os.ReadFile("example/03.lang")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	tokens := lexer.Tokenize(string(bytes))
	ast := parser.Parse(tokens)
	litter.Dump(ast)

}
