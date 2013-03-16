// Copyright 2013 errnoh. All rights reserved.
// Use of this source code is governed by
// MIT License that can be found in the LICENSE file.

package panel

import (
	termbox "github.com/nsf/termbox-go"
	"image"
)

type BorderedPanel interface {
	Bounds() image.Rectangle
	Border() termbox.Cell
}

// Draws borders _outside_ the 
func DrawBorder(p BorderedPanel) {
	r := p.Bounds()
	b := p.Border()

	var border []rune
	switch b.Ch {
	case 's', 'S':
		border = singleBorder
	case 'd', 'D':
		border = doubleBorder
	default:
		border = []rune{b.Ch, b.Ch, b.Ch, b.Ch, b.Ch, b.Ch}
	}

	termbox.SetCell(r.Min.X-1, r.Min.Y-1, borderClash(r.Min.X-1, r.Min.Y-1, border[4], 3), b.Fg, b.Bg)
	termbox.SetCell(r.Max.X, r.Min.Y-1, borderClash(r.Max.X, r.Min.Y-1, border[5], 6), b.Fg, b.Bg)
	termbox.SetCell(r.Min.X-1, r.Max.Y, borderClash(r.Min.X-1, r.Max.Y, border[3], 9), b.Fg, b.Bg)
	termbox.SetCell(r.Max.X, r.Max.Y, borderClash(r.Max.X, r.Max.Y, border[2], 12), b.Fg, b.Bg)

	for x := r.Min.X; x < r.Min.X+r.Dx(); x++ {
		termbox.SetCell(x, r.Min.Y-1, borderClash(x, r.Min.Y-1, border[0], 2), b.Fg, b.Bg)
		termbox.SetCell(x, r.Min.Y+r.Dy(), borderClash(x, r.Min.Y+r.Dy(), border[0], 8), b.Fg, b.Bg)
	}

	for y := r.Min.Y; y < r.Min.Y+r.Dy(); y++ {
		termbox.SetCell(r.Min.X-1, y, borderClash(r.Min.X-1, y, border[1], 1), b.Fg, b.Bg)
		termbox.SetCell(r.Min.X+r.Dx(), y, borderClash(r.Min.X+r.Dx(), y, border[1], 4), b.Fg, b.Bg)
	}
}

func borderClash(x, y int, char rune, side int) rune {
	m := MainScreen()
	orig := m.At(x, y).Ch
	if origVal, found := boxval[orig]; found {
		// NOTE: There are no symbols for all possible 3 way corners.
		switch side {
		case 1: // left
			origVal = origVal &^ (2 + 4 + 8 + 32 + 64 + 128)
		case 3:
			origVal = origVal &^ (4 + 8 + 64 + 128)
		case 2: // up
			origVal = origVal &^ (1 + 4 + 8 + 16 + 64 + 128)
		case 6:
			origVal = origVal &^ (1 + 8 + 16 + 128)
		case 4: // right
			origVal = origVal &^ (1 + 2 + 8 + 16 + 32 + 128)
		case 12:
			origVal = origVal &^ (1 + 2 + 16 + 32)
		case 8: // down
			origVal = origVal &^ (1 + 2 + 4 + 16 + 32 + 64)
		case 9:
			origVal = origVal &^ (2 + 4 + 32 + 64)
		}
		charVal := origVal | boxval[char]
		if newChar, ok := boxdrawing[charVal]; ok {
			char = newChar
		}
	}
	return char
}

var (
	// sideways, upwards, downright, downleft, upleft, upright
	singleBorder = []rune{'─', '│', '┘', '└', '┌', '┐'}
	doubleBorder = []rune{'═', '║', '╝', '╚', '╔', '╗'}
)

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

var boxval map[rune]uint8 = map[rune]uint8{
	'─': 1 + 4,
	'│': 2 + 8,
	'┌': 4 + 8,
	'┐': 1 + 8,
	'└': 2 + 4,
	'┘': 1 + 2,
	'├': 2 + 4 + 8,
	'┤': 1 + 2 + 8,
	'┬': 1 + 4 + 8,
	'┴': 1 + 2 + 4,
	'┼': 1 + 2 + 4 + 8,
	'═': 80,
	'║': 160,
	'╒': 72,
	'╓': 132,
	'╔': 192,
	'╕': 24,
	'╖': 129,
	'╗': 144,
	'╘': 66,
	'╙': 36,
	'╚': 96,
	'╛': 18,
	'╜': 33,
	'╝': 48,
	'╞': 74,
	'╟': 164,
	'╠': 224,
	'╡': 26,
	'╢': 161,
	'╣': 176,
	'╤': 88,
	'╥': 133,
	'╦': 208,
	'╧': 82,
	'╨': 37,
	'╩': 112,
	'╪': 90,
	'╫': 165,
	'╬': 240,
}
