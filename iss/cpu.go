package main

type Instruction uint32

func (i *Instruction) opcode() uint8 {
	return uint8(*i >> 26)
}

func (i *Instruction) rs() uint8 {
	return uint8((*i >> 21) & 0x0000001F)
}

func (i *Instruction) rt() uint8 {
	return uint8((*i >> 16) & 0x0000001F)
}

func (i *Instruction) rd() uint8 {
	return uint8((*i >> 11) & 0x0000001F)
}

func (i *Instruction) shift() uint8 {
	return uint8((*i >> 6) & 0x0000001F)
}

func (i *Instruction) function() uint8 {
	return uint8(*i & 0x0000003F)
}

func (i *Instruction) immediate() uint16 {
	return uint16(*i & 0x0000FFFF)
}

func (i *Instruction) address() uint32 {
	return uint32(*i & 0x03FFFFFF)
}

type Cpu struct {
	programCounter addr
	regfile        *RegFile
	instMemory     *Memory
	dataMemory     *Memory
}

type RegFile [32]data

func NewCpu(bin []byte) *Cpu {
	var cpu = Cpu{}
	var rf RegFile
	var imem, dmem Memory
	loadBinary(&imem, bin)

	cpu.programCounter = 0
	cpu.regfile = &rf
	cpu.instMemory = &imem
	cpu.dataMemory = &dmem

	return &cpu
}

func (c *Cpu) Execute() {
}

func (c *Cpu) Run() {
}

func loadBinary(m *Memory, b []byte) {
	for i, d := range b {
		m.WriteByte(addr(i), d)
	}
}
