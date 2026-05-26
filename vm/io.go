package vm

var _ Instruction = Write

// Write writes to stdOut of vm.
func Write(vm *VM, addr *uint32, count uint32, _ uint32, flags uint16) {
	reg := vm.GetRegion(flags)
	if reg == nil || vm.stdOut == nil {
		return
	}
	_, _ = vm.stdOut.Write(reg.Read(*addr, count))
}

var _ Instruction = Read

// Read reads from stdIn.
func Read(vm *VM, addr *uint32, count uint32, _ uint32, flags uint16) {
	reg := vm.GetRegion(flags)
	if reg == nil || vm.stdIn == nil {
		return
	}
	_, _ = vm.stdIn.Read(vm.stack.Push(count))
	reg.Write(*addr, vm.stack.Pop(count))
}

var _ Instruction = WriteStatus

// WriteStatus writes status.
func WriteStatus(vm *VM, dst *uint32, _ uint32, _ uint32, _ uint16) {
	vm.status = *dst
}

var _ Instruction = LoadStatus

// LoadStatus loads status.
func LoadStatus(vm *VM, dst *uint32, _ uint32, _ uint32, _ uint16) {
	*dst = vm.status
}
