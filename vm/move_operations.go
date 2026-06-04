package vm

var _ Instruction = Move

// Move moves data from src to dst.
func Move(_ *VM, dst *uint64, src, _ uint64, _ uint16, _ uint32) {
	*dst = src
}

var _ Instruction = MoveToAddrUint64

// MoveToAddrUint64 moves dst vale to addr.
func MoveToAddrUint64(vm *VM, dst *uint64, addr uint64, _ uint64, flags uint16, _ uint32) {
	if reg := vm.GetRegion(flags); reg != nil {
		reg.WriteUint64(addr, *dst)
	}
}

var _ Instruction = MoveToAddrUint32

func MoveToAddrUint32(vm *VM, dst *uint64, addr uint64, _ uint64, flags uint16, _ uint32) {
	if reg := vm.GetRegion(flags); reg != nil {
		reg.WriteUint32(addr, uint32(*dst))
	}
}

var _ Instruction = MoveToAddrUint16

func MoveToAddrUint16(vm *VM, dst *uint64, addr uint64, _ uint64, flags uint16, _ uint32) {
	if reg := vm.GetRegion(flags); reg != nil {
		reg.WriteUint16(addr, uint16(*dst))
	}
}

var _ Instruction = MoveToAddrByte

func MoveToAddrByte(vm *VM, dst *uint64, addr uint64, _ uint64, flags uint16, _ uint32) {
	if reg := vm.GetRegion(flags); reg != nil {
		reg.WriteByte(addr, byte(*dst))
	}
}
