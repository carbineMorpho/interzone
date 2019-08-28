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

func (e *entity) Move(p pos, w world) {

	tCollide := w.Get(p)
	collide, prs := w.CollideCheck(p)
	switch {
		case prs:
			e.Attack(collide)
		case tCollide.ch == '#':
			break
		default:
			e.p = p
	}
}

func playerInit(x, y int) entity{
	var player entity
	player.p.X = x
	player.p.Y = y
	player.ch = '@'
	player.hp = 10

	return player
}

func xeno(x, y int) entity{
	var xeno entity
	xeno.p.X = x
	xeno.p.Y = y
	xeno.ch = 'x'
	xeno.hp = 2
	
	return xeno
}
