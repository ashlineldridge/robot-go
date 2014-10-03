package main

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

var directions = map[Direction]string{
	North: "NORTH",
	East:  "EAST",
	South: "SOUTH",
	West:  "WEST",
}

func (d Direction) String() string {
	return directions[d]
}

func (d Direction) IsValid() bool {
	_, ok := directions[d]
	return ok
}
