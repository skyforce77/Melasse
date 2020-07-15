package core

type Instruction func()uint16

func (cpu *CPU) executeAdd(val byte) byte {
	nVal, overflow := AddByte(byte(cpu.registers.a), val)
	cpu.registers.f.zero = nVal == 0
	cpu.registers.f.subtract = false
	cpu.registers.f.carry = overflow
	cpu.registers.f.halfCarry = (byte(cpu.registers.a & 0xF)) + (val & 0xF) > 0xF
	return nVal
}

func (cpu *CPU) InstructionAdd(source LoadByteSource) Instruction {
	return func()uint16 {
		value := cpu.GetByte(source)
		nVal := cpu.executeAdd(value)
		cpu.registers.a = Register(nVal)
		return cpu.pc+1
	}
}

func (cpu *CPU) executeAddUint16(val uint16) uint16 {
	nVal, overflow := AddUint16(cpu.registers.GetHL(), val)
	cpu.registers.f.zero = nVal == 0
	cpu.registers.f.subtract = false
	cpu.registers.f.carry = overflow
	cpu.registers.f.halfCarry = (byte(cpu.registers.GetHL() & 0xF)) + byte(val & 0xF) > 0xF
	return nVal
}

func (cpu *CPU) InstructionAddHL(source LoadShortSource) Instruction {
	return func()uint16 {
		value := cpu.GetShort(source)
		nVal := cpu.executeAddUint16(uint16(value))
		cpu.registers.SetHL(nVal)
		return cpu.pc+1
	}
}
