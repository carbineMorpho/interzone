package main

import (
	"github.com/nsf/termbox-go"
	"time"
	"math/rand"
)

const SCREENX = 35
const SCREENY = 25

// return a random seed 
func seedInit() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}

// remove dead entities from memory
func hpCheck(e []entity) []entity {
	for i := 0; i < len(e); i++ {
		if deathCheck(e[i]) == true {
			log("%c died", e[i].ch)
			e[i] = e[len(e)-1]
			e = e[:len(e)-1]
		}
	}
	return e
}

// check if an entity is dead
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

// returns position relative to player
func mouseGet() (pos, bool){
	if ev := termbox.PollEvent(); ev.Key == termbox.MouseLeft {
		return pos{ev.MouseX -17, ev.MouseY -12}, true
	}
	return pos{0,0}, false
}

func main(){

	// termbox init
	err := termbox.Init()
	errorCheck(err)
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)

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
		w.Monitor("controls: (g)et (j)rop (u)se wasd/move X/quit") 
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
			w.Build(line(w.player.p, pos{40,40}), colorRandom())
		case 'g':
			w.player.PickUp(w)
		case 'j':
			if i, prs := inventory(w, &w.player); prs{
				w.prop[i].PosSet(w.player.p)
			}
		case 'u':
			if i, prs := inventory(w, &w.player); prs{
				w.prop[i].f()
			}
		case 'X':
			running = false
		}
		
		for i := 0; i < len(w.creature); i++ {
			w.creature[i].Hunt(w)
		}
	// end of game loop
	}
}
