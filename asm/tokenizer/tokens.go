package tokenizer

import (
	"text/scanner"
	"vm/vm"
)

type Parser interface {
	Add(tok rune, str string, registry *vm.Registry, pos scanner.Position) ([]Token, Parser, error)
}

type Token interface {
	Name() string
}

type Recognizer struct {
	ident string
}

func (r Recognizer) Add(tok rune, str string, registry *vm.Registry, pos scanner.Position) ([]Token, Parser, error) {
	// finish...
	if pos.Line == -1 {
		if r.ident != "" {
			return nil, nil, ErrorUnexpected{r.ident}
		}
		return nil, nil, nil
	}

	if r.ident == "" {
		switch str {
		case "var":
			return nil, &VariableDecl{}, nil
		case "const":
			return nil, &VariableDecl{Constant: true}, nil
		}

		if tok == scanner.Ident {
			r.ident = str
			return nil, r, nil
		}
		return nil, r, ErrorUnexpected{str}
	}

	switch tok {
	case ':':
		return []Token{Label{r.ident}}, Recognizer{}, nil
	case scanner.Ident:
		data, ok := registry.FromName(r.ident)
		if ok {
			return InstructionParser{name: r.ident, line: pos.Line, data: data, registersToParse: data.ParamCount, hasFlags: data.HasFlags}.Add(tok, str, registry, pos)
		}
	}
	return nil, nil, ErrorUnexpected{str}
}
