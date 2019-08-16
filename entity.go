package main

type entity struct{
        ch rune
        p pos
        hp int
        atk int
	goal pos
	f string // type descriptor [space, solid, entity]
}

func (e *entity) Attack(victim *entity ) {
// basic attack function, all ranges, no missile calculation

        victim.hp -= e.atk
	log("%c attacked %c for %d", e.ch, victim.ch, e.atk)
}

func (e *entity) Move(vector pos, collision map[pos]*entity) {
// check for collisions, decide responce, perform move

        x := e.p.x + vector.x
        y := e.p.y + vector.y

        collide, prs := collision[pos{x,y}]
        if prs != false {
                switch collide.f {
                case "solid":
			break
		case "entity":
			e.Attack(collide)
		case "space":
                        e.p.x = x
                        e.p.y = y
                }
        }

}

func (e *entity) AI(collision map[pos]*entity) {
// simple AI. will chase player if in sight

	// create virtual line
	point := line(e.p, e.goal)

	// check points along the line for sight blocking
	sight := true
	for i := 1; i < len(point)-1; i++ {
		if collision[point[i]].f == "solid" {
			sight = false
		}
	}

	// calculate direction to move
	vector := pos{point[1].x-e.p.x, point[1].y-e.p.y}
	if sight == true {
		e.Move(vector, collision)
	}
}
