package core

import (
	"github.com/veandco/go-sdl2/sdl"
)

type screen struct {
	board  [32][64]uint8
	window *sdl.Window
}

func (s *screen) clear() {
	for i := 0; i < 32; i++ {
		for j := 0; j < 64; j++ {
			s.board[i][j] = 0
		}
	}

	surface, err := s.window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)
	s.window.UpdateSurface()
}

func (s *screen) draw() {

	surface, err := s.window.GetSurface()
	if err != nil {
		panic(err)
	}

	for i := 0; i < 32; i++ {
		for j := 0; j < 64; j++ {
			rect := sdl.Rect{X: int32(j) * 15, Y: int32(i) * 15, W: 15, H: 15}
			if s.board[i][j] != 0 {
				pixel := sdl.MapRGBA(surface.Format, 182, 204, 204, 255)
				surface.FillRect(&rect, pixel)
			} else {
				pixel := sdl.MapRGBA(surface.Format, 0, 0, 0, 255)
				surface.FillRect(&rect, pixel)
			}
		}
	}

	s.window.UpdateSurface()
}
