package asm

import (
	"fmt"
	"strconv"
	"strings"
	"vm/vm"
)

type Parser struct {
	params []string
}

func (p Parser) Parse(data vm.InstructionData) (uint64, error) {
	var (
		regs           [3]byte
		flags          uint16
		additionalData uint32

		ok  bool
		err error
	)

	if len(p.params) < data.ParamCount {
		return 0, ErrInvalidInput{data}
	}

	for i := range data.ParamCount {
		regs[i], err = parseRegister(p.params[0])
		if err != nil {
			return 0, err
		}
		p.params = p.params[1:]
	}

	if len(p.params) > 0 && data.HasFlags {
		flags, ok, err = parseFlags(p.params[0])
		if err != nil {
			return 0, err
		}
		if ok {
			p.params = p.params[1:]
		}
	}

	if data.HasData {
		if len(p.params) != 1 {
			return 0, ErrInvalidInput{data}
		}
		val, err := strconv.ParseInt(p.params[0], 0, 33)
		if err != nil {
			return 0, err
		}
		additionalData = uint32(val)

		// check if data fits in 32 bits.
		if int64(additionalData) != val {
			return 0, fmt.Errorf("data integer overflow")
		}
		p.params = p.params[1:]
	}

	if len(p.params) != 0 {
		return 0, ErrInvalidInput{data}
	}

	return Encode(data.Id(), regs[0], regs[1], regs[2], flags, additionalData), err
}

func parseRegister(reg string) (byte, error) {
	regNum := strings.TrimPrefix(reg, "r")
	if regNum == reg {
		return 0, fmt.Errorf("invalid register: %s", reg)
	}
	val, err := strconv.ParseInt(regNum, 10, 5)
	if err != nil {
		return 0, fmt.Errorf("unable to parse register number (err: %w, number: %s)", err, regNum)
	}
	if val < 0 {
		return 0, fmt.Errorf("register number can't be negative")
	}
	return byte(val), nil
}

func parseFlags(str string) (uint16, bool, error) {
	if str == "" || str[0] != '$' {
		return 0, false, nil
	}
	str = str[1:]

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
