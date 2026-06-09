package linker

import (
	"bytes"
	"fmt"
	"vm/asm/tokenizer"
)

type Linker struct {
	symbols map[string]uint32

	instructionCount uint32
	variableBuff     []byte
	static, dynamic  *bytes.Buffer

	registry TypeRegistry
}

func New(registry TypeRegistry) *Linker {
	return &Linker{
		symbols: make(map[string]uint32, 10),

		static:       bytes.NewBuffer(nil),
		dynamic:      bytes.NewBuffer(nil),
		variableBuff: make([]byte, 1024),
		registry:     registry,
	}
}

func (l *Linker) Link(tokens []tokenizer.Token) (instructions []tokenizer.Instruction, static, dynamic []byte, err error) {
	// making symbols.
	for _, tok := range tokens {
		switch t := tok.(type) {
		case tokenizer.VariableDecl:
			err = l.handleVariable(t)
		case tokenizer.Instruction:
			l.instructionCount++
		case tokenizer.Label:
			// label is basically a reference to next instruction.
			l.symbols[t.Name()] = l.instructionCount - 1
		}
		if err != nil {
			return nil, nil, nil, err
		}
	}

	// actual linking.
	instructions = make([]tokenizer.Instruction, 0, len(tokens)) // this may be improved but idc.
	for _, tok := range tokens {
		inst, ok := tok.(tokenizer.Instruction)
		if !ok {
			continue
		}

		if inst.Data.NeedsValue() {
			if err = inst.Data.Link(l.symbols); err != nil {
				return nil, nil, nil, err
			}
		}
		instructions = append(instructions, inst)
	}

	return instructions, l.static.Bytes(), l.dynamic.Bytes(), nil
}

func (l *Linker) handleVariable(v tokenizer.VariableDecl) error {
	val, ok := l.registry.Get(v.Type)
	if !ok {
		return fmt.Errorf("unknown type: %s", v.Type)
	}
	if err := val.Parse(v.Value); err != nil {
		return err
	}

	if value, ok := val.Value(); ok {
		// directly storing value.
		l.symbols[v.Name()] = value
	} else {
		count, _ := val.Read(l.variableBuff)
		var addr int
		if v.Constant {
			addr = l.static.Len()
			l.static.Write(l.variableBuff[:count])
		} else {
			addr = l.dynamic.Len()
			l.dynamic.Write(l.variableBuff[:count])
		}
		// storing reference to value.
		l.symbols[v.Name()] = uint32(addr)
	}
	return nil
}
