package main

import "github.com/nsf/termbox-go"

const SCREEND = 28

// scrolling camera
func (w *world) Monitor() {

	var camera pos
	camera.X = w.player.p.X - (SCREEND/2)
	camera.Y = w.player.p.Y - (SCREEND/2)

	for y := 0; y < SCREEND; y++ {
		for x := 0; x < SCREEND; x++ {
			t := w.Get(pos{camera.X +x, camera.Y +y})
			termbox.SetCell(x, y, t.ch, t.fg, t.bg)
		}
	}

	for _, e := range w.creature {
		if (e.p.X > camera.X && e.p.X < camera.X + SCREEND) && (e.p.Y > camera.Y && e.p.Y < camera.Y + SCREEND){
			termbox.SetCell(e.p.X - camera.X, e.p.Y - camera.Y, e.ch, termbox.ColorWhite, termbox.ColorBlack)
		}
	}

	termbox.SetCell(w.player.p.X - camera.X, w.player.p.Y - camera.Y, w.player.ch, termbox.ColorWhite, termbox.ColorBlack)

	termbox.Flush()
}

// remove dead entities from memory
func hpCheck(e []entity) []entity {
	for i := 0; i < len(e); i++ {
		if deathCheck(e[i]) == true {
			e[i] = e[len(e)-1]
			e = e[:len(e)-1]
		}
	}
	return e
}

func deathCheck(e entity) bool {
	if e.hp <= 0 {
		return true
	}
	return false
}

func errorCheck(err error) {

	if err != nil {
		panic(err)
	}
}

// moves towards target if in sight
func (e *entity) AI(target pos, w world) {
	point := line(e.p, target)
	see := true	

	for i := 1; i < len(point); i++ {
		if w.Get(point[i]).ch == '#' {
			see = false
		}
	}

	if see && len(point) > 2{
		log("%+v", point)
		e.Move(point[1], w)
	}
}

func main(){

	// termbox init
	err := termbox.Init()
	errorCheck(err)
	defer termbox.Close()

	// world init
	var w world
	w.Init(100,100)

	// entity init
	w.player = playerInit(8,8)
	w.creature = append(w.creature, xeno(2,2))
	w.creature = append(w.creature, xeno(8,3))

	// main game loop
	running := true
	for running {
		w.creature = hpCheck(w.creature)
		running = !deathCheck(w.player)
		w.Monitor()
		
		switch termbox.PollEvent().Ch {
		case 'w':
			w.player.Move(w.player.p.UP(), w)		
		case 'a':
			w.player.Move(w.player.p.LEFT(), w)
		case 's':
			w.player.Move(w.player.p.DOWN(), w)
		case 'd':
			w.player.Move(w.player.p.RIGHT(), w)
		case 'l':
			w.Build(rorschach(w.player.p))
		case '0':
			running = false
		}
		
		for i := 0; i < len(w.creature); i++ {
			w.creature[i].AI(w.player.p, w)
		}
	// end of game loop
	}
}
