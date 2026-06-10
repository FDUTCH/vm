package tokenizer

import (
	"fmt"
	"text/scanner"
	"vm/vm"
)

type VariableDecl struct {
	Constant bool
	state    int
	name     string
	Type     string
	Value    string
}

func NewVariableDecl(constant bool, name string, Type string, value string) VariableDecl {
	return VariableDecl{Constant: constant, name: name, Type: Type, Value: value}
}

func (v VariableDecl) Name() string {
	return v.name
}

func (v VariableDecl) Add(tok rune, str string, _ *vm.Registry, pos scanner.Position) ([]Token, Parser, error) {
	// finish...
	if pos.Line == -1 {
		if v.state != 6 {
			return nil, nil, fmt.Errorf("uncompleted variable declaration")
		}
		return nil, nil, nil
	}

	v.state++
	var err error
	switch v.state {
	case 1: // name
		err = Expect(str, tok, scanner.Ident)
		v.name = str
	case 2: // equal char
		err = Expect(str, tok, '=')
	case 3: // type
		err = Expect(str, tok, scanner.Ident)
		v.Type = str
	case 4: // opening bracket
		err = Expect(str, tok, '(')
	case 5: // value
		err = Expect(str, tok, scanner.String, scanner.Int, scanner.Float, '-')
		if tok == '-' {
			v.state--
		}
		v.Value += str
	case 6: // closing bracket
		err = Expect(str, tok, ')')
		return []Token{v}, Recognizer{}, err
	}
	return nil, v, err
}
