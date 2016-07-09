package main

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
