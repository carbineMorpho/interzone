package main

import (
	"time"
	"math/rand"
)

type pos struct {
	X int
	Y int
}

func (p *pos) UP() pos {
	return pos{p.X, p.Y-1}
}

func (p *pos) DOWN() pos {
	return pos{p.X, p.Y+1}
}

func (p *pos) LEFT() pos {
	return pos{p.X-1, p.Y}
}

func (p *pos) RIGHT() pos {
	return pos{p.X+1, p.Y}
}

// return absolute value
func abs(n int) int{
	if n < 0 {
		return (n - 2*n)
	} else {
		return n
	}
}

// bresenham line drawing
func line(start, end pos) (point []pos) {
	
	// prepare for steep lines
	steep := abs(end.Y - start.Y) > abs(end.X - start.X)
	if steep {
		start.X, start.Y = start.Y, start.X
		end.X, end.Y = end.Y, end.X
	}
	
	// prepare for lines travelling left
	reverse := false
	if start.X > end.X {
		start.X, end.X = end.X, start.X
		start.Y, end.Y = end.Y, start.Y
		reverse = true
	}

	// prepare for lines with negative gradient
	neg := -1
	if start.Y < end.Y {
		neg = 1
	}

	deltaX := start.Y
	deltaY := abs(end.Y - start.Y)
	err := deltaX / 2

	y := start.Y
	for x:= start.X; x <= end.X; x++ {
		if steep {
			point = append(point, pos{y, x})
		} else {
			point = append(point, pos{x, y})
		}
		err -= deltaY
		if err < 0 {
			y+= neg
			err += deltaX
		}
	}

	if reverse {
		for i, j := 0, len(point)-1; i < j; i, j = i+1, j-1 {
			point[i], point[j] = point[j], point[i]
		}
	}

	return
}

func drunkGen(p pos, size int) (point []pos) {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for size > 0 {
		switch r.Intn(4) {
		case 0:
			p = p.UP()
		case 1:
			p = p.DOWN()
		case 2:
			p = p.LEFT()
		case 3:
			p = p.RIGHT()
		}
		point = append(point, p)
		size -= 1
	}
	return
}

func reflect(point []pos) (new []pos) {
	
	// find a center
	X := 0
	for _, p := range point {
		X += p.X
	}
	X  = X / len(point)

	// reflect shape
	for _, left :=  range point {
		right := left
		if left.X <= X {
			right.X += 2*(X - right.X)
			new = append(new, left)
			new = append(new, right)
		}
	}
	return
}

func rorschach(p pos) (new []pos) {
	new = drunkGen(p, 200)
	new = reflect(new)
	return
}
