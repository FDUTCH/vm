package asm

import (
	"fmt"
	"strings"
	"text/scanner"
	"vm/asm/tokenizer"
	"vm/vm"
)

type Assembler struct {
	registry *vm.Registry
}

func NewAssembler(registry *vm.Registry) *Assembler {
	return &Assembler{registry: registry}
}

func (a *Assembler) Lex(text string) (tokens []tokenizer.Token, err error) {
	var s scanner.Scanner
	s.Init(strings.NewReader(text))
	s.Mode = scanner.ScanIdents | scanner.ScanInts | scanner.ScanChars | scanner.ScanStrings | scanner.ScanFloats | scanner.ScanComments
	s.Whitespace = scanner.GoWhitespace

	var (
		parser       tokenizer.Parser = tokenizer.Recognizer{}
		resultTokens []tokenizer.Token
	)

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		if tok == scanner.Comment {
			continue
		}
		olaParser := parser
		if resultTokens, parser, err = parser.Add(tok, s.TokenText(), a.registry, s.Pos()); err != nil {
			return nil, fmt.Errorf("%T error at: %s (err: %w)", olaParser, s.Pos(), err)
		}
		if resultTokens != nil {
			tokens = append(tokens, resultTokens...)
		}
	}

	// end.
	if parser != nil && resultTokens == nil {
		resultTokens, _, err = parser.Add(0, "", nil, scanner.Position{Line: -1})
		tokens = append(tokens, resultTokens...)
	}

	return tokens, err
}
