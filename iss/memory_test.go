package main

import "testing"

func TestMemory(t *testing.T) {
	var m Memory
	var expect byte

	expect = byte(103)
	m.WriteByte(0, expect)
	var res = m.ReadByte(0)
	if res != expect {
		t.Fatalf("should be %v but %v", expect, res)
	}

	expect = byte(200)
	m.WriteByte(500, expect)
	res = m.ReadByte(500)
	if res != expect {
		t.Fatalf("should be %v but %v", expect, res)
	}
}

func TestMemoryBlock(t *testing.T) {
	var mb = NewMemoryBlock(0)
	var expect byte

	expect = byte(103)
	mb.WriteByte(0, expect)
	var d = mb.ReadByte(0)
	if d != expect {
		t.Fatalf("should be %v but %v:", expect, d)
	}

	expect = 11
	mb.WriteByte(255, expect)
	d = mb.ReadByte(255)
	if d != expect {
		t.Fatalf("should be %v but %v:", expect, d)
	}
}
