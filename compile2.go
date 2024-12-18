package main

import (
	"fmt"
	"os"
	"strconv"
)

func compile2(data []byte) {
	var res = `// auto-generated
package core

import (
	"math/rand"
)

type compiled struct {
emu *Chip8
funcs [` + strconv.Itoa(len(data)) + `]func(*cpu)
}

func (cmp *compiled) Run(pc uint16) {
c := cmp.emu.cpu
cmp.funcs[c.PC - 0x200](c)
}
func NewCompiled() *compiled {
res := new(compiled)

`
	pc := 0
	for pc < len(data) {
		res += "res.funcs[" + strconv.Itoa(pc) + "] = "
		res += parse_opcode2(uint16(data[pc])<<8+uint16(data[pc+1]), pc+0x200) + "\n"
		pc += 2
	}
	res += `return res
}
`
	os.WriteFile("core/compiled.go", []byte(res), 0700)
}

func parse_opcode2(opcode uint16, pc int) (code string) {
	// 0NNN
	// do nothing
	if opcode>>12 == 0 && opcode != 0x00E0 && opcode != 0xEE {
		code = fmt.Sprintf(`func(c *cpu) {
// 0NNN
c.PC += 2
}`)
		return
	}
	// 00E0
	if opcode == 0xE0 {

		code = fmt.Sprintf(`func(c *cpu) {
// 00E0
c.PC += 2
c.emu.screen.clear()
}`)
		return
	}
	// 00EE
	if opcode == 0xEE {

		code = fmt.Sprintf(`func(c *cpu) {
// 00EE
c.PC += 2
c.PC = c.emu.stack[c.emu.sp-1]
c.emu.sp--
}`)
		return
	}
	// 1NNN
	if opcode>>12 == 0x1 {
		nnn := opcode & 0xFFF

		code = fmt.Sprintf(`func(c *cpu) {
// 1NNN
c.PC += 2
c.PC = %d
}`, nnn)
		return
	}
	// 2NNN
	if opcode>>12 == 0x2 {
		nnn := opcode & 0xFFF

		code = fmt.Sprintf(`func(c *cpu) {
// 2NNN
c.PC += 2
c.emu.stack[c.emu.sp] = c.PC
c.emu.sp++
c.PC = %d
}`, nnn)
		return
	}
	// 3XNN
	if opcode>>12 == 0x3 {
		x := opcode >> 8 & 0xF
		nn := opcode & 0xFF

		code = fmt.Sprintf(`func(c *cpu) {
// 3XNN
c.PC += 2
if c.V[%d] == uint8(%d) {
c.PC += 2
}
}`, x, nn)
		return
	}
	// 4XNN
	if opcode>>12 == 0x4 {
		x := opcode >> 8 & 0xF
		nn := opcode & 0xFF

		code = fmt.Sprintf(`func(c *cpu) {
// 4XNN
c.PC += 2
if c.V[%d] != uint8(%d) {
c.PC += 2
}
}`, x, nn)
		return
	}
	// 5XY0
	if opcode>>12 == 0x5 && opcode&0xF == 0x0 {
		x := opcode >> 8 & 0xF
		y := opcode >> 4 & 0xF

		code = fmt.Sprintf(`func(c *cpu) {
// 5XY0
c.PC += 2
if c.V[%d] == c.V[%d] {
c.PC += 2
}
}`, x, y)
		return
	}
	// 6XNN
	if opcode>>12 == 0x6 {
		x := opcode >> 8 & 0xF
		nn := opcode & 0xFF

		code = fmt.Sprintf(`func(c *cpu) {
// 6XNN
c.PC += 2
c.V[%d] = uint8(%d)
}`, x, nn)
		return
	}
	// 7XNN
	if opcode>>12 == 0x7 {
		x := opcode >> 8 & 0xF
		nn := opcode & 0xFF

		code = fmt.Sprintf(`func(c *cpu) {
// 7XNN
c.PC += 2
c.V[%d] += uint8(%d)
}`, x, nn)
		return
	}
	// 8XY0
	if opcode>>12 == 0x8 && opcode&0xF == 0x0 {
		x := opcode >> 8 & 0xF
		y := opcode >> 4 & 0xF
		code = fmt.Sprintf(`func(c *cpu) {
// 8XY0
c.PC += 2
c.V[%d] = c.V[%d]
}`, x, y)
		return
	}
	// 8XY1
	if opcode>>12 == 0x8 && opcode&0xF == 0x1 {
		x := opcode >> 8 & 0xF
		y := opcode >> 4 & 0xF

		code = fmt.Sprintf(`func(c *cpu) {
// 8XY1
c.PC += 2
c.V[%d] |= c.V[%d]
}`, x, y)
		return
	}
	// 8XY2
	if opcode>>12 == 0x8 && opcode&0xF == 0x2 {
		x := opcode >> 8 & 0xF
		y := opcode >> 4 & 0xF

		code = fmt.Sprintf(`func(c *cpu) {
// 8XY2
c.PC += 2
c.V[%d] &= c.V[%d]
}`, x, y)
		return
	}
	// 8XY3
	if opcode>>12 == 0x8 && opcode&0xF == 0x3 {
		x := opcode >> 8 & 0xF
		y := opcode >> 4 & 0xF

		code = fmt.Sprintf(`func(c *cpu) {
// 8XY3
c.PC += 2
c.V[%d] ^= c.V[%d]
}`, x, y)
		return
	}
	// 8XY4
	if opcode>>12 == 0x8 && opcode&0xF == 0x4 {
		x := opcode >> 8 & 0xF
		y := opcode >> 4 & 0xF

		code = fmt.Sprintf(`func(c *cpu) {
// 8XY4
c.PC += 2
if c.V[%d] > 0xFF-c.V[%d] {
c.V[0xF] = 1
} else {
c.V[0xF] = 0
}
c.V[%d] += c.V[%d]
}`, x, y, x, y)
		return
	}
	// 8XY5
	if opcode>>12 == 0x8 && opcode&0xF == 0x5 {
		x := opcode >> 8 & 0xF
		y := opcode >> 4 & 0xF

		code = fmt.Sprintf(`func(c *cpu) {
// 8XY5
c.PC += 2
if c.V[%d] >= c.V[%d] {
c.V[0xF] = 1
} else {
c.V[0xF] = 0
}
c.V[%d] -= c.V[%d]
}`, x, y, x, y)
		return
	}
	// 8XY6
	if opcode>>12 == 0x8 && opcode&0xF == 0x6 {
		x := opcode >> 8 & 0xF

		code = fmt.Sprintf(`func(c *cpu) {
// 8XY6
c.PC += 2
c.V[0xF] = c.V[%d] & 0x1
c.V[%d] >>= 1
}`, x, x)
		return
	}
	// 8XY7
	if opcode>>12 == 0x8 && opcode&0xF == 0x7 {
		x := opcode >> 8 & 0xF
		y := opcode >> 4 & 0xF

		code = fmt.Sprintf(`func(c *cpu) {
// 8XY7
c.PC += 2
if c.V[%d] >= c.V[%d] {
c.V[0xF] = 1
} else {
c.V[0xF] = 0
}
c.V[%d] = c.V[%d] - c.V[%d]
}`, y, x, x, y, x)
		return
	}
	// 8XYE
	if opcode>>12 == 0x8 && opcode&0xF == 0xE {
		x := opcode >> 8 & 0xF

		code = fmt.Sprintf(`func(c *cpu) {
// 8XYE
c.PC += 2
c.V[0xF] = c.V[%d] >> 7
c.V[%d] <<= 1
}`, x, x)
		return
	}
	// 9XY0
	if opcode>>12 == 0x9 && opcode&0xF == 0x0 {
		x := opcode >> 8 & 0xF
		y := opcode >> 4 & 0xF

		code = fmt.Sprintf(`func(c *cpu) {
// 9XY0
c.PC += 2
if c.V[%d] != c.V[%d] {
c.PC += 2
}
}`, x, y)
		return
	}
	// ANNN
	if opcode>>12 == 0xA {
		nnn := opcode & 0xFFF

		code = fmt.Sprintf(`func(c *cpu) {
// ANNN
c.PC += 2
c.I = %d
}`, nnn)
		return
	}
	// BNNN
	if opcode>>12 == 0xB {
		nnn := opcode & 0xFFF

		code = fmt.Sprintf(`func(c *cpu) {
// BNNN
c.PC += 2
c.PC = uint16(c.V[0x0]) + %d
}`, nnn)
		return
	}
	// CXNN
	if opcode>>12 == 0xC {
		x := opcode >> 8 & 0xF
		nn := opcode & 0xFF

		code = fmt.Sprintf(`func(c *cpu) {
// CXNN
c.PC += 2
c.V[%d] = uint8(rand.Intn(256)) & uint8(%d)
}`, x, nn)
		return
	}
	// DXYN
	if opcode>>12 == 0xD {

		x := opcode >> 8 & 0xF
		y := opcode >> 4 & 0xF
		height := opcode & 0x000F

		code = fmt.Sprintf(`func(c *cpu) {
// DXYN
c.PC += 2
c.V[0xF] = 0
vx := c.V[%d]
vy := c.V[%d]

for y_line := uint8(0); y_line < uint8(%d); y_line++ {
in_mem := c.emu.memory.storage[c.I+uint16(y_line)]
for x_line := uint8(0); x_line < 8; x_line++ {
if in_mem>>(7-x_line)&1 != 0 {
if c.emu.screen.board[(vy+y_line)%%32][(vx+x_line)%%64] == 1 {
c.V[0xF] = 1
}
c.emu.screen.board[(vy+y_line)%%32][(vx+x_line)%%64] ^= 1
}
}
}
}`, x, y, height)
		return
	}
	// EX9E
	if opcode>>12 == 0xE && opcode&0xFF == 0x9E {
		x := opcode >> 8 & 0xF

		code = fmt.Sprintf(`func(c *cpu) {
// EX9E
c.PC += 2
if c.emu.key[c.V[%d]] != 0 {
c.PC += 2
}
}`, x)
		return
	}
	// EXA1
	if opcode>>12 == 0xE && opcode&0xFF == 0xA1 {
		x := opcode >> 8 & 0xF

		code = fmt.Sprintf(`func(c *cpu) {
// EXA1
c.PC += 2
if c.emu.key[c.V[%d]] == 0 {
c.PC += 2
}
}`, x)
		return
	}
	// FX07
	if opcode>>12 == 0xF && opcode&0xFF == 0x07 {
		x := opcode >> 8 & 0xF

		code = fmt.Sprintf(`func(c *cpu) {
// FX07
c.PC += 2
c.emu.delay_timer_mtx.RLock()
c.V[%d] = c.emu.delay_timer
c.emu.delay_timer_mtx.RUnlock()
}`, x)
		return
	}
	// FX0A
	if opcode>>12 == 0xF && opcode&0xFF == 0x0A {
		x := opcode >> 8 & 0xF

		code = fmt.Sprintf(`func(c *cpu) {
// FX0A
c.PC += 2
c.emu.wait_key_pressed = true
go func() {
c.V[%d] = <-c.emu.wait_key
}()
}`, x)
		return
	}
	// FX15
	if opcode>>12 == 0xF && opcode&0xFF == 0x15 {
		x := opcode >> 8 & 0xF

		code = fmt.Sprintf(`func(c *cpu) {
// FX15
c.PC += 2
c.emu.delay_timer_mtx.Lock()
c.emu.delay_timer = c.V[uint8(%d)]
c.emu.delay_timer_mtx.Unlock()
}`, x)
		return
	}
	// FX18
	if opcode>>12 == 0xF && opcode&0xFF == 0x18 {
		x := opcode >> 8 & 0xF

		code = fmt.Sprintf(`func(c *cpu) {
// FX18
c.PC += 2
c.emu.sound_timer_mtx.Lock()
c.emu.sound_timer = c.V[uint8(%d)]
c.emu.sound_timer_mtx.Unlock()
}`, x)
		return
	}
	// FX1E
	if opcode>>12 == 0xF && opcode&0xFF == 0x1E {
		x := opcode >> 8 & 0xF

		code = fmt.Sprintf(`func(c *cpu) {
// FX1E
c.PC += 2
c.I += uint16(c.V[%d])
}`, x)
		return
	}
	// FX29
	if opcode>>12 == 0xF && opcode&0xFF == 0x29 {
		x := opcode >> 8 & 0xF

		code = fmt.Sprintf(`func(c *cpu) {
// FX29
c.PC += 2
c.I = uint16(c.emu.memory.GetBuiltInFontAddr(uint8(c.V[%d])))
}`, x)
		return
	}
	// FX33
	if opcode>>12 == 0xF && opcode&0xFF == 0x33 {
		x := opcode >> 8 & 0xF

		code = fmt.Sprintf(`func(c *cpu) {
// FX33
c.PC += 2
vx := c.V[%d]
c.emu.memory.storage[c.I] = vx / 100
c.emu.memory.storage[c.I+1] = (vx / 10) %% 10
c.emu.memory.storage[c.I+2] = vx %% 10
}`, x)
		return
	}
	// FX55
	if opcode>>12 == 0xF && opcode&0xFF == 0x55 {
		x := opcode >> 8 & 0xF

		code = fmt.Sprintf(`func(c *cpu) {
// FX55
c.PC += 2
for i := uint16(0); i <= %d; i++ {
c.emu.memory.storage[c.I+i] = c.V[i]
}
}`, x)
		return
	}
	// FX65
	if opcode>>12 == 0xF && opcode&0xFF == 0x65 {
		x := opcode >> 8 & 0xF

		code = fmt.Sprintf(`func(c *cpu) {
// FX65
c.PC += 2
for i := uint16(0); i <= %d; i++ {
c.V[i] = c.emu.memory.storage[c.I+i]
}
}`, x)
		return
	}
	return
}
