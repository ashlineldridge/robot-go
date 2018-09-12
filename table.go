package main

type Table struct {
	Width  uint
	Height uint
}

func (t Table) IsValid(p Position) bool {
	return p.X < t.Width && p.Y < t.Height
}
