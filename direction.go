package main

import "fmt"

type Direction uint

const (
	North Direction = iota
	East
	South
	West
)

var directions = map[Direction]string{
	North: "NORTH", East: "EAST", South: "SOUTH", West: "WEST"}

type UnknownDirectionError string

func (e UnknownDirectionError) Error() string { return fmt.Sprintf("Unknown Direction %v", e) }

func (d Direction) String() string {
	if s, ok := directions[d]; ok {
		return s
	}
	return "UNKNOWN"
}

func (d Direction) Rotate(units int) Direction {
	return Direction((int(d) + units) % (int(West) + 1))
}

func NewDirection(s string) (*Direction, error) {
	for k, v := range directions {
		if v == s {
			return &k, nil
		}
	}
	return nil, UnknownDirectionError(s)
}
