// Copyright 2013 errnoh. All rights reserved.
// Use of this source code is governed by
// MIT License that can be found in the LICENSE file.

// Helper library for termbox-go to split screen into panels
package panel

import (
	termbox "github.com/nsf/termbox-go"
	"image"
)

type Panel interface {
	Bounds() image.Rectangle
	At(x, y int) termbox.Cell
	SetCell(x, y int, ch rune, fg, bg termbox.Attribute)
	Move(deltax, deltay int)
}

type Drawer interface {
	Draw()
}

func DrawAll(p ...Panel) {
	// NOTE: Used to be ...*Buffered, ...Panel is more similar to borders
	for _, panel := range p {
		if v, ok := panel.(Drawer); ok {
			v.Draw()
		}
	}
}

func Write(panel Panel, b []byte) (n int, err error) {
	// TODO: Fg & Bg attributes?
	var row, stop int
	s := string(b)
	for i, r := range s {
		if r == '\n' {
			stop = i + 1
			row++
			continue
		} else if r == '\r' {
			r = ' '
		}
		panel.SetCell(i-stop, row, r, 0, 0)
	}
	return 0, nil
}

// Panel that writes directly to termbox buffer.
// *Unbuffered methods can be used directly with *Buffered as well.
type Unbuffered struct {
	r image.Rectangle
}

func NewUnbuffered(r image.Rectangle) *Unbuffered {
	return &Unbuffered{
		r: r,
	}
}

func (panel *Unbuffered) SetCursor(x, y int) {
	if panel.Contains(x, y) {
		termbox.SetCursor(panel.r.Min.X+x, panel.r.Min.Y+y)
	}
}

// Returns unbuffered panel that contains area of 'panel' r specifies.
// NOTE: At current state *Unbuffered always writes to main termbox buffer.
func (panel *Unbuffered) Area(r image.Rectangle) *Unbuffered {
	newr := image.Rect(panel.r.Min.X+r.Min.X, panel.r.Min.Y+r.Min.Y, panel.r.Min.X+r.Max.X, panel.r.Min.Y+r.Max.Y)
	if newr.In(panel.r) {
		return NewUnbuffered(newr)
	}
	return nil

}

func (panel *Unbuffered) At(x, y int) termbox.Cell {
	if !panel.Contains(x, y) {
		return termbox.Cell{}
	}
	width, _ := termbox.Size()
	return termbox.CellBuffer()[(panel.r.Min.Y+y)*width+(panel.r.Min.X+x)]
}

func (panel *Unbuffered) Bounds() image.Rectangle {
	return panel.r
}

func (panel *Unbuffered) Contains(x, y int) bool {
	r := image.Rect(0, 0, panel.r.Dx(), panel.r.Dy())
	return image.Point{x, y}.In(r)
}

func (panel *Unbuffered) SetCell(x, y int, ch rune, fg, bg termbox.Attribute) {
	if panel.Contains(x, y) {
		termbox.SetCell(panel.r.Min.X+x, panel.r.Min.Y+y, ch, fg, bg)
	}
}

func (panel *Unbuffered) Write(b []byte) (n int, err error) {
	return Write(panel, b)
}

func (panel *Unbuffered) Move(deltax, deltay int) {
	panel.r = panel.r.Add(image.Point{deltax, deltay})
}

type Buffered struct {
	Unbuffered
	buffer []termbox.Cell
}

// Returns buffered struct of the main screen.
// NOTE: Resize creates new buffer?
func MainScreen() *Buffered {
	width, height := termbox.Size()
	return &Buffered{
		Unbuffered: Unbuffered{image.Rect(0, 0, width, height)},
		buffer:     termbox.CellBuffer(),
	}
}

func NewBuffered(r image.Rectangle) *Buffered {
	return &Buffered{
		Unbuffered: Unbuffered{r},
		buffer:     make([]termbox.Cell, r.Dx()*r.Dy()),
	}
}

func (panel *Buffered) SetCell(x, y int, ch rune, fg, bg termbox.Attribute) {
	if panel.Contains(x, y) {
		panel.buffer[y*panel.r.Dx()+x] = termbox.Cell{ch, fg, bg}
	}
}

func (panel *Buffered) Write(b []byte) (n int, err error) {
	return Write(panel, b)
}

func (panel *Buffered) Clear() *Buffered {
	for i := 0; i < len(panel.buffer); i++ {
		panel.buffer[i] = termbox.Cell{' ', 0, 0}
	}
	return panel
}

func (panel *Buffered) Draw() {
	main := MainScreen()
	target := main.r.Intersect(panel.r)
	if target.Empty() {
		return
	}

	// Which point of the panel is the starting point
	// (can't be negative, target must be inside panel rectangle)
	row := target.Min.Y - panel.r.Min.Y
	col := target.Min.X - panel.r.Min.X

	for y := 0; y < target.Dy(); y++ {
		copy(main.buffer[(target.Min.Y+y)*main.r.Dx()+target.Min.X:(target.Min.Y+y)*main.r.Dx()+target.Max.X], panel.buffer[(row+y)*panel.r.Dx()+col:(row+y)*panel.r.Dx()+col+target.Dx()])
	}
	return
}
