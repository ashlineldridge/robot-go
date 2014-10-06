package main

import "fmt"

type Robot struct {
	position *Position
	table    Table
}

func NewRobot(table Table) Robot {
	return Robot{position: nil, table: table}
}

func (r *Robot) Report() {
	if r.position != nil {
		fmt.Println(r.position)
	}
}

func (r *Robot) Place(p Position) {
	if r.table.IsValid(p) {
		if r.position == nil {
			r.position = new(Position)
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
