package main

import (
	"github.com/errnoh/termbox/helper"
	"github.com/errnoh/termbox/panel"
	termbox "github.com/nsf/termbox-go"
	"strings"
	"time"
)

func main() {
	termbox.Init()
	defer termbox.Close()
	m := panel.MainScreen()

	for _, s := range strings.Split(example, "\n") {
		helper.ScrollWriteUpwards(m, s)
		termbox.Flush()
		time.Sleep(time.Millisecond * 500)
        }
	for _, s := range strings.Split(example, "\n") {
		helper.ScrollWrite(m, s)
		termbox.Flush()
		time.Sleep(time.Millisecond * 500)
	}
	helper.ScrollWrite(m, "And that's\neverything\nfor now.")
	termbox.Flush()
	time.Sleep(time.Millisecond * 2000)
}

const example = `Sat Dec 17 00:01:04 1994 204.94.123.8 128.91.201.34    GET /focus/bsvs2a.mpg   HTTP/1.0
Sat Dec 17 00:01:36 1994 204.94.123.8 199.2.134.6   GET /focus/catalog.html HTTP/1.0
Sat Dec 17 00:01:45 1994 204.94.123.8 199.2.134.6   GET /focus/focLogo4.gif HTTP/1.0
Sat Dec 17 00:02:32 1994 204.94.123.8 199.2.134.6   GET /focus/foc_pt1.html HTTP/1.0
Sat Dec 17 00:02:42 1994 204.94.123.8 199.2.134.6   GET /focus/bst1sm.gif   HTTP/1.0
Sat Dec 17 00:02:44 1994 204.94.123.8 199.2.134.6   GET /focus/alp3sm.gif   HTTP/1.0
Sat Dec 17 00:02:53 1994 204.94.123.8 199.2.134.6   GET /focus/bo3sm.gif    HTTP/1.0
Sat Dec 17 00:03:00 1994 204.94.123.8 199.2.134.6   GET /focus/sos1sm.gif   HTTP/1.0
Sat Dec 17 00:03:06 1994 204.94.123.8 199.2.134.6   GET /focus/spl2sm.gif   HTTP/1.0
Sat Dec 17 00:03:10 1994 204.94.123.8 128.91.201.34 GET /focus/bst1.jpg HTTP/1.0
Sat Dec 17 00:04:00 1994 204.94.123.8 199.2.134.6   GET /focus/foc_pt1.html HTTP/1.0
Sat Dec 17 00:04:33 1994 204.94.123.8 128.91.201.34 GET /focus/alp3.jpg HTTP/1.0
Sat Dec 17 00:04:38 1994 204.94.123.8 199.2.134.6   GET /focus/foc_pt1.html HTTP/1.0
Sat Dec 17 00:05:04 1994 204.94.123.8 128.91.201.34 GET /focus/spl2.jpg HTTP/1.0
Sat Dec 17 00:05:26 1994 204.94.123.8 199.2.134.6   GET /focus/bsvs2a.mpg   HTTP/1.0
Sat Dec 17 00:05:43 1994 204.94.123.8 128.2.22.118  GET /focus/catalog.html HTTP/1.0
Sat Dec 17 00:05:44 1994 204.94.123.8 128.2.22.118  GET /focus/focLogo4.gif HTTP/1.0
Sat Dec 17 00:05:56 1994 204.94.123.8 128.2.22.118  GET /focus/foc_pt1.html HTTP/1.0
Sat Dec 17 00:08:58 1994 204.94.123.8 128.2.22.118  GET /focus/alp3.jpg HTTP/1.0`
