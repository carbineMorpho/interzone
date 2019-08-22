package main

import "github.com/nsf/termbox-go"

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

func collisionGet(e []entity, collision map[pos]*entity) (map[pos]*entity){
	for i := range(e) {
		collision[e[i].p] = &e[i]
	}
	return collision
}

func collisionGetAll(terrain []entity, creature []entity) (map[pos]*entity){
// data structure for potential collisions

	collision := make(map[pos]*entity)
	collision = collisionGet(terrain, collision)
	collision = collisionGet(creature, collision)
	return collision
}

func main() {
	
	// termbox init
	err := termbox.Init()
	errorCheck(err)
	defer termbox.Close()

	// world init
	terrain := make([]entity, 1)
	terrain = buildDemon(terrain)
	terrain = basicReflect(terrain)
	
	// entity init
	creature := make([]entity, 1)

	creature[0] = playerInit(pos{0,0})
	
	for i := 1; i < 2; i++ {
		creature = append(creature, xenoInit(pos{i+10,i}))
	}

	// game loop
	running := true
	for running == true{

		creature = hpCheck(creature)	
		running = playerCheck(creature[0])

		collision := collisionGetAll(terrain, creature)
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
			collision = collisionGetAll(terrain, creature)
			creature[i].goal = creature[0].p
			creature[i].AI(collision)
		}
	}
	// end of game loop
}
