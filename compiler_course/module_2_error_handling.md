# Module 2: Diagnostic Error Reporting (File, Line, Col)

In your current code, you have this utility function in [error.go](file:///Users/ajaygiri/code/golang/parser/src/utils/error.go):

```go
func Error(msg string) error {
	pc, file, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	return fmt.Errorf("%s:%d (%s): %s", file, line, fn.Name(), msg)
}
```

### ❌ The Common Mistake
`runtime.Caller(1)` prints the file name, line number, and function of the **Go Compiler program itself**.
If your parser encounters a syntax error in your source file (like `example/01.lang`), this utility will tell you that the error occurred in `parser.go:50`, which is useless for the user writing your programming language. 

A user needs to know: **"Where is the syntax error in my `.lang` file?"**

---

## 1. Tracking Position in the Lexer

To track positions, we must first add location coordinates to every `Token`.

### Step 1: Define `Position` and Update `Token`
Modify [tokens.go](file:///Users/ajaygiri/code/golang/parser/src/lexer/tokens.go) to track metadata:

```go
type Position struct {
	Filename string
	Line     int
	Col      int
}

type Token struct {
	Kind  TokenKind
	Value string
	Pos   Position // Source position where the token starts
}
```

### Step 2: Track Line and Col in the Lexer Loop
In your lexer in [lexer.go](file:///Users/ajaygiri/code/golang/parser/src/lexer/lexer.go), we need to track:
*   The current line (starting at 1)
*   The current column (starting at 1)

Whenever we advance our position in the source string, we must update our column and line numbers:

```go
type lexer struct {
	patterns []regexPattern
	Tokens   []Token
	source   string
	pos      int
	line     int
	col      int
	filename string
}
```

When consuming characters:
*   If we consume a newline character `\n`: we increment `line` and reset `col = 1`.
*   If we consume any other character: we increment `col` by `1`.

Here is an elegant way to advance the lexer and update coordinates:
```go
func (lex *lexer) advanceN(n int) {
	for i := 0; i < n; i++ {
		if lex.pos >= len(lex.source) {
			break
		}
		if lex.source[lex.pos] == '\n' {
			lex.line++
			lex.col = 1
		} else {
			lex.col++
		}
		lex.pos++
	}
}
```

---

## 2. Propagating Positions to the AST

For semantic analysis (type checking) and code generation, we need to report errors like *"Type mismatch: Cannot add string and number"* at a specific variable or expression.

To do this, every AST node must carry position information.

```go
package ast

import "github.com/imajaygiri/parser/src/lexer"

type Node interface {
	Position() lexer.Position
}

type NumberExpr struct {
	Value float64
	Pos   lexer.Position
}

func (n NumberExpr) Position() lexer.Position { return n.Pos }
```

When the parser parses an expression, it grabs the `Pos` of the current token and stores it in the returned AST node:

```go
func parse_primary_expr(p *parser) ast.Expr {
	token := p.currentToken() // Has Pos!
	switch token.Kind {
	case lexer.NUMBER:
		number, _ := strconv.ParseFloat(p.advance().Value, 64)
		return ast.NumberExpr{
			Value: number,
			Pos:   token.Pos, // Propagated!
		}
        ...
```

---

## 3. Formatting Rust-Style Diagnostic Errors

Instead of simple panics, you should print clean compiler messages. Here is how to write a pretty diagnostic printer:

```go
package utils

import (
	"fmt"
	"strings"
	"github.com/imajaygiri/parser/src/lexer"
)

func ReportError(filename string, sourceCode string, pos lexer.Position, message string) {
	lines := strings.Split(sourceCode, "\n")
	
	// Get the line with the error
	errorLine := ""
	if pos.Line - 1 < len(lines) {
		errorLine = lines[pos.Line - 1]
	}

	fmt.Printf("\033[1;31mError:\033[0m %s\n", message)
	fmt.Printf("  --> %s:%d:%d\n", filename, pos.Line, pos.Col)
	fmt.Printf("   |\n")
	fmt.Printf("%2d | %s\n", pos.Line, errorLine)
	
	// Print a caret pointer pointing to the error column
	pointer := strings.Repeat(" ", pos.Col - 1) + "^"
	fmt.Printf("   | %s\n", pointer)
	fmt.Printf("   |\n")
}
```

### Visual Example
If your compiler reads `example/01.lang` and runs into a syntax error at line 29, column 43:
```text
Error: Expected ->[CLOSE_PAREN] but received ->[COLON] instead.
  --> example/01.lang:29:43
   |
29 |   fn isFileRecent(creationTime: Time): boolean {
   |                                      ^
   |
```

---

## 🎓 MIT Professor's Note: Parser Recovery
> "A production-grade compiler does not halt at the first syntax error. It enters a state called **Error Recovery**. Once an error is detected, the parser discards tokens until it finds a **synchronization point** (usually a semicolon `;` or a closing curly brace `}`). It reports the error, recovers its state, and continues parsing the rest of the file to find more errors in a single compiler run."

## 💻 SWE Tips & Tricks
1.  **Keep it lightweight**: Don't store the entire source file in each token. Store a `Pos` struct containing only `Line`, `Col`, and a pointer or key to the file name.
2.  **Tabs vs Spaces**: Be careful when printing caret (`^`) pointers. A tab character `\t` might render as 4 or 8 spaces in a terminal. To keep your pointer aligned, replace tabs with spaces in the printed line or count tabs separately!
3.  **Parser Panic Handling**: If you use panics for syntax errors, you can recover from them at the top level of your parser using Go's `recover()`. This keeps your compiler's CLI execution clean without showing ugly Go runtime stack traces to your users.
