package main

import (
	"fmt"
)

type Target struct {
	x1, y1, x2, y2 int
}

type Point struct {
	x, y int
}

type Probe struct {
	position *Point
	velocity *Point
}

func (p *Probe) Move() {
	p.position.x += p.velocity.x
	p.position.y += p.velocity.y

	if p.velocity.x > 0 {
		p.velocity.x--
	} else if p.velocity.x < 0 {
		p.velocity.x++
	}
	p.velocity.y--
}

func (t *Target) Contains(p Point) bool {
	return p.x >= t.x1 && p.x <= t.x2 && p.y >= t.y2 && p.y <= t.y1
}

func (p *Probe) Simulate(t Target) (found bool, maxY int) {
	maxY = 0
	for p.position.x <= t.x2 && p.position.y >= t.y2 {
		p.Move()
		if p.position.y > maxY {
			maxY = p.position.y
		}
		if t.Contains(*p.position) {
			return true, maxY
		}
	}
	return false, maxY
}

func main() {
	res := make(map[Point]bool)
	x1, y1, x2, y2 := 88, -103, 125, -157

	t := Target{x1, y1, x2, y2}
	fmt.Println(t)
	var y int
	y = 0
	for {
		for x := 1; x <= x2; x++ {
			probe := Probe{&Point{0, 0}, &Point{x, y}}
			found, _ := probe.Simulate(t)
			if found {
				res[Point{x, y}] = true
			}
			probe = Probe{&Point{0, 0}, &Point{x, -y}}
			found, _ = probe.Simulate(t)
			if found {
				res[Point{x, -y}] = true
			}
		}
		fmt.Println(len(res))
		y++
	}
}
