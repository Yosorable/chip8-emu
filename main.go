package main

import (
	"chip8-emu/core"
	"log"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	emu := new(core.Chip8)
	emu.Initialize()
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		log.Panicf("Load sdl error: " + err.Error())
	}
	defer sdl.Quit()
	sdl.StopTextInput()

	// rom := "ROMS/PONG"
	rom := "ROMS/TETRIS"
	// rom := "ROMS/15PUZZLE"
	data, err := os.ReadFile(rom)
	if err != nil {
		log.Panicf("Load game error: " + err.Error())
	}

	emu.StartGame(data)
}
