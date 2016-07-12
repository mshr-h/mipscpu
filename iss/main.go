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

func main() {
	var binfile = "sample.o"
	if len(os.Args) >= 2 {
		binfile = os.Args[1]
	}

	bin, err := ioutil.ReadFile(binfile)
	if err != nil {
		panic(err)
	}

	var cpu = NewCpu(bin)
	for i := 0; i < 0xA; i++ {
		for j := 0; j < 0x10; j += 4 {
			var i = Instruction(cpu.instMemory.ReadWord(addr(0x10*i + j)))
			fmt.Printf("%08x %02x %2d %2d %2d\n", i, i.opcode(), i.rd(), i.rt(), i.rd())
		}
	}
}
