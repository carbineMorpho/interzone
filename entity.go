package main

import(
	"github.com/nsf/termbox-go"
	"strings"
)

type entity struct{
        ch rune
	fg termbox.Attribute
	bg termbox.Attribute
        p pos
        hp int
        atk int
	goal pos
	tags string // "space:solid:entity:player:entity"
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
	switch{
        case prs == false:
                e.p.x = x
                e.p.y = y
	case collide.Is("solid"):
		break
	case collide.Is("entity"):
		e.Attack(collide)
        }

}

func (e *entity) AI(collision map[pos]*entity) {
// simple AI. will chase player if in sight

	// create virtual line
	point := line(e.p, e.goal)

	// check points along the line for sight blocking
	sight := true
	for i := 1; i < len(point)-1; i++ {
		if _, prs := collision[point[i]]; prs {
			sight = false
		}
	}

	// calculate direction to move
	vector := pos{point[1].x-e.p.x, point[1].y-e.p.y}
	if sight == true {
		e.Move(vector, collision)
	}
}

func (e *entity) Is(property string) bool{
	tags := strings.Split(e.tags, ":")
	for i := range(tags) {
		if tags[i] == property {
			return true
		}
	}
	return false
}

func playerCheck(e entity) bool{
	if e.Is("player") == false {
		return false
	}
	return true
}

func playerInit(p pos) entity{
	var player entity
        player.ch = '@'
	player.fg = termbox.ColorWhite
	player.bg = termbox.ColorBlack
        player.p = p
        player.hp = 10
        player.atk = 2
        player.goal = pos{0,0}
        player.tags = "entity:player"
	return player
}

func xenoInit(p pos) entity{
	var xeno entity
	xeno.ch = 'x'
	xeno.fg = termbox.ColorGreen
	xeno.bg = termbox.ColorBlack
	xeno.p = p
	xeno.hp = 5
	xeno.atk = 2
	xeno.goal = pos{0,0}
	xeno.tags = "entity:enemy"
	return xeno
}

