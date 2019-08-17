package main

import (
	"github.com/nsf/termbox-go"
)

const COLOR = termbox.ColorDefault
const WORLDD = 128
const SCREEND = 16

func hpCheck(e []entity) []entity{
// remove dead entities from memory

	for i := 0; i < len(e); i++ {
		if e[i].hp <= 0 {
			e[i] = e[len(e)-1]
			e = e[:len(e)-1] 
		}
	}
	return e
}

func terrainInit() []entity {
// load static entities into memory
	
	terrain := make([]entity, WORLDD*WORLDD)
	i := 0

	for y := 0; y < WORLDD; y++ {
		for x := 0; x < WORLDD; x++ {
			terrain[i] = entity{
				ch: '.',
				p: pos{x,y},
				tags: "space",
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
	terrain := drunkGen()
	
	// entity init
	creature := make([]entity, 1)

	creature[0] = playerInit(pos{32,32})
	
/*
	for i := 1; i < 2; i++ {
		creature = append(creature, xenoInit(pos{i+10,i}))
	}
*/

	// game loop
	running := true
	for running == true{

		creature = hpCheck(creature)	
		running = playerCheck(creature[0])

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
		case 'l':
			squarePut(pos{9,8}, pos{90,40}, collision)

		}
		for i := 1; i < len(creature); i++{
			collision = collisionGet(terrain, creature)
			creature[i].goal = creature[0].p
			creature[i].AI(collision)
		}
	}
	// end of game loop
}
