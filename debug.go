package main

import (
	"github.com/nsf/termbox-go"
	"os"
	"fmt"
)

const LOG = "/home/louis/golog"

// format then append string to log file
func log(msg string, a ...interface{}) {
	f, err := os.OpenFile(LOG, os.O_APPEND|os.O_WRONLY, 0600)
	errorCheck(err)
	_, err = f.WriteString(fmt.Sprintf(msg, a...) + "\n")
	errorCheck(err)
}

// scrolling camera
func (w *world) Monitor(msg string) {

	var camera pos
	camera.X = w.player.p.X - (SCREENX/2)
	camera.Y = w.player.p.Y - (SCREENY/2)

	// terrain
	for y := 0; y < SCREENY; y++ {
		for x := 0; x < SCREENX; x++ {
			t := w.Get(pos{camera.X +x, camera.Y +y})
			termbox.SetCell(x, y, t.ch, t.fg, t.bg)
		}
	}

	// props
	for _, b := range w.prop {
		switch b.i.(type){
		case pos:
			p := b.i.(pos)
			if (p.X > camera.X && p.X < camera.X + SCREENX) && (p.Y > camera.Y && p.Y < camera.Y + SCREENY){
				termbox.SetCell(p.X - camera.X, p.Y - camera.Y, rune(b.name[0]), termbox.ColorCyan, termbox.ColorBlack)
			}
		}

	}

	// creatures
	for _, e := range w.creature {
		if (e.p.X > camera.X && e.p.X < camera.X + SCREENX) && (e.p.Y > camera.Y && e.p.Y < camera.Y + SCREENY){
			termbox.SetCell(e.p.X - camera.X, e.p.Y - camera.Y, e.ch, termbox.ColorRed, termbox.ColorBlack)
		}
	}

	// player
	termbox.SetCell(w.player.p.X - camera.X, w.player.p.Y - camera.Y, w.player.ch, termbox.ColorWhite, termbox.ColorBlack)

	// info
	menu := SCREENY+1
	menu += menuInfo(menu, "INTERZONE")
	menu += menuInfo(menu, msg)

	termbox.Flush()
	termbox.Clear(termbox.ColorBlack, termbox.ColorBlack)
}

// menu info
func menuInfo(y int, msg string, a ...interface{}) int{
	msg = fmt.Sprintf(msg, a...)
	for x, ch := range msg {
		termbox.SetCell(x +1, y, ch, termbox.ColorWhite, termbox.ColorBlack)
	}
	return 1
}

// returns index for selected tool
func inventory(w world, owner *entity) (int, bool){
	for i := range w.prop {
		if w.prop[i].i == owner {
			menuInfo(i+1, "%c - %s", rune(i+97), w.prop[i].name) 
		}
	}
	termbox.Flush()
	i := int(termbox.PollEvent().Ch-97)
	log("%d", i)
	if (i >= 0 && i <= len(w.prop)) && w.prop[i].i == owner{	
		return i, true
	}
	return 0, false
}
