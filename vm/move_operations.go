package vm

var _ Instruction = Move

// Move moves data from src to dst.
func Move(_ *VM, dst *uint32, src, _ uint32, _ uint16) {
	*dst = src
}

var _ Instruction = MoveToAddrUint32

// MoveToAddrUint32 moves dst vale to addr.
func MoveToAddrUint32(vm *VM, dst *uint32, addr uint32, _ uint32, flags uint16) {
	if reg := vm.GetRegion(flags); reg != nil {
		reg.WriteUint32(addr, *dst)
	}
}

var _ Instruction = MoveToAddrUint16

// MoveToAddrUint16 moves dst vale to addr.
func MoveToAddrUint16(vm *VM, dst *uint32, addr uint32, _ uint32, flags uint16) {
	if reg := vm.GetRegion(flags); reg != nil {
		reg.WriteUint16(addr, uint16(*dst))
	}
}

var _ Instruction = MoveToAddrByte

// MoveToAddrByte moves dst vale to addr.
func MoveToAddrByte(vm *VM, dst *uint32, addr uint32, _ uint32, flags uint16) {
	if reg := vm.GetRegion(flags); reg != nil {
		reg.WriteByte(addr, byte(*dst))
	}
}
