package asm

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
)

func Compile(text string) ([]byte, error) {
	var (
		out []byte

		writeBuff = make([]byte, 4)
		line      int
	)

	scanner := bufio.NewScanner(bytes.NewBufferString(text))
	for scanner.Scan() {
		line++
		op, ok, err := ParseLine(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("error at line: %d (err: %w)", line)
		}
		if ok {
			binary.LittleEndian.PutUint32(writeBuff, op)
			out = append(out, writeBuff...)
		}
	}
	return out, nil
}
