package main

type entity struct {
	p pos
	ch rune
	hp int
}

func (e *entity) Attack(victim *entity) {
	log("%c attacked %c for %d", e.ch, victim.ch, 1)
	victim.hp -= 1
}

func (e *entity) PickUp(w world) {
	t, prs := w.PropCheck(e.p)
	if prs {
		t.PosSet(e)
	}
}

func (e *entity) Move(p pos, w world) {

	tCollide := w.Get(p)
	collide := w.CollideCheck(p)
	switch {
		case collide != nil:
			e.Attack(collide)
		case tCollide.ch == '#':
			break
		default:
			e.p = p
	}
}

// moves towards goal if in sight
func (e *entity) Hunt(w world) {
	point := line(e.p, w.player.p)
	if w.See(e.p, w.player.p) && len(point) > 2{
		e.Move(point[1], w)
	}
}


func playerInit(x, y int) entity{
	var player entity
	player.ch = '@'
	player.p.X = x
	player.p.Y = y
	player.hp = 10

	return player
}

func xeno(x, y int) entity{
	var xeno entity
	xeno.ch = 'x'
	xeno.p.X = x
	xeno.p.Y = y
	xeno.hp = 2
	
	return xeno
}