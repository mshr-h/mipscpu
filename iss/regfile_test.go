package main

import "testing"

func TestRegFile(t *testing.T) {
	var r RegFile
	var expect uint32

	r.Write(0, 100)
	if res := r.Read(0); res != 0 {
		t.Fatalf("should be %v but %v", 0, res)
	}

	expect = 200
	r.Write(1, expect)
	if res := r.Read(1); res != expect {
		t.Fatalf("should be %v but %v", expect, res)
	}

	expect = 500
	r.Write(31, expect)
	if res := r.Read(31); res != expect {
		t.Fatalf("should be %v but %v", expect, res)
	}

}
