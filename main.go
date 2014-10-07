package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/ashlineldridge/robot-go/robot"
	"github.com/ashlineldridge/robot-go/table"
)

const (
	tableWidth  int = 5
	tableHeight int = 5
)

func main() {
	table := table.Table{Width: tableWidth, Height: tableHeight}
	robot := robot.NewRobot(table)
	processInput(&robot, os.Stdin)
}

func processInput(robot *robot.Robot, reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		executeCommand(robot, scanner.Text())
	}
}

func executeCommand(robot *robot.Robot, command string) {
	var instruction, d string
	var x, y int
	fmt.Sscanf(command, "%s %d,%d,%s\n", &instruction, &x, &y, &d)
	switch instruction {
	case "PLACE":
		robot.Place(&table.Position{X: x, Y: y, D: table.NewDirection(d)})
	case "MOVE":
		robot.Move()
	case "LEFT":
		robot.Left()
	case "RIGHT":
		robot.Right()
	case "REPORT":
		robot.Report()
	}
}
