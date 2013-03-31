// Copyright 2013 errnoh. All rights reserved.
// Use of this source code is governed by
// MIT License that can be found in the LICENSE file.

package helper

import (
	"github.com/errnoh/termbox/panel"
	termbox "github.com/nsf/termbox-go"
	"image"
	"strings"
)

// Function names will probably change or they might be combined into a single function.

func ScrollWrite(panel *panel.Buffered, s string) {
	lines := strings.Split(s, "\n")
	buf := panel.Buffer()
	r := panel.Bounds()

	var (
		length     = r.Dx() * len(lines)
		START      = 0
		STARTLINES = START + length
		END        = len(buf)
		ENDLINES   = END - length
	)

	src := buf[:ENDLINES]
	target := buf[STARTLINES:]

	width := r.Dx()
	for i := len(target) - width; i >= 0; i -= width {
		copy(target[i:i+width], src[i:i+width])
	}

	for i := 0; i < len(lines); i++ {
		clearRow(panel, i)
	}

	panel.Write([]byte(s))
}

// XXX: Currently only usable on main screen. Unbuffered needs some love later.
func ScrollWriteUpwards(panel *panel.Buffered, s string) {
	lines := strings.Split(s, "\n")
	buf := panel.Buffer()
	r := panel.Bounds()

	var (
		length     = r.Dx() * len(lines)
		START      = 0
		STARTLINES = START + length
		END        = len(buf)
		ENDLINES   = END - length
	)

	src := buf[STARTLINES:]
	target := buf[:ENDLINES]

	width := r.Dx()
	for i := 0; i < len(target); i += width {
		copy(target[i:i+width], src[i:i+width])
	}

	for i := 0; i < len(lines); i++ {
		clearRow(panel, r.Dy()-1-i)
	}

	targetr := panel.Area(image.Rect(0, r.Dy()-len(lines), width, r.Dy()))
	targetr.Write([]byte(s))
}

func clearRow(panel *panel.Buffered, row int) {
	// TODO: Possibly change to copying []Cell{' ', 0, 0...} instead, speed vs allocating more memory.
	//       Consider using unexported big enough []Cell ready just for fast wiping.
	r := panel.Bounds()
	if row < 0 || row >= r.Dy() {
		return
	}

	buf := panel.Buffer()
	for i := row * r.Dx(); i < row*r.Dx()+r.Dx(); i++ {
		buf[i] = termbox.Cell{' ', 0, 0}
	}
}
