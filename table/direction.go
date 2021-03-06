package table

type Direction int

const (
	North Direction = iota
	East
	South
	West
	Invalid
)

var directions = map[Direction]string{
	North: "NORTH",
	East:  "EAST",
	South: "SOUTH",
	West:  "WEST",
}

func NewDirection(s string) Direction {
	for k, v := range directions {
		if v == s {
			return k
		}
	}
	return Invalid
}

func (d Direction) String() string {
	return directions[d]
}

func (d Direction) IsValid() bool {
	_, ok := directions[d]
	return ok
}

func (d Direction) Rotate(delta int) Direction {
	const count int = 4
	return Direction((int(d) + delta + count) % count)
}
