package vm

import (
	"reflect"
	"runtime"
	"strings"
)

type Registry struct {
	ids          map[uintptr]byte
	byName       map[string]InstructionData
	names        [256]string
	instructions [256]Instruction
}

func NewRegistry(data ...any) *Registry {
	var (
		r = &Registry{ids: make(map[uintptr]byte, 256), byName: make(map[string]InstructionData)}

		parseArgsCount bool
		instruction    Instruction
	)

	for _, val := range data {
		if !parseArgsCount {
			instruction = val.(Instruction)
		} else {
			r.Register(instruction, val.(InstructionData))
		}
		parseArgsCount = !parseArgsCount
	}
	return r
}

func DefaultRegistry() *Registry {
	return NewRegistry(
		Inc, InstructionData{ParamCount: 1},
		Dec, InstructionData{ParamCount: 1},
		LoadImmediate, InstructionData{ParamCount: 1, HasData: true},
		LoadUint64, InstructionData{ParamCount: 2, HasFlags: true},
		LoadUint32, InstructionData{ParamCount: 2, HasFlags: true},
		LoadUint16, InstructionData{ParamCount: 2, HasFlags: true},
		LoadByte, InstructionData{ParamCount: 2, HasFlags: true},
		Move, InstructionData{ParamCount: 2},
		MoveToAddrUint64, InstructionData{ParamCount: 2, HasFlags: true},
		MoveToAddrUint32, InstructionData{ParamCount: 2, HasFlags: true},
		MoveToAddrUint16, InstructionData{ParamCount: 2, HasFlags: true},
		MoveToAddrByte, InstructionData{ParamCount: 2, HasFlags: true},
		AddImmediate, InstructionData{ParamCount: 1, HasData: true},
		SubImmediate, InstructionData{ParamCount: 1, HasData: true},
		Add, InstructionData{ParamCount: 3, HasFlags: true},
		Sub, InstructionData{ParamCount: 3, HasFlags: true},
		Mul, InstructionData{ParamCount: 3, HasFlags: true},
		Div, InstructionData{ParamCount: 3, HasFlags: true},
		Jump, InstructionData{ParamCount: 1},
		JumpImmediate, InstructionData{ParamCount: 0, HasData: true},
		JumpRelImmediate, InstructionData{ParamCount: 0, HasFlags: true, HasData: true},
		JumpRel, InstructionData{ParamCount: 1, HasFlags: true},
		JumpConditional, InstructionData{ParamCount: 2, HasFlags: true},
		JumpRelConditional, InstructionData{ParamCount: 2, HasFlags: true},
		JumpConditionalImmediate, InstructionData{ParamCount: 1, HasFlags: true, HasData: true},
		JumpRelConditionalImmediate, InstructionData{ParamCount: 1, HasFlags: true, HasData: true},
		Equal, InstructionData{ParamCount: 3},
		Greater, InstructionData{ParamCount: 3, HasFlags: true},
		Call, InstructionData{ParamCount: 0, HasData: true},
		CallAddr, InstructionData{ParamCount: 1},
		Ret, InstructionData{ParamCount: 0},
		PushImmediate, InstructionData{ParamCount: 1, HasData: true},
		PushByte, InstructionData{ParamCount: 2},
		PopByte, InstructionData{ParamCount: 1},
		PushUint16, InstructionData{ParamCount: 2},
		PopUint16, InstructionData{ParamCount: 1},
		PushUint32, InstructionData{ParamCount: 2},
		PopUint32, InstructionData{ParamCount: 1},
		PushUint64, InstructionData{ParamCount: 2},
		PopUint64, InstructionData{ParamCount: 1},
		Push, InstructionData{ParamCount: 3, HasFlags: true},
		Pop, InstructionData{ParamCount: 2, HasFlags: true},
		Write, InstructionData{ParamCount: 2, HasFlags: true},
		Read, InstructionData{ParamCount: 2, HasFlags: true},
		WriteStatus, InstructionData{ParamCount: 1},
		LoadStatus, InstructionData{ParamCount: 1},
	)
}

func (r *Registry) Register(instruction Instruction, data InstructionData) {
	l := len(r.ids)
	id := byte(l)
	ptr := reflect.ValueOf(instruction).Pointer()
	name := strings.Split(runtime.FuncForPC(ptr).Name(), ".")[1]

	data.id = id
	r.byName[name] = data
	r.ids[ptr] = id
	r.instructions[l] = instruction
	r.names[l] = name
	if len(r.instructions) == l || len(r.byName) == l {
		panic("registering same instruction twice")
	}
}

func (r *Registry) GetID(instruction Instruction) byte {
	return r.ids[reflect.ValueOf(instruction).Pointer()]
}

func (r *Registry) FromName(name string) (InstructionData, bool) {
	data, ok := r.byName[name]
	return data, ok
}

func (r *Registry) GetName(id byte) string {
	return r.names[id]
}

func (r *Registry) Instructions() [256]Instruction {
	return r.instructions
}
