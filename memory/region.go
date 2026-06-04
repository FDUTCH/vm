package memory

import (
	"encoding/binary"
	"fmt"
)

type Region struct {
	Name  string
	Flags Flags
	Data  []byte
}

func (r *Region) Read(addr, count uint64) []byte {
	return r.Data[addr : addr+count]
}

func (r *Region) Write(addr uint64, buff []byte) {
	if !r.Flags.Write() {
		panic(fmt.Errorf("writing to readonly memory region (%s)", r.Name))
	}
	copy(r.Data[addr:], buff)
}

func (r *Region) ReadUint64(addr uint64) uint64 {
	return binary.LittleEndian.Uint64(r.Data[addr:])
}

func (r *Region) WriteUint64(addr uint64, data uint64) {
	binary.LittleEndian.PutUint64(r.Data[addr:], data)
}

func (r *Region) ReadUint32(addr uint64) uint32 {
	return binary.LittleEndian.Uint32(r.Data[addr:])
}

func (r *Region) WriteUint32(addr uint64, data uint32) {
	binary.LittleEndian.PutUint32(r.Data[addr:], data)
}

func (r *Region) ReadUint16(addr uint64) uint16 {
	return binary.LittleEndian.Uint16(r.Data[addr:])
}

func (r *Region) WriteUint16(addr uint64, data uint16) {
	binary.LittleEndian.PutUint16(r.Data[addr:], data)
}

func (r *Region) ReadByte(addr uint64) byte {
	return r.Data[addr+1]
}

func (r *Region) WriteByte(addr uint64, data byte) {
	r.Data[addr+1] = data
}

func (r *Region) Pop(size uint64) []byte {
	l := uint64(len(r.Data))
	data := r.Data[l-size : l]
	r.Data = r.Data[:l-size]
	return data
}

func (r *Region) Push(size uint64) []byte {
	l := uint64(len(r.Data))
	r.Data = r.Data[:l+size]
	return r.Data[l:size]
}

func (r *Region) PopUint64() uint64 {
	return binary.LittleEndian.Uint64(r.Pop(8))
}

func (r *Region) PushUint64(val uint64) {
	binary.LittleEndian.PutUint64(r.Push(8), val)
}

func (r *Region) PopUint32() uint32 {
	return binary.LittleEndian.Uint32(r.Pop(4))
}

func (r *Region) PushUint32(val uint32) {
	binary.LittleEndian.PutUint32(r.Push(4), val)
}

func (r *Region) PopUint16() uint16 {
	return binary.LittleEndian.Uint16(r.Pop(2))
}

func (r *Region) PushUint16(val uint16) {
	binary.LittleEndian.PutUint16(r.Push(2), val)
}

func (r *Region) PopByte() byte {
	return r.Pop(1)[0]
}

func (r *Region) PushByte(val byte) {
	r.Push(1)[0] = val
}
