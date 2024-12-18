package core

import (
	"log"
	"sync"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type Chip8 struct {
	cpu    *cpu
	memory *memory
	screen *screen

	delay_timer     uint8
	delay_timer_mtx sync.RWMutex
	sound_timer     uint8
	sound_timer_mtx sync.RWMutex

	stack [16]uint16
	sp    uint8
	key   [16]uint8

	wait_key         chan uint8
	wait_key_pressed bool
}

func (emu *Chip8) Initialize() {
	if emu.cpu == nil {
		emu.cpu = new(cpu)
		emu.cpu.emu = emu
	}
	emu.cpu.Initialize()

	if emu.memory == nil {
		emu.memory = new(memory)
	}
	emu.memory.Initialize()

	if emu.screen == nil {
		emu.screen = new(screen)
		window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 64*15, 32*15, sdl.WINDOW_SHOWN)
		if err != nil {
			log.Panicln("Create window error: " + err.Error())
		}
		emu.screen.window = window
	}
	emu.screen.clear()

}

func (emu *Chip8) StartGame(rom []uint8) {
	copy(emu.memory.storage[0x200:], rom)
	emu.cpu.PC = 0x200

	running := true

	timer_ticker := time.NewTicker(time.Second / 60)
	frame_ticker := time.NewTicker(time.Second / 60)

	go func() {
		for {
			<-timer_ticker.C
			emu.delay_timer_mtx.Lock()
			if emu.delay_timer > 0 {
				emu.delay_timer--
			}
			emu.delay_timer_mtx.Unlock()
			emu.sound_timer_mtx.Lock()
			if emu.sound_timer > 0 {
				emu.sound_timer--
			}
			emu.sound_timer_mtx.Unlock()
		}
	}()
	for frame := 0; running; frame++ {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event := event.(type) {
			case *sdl.QuitEvent:
				log.Println("Quit")
				running = false
			case *sdl.KeyboardEvent:
				emu.SetKey(int(event.Keysym.Sym), event.Type == sdl.KEYDOWN)
			}
		}
		if !emu.wait_key_pressed {
			emu.cpu.ExecuteCode()
			if emu.cpu.PC == uint16(len(rom)) {
				break
			}
		}
		emu.screen.draw()
		<-frame_ticker.C
	}
}

func (emu *Chip8) SetKey(origin int, pressed bool) {
	key := uint8(0)
	switch origin {
	case sdl.K_1:
		key = 0
	case sdl.K_2:
		key = 1
	case sdl.K_3:
		key = 2
	case sdl.K_4:
		key = 3
	case sdl.K_q:
		key = 4
	case sdl.K_w:
		key = 5
	case sdl.K_e:
		key = 6
	case sdl.K_r:
		key = 7
	case sdl.K_a:
		key = 8
	case sdl.K_s:
		key = 9
	case sdl.K_d:
		key = 10
	case sdl.K_f:
		key = 11
	case sdl.K_z:
		key = 12
	case sdl.K_x:
		key = 13
	case sdl.K_c:
		key = 14
	case sdl.K_v:
		key = 15
	default:
		return
	}
	if pressed {
		if emu.wait_key_pressed {
			emu.wait_key <- key
			emu.wait_key_pressed = false
		}
		emu.key[key] = 1
		log.Printf("[keyboard] %d pressed", key)
	} else {
		emu.key[key] = 0
		log.Printf("[keyboard] %d released", key)
	}
}
