package main

type Memory []*MemoryBlock

type MemoryBlock struct {
	baseaddress addr
	data        [^BASE_MASK + 1]byte
}

func findMemoryBlock(m *Memory, address addr) *MemoryBlock {
	for _, mb := range *m {
		if mb.baseaddress <= address && address < mb.baseaddress+addr(len(mb.data)) {
			return mb
		}
	}
	return nil
}

func (m *Memory) ReadByte(address addr) byte {
	if mb := findMemoryBlock(m, address); mb != nil {
		return mb.ReadByte(address - mb.baseaddress)
	}
	return 0xA5
}

func (m *Memory) ReadWord(address addr) uint32 {
	if address%4 != 0 {
		panic("unaligned address")
	}

	d1 := uint32(m.ReadByte(address))
	d2 := uint32(m.ReadByte(address + 1))
	d3 := uint32(m.ReadByte(address + 2))
	d4 := uint32(m.ReadByte(address + 3))

	return (d1 << 24) | (d2 << 16) | (d3 << 8) | d4
}

func (m *Memory) WriteByte(address addr, data byte) {
	if mb := findMemoryBlock(m, address); mb != nil {
		mb.WriteByte(address-mb.baseaddress, data)
		return
	}

	new := NewMemoryBlock(address & BASE_MASK)
	new.WriteByte(address-new.baseaddress, data)
	*m = append(*m, new)
}

func (m *Memory) WriteWord(address addr, data uint32) {
	if address%4 != 0 {
		panic("unaligned address")
	}

	m.WriteByte(address, uint8(data>>24))
	m.WriteByte(address+1, uint8((data>>16)&0xFF))
	m.WriteByte(address+2, uint8((data>>8)&0xFF))
	m.WriteByte(address+3, uint8(data&0xFF))
}

func NewMemoryBlock(baseaddress addr) *MemoryBlock {
	var mb = MemoryBlock{baseaddress: baseaddress}
	return &mb
}

func (mb *MemoryBlock) ReadByte(address addr) byte {
	return mb.data[address]
}

func (mb *MemoryBlock) WriteByte(address addr, data byte) {
	mb.data[address] = data
}
