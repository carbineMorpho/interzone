package main

import(
	"github.com/nsf/termbox-go"
	"time"
	"math/rand"
)

const SCREEND = 48

func abs(n int) int{
// absolute value
        if n < 0 {
                return (n - 2*n)
        } else {
                return n
        }
}

func colorRandom() (color termbox.Attribute){
	r := rand.New(rand.NewSource(time.Now().UnixNano())) 
	switch r.Intn(4) {
	case 0:
		color = termbox.ColorRed
	case 1:
		color = termbox.ColorGreen
	case 2:
		color = termbox.ColorBlue
	case 3:
		color = termbox.ColorYellow
	}
	return
}

func monitor(camera pos, screen map[pos]*entity) {
// outputs the collision map

        // goto top left of screen
        camera.x = camera.x - (SCREEND/2)
        camera.y = camera.y - (SCREEND/2)

        // iterate over the screen cells
        for y := 0; y < SCREEND; y++ {
                for x := 0; x < SCREEND; x++ {
                        e, prs := screen[pos{camera.x+x, camera.y+y}]
                        if prs == true {
                                termbox.SetCell(x, y, e.ch, e.fg, e.bg)
                        } else {
                                termbox.SetCell(x, y, ' ', termbox.ColorBlack, termbox.ColorBlack)
                        }
                }
        }

        // display the backbuffer then clear it
        termbox.Flush()
        termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
}
