package asm

import (
	"fmt"
	"slices"
	"strings"
)

func MustParseLine(line string) (uint32, bool) {
	code, ok, err := ParseLine(line)
	if err != nil {
		panic(err)
	}
	return code, ok
}

func ParseLine(line string) (uint32, bool, error) {
	fields := strings.Fields(strings.ToLower(line))
	idx := slices.IndexFunc(fields, func(s string) bool {
		return strings.HasPrefix(s, "//")
	})
	if idx != -1 {
		fields = fields[:idx]
	}

	if len(fields) == 0 {
		return 0, false, nil
	}
	name := fields[0]
	opcode, ok := instructions[name]
	if !ok {
		return 0, false, fmt.Errorf("unknown opcode: %s", name)
	}
	op, err := opcode.Parse(fields[1:])
	return op, true, err
}
