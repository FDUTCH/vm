package tokenizer

import (
	"fmt"
	"strconv"
	"text/scanner"
	"vm/vm"
)

type InstructionParser struct {
	name                                     string
	line                                     int
	data                                     vm.InstructionData
	registersToParse                         int
	hasFlags, flagSym, dataParsed, dottedRef bool
	Dst                                      byte
	Src1                                     byte
	Src2                                     byte
	Flags                                    uint16
	Data                                     Reference[uint32]
}

func (i InstructionParser) Add(tok rune, str string, registry *vm.Registry, pos scanner.Position) ([]Token, Parser, error) {
	if pos.Line != i.line {
		if i.registersToParse != 0 || (!i.dataParsed && i.data.HasData) {
			return nil, nil, fmt.Errorf("invalid param count for %s", i.name)
		}

		tokens, parser, err := Recognizer{}.Add(tok, str, registry, pos)
		return append([]Token{i.Instruction()}, tokens...), parser, err
	}

	var err error
	if i.registersToParse > 0 {
		i.registersToParse--
		switch i.data.ParamCount - i.registersToParse {
		case 1:
			i.Dst, err = parseRegister(str)
		case 2:
			i.Src1, err = parseRegister(str)
		case 3:
			i.Src2, err = parseRegister(str)
		}
		goto end
	}

	// it means we left part of a data after dot.
	if i.data.HasData && i.dataParsed && tok == '.' && !i.dottedRef && i.Data.NeedsValue() {
		i.dottedRef = true
		goto end
	}

	// adding leftover part of the data.
	if i.dottedRef {
		i.Data = NewReferenceWithName[uint32](i.Data.Name() + "." + str)
		goto end
	}

	// it means that we have a flag to parse.
	if i.hasFlags && tok == '$' {
		i.hasFlags = false
		i.flagSym = true
		goto end
	}

	// parsing flag.
	if i.flagSym {
		i.flagSym = false
		var ok bool
		i.Flags, ok, err = parseFlags(str)
		if err != nil || ok {
			goto end
		}
	}

	// parsing data.
	if i.data.HasData && !i.dataParsed {
		i.dataParsed = true
		i.Data, err = parseData(str)
		goto end
	}

	// if str doesn't match any condition...
	return nil, nil, ErrorUnexpected{str}
end:
	return nil, i, err
}

func (i InstructionParser) Instruction() Instruction {
	return Instruction{
		name:  i.name,
		ID:    i.data.Id(),
		Dst:   i.Dst,
		Src1:  i.Src1,
		Src2:  i.Src2,
		Flags: i.Flags,
		Data:  i.Data,
	}
}

func parseRegister(str string) (byte, error) {
	if str[0] != 'r' {
		return 0, fmt.Errorf("unable to parse register")
	}
	str = str[1:]
	val, err := strconv.ParseUint(str, 0, 4)
	return byte(val), err
}

func parseFlags(str string) (uint16, bool, error) {
	val, err := strconv.ParseUint(str, 0, 12)
	if err == nil {
		return uint16(val), true, nil
	}
	var flags uint16
	for _, char := range str {
		switch char {
		case 'S':
			flags |= vm.FlagStaticRegion
		case 'N':
			flags |= vm.FlagNoneRegion
		case 'K':
			flags |= vm.FlagStackRegion
		case 'U':
			flags |= vm.FlagUnsigned
		case 'F':
			flags |= vm.FlagFloat
		case 'I':
			flags |= vm.FlagInvert
		default:
			return 0, false, fmt.Errorf("unknown flag: %c", char)
		}
	}
	return flags, true, nil
}

func parseData(str string) (Reference[uint32], error) {
	val, err := strconv.ParseInt(str, 0, 33)
	if err != nil {
		return NewReferenceWithName[uint32](str), nil
	}

	data := uint32(val)
	if int64(data) != val {
		return Reference[uint32]{}, fmt.Errorf("integer data overflow")
	}
	return NewReferenceWithValue(data), nil
}

type Instruction struct {
	name  string
	ID    byte
	Dst   byte
	Src1  byte
	Src2  byte
	Flags uint16
	Data  Reference[uint32]
}

func (a Instruction) Name() string {
	return a.name
}
