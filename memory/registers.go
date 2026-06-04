package memory

import (
	"fmt"
	"strings"
)

// Registers contains all register for vm (register at index 0 is program counter)
type Registers [16]uint64

func (r Registers) String() string {
	var builder strings.Builder
	for i, val := range r {
		fmt.Fprintf(&builder, "[r%d:%d] ", i, val)
	}
	return builder.String()
}

const ProgramCounterReg = 0
