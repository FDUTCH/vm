package vm

import (
	"io"
	"os"
	"vm/memory"
)

type VM struct {
	registers memory.Registers

	program []uint64

	status uint64

	stack   memory.Region
	static  memory.Region
	dynamic memory.Region

	registry *Registry
	opcodes  [256]Instruction

	stdIn  io.Reader
	stdOut io.Writer
}

func NewVM(program []uint64, regions []memory.Region, options ...Option) (*VM, error) {
	var (
		static  memory.Region
		dynamic memory.Region
		stack   memory.Region
	)

	for _, rg := range regions {
		switch {
		case rg.Flags.Write():
			dynamic = rg
		case rg.Flags.Stack():
			stack = rg
		default:
			static = rg
		}
	}

	vm := &VM{
		program: program,
		stack:   stack,
		static:  static,
		dynamic: dynamic,
		stdIn:   os.Stdin,
		stdOut:  os.Stdout,
	}

	for _, op := range options {
		op(vm)
	}

	// setting instruction set if it is nil.
	if vm.registry == nil {
		OptionInstructionSet(DefaultRegistry())(vm)
	}
	return vm, nil
}

func (vm *VM) Tick() {
	// fetching instruction.
	word := vm.program[vm.registers[memory.ProgramCounterReg]]

	// decoding instruction.
	op, d, s1, s2, flags, data := Decode(word)

	// setting source & destination registers.
	dst := &vm.registers[d]
	src1 := vm.registers[s1]
	src2 := vm.registers[s2]

	// executing instruction.
	vm.opcodes[op](vm, dst, src1, src2, flags, data)

	// incrementing program counter.
	vm.registers[memory.ProgramCounterReg]++
}

func (vm *VM) Status() *uint64 {
	return &vm.status
}

func (vm *VM) GetRegion(flags uint16) *memory.Region {
	switch {
	case flags&FlagNoneRegion != 0:
		return nil
	case flags&FlagStaticRegion != 0:
		return &vm.static
	case flags&FlagStackRegion != 0:
		return &vm.stack
	}
	return &vm.dynamic
}
