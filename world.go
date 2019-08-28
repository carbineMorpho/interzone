package main

import "github.com/nsf/termbox-go"

type tile struct {
	ch rune
	fg termbox.Attribute
	bg termbox.Attribute
}

type world struct {
	size pos
	terrain []*tile
	player entity
	creature []entity
}

// put a tile on the map. overwrites and handles bounds
func (w *world) Put(p pos, t tile) {

	if (p.X >= 0 && p.X < w.size.X) && (p.Y >= 0 && p.Y < w.size.Y) {
		i := p.Y*w.size.X + p.X
		w.terrain[i] = &t
	}
}

// get a tile from the map. handles bounds
func (w *world) Get(p pos) tile {

	if (p.X >= 0 && p.X < w.size.X) && (p.Y >= 0 && p.Y < w.size.Y) {
		i := p.Y*w.size.X + p.X
		return *w.terrain[i]
	}
	return tile{'#', termbox.ColorWhite, termbox.ColorWhite}
}

// checks for entity at given point
func (w *world) CollideCheck(p pos) (*entity, bool){
	
	for i := range w.creature {
		e := &w.creature[i]
		if p.X == e.p.X && p.Y == e.p.Y {
			return e, true
		}
	}
	return nil, false
}

// places an array of points onto terrain
func (w *world) Build(point []pos, color termbox.Attribute) {
	var block tile
	block.ch = '#'
	block.fg = color
	block.bg = color

	for i := range point {
		w.Put(point[i], block)
	}
}

// initialise world terrain
func (w *world) Init(width, height int) {
	air := tile{' ', termbox.ColorBlack, termbox.ColorBlack}

	w.size.X = width
	w.size.Y = height
	
	for i := 0; i < width*height; i++ {
		w.terrain = append(w.terrain, &air)
	}
	
}
