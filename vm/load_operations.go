package vm

var _ Instruction = LoadUInt16Immediate

// LoadUInt16Immediate loads uint16 from data to dst.
func LoadUInt16Immediate(_ *VM, dst *uint32, _, _ uint32, data uint16) {
	*dst = uint32(data)
}

var _ Instruction = LoadUint32

// LoadUint32 loads uint32 by address src to dst (flags determine which memory region gonna be used).
func LoadUint32(vm *VM, dst *uint32, src, _ uint32, flags uint16) {
	region := vm.GetRegion(flags)
	if region != nil {
		*dst = region.ReadUint32(src)
	}
}

var _ Instruction = LoadUint16

func LoadUint16(vm *VM, dst *uint32, src, _ uint32, flags uint16) {
	region := vm.GetRegion(flags)
	if region != nil {
		*dst = uint32(region.ReadUint16(src))
	}
}

var _ Instruction = LoadByte

func LoadByte(vm *VM, dst *uint32, src, _ uint32, flags uint16) {
	region := vm.GetRegion(flags)
	if region != nil {
		*dst = uint32(region.ReadByte(src))
	}
}
