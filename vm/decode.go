package vm

func Decode(word uint64) (byte, byte, byte, byte, uint16, uint32) {
	op := byte(word >> 56)                // 8 bits
	dst := byte((word >> 52) & 0x0F)      // 4 bits
	src1 := byte((word >> 48) & 0x0F)     // 4 bits
	src2 := byte((word >> 44) & 0x0F)     // 4 bits
	flags := uint16((word >> 32) & 0xFFF) // 12 bits
	data := uint32(word)
	return op, dst, src1, src2, flags, data
}
