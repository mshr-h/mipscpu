package main

const (
	IGNORE = uint8(0xFF)
)

type Instruction uint32

type InstructionFormat struct {
	opcode   uint8
	function uint8
}

func initInstructionTable() map[string]InstructionFormat {
	table := make(map[string]InstructionFormat)

	table["ADD"] = InstructionFormat{opcode: 0x00, function: 0x20}
	table["ADDU"] = InstructionFormat{opcode: 0x00, function: 0x21}
	table["SUB"] = InstructionFormat{opcode: 0x00, function: 0x22}
	table["SUBU"] = InstructionFormat{opcode: 0x00, function: 0x23}
	table["ADDI"] = InstructionFormat{opcode: 0x08, function: IGNORE}
	table["ADDIU"] = InstructionFormat{opcode: 0x09, function: IGNORE}
	table["MULT"] = InstructionFormat{opcode: 0x00, function: 0x18}
	table["MULTU"] = InstructionFormat{opcode: 0x00, function: 0x19}
	table["DIV"] = InstructionFormat{opcode: 0x00, function: 0x1A}
	table["DIVU"] = InstructionFormat{opcode: 0x00, function: 0x1B}
	table["LW"] = InstructionFormat{opcode: 0x23, function: IGNORE}
	table["LH"] = InstructionFormat{opcode: 0x21, function: IGNORE}
	table["LHU"] = InstructionFormat{opcode: 0x25, function: IGNORE}
	table["LB"] = InstructionFormat{opcode: 0x20, function: IGNORE}
	table["LBU"] = InstructionFormat{opcode: 0x24, function: IGNORE}
	table["SW"] = InstructionFormat{opcode: 0x2B, function: IGNORE}
	table["SH"] = InstructionFormat{opcode: 0x29, function: IGNORE}
	table["SB"] = InstructionFormat{opcode: 0x28, function: IGNORE}
	table["LUI"] = InstructionFormat{opcode: 0x0F, function: IGNORE}
	table["MFHI"] = InstructionFormat{opcode: 0x00, function: 0x10}
	table["MFLO"] = InstructionFormat{opcode: 0x00, function: 0x12}
	table["MFCZ"] = InstructionFormat{opcode: 0x00, function: IGNORE}
	table["MTCZ"] = InstructionFormat{opcode: 0x00, function: IGNORE}
	table["AND"] = InstructionFormat{opcode: 0x00, function: 0x24}
	table["ANDI"] = InstructionFormat{opcode: 0x0C, function: IGNORE}
	table["OR"] = InstructionFormat{opcode: 0x00, function: 0x25}
	table["ORI"] = InstructionFormat{opcode: 0x0D, function: IGNORE}
	table["XOR"] = InstructionFormat{opcode: 0x00, function: 0x26}
	table["NOR"] = InstructionFormat{opcode: 0x00, function: 0x27}
	table["SLT"] = InstructionFormat{opcode: 0x00, function: 0x2A}
	table["SLTU"] = InstructionFormat{opcode: 0x00, function: 0x2B}
	table["SLTI"] = InstructionFormat{opcode: 0x0A, function: IGNORE}
	table["SLL"] = InstructionFormat{opcode: 0x00, function: 0x00}
	table["SRL"] = InstructionFormat{opcode: 0x00, function: 0x02}
	table["SRA"] = InstructionFormat{opcode: 0x00, function: 0x03}
	table["SLLV"] = InstructionFormat{opcode: 0x00, function: 0x04}
	table["SRLV"] = InstructionFormat{opcode: 0x00, function: 0x06}
	table["SRAV"] = InstructionFormat{opcode: 0x00, function: 0x07}
	table["BEQ"] = InstructionFormat{opcode: 0x04, function: IGNORE}
	table["BNE"] = InstructionFormat{opcode: 0x05, function: IGNORE}
	table["J"] = InstructionFormat{opcode: 0x02, function: IGNORE}
	table["JR"] = InstructionFormat{opcode: 0x00, function: 0x08}
	table["JAL"] = InstructionFormat{opcode: 0x03, function: IGNORE}

	return table
}

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
	instTable      *map[string]InstructionFormat
}

func NewCpu(bin []byte) *Cpu {
	var cpu = Cpu{}
	var rf RegFile
	var imem, dmem Memory
	var instTable = initInstructionTable()
	loadBinary(&imem, bin)

	cpu.programCounter = 0
	cpu.regfile = &rf
	cpu.instMemory = &imem
	cpu.dataMemory = &dmem
	cpu.instTable = &instTable

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
