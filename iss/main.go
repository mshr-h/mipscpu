package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type addr uint16
type data uint32

const (
	BASE_MASK   = addr(0x7F00)
	INST_TYPE_R = iota
	INST_TYPE_I
	INST_TYPE_J
)

type Instruction uint32

func (i *Instruction) opcode() uint8 {
	return 0
}

func (i *Instruction) rs() uint8 {
	return 0
}

func (i *Instruction) rd() uint8 {
	return 0
}

func (i *Instruction) shift() uint8 {
	return 0
}

func (i *Instruction) function() uint8 {
	return 0
}

func (i *Instruction) immediate() uint16 {
	return 0
}

func (i *Instruction) address() uint32 {
	return 0
}

func main() {
	var binfile = "a.out"
	if len(os.Args) >= 2 {
		binfile = os.Args[1]
	}

	bin, err := ioutil.ReadFile(binfile)
	if err != nil {
		panic(err)
	}

	var cpu = NewCpu(bin)
	for i := 0; i < 0x1F; i++ {
		fmt.Printf("%02x\n", cpu.instMemory.ReadByte(addr(i)))
	}
}
