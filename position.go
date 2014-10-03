package main

import "fmt"

type CardinalDirection int

const (
	North CardinalDirection = iota
	East
	South
	West
)

type Position struct {
	x int
	y int
	f CardinalDirection
}

func (c CardinalDirection) String() string {
	switch c {
	case North:
		return "NORTH"
	case East:
		return "EAST"
	case South:
		return "SOUTH"
	case West:
		return "WEST"
	}
	return ""
}

func (p Position) String() string {
	return fmt.Sprintf("%d,%d,%s", p.x, p.y, p.f)
}
