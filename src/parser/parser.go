package parser

import (
	"github.com/imajaygiri/parser/src/ast"
	"github.com/imajaygiri/parser/src/lexer"
)

type parser struct {
	// handling error in futures
	// errors []error
	tokens []lexer.Token
	pos    int
}

func createParser(tokens []lexer.Token) *parser {
	return &parser{tokens: tokens, pos: 0}
}

func (p *parser) currentToken() lexer.Token {
	return p.tokens[p.pos]
}
func (p *parser) currentTokenKind() lexer.TokenKind {
	return p.tokens[p.pos].Kind
}

// returns token and moves cursor forward by one
func (p *parser) advance() lexer.Token {
	tk := p.currentToken()
	p.pos++
	return tk
}

func (p *parser) hasTokens() bool {
	return p.pos < len(p.tokens) && p.currentTokenKind() != lexer.EOF
}

func Parse(tokens []lexer.Token) ast.BlockStmt {
	Body := make([]ast.Stmt, 0)
	createTokenLookups()
	p := createParser(tokens)
	for p.hasTokens() {
		Body = append(Body, parse_stmt(p))
	}
	return ast.BlockStmt{
		Body: Body,
	}
}
