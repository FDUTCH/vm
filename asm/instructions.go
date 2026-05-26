package asm

import (
	"fmt"
	"strconv"
	"strings"
	"vm/vm"
)

type Instruction struct {
	ID     byte
	Params int
}

var instructions = map[string]Instruction{
	"inc":     {vm.OpInc, 1},
	"dec":     {vm.OpDec, 1},
	"ldi":     {vm.OpLoadUInt16Immediate, 1},
	"ld":      {vm.OpLoadUint32, 2},
	"lds":     {vm.OpLoadUint16, 2},
	"ldb":     {vm.OpLoadByte, 2},
	"mov":     {vm.OpMove, 2},
	"mova":    {vm.OpMoveToAddrUint32, 2},
	"movas":   {vm.OpMoveToAddrUint16, 2},
	"movab":   {vm.OpMoveToAddrByte, 2},
	"addi":    {vm.OpAddUint12Immediate, 1},
	"subi":    {vm.OpSubUint12Immediate, 1},
	"add":     {vm.OpAdd, 3},
	"sub":     {vm.OpSub, 3},
	"mul":     {vm.OpMul, 3},
	"div":     {vm.OpDiv, 3},
	"jump":    {vm.OpJump, 1},
	"jumpc":   {vm.OpJumpConditional, 2},
	"equal":   {vm.OpEqual, 3},
	"grt":     {vm.OpGreater, 3},
	"call":    {vm.OpCall, 1},
	"ret":     {vm.OpRet, 0},
	"pushi":   {vm.OpPushImmediate, 1},
	"pushs":   {vm.OpPushUint16, 2},
	"pops":    {vm.OpPopUint16, 1},
	"push":    {vm.OpPushUint32, 2},
	"pop":     {vm.OpPopUint32, 1},
	"pushv":   {vm.OpPush, 3},
	"popv":    {vm.OpPop, 2},
	"write":   {vm.OpWrite, 2},
	"read":    {vm.OpRead, 2},
	"wstatus": {vm.OpWriteStatus, 1},
	"lstatus": {vm.OpLoadStatus, 1},
}

func (in Instruction) Parse(params []string) (uint32, error) {
	var (
		regs  [3]byte
		flags uint16

		err error
	)

	if len(params) < in.Params {
		return 0, fmt.Errorf("not enough params (expected count: %d, actual count: %d)", in.Params, len(params))
	}

	for i := range in.Params {
		regs[i], err = parseRegister(params[0])
		if err != nil {
			return 0, err
		}
		params = params[1:]
	}

	switch len(params) {
	case 0:
		// no flags.
	case 1:
		flags, err = parseFlags(params[0])
	default:
		return 0, fmt.Errorf("invalid param count")
	}

	return Encode(in.ID, regs[0], regs[1], regs[2], flags), err
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

func parseFlags(str string) (uint16, error) {
	val, err := strconv.ParseInt(str, 10, 13)
	if err == nil {
		if val < 0 {
			return 0, fmt.Errorf("flags should not be negative")
		}
		return uint16(val), nil
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
			return 0, fmt.Errorf("unknown flag: %c", char)
		}
	}
	return flags, nil
}
