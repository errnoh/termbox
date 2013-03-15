// Copyright 2013 errnoh. All rights reserved.
// Use of this source code is governed by
// MIT License that can be found in the LICENSE file.

package main

import (
	"github.com/errnoh/termbox/panel"
	"github.com/errnoh/termbox/panel/border"
	termbox "github.com/nsf/termbox-go"
	"image"
)

var panels []panel.Panel

func newOcto() *panel.Buffered {
	p := panel.NewBuffered(image.Rect(10, 30, 50, 54)).Clear()
	p.Write(octo)
	panels = append(panels, p)
	return p
}

func main() {
	termbox.Init()
	defer termbox.Close()
	m := panel.MainScreen()
	var p panel.Panel = newOcto()
loop:
	for {
		m.Clear().Write(help)
		border.DrawAll(0, 0, panels...)
		if len(panels) > 0 {
			border.Add(2, 0, p)
		}
		termbox.Flush()

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				break loop
			case termbox.KeyTab:
				p = panels[0]
				Top(p)
			case termbox.KeyArrowDown:
				p.Move(0, 1)
			case termbox.KeyArrowUp:
				p.Move(0, -1)
			case termbox.KeyArrowLeft:
				p.Move(-1, 0)
			case termbox.KeyArrowRight:
				p.Move(1, 0)
			case termbox.KeyInsert:
				p = newOcto()
				Top(p)
			case termbox.KeyDelete:
				if len(panels) > 0 {
					panels = panels[:len(panels)-1]
					if len(panels) > 0 {
						p = panels[0]
						Top(p)
					}
				}
			}
		}
	}
}

func Top(p panel.Panel) {
	if panels[len(panels)-1] == p {
		return
	}

	var i int
	for i = IndexOf(p) + 1; i < len(panels); i++ {
		panels[i-1] = panels[i]
	}
	panels[i-1] = p
}

func IndexOf(p panel.Panel) int {
	for i := 0; i < len(panels); i++ {
		if panels[i] == p {
			return i
		}
	}
	return -1
}

// Octocat, base image from http://octodex.github.com/original/
var octo []byte = []byte(`                                        
         MMMM            .MMM           
         MMMMMMMMMMMMMMMMMMMM           
         MMMMMMMMMMMMMMMMMMMM           
         MMMMMMMMMMMMMMMMMMMMM          
        MMMMMMMMMMMMMMMMMMMMMMM         
       MMMMMMMMMMMMMMMMMMMMMMMM         
       MMMMM:::::::::::::::MMMM         
       MMMM::.7.:::::::.7.::MMM         
        MM~:~777~::::::777~:MMM         
   .  MMMMM:: . :::+::: . ::MM7MM ..    
         .MM::::::7:?::::::MM.          
            MMMM~::::::MMMM             
        MM      MMMMMMM                 
         M+    MMMMMMMMM                
          MMMMMMMMM MMMM                
               MMMM MMMM                
               MMMM MMMM                
            .~~MMMM~MMMM~~.             
         ~~~~MM:~MM~MM~:MM~~~~          
        ~~~~~~====~~~====~~~~~~         
         :~~~~~====~====~~~~~~          
             :~====~====~~              
                                        `)

var help []byte = []byte(`arrow keys: move
tab:        change focus
insert:     create panel
delete:     delete panel
esc:        quit`)
