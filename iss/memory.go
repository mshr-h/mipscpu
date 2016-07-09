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

func (m *Memory) WriteByte(address addr, data byte) {
	if mb := findMemoryBlock(m, address); mb != nil {
		mb.WriteByte(address-mb.baseaddress, data)
		return
	}

	new := NewMemoryBlock(address & BASE_MASK)
	new.WriteByte(address-new.baseaddress, data)
	*m = append(*m, new)
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
