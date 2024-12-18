package main

import (
	"fmt"
	"os"
)

func compile(data []byte) {
	var res = `// auto-generated
package core

import (
	"log"
	"math/rand"
)

type compiled struct {
emu *Chip8
}

func NewCompiled() *compiled {
return new(compiled)
}

func (cmp *compiled) Run(pc uint16) {
c := cmp.emu.cpu
c.PC += 2

`
	pc := 0
	for pc < len(data) {
		res += parse_opcode(uint16(data[pc])<<8+uint16(data[pc+1]), pc+0x200) + "\n"
		pc += 2
	}
	res += `log.Println("UNKNOWN PC")
}`
	os.WriteFile("core/compiled.go", []byte(res), 0700)
}

func parse_opcode(opcode uint16, pc int) (code string) {
	// 0NNN
	// do nothing
	if opcode>>12 == 0 && opcode != 0x00E0 && opcode != 0xEE {
		code = fmt.Sprintf(`if pc == %d { // 0NNN
return
}`, pc)
		return
	}
	// 00E0
	if opcode == 0xE0 {

		code = fmt.Sprintf(`if pc == %d { // 00E0
c.emu.screen.clear()
return
}`, pc)
		return
	}
	// 00EE
	if opcode == 0xEE {

		code = fmt.Sprintf(`if pc == %d { // 00EE
c.PC = c.emu.stack[c.emu.sp-1]
c.emu.sp--
return
}`, pc)
		return
	}
	// 1NNN
	if opcode>>12 == 0x1 {
		nnn := opcode & 0xFFF

		code = fmt.Sprintf(`if pc == %d { // 1NNN
c.PC = %d
return
}`, pc, nnn)
		return
	}
	// 2NNN
	if opcode>>12 == 0x2 {
		nnn := opcode & 0xFFF

		code = fmt.Sprintf(`if pc == %d { // 2NNN
c.emu.stack[c.emu.sp] = c.PC
c.emu.sp++
c.PC = %d
return
}`, pc, nnn)
		return
	}
	// 3XNN
	if opcode>>12 == 0x3 {
		x := opcode >> 8 & 0xF
		nn := opcode & 0xFF

		code = fmt.Sprintf(`if pc == %d { // 3XNN
if c.V[%d] == uint8(%d) {
c.PC += 2
}
return
}`, pc, x, nn)
		return
	}
	// 4XNN
	if opcode>>12 == 0x4 {
		x := opcode >> 8 & 0xF
		nn := opcode & 0xFF

		code = fmt.Sprintf(`if pc == %d { // 4XNN
if c.V[%d] != uint8(%d) {
c.PC += 2
}
return
}`, pc, x, nn)
		return
	}
	// 5XY0
	if opcode>>12 == 0x5 && opcode&0xF == 0x0 {
		x := opcode >> 8 & 0xF
		y := opcode >> 4 & 0xF

		code = fmt.Sprintf(`if pc == %d { // 5XY0
if c.V[%d] == c.V[%d] {
c.PC += 2
}
return
}`, pc, x, y)
		return
	}
	// 6XNN
	if opcode>>12 == 0x6 {
		x := opcode >> 8 & 0xF
		nn := opcode & 0xFF

		code = fmt.Sprintf(`if pc == %d { // 6XNN
c.V[%d] = uint8(%d)
return
}`, pc, x, nn)
		return
	}
	// 7XNN
	if opcode>>12 == 0x7 {
		x := opcode >> 8 & 0xF
		nn := opcode & 0xFF

		code = fmt.Sprintf(`if pc == %d { // 7XNN
c.V[%d] += uint8(%d)
return
}`, pc, x, nn)
		return
	}
	// 8XY0
	if opcode>>12 == 0x8 && opcode&0xF == 0x0 {
		x := opcode >> 8 & 0xF
		y := opcode >> 4 & 0xF
		code = fmt.Sprintf(`if pc == %d { // 8XY0
c.V[%d] = c.V[%d]
return
}`, pc, x, y)
		return
	}
	// 8XY1
	if opcode>>12 == 0x8 && opcode&0xF == 0x1 {
		x := opcode >> 8 & 0xF
		y := opcode >> 4 & 0xF

		code = fmt.Sprintf(`if pc == %d { // 8XY1
c.V[%d] |= c.V[%d]
return
}`, pc, x, y)
		return
	}
	// 8XY2
	if opcode>>12 == 0x8 && opcode&0xF == 0x2 {
		x := opcode >> 8 & 0xF
		y := opcode >> 4 & 0xF

		code = fmt.Sprintf(`if pc == %d { // 8XY2
c.V[%d] &= c.V[%d]
return
}`, pc, x, y)
		return
	}
	// 8XY3
	if opcode>>12 == 0x8 && opcode&0xF == 0x3 {
		x := opcode >> 8 & 0xF
		y := opcode >> 4 & 0xF

		code = fmt.Sprintf(`if pc == %d { // 8XY3
c.V[%d] ^= c.V[%d]
return
}`, pc, x, y)
		return
	}
	// 8XY4
	if opcode>>12 == 0x8 && opcode&0xF == 0x4 {
		x := opcode >> 8 & 0xF
		y := opcode >> 4 & 0xF

		code = fmt.Sprintf(`if pc == %d { // 8XY4
if c.V[%d] > 0xFF-c.V[%d] {
c.V[0xF] = 1
} else {
c.V[0xF] = 0
}
c.V[%d] += c.V[%d]
return
}`, pc, x, y, x, y)
		return
	}
	// 8XY5
	if opcode>>12 == 0x8 && opcode&0xF == 0x5 {
		x := opcode >> 8 & 0xF
		y := opcode >> 4 & 0xF

		code = fmt.Sprintf(`if pc == %d { // 8XY5
if c.V[%d] >= c.V[%d] {
c.V[0xF] = 1
} else {
c.V[0xF] = 0
}
c.V[%d] -= c.V[%d]
return
}`, pc, x, y, x, y)
		return
	}
	// 8XY6
	if opcode>>12 == 0x8 && opcode&0xF == 0x6 {
		x := opcode >> 8 & 0xF

		code = fmt.Sprintf(`if pc == %d { // 8XY6
c.V[0xF] = c.V[%d] & 0x1
c.V[%d] >>= 1
return
}`, pc, x, x)
		return
	}
	// 8XY7
	if opcode>>12 == 0x8 && opcode&0xF == 0x7 {
		x := opcode >> 8 & 0xF
		y := opcode >> 4 & 0xF

		code = fmt.Sprintf(`if pc == %d { // 8XY7
if c.V[%d] >= c.V[%d] {
c.V[0xF] = 1
} else {
c.V[0xF] = 0
}
c.V[%d] = c.V[%d] - c.V[%d]
return
}`, pc, y, x, x, y, x)
		return
	}
	// 8XYE
	if opcode>>12 == 0x8 && opcode&0xF == 0xE {
		x := opcode >> 8 & 0xF

		code = fmt.Sprintf(`if pc == %d { // 8XYE
c.V[0xF] = c.V[%d] >> 7
c.V[%d] <<= 1
return
}`, pc, x, x)
		return
	}
	// 9XY0
	if opcode>>12 == 0x9 && opcode&0xF == 0x0 {
		x := opcode >> 8 & 0xF
		y := opcode >> 4 & 0xF

		code = fmt.Sprintf(`if pc == %d { // 9XY0
if c.V[%d] != c.V[%d] {
c.PC += 2
}
return
}`, pc, x, y)
		return
	}
	// ANNN
	if opcode>>12 == 0xA {
		nnn := opcode & 0xFFF

		code = fmt.Sprintf(`if pc == %d { // ANNN
c.I = %d
return
}`, pc, nnn)
		return
	}
	// BNNN
	if opcode>>12 == 0xB {
		nnn := opcode & 0xFFF

		code = fmt.Sprintf(`if pc == %d { // BNNN
c.PC = uint16(c.V[0x0]) + %d
return
}`, pc, nnn)
		return
	}
	// CXNN
	if opcode>>12 == 0xC {
		x := opcode >> 8 & 0xF
		nn := opcode & 0xFF

		code = fmt.Sprintf(`if pc == %d { // CXNN
c.V[%d] = uint8(rand.Intn(256)) & uint8(%d)
return
}`, pc, x, nn)
		return
	}
	// DXYN
	if opcode>>12 == 0xD {

		x := opcode >> 8 & 0xF
		y := opcode >> 4 & 0xF
		height := opcode & 0x000F

		code = fmt.Sprintf(`if pc == %d { // DXYN
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
return
}`, pc, x, y, height)
		return
	}
	// EX9E
	if opcode>>12 == 0xE && opcode&0xFF == 0x9E {
		x := opcode >> 8 & 0xF

		code = fmt.Sprintf(`if pc == %d { // EX9E
if c.emu.key[c.V[%d]] != 0 {
c.PC += 2
}
return
}`, pc, x)
		return
	}
	// EXA1
	if opcode>>12 == 0xE && opcode&0xFF == 0xA1 {
		x := opcode >> 8 & 0xF

		code = fmt.Sprintf(`if pc == %d { // EXA1
if c.emu.key[c.V[%d]] == 0 {
c.PC += 2
}
return
}`, pc, x)
		return
	}
	// FX07
	if opcode>>12 == 0xF && opcode&0xFF == 0x07 {
		x := opcode >> 8 & 0xF

		code = fmt.Sprintf(`if pc == %d { // FX07
c.emu.delay_timer_mtx.RLock()
c.V[%d] = c.emu.delay_timer
c.emu.delay_timer_mtx.RUnlock()
return
}`, pc, x)
		return
	}
	// FX0A
	if opcode>>12 == 0xF && opcode&0xFF == 0x0A {
		x := opcode >> 8 & 0xF

		code = fmt.Sprintf(`if pc == %d { // FX0A
c.emu.wait_key_pressed = true
go func() {
c.V[%d] = <-c.emu.wait_key
}()
return
}`, pc, x)
		return
	}
	// FX15
	if opcode>>12 == 0xF && opcode&0xFF == 0x15 {
		x := opcode >> 8 & 0xF

		code = fmt.Sprintf(`if pc == %d { // FX15
c.emu.delay_timer_mtx.Lock()
c.emu.delay_timer = c.V[uint8(%d)]
c.emu.delay_timer_mtx.Unlock()
return
}`, pc, x)
		return
	}
	// FX18
	if opcode>>12 == 0xF && opcode&0xFF == 0x18 {
		x := opcode >> 8 & 0xF

		code = fmt.Sprintf(`if pc == %d { // FX18
c.emu.sound_timer_mtx.Lock()
c.emu.sound_timer = c.V[uint8(%d)]
c.emu.sound_timer_mtx.Unlock()
return
}`, pc, x)
		return
	}
	// FX1E
	if opcode>>12 == 0xF && opcode&0xFF == 0x1E {
		x := opcode >> 8 & 0xF

		code = fmt.Sprintf(`if pc == %d { // FX1E
c.I += uint16(c.V[%d])
return
}`, pc, x)
		return
	}
	// FX29
	if opcode>>12 == 0xF && opcode&0xFF == 0x29 {
		x := opcode >> 8 & 0xF

		code = fmt.Sprintf(`if pc == %d { // FX29
c.I = uint16(c.emu.memory.GetBuiltInFontAddr(uint8(c.V[%d])))
return
}`, pc, x)
		return
	}
	// FX33
	if opcode>>12 == 0xF && opcode&0xFF == 0x33 {
		x := opcode >> 8 & 0xF

		code = fmt.Sprintf(`if pc == %d { // FX33
vx := c.V[%d]
c.emu.memory.storage[c.I] = vx / 100
c.emu.memory.storage[c.I+1] = (vx / 10) %% 10
c.emu.memory.storage[c.I+2] = vx %% 10
return
}`, pc, x)
		return
	}
	// FX55
	if opcode>>12 == 0xF && opcode&0xFF == 0x55 {
		x := opcode >> 8 & 0xF

		code = fmt.Sprintf(`if pc == %d { // FX55
for i := uint16(0); i <= %d; i++ {
c.emu.memory.storage[c.I+i] = c.V[i]
}
return
}`, pc, x)
		return
	}
	// FX65
	if opcode>>12 == 0xF && opcode&0xFF == 0x65 {
		x := opcode >> 8 & 0xF

		code = fmt.Sprintf(`if pc == %d { // FX65
for i := uint16(0); i <= %d; i++ {
c.V[i] = c.emu.memory.storage[c.I+i]
}
return
}`, pc, x)
		return
	}
	return
}
