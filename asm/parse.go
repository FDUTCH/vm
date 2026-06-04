package asm

import (
	"fmt"
	"slices"
	"strings"
	"vm/vm"
)

func MustParseLine(line string, registry *vm.Registry) (uint64, bool) {
	code, ok, err := ParseLine(line, registry)
	if err != nil {
		panic(err)
	}
	return code, ok
}

func ParseLine(line string, registry *vm.Registry) (uint64, bool, error) {
	fields := strings.Fields(line)
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
	data, ok := registry.FromName(name)
	if !ok {
		return 0, false, fmt.Errorf("unknown opcode: %s", name)
	}
	word, err := Parser{fields[1:]}.Parse(data)
	return word, true, err
}
