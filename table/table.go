package table

type Table struct {
	Width  int
	Height int
}

func (t *Table) IsValid(p *Position) bool {
	return p.X >= 0 && p.X < t.Width && p.Y >= 0 && p.Y < t.Height && p.D.IsValid()
}
