package main

import "fmt"

type Position struct {
	x int
	y int
	d Direction
}

func (p *Position) String() string {
	return fmt.Sprintf("%d,%d,%s", p.x, p.y, p.d)
}

func (p *Position) Rotate(delta int) Position {
	return Position{x: p.x, y: p.y, d: p.d.Rotate(delta)}
}
