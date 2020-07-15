package core

type LoadByteTarget uint8
const (
	TRegisterA LoadByteTarget = iota
	TRegisterB
	TRegisterC
	TRegisterD
	TRegisterE
	TRegisterH
	TRegisterL
	TPointerHL
)

type LoadShortTarget uint8
const (
	TRegisterBC LoadShortTarget = iota
	TRegisterDE
	TRegisterHL
	TRegisterSP
)

type LoadByteSource uint8
const (
	SRegisterA LoadByteSource = iota
	SRegisterB
	SRegisterC
	SRegisterD
	SRegisterE
	SRegisterH
	SRegisterL
	SPointerHL
	SDirect8
)

type LoadShortSource uint8
const (
	SRegisterBC LoadShortSource = iota
	SRegisterDE
	SRegisterHL
	SRegisterSP
	SDirect16
)

type Register byte


type FlagsRegister struct {
	zero bool
	subtract bool
	halfCarry bool
	carry bool
}

func (fr *FlagsRegister) GetByte() uint8 {
	var result uint8 = 0
	if fr.zero {
		result |= 1 << 7
	}
	if fr.subtract {
		result |= 1 << 6
	}
	if fr.halfCarry {
		result |= 1 << 5
	}
	if fr.carry {
		result |= 1 << 4
	}
	return result
}

func (fr *FlagsRegister) SetByte(val uint8) {
	fr.zero = ((val >> 7) & 0B1) != 0
	fr.subtract = ((val >> 6) & 0B1) != 0
	fr.halfCarry = ((val >> 5) & 0B1) != 0
	fr.carry = ((val >> 4) & 0B1) != 0
}


type Registers struct {
	a Register
	b Register
	c Register
	d Register
	e Register
	f FlagsRegister
	h Register
	l Register
}

func (cpu *CPU) GetByte(source LoadByteSource) byte {
	switch source {
	case SRegisterA:
		return byte(cpu.registers.a)
	case SRegisterB:
		return byte(cpu.registers.b)
	case SRegisterC:
		return byte(cpu.registers.c)
	case SRegisterD:
		return byte(cpu.registers.d)
	case SRegisterE:
		return byte(cpu.registers.e)
	case SRegisterH:
		return byte(cpu.registers.h)
	case SRegisterL:
		return byte(cpu.registers.l)
	case SPointerHL:
		return cpu.memory.ReadByte(cpu.registers.GetHL())
	case SDirect8:
		return cpu.memory.ReadByte(cpu.pc+1)
	}
	return 0
}

func (cpu *CPU) SetByte(target LoadByteTarget, val byte) {
	switch target {
	case TRegisterA:
		cpu.registers.a = Register(val)
	case TRegisterB:
		cpu.registers.b = Register(val)
	case TRegisterC:
		cpu.registers.c = Register(val)
	case TRegisterD:
		cpu.registers.d = Register(val)
	case TRegisterE:
		cpu.registers.e = Register(val)
	case TRegisterH:
		cpu.registers.h = Register(val)
	case TRegisterL:
		cpu.registers.l = Register(val)
	case TPointerHL:
		cpu.memory.SetByte(cpu.registers.GetHL(), val)
	}
}

func (cpu *CPU) GetShort(source LoadShortSource) uint16 {
	switch source {
	case SRegisterBC:
		return cpu.registers.GetBC()
	case SRegisterDE:
		return cpu.registers.GetDE()
	case SRegisterHL:
		return cpu.registers.GetHL()
	case SRegisterSP:
		return cpu.sp
	case SDirect16:
		return cpu.memory.ReadShort(cpu.pc+1)
	}
	return 0
}

func (cpu *CPU) SetShort(target LoadShortTarget, val uint16) {
	switch target {
	case TRegisterBC:
		cpu.registers.SetBC(val)
	case TRegisterDE:
		cpu.registers.SetDE(val)
	case TRegisterHL:
		cpu.registers.SetHL(val)
	case TRegisterSP:
		cpu.sp = val
	}
}

func (r *Registers) GetAF() uint16 {
	return uint16(r.a) << 8 | uint16(r.f.GetByte())
}

func (r *Registers) SetAF(val uint16) {
	r.a = Register((val & 0xFF00) >> 8)
	r.f.SetByte(uint8(val & 0xFF))
}

func (r *Registers) GetBC() uint16 {
	return uint16(r.b) << 8 | uint16(r.c)
}

func (r *Registers) SetBC(val uint16) {
	r.b = Register((val & 0xFF00) >> 8)
	r.c = Register(val & 0xFF)
}

func (r *Registers) GetDE() uint16 {
	return uint16(r.d) << 8 | uint16(r.e)
}

func (r *Registers) SetDE(val uint16) {
	r.d = Register((val & 0xFF00) >> 8)
	r.e = Register(val & 0xFF)
}

func (r *Registers) GetHL() uint16 {
	return uint16(r.h) << 8 | uint16(r.l)
}

func (r *Registers) SetHL(val uint16) {
	r.h = Register((val & 0xFF00) >> 8)
	r.l = Register(val & 0xFF)
}