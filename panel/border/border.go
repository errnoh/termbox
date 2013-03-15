// Copyright 2013 errnoh. All rights reserved.
// Use of this source code is governed by
// MIT License that can be found in the LICENSE file.

package border

import (
	"github.com/errnoh/termbox/panel"
	termbox "github.com/nsf/termbox-go"
)

func DrawAll(fg, bg termbox.Attribute, p ...panel.Panel) {
	for _, p2 := range p {
		if v, ok := p2.(panel.Drawer); ok {
			v.Draw()
		}
		Add(fg, bg, p2)
	}
}

// Draws borders _outside_ the panel.
func Add(fg, bg termbox.Attribute, p panel.Panel) panel.Panel {
	r := p.Bounds()

	termbox.SetCell(r.Min.X-1, r.Min.Y-1, boxdrawing[12], fg, bg)
	termbox.SetCell(r.Max.X, r.Min.Y-1, boxdrawing[9], fg, bg)
	termbox.SetCell(r.Min.X-1, r.Max.Y, boxdrawing[6], fg, bg)
	termbox.SetCell(r.Max.X, r.Max.Y, boxdrawing[3], fg, bg)

	for x := r.Min.X; x < r.Min.X+r.Dx(); x++ {
		termbox.SetCell(x, r.Min.Y-1, boxdrawing[5], fg, bg)
		termbox.SetCell(x, r.Min.Y+r.Dy(), boxdrawing[5], fg, bg)
	}

	for y := r.Min.Y; y < r.Min.Y+r.Dy(); y++ {
		termbox.SetCell(r.Min.X-1, y, boxdrawing[10], fg, bg)
		termbox.SetCell(r.Min.X+r.Dx(), y, boxdrawing[10], fg, bg)
	}

	return p
}

func borderClash(x, y int, char rune, side int) rune {
	m := panel.MainScreen()
	orig := m.At(x, y).Ch
	// Does it clash with other border?
	if orig < '─' || orig > '╿' {
		return char
	}

	// origia ei pitäisi hatuttaa vaan se origin vastaava numero.
	switch side {
	case 1: // left
		orig = orig ^ (1 + 16)
	case 2: // up
		orig = orig ^ (2 + 32)
	case 4: // right
		orig = orig ^ (4 + 64)
	case 8: // down
		orig = orig ^ (8 + 128)
	}
	return orig
}

var boxdrawing map[uint8]rune = map[uint8]rune{
	1 + 4:         '─',
	2 + 8:         '│',
	4 + 8:         '┌',
	1 + 8:         '┐',
	2 + 4:         '└',
	1 + 2:         '┘',
	2 + 4 + 8:     '├',
	1 + 2 + 8:     '┤',
	1 + 4 + 8:     '┬',
	1 + 2 + 4:     '┴',
	1 + 2 + 4 + 8: '┼',
	80:            '═',
	160:           '║',
	72:            '╒',
	132:           '╓',
	192:           '╔',
	24:            '╕',
	129:           '╖',
	144:           '╗',
	66:            '╘',
	36:            '╙',
	96:            '╚',
	18:            '╛',
	33:            '╜',
	48:            '╝',
	74:            '╞',
	164:           '╟',
	224:           '╠',
	26:            '╡',
	161:           '╢',
	176:           '╣',
	88:            '╤',
	133:           '╥',
	208:           '╦',
	82:            '╧',
	37:            '╨',
	112:           '╩',
	90:            '╪',
	165:           '╫',
	240:           '╬',
}
