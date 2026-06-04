package asm

import (
	"fmt"
	"vm/vm"
)

type ErrInvalidInput struct {
	Data vm.InstructionData
}

func (i ErrInvalidInput) Error() string {
	return fmt.Sprintf("invalid instruction input, (expected format: %s)", i.Data.InputExample())
}
