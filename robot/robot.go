package robot

import (
	"fmt"

	"github.com/ashlineldridge/robot-go/table"
)

type Robot struct {
	position *table.Position
	table    table.Table
}

func NewRobot(table table.Table) Robot {
	return Robot{position: nil, table: table}
}

func (r *Robot) Report() {
	if r.position != nil {
		fmt.Println(r.position)
	}
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
