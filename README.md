# MIPS CPU implementation and its Instruction Set Simulator

## CPU architecture

- havard architecture
- byte access memory
- address space: 0x0000 - 0x7FFF
- 32 integer register
- $0 is always 0
- 3 types of instructions(R, I, J)

## Instruction format

```text
Type 31                                                                  0
      +------------+--------+--------+--------+-----------+--------------+
 R    | opcode (6) | rs (5) | rt (5) | rd (5) | shift (5) | function (6) |
      +------------+--------+--------+--------+-----------+--------------+
 I    | opcode (6) | rs (5) | rt (5) |           immediate (16)          |
      +------------+--------+--------+-----------------------------------+
 J    | opcode (6) |                     address (26)                    |
      +------------+-----------------------------------------------------+
```

