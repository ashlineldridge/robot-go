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

func (r *Robot) Place(p table.Position) {
	if r.table.IsValid(p) {
		if r.position == nil {
			r.position = new(table.Position)
		}
		*r.position = p
	}
}

func (r *Robot) Left() {
	if r.position != nil {
		*r.position = r.position.Rotate(-1)
	}
}

func (r *Robot) Right() {
	if r.position != nil {
		*r.position = r.position.Rotate(1)
	}
}
