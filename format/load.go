package format

import (
	"os"
	"vm/memory"
)

func Load(path string) (regions []memory.Region, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := NewReader(file)
	h, err := r.ReadHeader()
	if err != nil {
		return nil, err
	}

	for _, rg := range h.Layout.Regions {
		buff := make([]byte, rg.Size, rg.Capacity)
		_, err = r.Read(buff)
		if err != nil {
			return nil, err
		}
		regions = append(regions, memory.Region{
			Name:  rg.Name,
			Flags: rg.Flags,
			Data:  buff,
		})
	}
	return regions, err
}
