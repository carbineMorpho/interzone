package main

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

// return a random direction
func (p pos) Random() pos {
	r := seedInit()
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
	return p
}


// return absolute value
func abs(n int) int{
	if n < 0 {
		return -n
	}
	return n
}

// linear interpolation
func lerp(start, end int, t float64) int{
	return int(float64(start) + t * float64(end-start))
}

// lerp between two positions
func lerpPos(start, end pos, t float64) pos{
	return pos{lerp(start.X, end.X, t), lerp(start.Y, end.Y, t)}
}

// create a line through lerp
func line(start, end pos) (point []pos){
	deltaX := abs(end.X - start.X)
	deltaY := abs(end.Y - start.Y)

	n := deltaX
	if deltaY > deltaX {
		n = deltaY
	}

	for i := 0; i <= n; i++ {
		t := float64(i)/float64(n)
		point = append(point, lerpPos(start, end, t))
	}
	return
}

// random walker algorithm
func drunkGen(p pos, size int) (point []pos) {
	for size > 0 {
		p = p.Random()
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
	for _, right :=  range point {
		left := right
		if right.X > X {
			left.X += 2*(X - right.X)
			new = append(new, left)
			new = append(new, right)
		}
	}
	return
}

// randomly generated rorschach test
func rorschach(p pos) (point []pos) {
	point = drunkGen(p, 200)
	point = reflect(point)
	return
}

func square(start, end pos) (point []pos) {
	for Y := start.Y; Y <= end.Y; Y++ {
		for X := start.X; X <= end.X; X++ {
			point = append(point, pos{X,Y})
		}
	}
	return
}

func dig(positive, negative []pos) (point []pos) {
	pointMap := make(map[pos]bool)

	for _, key := range positive {
		pointMap[key] = true
	}

	for _, key := range negative {
		pointMap[key] = false
	}

	for key, val := range pointMap {
		if val {
			point = append(point, key)
		}
	}
	return	
}

func room(start, end pos) (point []pos) {
	point = merge(point, line(start, pos{end.X, start.Y}))
	point = merge(point, line(start, pos{start.X, end.Y}))
	point = merge(point, line(end, pos{end.X, start.Y}))
	point = merge(point, line(end, pos{start.X, end.Y}))
	return point
}

func merge(a, b []pos) []pos {
	for _, val := range b {
		a = append(a, val)
	}
	return a
}


// creates a structure through binary space partitioning
func house(start, end pos) (point []pos) {

	// initial room
	positive := room(start, end)
	negative := []pos{pos{end.X-2, end.Y}}

	for i := 0; i < 3; i++ {
		width, length := (end.X - start.X), (end.Y - start.Y)
		if width > length {
			// horizontal split
			end.X -= width / 2
			positive = merge(positive, room(start, end))
			negative = append(negative, pos{end.X, end.Y-2})
		} else {
			// vertical split
			end.Y -= length / 2
			positive = merge(positive, room(start, end))
		}
			negative = append(negative, pos{end.X-2, end.Y})
	}
	return dig(positive, negative)
}

