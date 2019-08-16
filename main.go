package main

import (
	"github.com/nsf/termbox-go"
)

const COLOR = termbox.ColorDefault
const WORLDD = 128
const SCREEND = 16

func hpCheck(e []entity) []entity {
// remove dead entities from memory

	for i := 0; i < len(e); i++ {
		if e[i].hp <= 0 {
			e[i] = e[len(e)-1]
			e = e[:len(e)-1] 
		}
	}
	return e
}

func terrainInit() map[pos]entity {
// load static entities into memory
	
	terrain := make(map[pos]entity, WORLDD*WORLDD)
	i := 0

	for y := 0; y < WORLDD; y++ {
		for x := 0; x < WORLDD; x++ {
			terrain[pos{x,y}] = entity{
				ch: '.',
				f: 'space',
			}
			i += 1
		}
	}

	return terrain
}


func collisionGet(terrain []entity, creature []entity) map[pos]*entity{
// data structure for potential collisions

	collision := make(map[pos]*entity)

	for i := range terrain{
		collision[terrain[i].p] = &terrain[i]
	}

	for i := range creature{
		if creature[i].hp > 0 {
			collision[creature[i].p] = &creature[i]
		}
	}
	return collision
}

func abs(n int) int{
// absolute value
	if n < 0 {
		return (n - 2*n)
	} else {
		return n
	}
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
			if prs != false {
				termbox.SetCell(x, y, e.ch, COLOR, COLOR)
			}
		}
	}
	
	// display the backbuffer then clear it
	termbox.Flush()
	termbox.Clear(COLOR, COLOR)
}

func main() {
	
	// termbox init
	err := termbox.Init()
	errorCheck(err)
	defer termbox.Close()

	// world init
	terrain := terrainInit()
	terrain = linePut(pos{10,20}, pos{60,90}, terrain)
	
	// entity init
	creature := make([]entity, 1)

	o := pos{9,9}
	player := entity{
		ch: '@',
		p: o,
		hp: 10,
		atk: 3,
		goal: o,
		f: "entity",
	}
	creature[0] = player
	
	xenomorph := entity{
		ch: 'x',
		p: o,
		hp: 5,
		atk: 1,
		goal: o,
		f: "entity",
	}
	for i := 1; i < 2; i++ {
		xenomorph.p = pos{i,i}
		creature = append(creature, xenomorph)
	}

	// game loop
	running := true
	for running == true{
		creature = hpCheck(creature)	
		collision := collisionGet(terrain, creature)
		monitor(creature[0].p, collision)

		input := termbox.PollEvent().Ch
		switch input {
		case '0':
			running = false
		case 'w':
			creature[0].Move(pos{0,-1}, collision)
		case 'a':
			creature[0].Move(pos{-1,0}, collision)
		case 's':
			creature[0].Move(pos{0, 1}, collision)
		case 'd':
			creature[0].Move(pos{1, 0}, collision)
		}
		for i := 1; i < len(creature); i++{
			collision = collisionGet(terrain, creature)
			creature[i].goal = creature[0].p
			creature[i].AI(collision)
		}
	}
	// end of game loop
}
