package vm

import "strings"

type Instruction = func(vm *VM, dst *uint64, src1, src2 uint64, flags uint16, data uint32)

type InstructionData struct {
	id         byte
	ParamCount int
	HasFlags   bool
	HasData    bool
}

func (i InstructionData) Id() byte {
	return i.id
}

func (i InstructionData) InputExample() string {
	builder := strings.Builder{}
	for p := range i.ParamCount {
		switch p {
		case 0:
			builder.WriteString("<dst>")
		case 1:
			builder.WriteString(" <src1>")
		case 2:
			builder.WriteString(" <src2>")
		}
	}
	if i.HasFlags {
		builder.WriteString(" [flags]")
	}
	if i.HasData {
		builder.WriteString(" <data>")
	}
	return builder.String()
}
