package main

import "fmt"

type Position struct {
	X, Y   uint
	Facing Direction
}

func (p *Position) Next() Position {
	next := *p
	switch p.Facing {
	case North:
		next.Y++
	case East:
		next.X++
	case South:
		next.Y--
	case West:
		next.X--
	}
	return next
}

func (p *Position) String() string {
	return fmt.Sprintf("%v, %v, %v", p.X, p.Y, p.Facing)
}
