package format

import "vm/memory"

type RegionInfo struct {
	Name     string
	Flags    memory.Flags
	Size     uint32
	Capacity uint32
}
