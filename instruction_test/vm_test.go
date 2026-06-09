package instruction

import (
	"os"
	"testing"
	"vm/asm"
	"vm/asm/linker"
	"vm/asm/tokenizer"
	"vm/memory"
	"vm/vm"
)

func TestRunVM(t *testing.T) {
	data := `
		LoadImmediate r2 hello
Cycle:
		LoadByte r4 r2 $S // loading char from static region
		Inc r2 // incrementing pointer
		Equal r5 r4 r1 // comparing chars
		JumpConditionalImmediate r5 Cycle // jumping if r5 != 0
		Write r6 r2 $S // writing to stdout 
		LoadImmediate r1 1 // загружаем 1 в r1
		WriteStatus r1 // пишем r1 как статус
	`

	reg := vm.DefaultRegistry()

	tokens, err := asm.NewAssembler(reg).Lex(data)
	if err != nil {
		t.Fatal(err)
	}

	l := linker.New(linker.DefaultTypeRegistry())
	instructions, static, dynamic, err := l.Link(append(tokens, tokenizer.NewVariableDecl(
		true,
		"hello",
		"str",
		"\"Hello World!!!\n\u0000\"",
	)))
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
