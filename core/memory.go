package core

type Memory struct {
	data [0xFFFF]byte
}

func (mem *Memory) ReadByte(addr uint16) byte {
	return mem.data[addr]
}

func (mem *Memory) SetByte(addr uint16, val byte) {
	mem.data[addr] = val
}

func (mem *Memory) ReadShort(addr uint16) uint16 {
	return uint16(mem.data[addr]) << 8 | uint16(mem.data[addr+1])
}

func (mem *Memory) SetShort(addr uint16, val uint16) {
	mem.data[addr] = byte((val & 0xFF00) >> 8)
	mem.data[addr+1] = byte(val & 0xFF)
}