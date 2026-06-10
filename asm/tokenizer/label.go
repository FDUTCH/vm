package tokenizer

import (
	"fmt"
	"text/scanner"
	"vm/vm"
)

type DottedLabelParser struct {
	first  string
	second string
}

func (l DottedLabelParser) Add(tok rune, str string, _ *vm.Registry, pos scanner.Position) ([]Token, Parser, error) {
	// finish...
	if pos.Line == -1 {
		return nil, nil, fmt.Errorf("uncompleted label declaration")
	}

	switch tok {
	case scanner.Ident:
		if l.second == "" {
			l.second = str
			return nil, l, nil
		}
	case ':':
		if l.second != "" {
			return []Token{Label{l.first + "." + l.second}}, Recognizer{}, nil
		}
	}
	return nil, nil, ErrorUnexpected{str}
}

type Label struct {
	name string
}

func (l Label) Name() string {
	return l.name
}
