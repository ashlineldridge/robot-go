package main

import "fmt"

type Robot struct {
	position *Position
	table    *Table
}

func NewRobot(table *Table) Robot {
	return Robot{position: nil, table: table}
}

func (r Robot) Report() {
	if r.position != nil {
		fmt.Println(r.position)
	}
}
