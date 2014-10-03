package main

type Table struct {
	width  int
	height int
}

func (t Table) IsValid(p *Position) bool {
	return p.x >= 0 && p.x < t.width && p.y >= 0 && p.y < t.height && p.d.IsValid()
}
