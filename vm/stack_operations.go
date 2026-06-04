package vm

import "vm/memory"

var _ Instruction = Call

// Call calls function from dst addr.
func Call(vm *VM, dst *uint64, _, _ uint64, _ uint16, _ uint32) {
	currentAddr := vm.registers[memory.ProgramCounterReg]
	vm.stack.PushUint64(currentAddr)
	vm.registers[memory.ProgramCounterReg] = *dst
}

var _ Instruction = Ret

// Ret returns from the function.
func Ret(vm *VM, _ *uint64, _, _ uint64, _ uint16, _ uint32) {
	vm.registers[memory.ProgramCounterReg] = vm.stack.PopUint64()
}

var _ Instruction = PushImmediate

// PushImmediate pushes uint32 data to the stack (writing address to dst).
func PushImmediate(vm *VM, dst *uint64, _, _ uint64, _ uint16, data uint32) {
	*dst = uint64(len(vm.stack.Data))
	vm.stack.PushUint32(data)
}

var _ Instruction = PushByte

// PushByte pushes uint16 src to the stack (writing address to dst).
func PushByte(vm *VM, dst *uint64, src, _ uint64, _ uint16, _ uint32) {
	*dst = uint64(len(vm.stack.Data))
	vm.stack.PushUint16(uint16(src))
}

var _ Instruction = PopByte

// PopByte pops uint16 into dst.
func PopByte(vm *VM, dst *uint64, _, _ uint64, _ uint16, _ uint32) {
	*dst = uint64(vm.stack.PopUint16())
}

var _ Instruction = PushUint16

// PushUint16 pushes uint16 src to the stack (writing address to dst).
func PushUint16(vm *VM, dst *uint64, src, _ uint64, _ uint16, _ uint32) {
	*dst = uint64(len(vm.stack.Data))
	vm.stack.PushUint16(uint16(src))
}

var _ Instruction = PopUint16

// PopUint16 pops uint16 into dst.
func PopUint16(vm *VM, dst *uint64, _, _ uint64, _ uint16, _ uint32) {
	*dst = uint64(vm.stack.PopUint16())
}

var _ Instruction = PushUint32

// PushUint32 pushes dst to stack (writing address to dst).
func PushUint32(vm *VM, dst *uint64, src, _ uint64, _ uint16, _ uint32) {
	*dst = uint64(len(vm.stack.Data))
	vm.stack.PushUint32(uint32(src))
}

var _ Instruction = PopUint32

// PopUint32 pops uint32 into dst.
func PopUint32(vm *VM, dst *uint64, _, _ uint64, _ uint16, _ uint32) {
	*dst = uint64(vm.stack.PopUint32())
}

var _ Instruction = PushUint64

// PushUint64 pushes dst to stack (writing address to dst).
func PushUint64(vm *VM, dst *uint64, src, _ uint64, _ uint16, _ uint32) {
	*dst = uint64(len(vm.stack.Data))
	vm.stack.PushUint64(src)
}

var _ Instruction = PopUint64

// PopUint64 pops uint64 into dst.
func PopUint64(vm *VM, dst *uint64, _, _ uint64, _ uint16, _ uint32) {
	*dst = uint64(vm.stack.PopUint64())
}

var _ Instruction = Push

// Push pushes data from memory.Region resolved from flags and writing address to addrRet.
func Push(vm *VM, addrRet *uint64, addr uint64, count uint64, flags uint16, _ uint32) {
	*addrRet = uint64(len(vm.stack.Data))
	dst := vm.stack.Push(count)
	if reg := vm.GetRegion(flags); reg != nil {
		src := reg.Read(addr, count)
		copy(dst, src)
	}
}

var _ Instruction = Pop

// Pop pops stack and writes to memory.Region resolved from flags.
func Pop(vm *VM, addr *uint64, count, _ uint64, flags uint16, _ uint32) {
	src := vm.stack.Pop(count)
	if reg := vm.GetRegion(flags); reg != nil {
		reg.Write(*addr, src)
	}
}
