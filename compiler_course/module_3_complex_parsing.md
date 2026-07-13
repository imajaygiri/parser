# Module 3: Parsing Complex Structures (Functions, Methods, & Classes)

While expressions (like arithmetic and member access) are best parsed using a **Pratt Parser**, top-level declarations and statements (like functions, classes, and loops) are best parsed using **Recursive Descent**. 

Why? Because declarations always start with a clear keyword (`fn`, `class`, `if`, `while`) and do not have precedence/binding-power conflicts.

Let's design and parse these complex structures.

---

## 1. Grammars for Functions and Classes

Before writing parser code, we must specify our syntax using Extended Backus-Naur Form (EBNF):

```ebnf
FunctionDecl ::= "fn" Identifier "(" ParameterList? ")" (":" Type)? BlockStmt
ParameterList ::= Parameter ("," Parameter)*
Parameter     ::= Identifier ":" Type

ClassDecl    ::= "class" Identifier "{" ClassMember* "}"
ClassMember  ::= ("let" PropertyName ":" Type ";") 
               | ("static"? "fn" MethodName "(" ParameterList? ")" (":" Type)? BlockStmt)
```

---

## 2. Defining AST Nodes

Let's translate these grammar rules into Go structs.

### Functions and Parameters
```go
package ast

type Parameter struct {
	Name string
	Type Type
}

type FunctionDeclStmt struct {
	Name       string
	Parameters []Parameter
	ReturnType Type // nil if void
	Body       BlockStmt
}

func (n FunctionDeclStmt) stmt() {}
```

### Classes and Methods
```go
type ClassProperty struct {
	Name     string
	Type     Type
	IsStatic bool
}

type ClassMethod struct {
	Name       string
	Parameters []Parameter
	ReturnType Type
	Body       BlockStmt
	IsStatic   bool
}

type ClassDeclStmt struct {
	Name       string
	Properties []ClassProperty
	Methods    []ClassMethod
}

func (n ClassDeclStmt) stmt() {}
```

---

## 3. Writing the Parser Functions

Let's write the parsing logic for these nodes.

### Parsing Functions (`fn add(a: number, b: number): number { ... }`)
```go
func parse_function_decl(p *parser) ast.Stmt {
	p.expect(lexer.FN) // Consume 'fn'
	
	funcName := p.expect(lexer.INDENTIFIER).Value
	p.expect(lexer.OPEN_PAREN)
	
	parameters := []ast.Parameter{}
	for p.currentTokenKind() != lexer.CLOSE_PAREN {
		paramName := p.expect(lexer.INDENTIFIER).Value
		p.expect(lexer.COLON)
		paramType := parse_type(p, default_bp)
		
		parameters = append(parameters, ast.Parameter{
			Name: paramName,
			Type: paramType,
		})
		
		if p.currentTokenKind() != lexer.CLOSE_PAREN {
			p.expect(lexer.COMMA)
		}
	}
	p.expect(lexer.CLOSE_PAREN)
	
	var returnType ast.Type
	if p.currentTokenKind() == lexer.COLON {
		p.advance() // Consume ':'
		returnType = parse_type(p, default_bp)
	}
	
	body := parse_block_statement(p)
	
	return ast.FunctionDeclStmt{
		Name:       funcName,
		Parameters: parameters,
		ReturnType: returnType,
		Body:       body,
	}
}
```

### Parsing Classes (`class DirectoryReader { ... }`)
```go
func parse_class_decl(p *parser) ast.Stmt {
	p.expect(lexer.CLASS) // Consume 'class'
	
	className := p.expect(lexer.INDENTIFIER).Value
	p.expect(lexer.OPEN_CURLY)
	
	var properties []ast.ClassProperty
	var methods []ast.ClassMethod
	
	for p.currentTokenKind() != lexer.CLOSE_CURLY {
		isStatic := false
		if p.currentTokenKind() == lexer.STATIC {
			p.advance()
			isStatic = true
		}
		
		// If it's a property: let x: number;
		if p.currentTokenKind() == lexer.LET {
			p.advance() // Consume 'let'
			propName := p.expect(lexer.INDENTIFIER).Value
			p.expect(lexer.COLON)
			propType := parse_type(p, default_bp)
			p.expect(lexer.SEMI_COLON)
			
			properties = append(properties, ast.ClassProperty{
				Name:     propName,
				Type:     propType,
				IsStatic: isStatic,
			})
		} else if p.currentTokenKind() == lexer.FN {
			// If it's a method: fn hello() { ... }
			methodStmt := parse_function_decl(p).(ast.FunctionDeclStmt)
			methods = append(methods, ast.ClassMethod{
				Name:       methodStmt.Name,
				Parameters: methodStmt.Parameters,
				ReturnType: methodStmt.ReturnType,
				Body:       methodStmt.Body,
				IsStatic:   isStatic,
			})
		} else {
			panic(fmt.Sprintf("Unexpected token in class body: %s", p.currentTokenKind().ToString()))
		}
	}
	
	p.expect(lexer.CLOSE_CURLY)
	return ast.ClassDeclStmt{
		Name:       className,
		Properties: properties,
		Methods:    methods,
	}
}
```

---

## 🎓 MIT Professor's Note: Blocks and Scopes
> "When parsing block statements `{ ... }`, you are creating nested lexical scopes. In your AST, a block statement contains a list of statements (`[]Stmt`). Later, during semantic analysis, each block statement node will trigger the creation of a new symbol table parented to the outer scope's symbol table. This ensures variables defined inside a block don't leak out."

## 💻 SWE Tips & Tricks
1.  **Reusability**: Notice how `parse_class_decl` doesn't re-implement function parsing. It calls `parse_function_decl` and casts the returned statement to `ast.FunctionDeclStmt` to create a `ClassMethod`. This avoids duplicated parsing logic.
2.  **Semicolons**: Decide early whether semicolons are required after function definitions. In languages like JS/TS, semicolons are optional after functions/methods. In our parser, we enforce semicolons for properties (`let x: string;`) but omit them after blocks (`fn hello() {}`).
3.  **Forward Parsing**: When parsing nested arrays or types (like `[][]number`), your type parser should run recursively. Your current [types.go](file:///Users/ajaygiri/code/golang/parser/src/parser/types.go#L64) already uses recursive calls to `parse_type` nicely to achieve this.
