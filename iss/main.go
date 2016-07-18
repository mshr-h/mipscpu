package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type addr uint16
type data uint32

const (
	BASE_MASK = addr(0x7F00)
)

func findInstructionName(i Instruction, table *map[string]InstructionFormat) string {
	for k, v := range *table {
		if v.opcode == i.opcode() {
			if v.function != IGNORE && v.function == i.function() {
				return k
			} else if v.function == IGNORE {
				return k
			}
		}
	}
	return ""
}

func main() {
	var binfile = "fibonacci.o"
	if len(os.Args) >= 2 {
		binfile = os.Args[1]
	}

	bin, err := ioutil.ReadFile(binfile)
	if err != nil {
		panic(err)
	}

	var cpu = NewCpu(bin)
	fmt.Printf("  inst    opcd rd rt rs   imm  address\n")
	for i := 0; i < 0xA; i++ {
		for j := 0; j < 0x10; j += 4 {
			var i = Instruction(cpu.instMemory.ReadWord(addr(0x10*i + j)))
			var mnemonic = findInstructionName(i, cpu.instTable)
			fmt.Printf("%08x %5s %2d %2d %2d %5d %8d\n", i, mnemonic, i.rd(), i.rt(), i.rs(), i.immediate(), i.address())
		}
	}
}
