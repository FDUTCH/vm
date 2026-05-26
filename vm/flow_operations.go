package vm

import (
	"vm/memory"
)

var _ Instruction = Jump

// Jump performs jump (by dst value).
func Jump(vm *VM, dst *uint32, _, _ uint32, flags uint16) {
	if (flags & FlagUnsigned) != 0 {
		vm.registers[memory.ProgramCounterReg] += *dst
		return
	}
	*signed(&vm.registers[memory.ProgramCounterReg]) += *signed(dst)
}

var _ Instruction = JumpConditional

// JumpConditional performs conditional jump (checking value of src).
func JumpConditional(vm *VM, dst *uint32, src, _ uint32, flags uint16) {
	if ((flags & FlagInvert) != 0) != (src == 0) {
		return
	}

	if (flags & FlagUnsigned) != 0 {
		vm.registers[memory.ProgramCounterReg] += *dst
		return
	}
	*signed(&vm.registers[memory.ProgramCounterReg]) += *signed(dst)
}
