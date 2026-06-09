package vm

import (
	"vm/memory"
)

var _ Instruction = JumpRel

// JumpRel performs relative jump (by dst value).
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

var _ Instruction = Jump

// JumpImmediate performs jump to dst data.
func JumpImmediate(vm *VM, _ *uint64, _, _ uint64, _ uint16, data uint32) {
	vm.registers[memory.ProgramCounterReg] = uint64(data)
}

var _ Instruction = JumpRelImmediate

// JumpRelImmediate performs relative jump by data value.
func JumpRelImmediate(vm *VM, _ *uint64, _, _ uint64, flags uint16, data uint32) {
	if (flags & FlagUnsigned) != 0 {
		vm.registers[memory.ProgramCounterReg] += uint64(data)
		return
	}
	vm.registers[memory.ProgramCounterReg] = uint64(int64(vm.registers[memory.ProgramCounterReg]) + int64(data))
}

var _ Instruction = JumpConditional

// JumpConditional performs conditional jump (checking value of src).
func JumpConditional(vm *VM, dst *uint64, src, _ uint64, flags uint16, _ uint32) {
	if ((flags & FlagInvert) != 0) != (src == 0) {
		return
	}

	vm.registers[memory.ProgramCounterReg] = *dst
}

var _ Instruction = JumpRelConditional

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

var _ Instruction = JumpConditionalImmediate

// JumpConditionalImmediate performs conditional jump to data value (checking value of dst).
func JumpConditionalImmediate(vm *VM, dst *uint64, _, _ uint64, flags uint16, data uint32) {
	if ((flags & FlagInvert) != 0) != (*dst == 0) {
		return
	}

	vm.registers[memory.ProgramCounterReg] = uint64(data)
}

var _ Instruction = JumpRelConditionalImmediate

// JumpRelConditionalImmediate performs conditional jump by data value (checking value of src).
func JumpRelConditionalImmediate(vm *VM, dst *uint64, _, _ uint64, flags uint16, data uint32) {
	if ((flags & FlagInvert) != 0) != (*dst == 0) {
		return
	}

	if (flags & FlagUnsigned) != 0 {
		vm.registers[memory.ProgramCounterReg] += uint64(data)
		return
	}
	vm.registers[memory.ProgramCounterReg] = uint64(int64(vm.registers[memory.ProgramCounterReg]) + int64(data))
}
