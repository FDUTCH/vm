package asm

func Encode(op byte, dst byte, src1 byte, src2 byte, data uint16) uint32 {
	word := uint32(0)
	word |= uint32(op) << 24
	word |= uint32(dst&0xf) << 20
	word |= uint32(src1&0xf) << 16
	word |= uint32(src2&0xf) << 12
	word |= uint32(data & 0xFFF)
	return word
}
