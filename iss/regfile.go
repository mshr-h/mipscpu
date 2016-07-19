package main

type RegFile [32]uint32

func (r *RegFile) Read(address uint8) uint32 {
	return r[address]
}

func (r *RegFile) Write(address uint8, data uint32) {
	if address != 0 {
		r[address] = data
	}
}
