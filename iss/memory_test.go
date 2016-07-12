package main

import "testing"

func TestMemory(t *testing.T) {
	var m Memory
	var expect byte

	expect = byte(103)
	m.WriteByte(0, expect)
	if res := m.ReadByte(0); res != expect {
		t.Fatalf("should be %v but %v", expect, res)
	}

	expect = byte(200)
	m.WriteByte(500, expect)
	if res := m.ReadByte(500); res != expect {
		t.Fatalf("should be %v but %v", expect, res)
	}

	var word uint32

	word = uint32(0xFF00FF00)
	m.WriteWord(0, word)
	if res := m.ReadWord(0); res != word {
		t.Fatalf("should be %v but %v", word, res)
	}

	word = uint32(0x1234ABCD)
	m.WriteWord(1024, word)
	if res := m.ReadWord(1024); res != word {
		t.Fatalf("should be %v but %v", word, res)
	}
}

func TestMemoryBlock(t *testing.T) {
	var mb = NewMemoryBlock(0)
	var expect byte

	expect = byte(103)
	mb.WriteByte(0, expect)
	if res := mb.ReadByte(0); res != expect {
		t.Fatalf("should be %v but %v:", expect, res)
	}

	expect = 11
	mb.WriteByte(255, expect)
	if res := mb.ReadByte(255); res != expect {
		t.Fatalf("should be %v but %v:", expect, res)
	}
}
