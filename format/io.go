package format

import (
	"encoding/binary"
	"io"
	"vm/memory"
)

type Reader struct {
	io.Reader
	smallReadBuff, bigReadBuff []byte
}

func NewReader(r io.Reader) *Reader {
	return &Reader{Reader: r, smallReadBuff: make([]byte, 4), bigReadBuff: make([]byte, 256)}
}

func (r *Reader) ReadUint32() (uint32, error) {
	_, err := r.Read(r.smallReadBuff)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint32(r.smallReadBuff), nil
}

func (r *Reader) ReadString() (string, error) {
	l, err := r.ReadUint32()
	if err != nil {
		return "", err
	}
	buff := r.bigReadBuff[:l]
	_, err = r.Read(buff)
	if err != nil {
		return "", err
	}
	return string(buff), nil
}

func (r *Reader) ReadByte() (byte, error) {
	buff := r.smallReadBuff[:1]
	_, err := r.Read(buff)
	return buff[0], err
}

func (r *Reader) ReadRegionInfo() (RegionInfo, error) {
	var (
		info  RegionInfo
		flags byte
		err   error
	)
	if info.Name, err = r.ReadString(); err != nil {
		return info, err
	}
	if flags, err = r.ReadByte(); err != nil {
		return info, err
	}
	info.Flags = memory.Flags(flags)
	if info.Size, err = r.ReadUint32(); err != nil {
		return info, err
	}
	info.Capacity = info.Size
	if info.Flags&memory.FlagStack != 0 {
		info.Size = 0
	}
	return info, nil
}

func (r *Reader) ReadHeader() (Header, error) {
	var (
		h           Header
		regionCount byte
		err         error
	)
	if h.Version, err = r.ReadString(); err != nil {
		return h, err
	}
	if regionCount, err = r.ReadByte(); err != nil {
		return h, err
	}
	for range regionCount {
		var info RegionInfo
		if info, err = r.ReadRegionInfo(); err == nil {
			h.Layout.Regions = append(h.Layout.Regions, info)
		}
		return h, err
	}
	return h, nil
}
