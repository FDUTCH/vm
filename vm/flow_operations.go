package vm

import (
	"vm/memory"
)

var _ Instruction = JumpRel

// JumpRel performs jump (by dst value).
func JumpRel(vm *VM, dst *uint64, _, _ uint64, flags uint16, _ uint32) {
	if (flags & FlagUnsigned) != 0 {
		vm.registers[memory.ProgramCounterReg] += *dst
		return
	}
	vm.registers[memory.ProgramCounterReg] = uint64(int64(vm.registers[memory.ProgramCounterReg]) + int64(*dst))
}

var _ Instruction = Jump

// Jump performs jump (to dst value).
func Jump(vm *VM, dst *uint64, _, _ uint64, _ uint16, _ uint32) {
	vm.registers[memory.ProgramCounterReg] = *dst
}

var _ Instruction = JumpConditional

// JumpConditional performs conditional jump (checking value of src).
func JumpConditional(vm *VM, dst *uint64, src, _ uint64, flags uint16, _ uint32) {
	if ((flags & FlagInvert) != 0) != (src == 0) {
		return
	}

	vm.registers[memory.ProgramCounterReg] = *dst
}

var _ Instruction = JumpConditional

// JumpRelConditional performs conditional jump by dst value (checking value of src).
func JumpRelConditional(vm *VM, dst *uint64, src, _ uint64, flags uint16, _ uint32) {
	if ((flags & FlagInvert) != 0) != (src == 0) {
		return
	}

	if (flags & FlagUnsigned) != 0 {
		vm.registers[memory.ProgramCounterReg] += *dst
		return
	}
	vm.registers[memory.ProgramCounterReg] = uint64(int64(vm.registers[memory.ProgramCounterReg]) + int64(*dst))
}
