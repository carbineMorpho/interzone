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

func lerp(start, end int, t float64) int{
	return int(float64(start) + t * float64(end-start))
}

func lerpPos(start, end pos, t float64) pos{
	return pos{lerp(start.X, end.X, t), lerp(start.Y, end.Y, t)}
}

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

// bresenham line drawing
func line2(start, end pos) (point []pos) {
	
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
