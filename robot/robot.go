package robot

import (
	"fmt"
	"io"

	"github.com/ashlineldridge/robot-go/table"
)

type Robot struct {
	position *table.Position
	table    table.Table
	writer   io.Writer
}

func NewRobot(table table.Table, writer io.Writer) Robot {
	return Robot{position: nil, table: table, writer: writer}
}

func (r *Robot) Place(p *table.Position) *table.Position {
	if r.table.IsValid(p) {
		if r.position == nil {
			r.position = new(table.Position)
		}
		*r.position = *p
	}
	return r.position
}

func (r *Robot) Left() *table.Position {
	if r.position != nil {
		*r.position = r.position.Rotate(-1)
	}
	return r.position
}

func (r *Robot) Right() *table.Position {
	if r.position != nil {
		*r.position = r.position.Rotate(1)
	}
	return r.position
}

func (r *Robot) Move() *table.Position {
	if r.position != nil {
		xDelta, yDelta := 0, 0
		switch r.position.D {
		case table.North:
			yDelta = 1
		case table.East:
			xDelta = 1
		case table.South:
			yDelta = -1
		case table.West:
			xDelta = -1
		}
		moved := r.position.Add(xDelta, yDelta)
		if r.table.IsValid(&moved) {
			*r.position = moved
		}
	}
	return r.position
}

func (r *Robot) Report() {
	if r.position != nil {
		fmt.Fprintln(r.writer, r.position)
	}
}
