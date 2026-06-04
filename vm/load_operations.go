package vm

var _ Instruction = LoadImmediate

// LoadImmediate loads uint16 from data to dst.
func LoadImmediate(_ *VM, dst *uint64, _, _ uint64, _ uint16, data uint32) {
	*dst = uint64(data)
}

var _ Instruction = LoadUint64

// LoadUint64 loads uint64 by address src to dst (flags determine which memory region gonna be used).
func LoadUint64(vm *VM, dst *uint64, src, _ uint64, flags uint16, _ uint32) {
	region := vm.GetRegion(flags)
	if region != nil {
		*dst = region.ReadUint64(src)
	}
}

var _ Instruction = LoadUint32

// LoadUint32 loads uint32 by address src to dst (flags determine which memory region gonna be used).
func LoadUint32(vm *VM, dst *uint64, src, _ uint64, flags uint16, _ uint32) {
	region := vm.GetRegion(flags)
	if region != nil {
		*dst = uint64(region.ReadUint32(src))
	}
}

var _ Instruction = LoadUint16

func LoadUint16(vm *VM, dst *uint64, src, _ uint64, flags uint16, _ uint32) {
	region := vm.GetRegion(flags)
	if region != nil {
		*dst = uint64(region.ReadUint16(src))
	}
}

var _ Instruction = LoadByte

func LoadByte(vm *VM, dst *uint64, src, _ uint64, flags uint16, _ uint32) {
	region := vm.GetRegion(flags)
	if region != nil {
		*dst = uint64(region.ReadByte(src))
	}
}
