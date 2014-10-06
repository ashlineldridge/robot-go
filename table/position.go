package table

import "fmt"

type Position struct {
	X int
	Y int
	D Direction
}

func (p *Position) String() string {
	return fmt.Sprintf("%d,%d,%s", p.X, p.Y, p.D)
}

func (p *Position) Rotate(delta int) Position {
	return Position{X: p.X, Y: p.Y, D: p.D.Rotate(delta)}
}
