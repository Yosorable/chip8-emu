package core

import (
	"log"
	"math/rand"
)

type cpu struct {
	V  [16]uint8 // common registers
	I  uint16    // index register
	PC uint16    // program counter

	emu *Chip8
}

func (c *cpu) Initialize() {
	for i := 0; i < 16; i++ {
		c.V[i] = 0
	}
	c.I = 0
	c.PC = 0
}

func logCode(opcode uint16, tp string) {
	debug := true
	if debug {
		log.Printf("[opcode] code: 0x%04X, type: %s", opcode, tp)
	}
}

func (c *cpu) ExecuteCode() {
	var opcode = uint16(c.emu.memory.storage[c.emu.cpu.PC])<<8 + uint16(c.emu.memory.storage[c.emu.cpu.PC+1])

	// next instruction
	c.PC += 2

	// 0NNN
	// do nothing
	if opcode>>12 == 0 && opcode != 0x00E0 && opcode != 0xEE {
		logCode(opcode, "0NNN")
		return
	}
	// 00E0
	if opcode == 0xE0 {
		logCode(opcode, "00E0")
		c.emu.screen.clear()
		return
	}
	// 00EE
	if opcode == 0xEE {
		logCode(opcode, "00EE")
		c.PC = c.emu.stack[c.emu.sp-1]
		c.emu.sp--
		return
	}
	// 1NNN
	if opcode>>12 == 0x1 {
		logCode(opcode, "1NNN")
		nnn := opcode & 0xFFF
		c.PC = nnn
		return
	}
	// 2NNN
	if opcode>>12 == 0x2 {
		logCode(opcode, "2NNN")
		nnn := opcode & 0xFFF
		c.emu.stack[c.emu.sp] = c.PC
		c.emu.sp++
		c.PC = nnn
		return
	}
	// 3XNN
	if opcode>>12 == 0x3 {
		logCode(opcode, "3XNN")
		x := opcode >> 8 & 0xF
		nn := opcode & 0xFF
		if c.V[x] == uint8(nn) {
			c.PC += 2
		}
		return
	}
	// 4XNN
	if opcode>>12 == 0x4 {
		logCode(opcode, "4XNN")
		x := opcode >> 8 & 0xF
		nn := opcode & 0xFF
		if c.V[x] != uint8(nn) {
			c.PC += 2
		}
		return
	}
	// 5XY0
	if opcode>>12 == 0x5 && opcode&0xF == 0x0 {
		logCode(opcode, "5XY0")
		x := opcode >> 8 & 0xF
		y := opcode >> 4 & 0xF
		if c.V[x] == c.V[y] {
			c.PC += 2
		}
		return
	}
	// 6XNN
	if opcode>>12 == 0x6 {
		logCode(opcode, "6XNN")
		x := opcode >> 8 & 0xF
		nn := opcode & 0xFF
		c.V[x] = uint8(nn)
		return
	}
	// 7XNN
	if opcode>>12 == 0x7 {
		logCode(opcode, "7XNN")
		x := opcode >> 8 & 0xF
		nn := opcode & 0xFF
		c.V[x] += uint8(nn)
		return
	}
	// 8XY0
	if opcode>>12 == 0x8 && opcode&0xF == 0x0 {
		logCode(opcode, "8XY0")
		x := opcode >> 8 & 0xF
		y := opcode >> 4 & 0xF
		c.V[x] = c.V[y]
		return
	}
	// 8XY1
	if opcode>>12 == 0x8 && opcode&0xF == 0x1 {
		logCode(opcode, "8XY1")
		x := opcode >> 8 & 0xF
		y := opcode >> 4 & 0xF
		c.V[x] |= c.V[y]
		return
	}
	// 8XY2
	if opcode>>12 == 0x8 && opcode&0xF == 0x2 {
		logCode(opcode, "8XY2")
		x := opcode >> 8 & 0xF
		y := opcode >> 4 & 0xF
		c.V[x] &= c.V[y]
		return
	}
	// 8XY3
	if opcode>>12 == 0x8 && opcode&0xF == 0x3 {
		logCode(opcode, "8XY3")
		x := opcode >> 8 & 0xF
		y := opcode >> 4 & 0xF
		c.V[x] ^= c.V[y]
		return
	}
	// 8XY4
	if opcode>>12 == 0x8 && opcode&0xF == 0x4 {
		logCode(opcode, "8XY4")
		x := opcode >> 8 & 0xF
		y := opcode >> 4 & 0xF
		if c.V[x] > 0xFF-c.V[y] {
			c.V[0xF] = 1
		} else {
			c.V[0xF] = 0
		}
		c.V[x] += c.V[y]
		return
	}
	// 8XY5
	if opcode>>12 == 0x8 && opcode&0xF == 0x5 {
		logCode(opcode, "8XY5")
		x := opcode >> 8 & 0xF
		y := opcode >> 4 & 0xF
		if c.V[x] >= c.V[y] {
			c.V[0xF] = 1
		} else {
			c.V[0xF] = 0
		}
		c.V[x] -= c.V[y]
		return
	}
	// 8XY6
	if opcode>>12 == 0x8 && opcode&0xF == 0x6 {
		logCode(opcode, "8XY6")
		x := opcode >> 8 & 0xF
		c.V[0xF] = c.V[x] & 0x1
		c.V[x] >>= 1
		return
	}
	// 8XY7
	if opcode>>12 == 0x8 && opcode&0xF == 0x7 {
		logCode(opcode, "8XY7")
		x := opcode >> 8 & 0xF
		y := opcode >> 4 & 0xF
		if c.V[y] >= c.V[x] {
			c.V[0xF] = 1
		} else {
			c.V[0xF] = 0
		}
		c.V[x] = c.V[y] - c.V[x]
		return
	}
	// 8XYE
	if opcode>>12 == 0x8 && opcode&0xF == 0xE {
		logCode(opcode, "8XYE")
		x := opcode >> 8 & 0xF
		c.V[0xF] = c.V[x] >> 7
		c.V[x] <<= 1
		return
	}
	// 9XY0
	if opcode>>12 == 0x9 && opcode&0xF == 0x0 {
		logCode(opcode, "9XY0")
		x := opcode >> 8 & 0xF
		y := opcode >> 4 & 0xF
		if c.V[x] != c.V[y] {
			c.PC += 2
		}
		return
	}
	// ANNN
	if opcode>>12 == 0xA {
		logCode(opcode, "ANNN")
		nnn := opcode & 0xFFF
		c.I = nnn
		return
	}
	// BNNN
	if opcode>>12 == 0xB {
		logCode(opcode, "BNNN")
		nnn := opcode & 0xFFF
		c.PC = uint16(c.V[0x0]) + nnn
		return
	}
	// CXNN
	if opcode>>12 == 0xC {
		logCode(opcode, "CXNN")
		x := opcode >> 8 & 0xF
		nn := opcode & 0xFF
		c.V[x] = uint8(rand.Intn(256)) & uint8(nn)
		return
	}
	// DXYN
	if opcode>>12 == 0xD {
		logCode(opcode, "DXYN")
		c.V[0xF] = 0

		x := opcode >> 8 & 0xF
		y := opcode >> 4 & 0xF
		height := opcode & 0x000F

		vx := c.V[x]
		vy := c.V[y]

		for y_line := uint8(0); y_line < uint8(height); y_line++ {
			in_mem := c.emu.memory.storage[c.I+uint16(y_line)]
			for x_line := uint8(0); x_line < 8; x_line++ {
				if in_mem>>(7-x_line)&1 != 0 {
					if c.emu.screen.board[(vy+y_line)%32][(vx+x_line)%64] == 1 {
						c.V[0xF] = 1
					}
					c.emu.screen.board[(vy+y_line)%32][(vx+x_line)%64] ^= 1
				}
			}
		}

		return
	}
	// EX9E
	if opcode>>12 == 0xE && opcode&0xFF == 0x9E {
		logCode(opcode, "EX9E")
		x := opcode >> 8 & 0xF
		if c.emu.key[c.V[x]] != 0 {
			c.PC += 2
		}
		return
	}
	// EXA1
	if opcode>>12 == 0xE && opcode&0xFF == 0xA1 {
		logCode(opcode, "EXA1")
		x := opcode >> 8 & 0xF
		if c.emu.key[c.V[x]] == 0 {
			c.PC += 2
		}
		return
	}
	// FX07
	if opcode>>12 == 0xF && opcode&0xFF == 0x07 {
		logCode(opcode, "FX07")
		x := opcode >> 8 & 0xF
		c.emu.delay_timer_mtx.RLock()
		c.V[x] = c.emu.delay_timer
		c.emu.delay_timer_mtx.RUnlock()
		return
	}
	// FX0A
	if opcode>>12 == 0xF && opcode&0xFF == 0x0A {
		logCode(opcode, "FX0A")
		x := opcode >> 8 & 0xF
		c.emu.wait_key_pressed = true
		go func() {
			c.V[x] = <-c.emu.wait_key
		}()
		return
	}
	// FX15
	if opcode>>12 == 0xF && opcode&0xFF == 0x15 {
		logCode(opcode, "FX15")
		c.emu.delay_timer_mtx.Lock()
		c.emu.delay_timer = c.V[uint8(opcode>>8&0xF)]
		c.emu.delay_timer_mtx.Unlock()
		return
	}
	// FX18
	if opcode>>12 == 0xF && opcode&0xFF == 0x18 {
		logCode(opcode, "FX18")
		c.emu.sound_timer_mtx.Lock()
		c.emu.sound_timer = c.V[uint8(opcode>>8&0xF)]
		c.emu.sound_timer_mtx.Unlock()
		return
	}
	// FX1E
	if opcode>>12 == 0xF && opcode&0xFF == 0x1E {
		logCode(opcode, "FX1E")
		x := opcode >> 8 & 0xF
		c.I += uint16(c.V[x])
		return
	}
	// FX29
	if opcode>>12 == 0xF && opcode&0xFF == 0x29 {
		logCode(opcode, "FX29")
		x := opcode >> 8 & 0xF
		c.I = uint16(c.emu.memory.GetBuiltInFontAddr(uint8(c.V[x])))
		return
	}
	// FX33
	if opcode>>12 == 0xF && opcode&0xFF == 0x33 {
		logCode(opcode, "FX33")
		x := opcode >> 8 & 0xF
		vx := c.V[x]
		c.emu.memory.storage[c.I] = vx / 100
		c.emu.memory.storage[c.I+1] = (vx / 10) % 10
		c.emu.memory.storage[c.I+2] = vx % 10
		return
	}
	// FX55
	if opcode>>12 == 0xF && opcode&0xFF == 0x55 {
		logCode(opcode, "FX55")
		x := opcode >> 8 & 0xF
		for i := uint16(0); i <= x; i++ {
			c.emu.memory.storage[c.I+i] = c.V[i]
		}
		return
	}
	// FX65
	if opcode>>12 == 0xF && opcode&0xFF == 0x65 {
		logCode(opcode, "FX65")
		x := opcode >> 8 & 0xF
		for i := uint16(0); i <= x; i++ {
			c.V[i] = c.emu.memory.storage[c.I+i]
		}
		return
	}
}
