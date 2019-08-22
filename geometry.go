package main

import(
	"time"
	"math/rand"
)

type pos struct{
	x int
	y int
}

func line(start, end pos) (point []pos){
// bresenham line drawing

	// prepare for line gradients > 1
        steep := abs(end.y - start.y) > abs(end.x - start.x)
        if steep {
                start.x, start.y = start.y, start.x
                end.x, end.y = end.y, end.x
        }

	// prepare for lines that travel left
        reverse := false
        if start.x > end.x {
                start.x, end.x = end.x, start.x
                start.y, end.y = end.y, start.y
                reverse = true
        }

	// prepare for lines that travel down
        neg := -1
        if start.y < end.y {
                neg = 1
        }

        deltaX := end.x - start.x
        deltaY := abs(end.y - start.y)
        err := deltaX / 2

        y := start.y
        for x := start.x; x < end.x+1; x++ {
                if steep {
                        point = append(point, pos{y, x})
                } else {
                        point = append(point, pos{x, y})
                }
                err -= deltaY
                if err < 0 {
                        y += neg
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

func linePut(start, end pos, e entity, terrain map[pos]entity) map[pos]entity {
	point := line(start, end)
	for i := range(point) {
		e.p = point[i]
		terrain[point[i]] = e
	}
	return terrain
}

func square(start, end pos) (point []pos){
	for y := start.y; y < end.y; y++ {
		for x := start.x; x < end.x; x++ {
			point = append(point, pos{x,y})
		}
	}
	return
}

func squarePut(start, end pos, e entity, terrain map[pos]entity) map[pos]entity {
	point := square(start, end)
	for i := range(point) {
		e.p = point[i]
		terrain[e.p] = e
	}
	return terrain
}

func basicReflect(terrain []entity) []entity {

	// find point of reflection, mean of X coords	
	x := 0
	for _, e := range terrain {
		x += e.p.x
	}
	x = x / len(terrain)

	// reflect the left side onto right
	nextTerrain := make([]entity, len(terrain))
	for _, left := range terrain {
		right := left
		if left.p.x < x {
			right.p.x += 2*(x - right.p.x)
			nextTerrain = append(nextTerrain, left)
			nextTerrain = append(nextTerrain, right)
		}
	}
	return nextTerrain
}

func basicGen(complexity int) []entity {

	var tile entity
	tile.ch = ' '
	tile.fg = colorRandom()
	tile.bg = colorRandom()
	tile.tags = "solid"

	terrainMap := make(map[pos]entity)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// drunk walk algorithm, weighted towards previous direction
	walker := pos{0,0}
	for complexity > 0 {
		switch r.Intn(4) {
			case 0: // up
				walker = pos{walker.x, walker.y -1}
			case 1: // left
                                walker = pos{walker.x -1, walker.y}
			case 2: // down
                                walker = pos{walker.x, walker.y +1}
			case 3: // right
                                walker = pos{walker.x +1, walker.y}
		}
		tile.p = walker
		terrainMap[walker] = tile
		complexity -= 1
	}

	// draw square rooms randomly in map
	for k := range terrainMap {
		if r.Intn(10) == 1 {
			c := terrainMap[k]
			size := r.Intn(4)+2
			terrainMap = squarePut(c.p, pos{c.p.x + size, c.p.y + size}, tile, terrainMap)
		}
	}

	// turn map into slice
	terrain := make([]entity, 0)
	for k := range terrainMap {
		terrain = append(terrain, terrainMap[k])
	}

	return terrain
}

func buildDemon(terrain []entity) []entity {

	terrain = basicGen(300)
	terrain = basicReflect(terrain)
	return terrain
}
