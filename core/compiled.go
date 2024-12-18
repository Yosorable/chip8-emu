// auto-generated
package core

import (
	"math/rand"
)

type compiled struct {
emu *Chip8
funcs [246]func(*cpu)
}

func (cmp *compiled) Run(pc uint16) {
c := cmp.emu.cpu
cmp.funcs[c.PC - 0x200](c)
}
func NewCompiled() *compiled {
res := new(compiled)

res.funcs[0] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[10] = uint8(2)
}
res.funcs[2] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[11] = uint8(12)
}
res.funcs[4] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[12] = uint8(63)
}
res.funcs[6] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[13] = uint8(12)
}
res.funcs[8] = func(c *cpu) {
// ANNN
c.PC += 2
c.I = 746
}
res.funcs[10] = func(c *cpu) {
// DXYN
c.PC += 2
c.V[0xF] = 0
vx := c.V[10]
vy := c.V[11]

for y_line := uint8(0); y_line < uint8(6); y_line++ {
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
}
res.funcs[12] = func(c *cpu) {
// DXYN
c.PC += 2
c.V[0xF] = 0
vx := c.V[12]
vy := c.V[13]

for y_line := uint8(0); y_line < uint8(6); y_line++ {
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
}
res.funcs[14] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[14] = uint8(0)
}
res.funcs[16] = func(c *cpu) {
// 2NNN
c.PC += 2
c.emu.stack[c.emu.sp] = c.PC
c.emu.sp++
c.PC = 724
}
res.funcs[18] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[6] = uint8(3)
}
res.funcs[20] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[8] = uint8(2)
}
res.funcs[22] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[0] = uint8(96)
}
res.funcs[24] = func(c *cpu) {
// FX15
c.PC += 2
c.emu.delay_timer_mtx.Lock()
c.emu.delay_timer = c.V[uint8(0)]
c.emu.delay_timer_mtx.Unlock()
}
res.funcs[26] = func(c *cpu) {
// FX07
c.PC += 2
c.emu.delay_timer_mtx.RLock()
c.V[0] = c.emu.delay_timer
c.emu.delay_timer_mtx.RUnlock()
}
res.funcs[28] = func(c *cpu) {
// 3XNN
c.PC += 2
if c.V[0] == uint8(0) {
c.PC += 2
}
}
res.funcs[30] = func(c *cpu) {
// 1NNN
c.PC += 2
c.PC = 538
}
res.funcs[32] = func(c *cpu) {
// CXNN
c.PC += 2
c.V[7] = uint8(rand.Intn(256)) & uint8(23)
}
res.funcs[34] = func(c *cpu) {
// 7XNN
c.PC += 2
c.V[7] += uint8(8)
}
res.funcs[36] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[9] = uint8(255)
}
res.funcs[38] = func(c *cpu) {
// ANNN
c.PC += 2
c.I = 752
}
res.funcs[40] = func(c *cpu) {
// DXYN
c.PC += 2
c.V[0xF] = 0
vx := c.V[6]
vy := c.V[7]

for y_line := uint8(0); y_line < uint8(1); y_line++ {
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
}
res.funcs[42] = func(c *cpu) {
// ANNN
c.PC += 2
c.I = 746
}
res.funcs[44] = func(c *cpu) {
// DXYN
c.PC += 2
c.V[0xF] = 0
vx := c.V[10]
vy := c.V[11]

for y_line := uint8(0); y_line < uint8(6); y_line++ {
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
}
res.funcs[46] = func(c *cpu) {
// DXYN
c.PC += 2
c.V[0xF] = 0
vx := c.V[12]
vy := c.V[13]

for y_line := uint8(0); y_line < uint8(6); y_line++ {
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
}
res.funcs[48] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[0] = uint8(1)
}
res.funcs[50] = func(c *cpu) {
// EXA1
c.PC += 2
if c.emu.key[c.V[0]] == 0 {
c.PC += 2
}
}
res.funcs[52] = func(c *cpu) {
// 7XNN
c.PC += 2
c.V[11] += uint8(254)
}
res.funcs[54] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[0] = uint8(4)
}
res.funcs[56] = func(c *cpu) {
// EXA1
c.PC += 2
if c.emu.key[c.V[0]] == 0 {
c.PC += 2
}
}
res.funcs[58] = func(c *cpu) {
// 7XNN
c.PC += 2
c.V[11] += uint8(2)
}
res.funcs[60] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[0] = uint8(31)
}
res.funcs[62] = func(c *cpu) {
// 8XY2
c.PC += 2
c.V[11] &= c.V[0]
}
res.funcs[64] = func(c *cpu) {
// DXYN
c.PC += 2
c.V[0xF] = 0
vx := c.V[10]
vy := c.V[11]

for y_line := uint8(0); y_line < uint8(6); y_line++ {
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
}
res.funcs[66] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[0] = uint8(12)
}
res.funcs[68] = func(c *cpu) {
// EXA1
c.PC += 2
if c.emu.key[c.V[0]] == 0 {
c.PC += 2
}
}
res.funcs[70] = func(c *cpu) {
// 7XNN
c.PC += 2
c.V[13] += uint8(254)
}
res.funcs[72] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[0] = uint8(13)
}
res.funcs[74] = func(c *cpu) {
// EXA1
c.PC += 2
if c.emu.key[c.V[0]] == 0 {
c.PC += 2
}
}
res.funcs[76] = func(c *cpu) {
// 7XNN
c.PC += 2
c.V[13] += uint8(2)
}
res.funcs[78] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[0] = uint8(31)
}
res.funcs[80] = func(c *cpu) {
// 8XY2
c.PC += 2
c.V[13] &= c.V[0]
}
res.funcs[82] = func(c *cpu) {
// DXYN
c.PC += 2
c.V[0xF] = 0
vx := c.V[12]
vy := c.V[13]

for y_line := uint8(0); y_line < uint8(6); y_line++ {
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
}
res.funcs[84] = func(c *cpu) {
// ANNN
c.PC += 2
c.I = 752
}
res.funcs[86] = func(c *cpu) {
// DXYN
c.PC += 2
c.V[0xF] = 0
vx := c.V[6]
vy := c.V[7]

for y_line := uint8(0); y_line < uint8(1); y_line++ {
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
}
res.funcs[88] = func(c *cpu) {
// 8XY4
c.PC += 2
if c.V[6] > 0xFF-c.V[8] {
c.V[0xF] = 1
} else {
c.V[0xF] = 0
}
c.V[6] += c.V[8]
}
res.funcs[90] = func(c *cpu) {
// 8XY4
c.PC += 2
if c.V[7] > 0xFF-c.V[9] {
c.V[0xF] = 1
} else {
c.V[0xF] = 0
}
c.V[7] += c.V[9]
}
res.funcs[92] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[0] = uint8(63)
}
res.funcs[94] = func(c *cpu) {
// 8XY2
c.PC += 2
c.V[6] &= c.V[0]
}
res.funcs[96] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[1] = uint8(31)
}
res.funcs[98] = func(c *cpu) {
// 8XY2
c.PC += 2
c.V[7] &= c.V[1]
}
res.funcs[100] = func(c *cpu) {
// 4XNN
c.PC += 2
if c.V[6] != uint8(2) {
c.PC += 2
}
}
res.funcs[102] = func(c *cpu) {
// 1NNN
c.PC += 2
c.PC = 632
}
res.funcs[104] = func(c *cpu) {
// 4XNN
c.PC += 2
if c.V[6] != uint8(63) {
c.PC += 2
}
}
res.funcs[106] = func(c *cpu) {
// 1NNN
c.PC += 2
c.PC = 642
}
res.funcs[108] = func(c *cpu) {
// 4XNN
c.PC += 2
if c.V[7] != uint8(31) {
c.PC += 2
}
}
res.funcs[110] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[9] = uint8(255)
}
res.funcs[112] = func(c *cpu) {
// 4XNN
c.PC += 2
if c.V[7] != uint8(0) {
c.PC += 2
}
}
res.funcs[114] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[9] = uint8(1)
}
res.funcs[116] = func(c *cpu) {
// DXYN
c.PC += 2
c.V[0xF] = 0
vx := c.V[6]
vy := c.V[7]

for y_line := uint8(0); y_line < uint8(1); y_line++ {
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
}
res.funcs[118] = func(c *cpu) {
// 1NNN
c.PC += 2
c.PC = 554
}
res.funcs[120] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[8] = uint8(2)
}
res.funcs[122] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[3] = uint8(1)
}
res.funcs[124] = func(c *cpu) {
// 8XY0
c.PC += 2
c.V[0] = c.V[7]
}
res.funcs[126] = func(c *cpu) {
// 8XY5
c.PC += 2
if c.V[0] >= c.V[11] {
c.V[0xF] = 1
} else {
c.V[0xF] = 0
}
c.V[0] -= c.V[11]
}
res.funcs[128] = func(c *cpu) {
// 1NNN
c.PC += 2
c.PC = 650
}
res.funcs[130] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[8] = uint8(254)
}
res.funcs[132] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[3] = uint8(10)
}
res.funcs[134] = func(c *cpu) {
// 8XY0
c.PC += 2
c.V[0] = c.V[7]
}
res.funcs[136] = func(c *cpu) {
// 8XY5
c.PC += 2
if c.V[0] >= c.V[13] {
c.V[0xF] = 1
} else {
c.V[0xF] = 0
}
c.V[0] -= c.V[13]
}
res.funcs[138] = func(c *cpu) {
// 3XNN
c.PC += 2
if c.V[15] == uint8(1) {
c.PC += 2
}
}
res.funcs[140] = func(c *cpu) {
// 1NNN
c.PC += 2
c.PC = 674
}
res.funcs[142] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[1] = uint8(2)
}
res.funcs[144] = func(c *cpu) {
// 8XY5
c.PC += 2
if c.V[0] >= c.V[1] {
c.V[0xF] = 1
} else {
c.V[0xF] = 0
}
c.V[0] -= c.V[1]
}
res.funcs[146] = func(c *cpu) {
// 3XNN
c.PC += 2
if c.V[15] == uint8(1) {
c.PC += 2
}
}
res.funcs[148] = func(c *cpu) {
// 1NNN
c.PC += 2
c.PC = 698
}
res.funcs[150] = func(c *cpu) {
// 8XY5
c.PC += 2
if c.V[0] >= c.V[1] {
c.V[0xF] = 1
} else {
c.V[0xF] = 0
}
c.V[0] -= c.V[1]
}
res.funcs[152] = func(c *cpu) {
// 3XNN
c.PC += 2
if c.V[15] == uint8(1) {
c.PC += 2
}
}
res.funcs[154] = func(c *cpu) {
// 1NNN
c.PC += 2
c.PC = 712
}
res.funcs[156] = func(c *cpu) {
// 8XY5
c.PC += 2
if c.V[0] >= c.V[1] {
c.V[0xF] = 1
} else {
c.V[0xF] = 0
}
c.V[0] -= c.V[1]
}
res.funcs[158] = func(c *cpu) {
// 3XNN
c.PC += 2
if c.V[15] == uint8(1) {
c.PC += 2
}
}
res.funcs[160] = func(c *cpu) {
// 1NNN
c.PC += 2
c.PC = 706
}
res.funcs[162] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[0] = uint8(32)
}
res.funcs[164] = func(c *cpu) {
// FX18
c.PC += 2
c.emu.sound_timer_mtx.Lock()
c.emu.sound_timer = c.V[uint8(0)]
c.emu.sound_timer_mtx.Unlock()
}
res.funcs[166] = func(c *cpu) {
// 2NNN
c.PC += 2
c.emu.stack[c.emu.sp] = c.PC
c.emu.sp++
c.PC = 724
}
res.funcs[168] = func(c *cpu) {
// 8XY4
c.PC += 2
if c.V[14] > 0xFF-c.V[3] {
c.V[0xF] = 1
} else {
c.V[0xF] = 0
}
c.V[14] += c.V[3]
}
res.funcs[170] = func(c *cpu) {
// 2NNN
c.PC += 2
c.emu.stack[c.emu.sp] = c.PC
c.emu.sp++
c.PC = 724
}
res.funcs[172] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[6] = uint8(62)
}
res.funcs[174] = func(c *cpu) {
// 3XNN
c.PC += 2
if c.V[3] == uint8(1) {
c.PC += 2
}
}
res.funcs[176] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[6] = uint8(3)
}
res.funcs[178] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[8] = uint8(254)
}
res.funcs[180] = func(c *cpu) {
// 3XNN
c.PC += 2
if c.V[3] == uint8(1) {
c.PC += 2
}
}
res.funcs[182] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[8] = uint8(2)
}
res.funcs[184] = func(c *cpu) {
// 1NNN
c.PC += 2
c.PC = 534
}
res.funcs[186] = func(c *cpu) {
// 7XNN
c.PC += 2
c.V[9] += uint8(255)
}
res.funcs[188] = func(c *cpu) {
// 4XNN
c.PC += 2
if c.V[9] != uint8(254) {
c.PC += 2
}
}
res.funcs[190] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[9] = uint8(255)
}
res.funcs[192] = func(c *cpu) {
// 1NNN
c.PC += 2
c.PC = 712
}
res.funcs[194] = func(c *cpu) {
// 7XNN
c.PC += 2
c.V[9] += uint8(1)
}
res.funcs[196] = func(c *cpu) {
// 4XNN
c.PC += 2
if c.V[9] != uint8(2) {
c.PC += 2
}
}
res.funcs[198] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[9] = uint8(1)
}
res.funcs[200] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[0] = uint8(4)
}
res.funcs[202] = func(c *cpu) {
// FX18
c.PC += 2
c.emu.sound_timer_mtx.Lock()
c.emu.sound_timer = c.V[uint8(0)]
c.emu.sound_timer_mtx.Unlock()
}
res.funcs[204] = func(c *cpu) {
// 7XNN
c.PC += 2
c.V[6] += uint8(1)
}
res.funcs[206] = func(c *cpu) {
// 4XNN
c.PC += 2
if c.V[6] != uint8(64) {
c.PC += 2
}
}
res.funcs[208] = func(c *cpu) {
// 7XNN
c.PC += 2
c.V[6] += uint8(254)
}
res.funcs[210] = func(c *cpu) {
// 1NNN
c.PC += 2
c.PC = 620
}
res.funcs[212] = func(c *cpu) {
// ANNN
c.PC += 2
c.I = 754
}
res.funcs[214] = func(c *cpu) {
// FX33
c.PC += 2
vx := c.V[14]
c.emu.memory.storage[c.I] = vx / 100
c.emu.memory.storage[c.I+1] = (vx / 10) % 10
c.emu.memory.storage[c.I+2] = vx % 10
}
res.funcs[216] = func(c *cpu) {
// FX65
c.PC += 2
for i := uint16(0); i <= 2; i++ {
c.V[i] = c.emu.memory.storage[c.I+i]
}
}
res.funcs[218] = func(c *cpu) {
// FX29
c.PC += 2
c.I = uint16(c.emu.memory.GetBuiltInFontAddr(uint8(c.V[1])))
}
res.funcs[220] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[4] = uint8(20)
}
res.funcs[222] = func(c *cpu) {
// 6XNN
c.PC += 2
c.V[5] = uint8(0)
}
res.funcs[224] = func(c *cpu) {
// DXYN
c.PC += 2
c.V[0xF] = 0
vx := c.V[4]
vy := c.V[5]

for y_line := uint8(0); y_line < uint8(5); y_line++ {
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
}
res.funcs[226] = func(c *cpu) {
// 7XNN
c.PC += 2
c.V[4] += uint8(21)
}
res.funcs[228] = func(c *cpu) {
// FX29
c.PC += 2
c.I = uint16(c.emu.memory.GetBuiltInFontAddr(uint8(c.V[2])))
}
res.funcs[230] = func(c *cpu) {
// DXYN
c.PC += 2
c.V[0xF] = 0
vx := c.V[4]
vy := c.V[5]

for y_line := uint8(0); y_line < uint8(5); y_line++ {
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
}
res.funcs[232] = func(c *cpu) {
// 00EE
c.PC += 2
c.PC = c.emu.stack[c.emu.sp-1]
c.emu.sp--
}
res.funcs[234] = func(c *cpu) {
// 8XY0
c.PC += 2
c.V[0] = c.V[8]
}
res.funcs[236] = func(c *cpu) {
// 8XY0
c.PC += 2
c.V[0] = c.V[8]
}
res.funcs[238] = func(c *cpu) {
// 8XY0
c.PC += 2
c.V[0] = c.V[8]
}
res.funcs[240] = func(c *cpu) {
// 8XY0
c.PC += 2
c.V[0] = c.V[0]
}
res.funcs[242] = func(c *cpu) {
// 0NNN
c.PC += 2
}
res.funcs[244] = func(c *cpu) {
// 0NNN
c.PC += 2
}
return res
}
