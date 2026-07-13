package main

import (
	"fmt"
	"github.com/imajaygiri/parser/src/lexer"
	"github.com/imajaygiri/parser/src/parser"
	"github.com/sanity-io/litter"
	"os"
)

func main() {
	bytes, err := os.ReadFile("example/05.lang")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf(string(bytes))
	tokens := lexer.Tokenize(string(bytes))
	ast := parser.Parse(tokens)
	litter.Dump(ast)

}
