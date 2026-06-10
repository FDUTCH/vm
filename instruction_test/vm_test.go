package instruction

import (
	_ "embed"
	"os"
	"testing"
	"vm/asm"
	"vm/asm/linker"
	"vm/memory"
	"vm/vm"
)

//go:embed hello_world.some
var data string

func TestRunVM(t *testing.T) {

	reg := vm.DefaultRegistry()

	tokens, err := asm.NewAssembler(reg).Lex(data)
	if err != nil {
		t.Fatal(err)
	}

	l := linker.New(linker.DefaultTypeRegistry())
	instructions, static, dynamic, err := l.Link(tokens)
	if err != nil {
		t.Fatal(err)
	}

	staticRegion := memory.Region{
		Name: "static",
		Data: static,
	}

	dynamicRegion := memory.Region{
		Name:  "dynamic",
		Flags: memory.FlagWrite,
		Data:  dynamic,
	}

	stack := memory.Region{
		Name:  "stack",
		Flags: memory.FlagStack,
		Data:  make([]byte, 0, 1024),
	}

	program := make([]uint64, 0, len(instructions))
	for _, i := range instructions {
		program = append(program, asm.EncodeInstructionToken(i))
	}

	machine, err := vm.NewVM(program, []memory.Region{staticRegion, dynamicRegion, stack}, vm.OptionStdOut(os.Stdout))
	if err != nil {
		panic(err)
	}
	status := machine.Status()

	for *status == 0 {
		machine.Tick()
	}
}
