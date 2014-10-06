package main

import (
	"github.com/ashlineldridge/robot-go/robot"
	"github.com/ashlineldridge/robot-go/table"
)

const (
	tableWidth  int = 5
	tableHeight int = 5
)

func main() {
	t := table.Table{Width: tableWidth, Height: tableHeight}
	r := robot.NewRobot(t)
	r.Report()
	r.Place(&table.Position{X: 1, Y: 1, D: table.North})
	r.Report()
	r.Move()
	r.Report()
}
