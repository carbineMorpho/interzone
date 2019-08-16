package main

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

func linePut(start, end pos, collision map[pos]*entity) {
	point := line(start, end)
	for i := range(point) {
		collision[(point[i])].ch = '#'
		collision[(point[i])].tags = "solid"
	}
}
