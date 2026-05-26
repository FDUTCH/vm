package vm

import "vm/memory"

var _ Instruction = Call

// Call calls function from dst addr.
func Call(vm *VM, dst *uint32, _ uint32, _ uint32, _ uint16) {
	currentAddr := vm.registers[memory.ProgramCounterReg]
	vm.stack.PushUint32(currentAddr)
	vm.registers[memory.ProgramCounterReg] = *dst
}

var _ Instruction = Ret

// Ret returns from the function.
func Ret(vm *VM, _ *uint32, _ uint32, _ uint32, _ uint16) {
	vm.registers[memory.ProgramCounterReg] = vm.stack.PopUint32()
}

var _ Instruction = PushImmediate

// PushImmediate pushes uint12 data to the stack (writing address to dst).
func PushImmediate(vm *VM, dst *uint32, _ uint32, _ uint32, data uint16) {
	*dst = uint32(len(vm.stack.Data))
	vm.stack.PushUint16(data)
}

var _ Instruction = PushUint16

// PushUint16 pushes uint16 src to the stack (writing address to dst).
func PushUint16(vm *VM, dst *uint32, src uint32, _ uint32, _ uint16) {
	*dst = uint32(len(vm.stack.Data))
	vm.stack.PushUint16(uint16(src))
}

var _ Instruction = PopUint16

// PopUint16 pops uint16 into dst.
func PopUint16(vm *VM, dst *uint32, _ uint32, _ uint32, _ uint16) {
	*dst = uint32(vm.stack.PopUint16())
}

var _ Instruction = PushUint32

// PushUint32 pushes dst to stack (writing address to dst).
func PushUint32(vm *VM, dst *uint32, src uint32, _ uint32, _ uint16) {
	*dst = uint32(len(vm.stack.Data))
	vm.stack.PushUint32(src)
}

var _ Instruction = PopUint32

// PopUint32 pops uint32 into dst.
func PopUint32(vm *VM, dst *uint32, _ uint32, _ uint32, _ uint16) {
	*dst = vm.stack.PopUint32()
}

var _ Instruction = Push

// Push pushes data from memory.Region resolved from flags and writing address to addrRet.
func Push(vm *VM, addrRet *uint32, addr uint32, count uint32, flags uint16) {
	*addrRet = uint32(len(vm.stack.Data))
	dst := vm.stack.Push(count)
	if reg := vm.GetRegion(flags); reg != nil {
		src := reg.Read(addr, count)
		copy(dst, src)
	}
}

var _ Instruction = Pop

// Pop pops stack and writes to memory.Region resolved from flags.
func Pop(vm *VM, addr *uint32, count uint32, _ uint32, flags uint16) {
	src := vm.stack.Pop(count)
	if reg := vm.GetRegion(flags); reg != nil {
		reg.Write(*addr, src)
	}
}
