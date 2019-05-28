package parser

import (
	"github.com/akif999/prototypes/go/interpreter_in_go/ast"
	"github.com/akif999/prototypes/go/interpreter_in_go/lexer"
	"github.com/akif999/prototypes/go/interpreter_in_go/token"
)

type Parser struct {
	l *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	// set curToken and peekToken
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
