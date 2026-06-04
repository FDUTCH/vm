package asm

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"vm/vm"
)

func Compile(text string, registry *vm.Registry) ([]byte, error) {
	var (
		out []byte

		writeBuff = make([]byte, 8)
		line      int
	)

	scanner := bufio.NewScanner(bytes.NewBufferString(text))
	for scanner.Scan() {
		line++
		op, ok, err := ParseLine(scanner.Text(), registry)
		if err != nil {
			return nil, fmt.Errorf("error at line: %d (err: %w)", line, err)
		}
		if ok {
			binary.LittleEndian.PutUint64(writeBuff, op)
			out = append(out, writeBuff...)
		}
	}
	return out, nil
}
