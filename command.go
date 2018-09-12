package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var placeRegex, _ = regexp.Compile("(?i)PLACE\\s+(\\d+),\\s*(\\d+),\\s*([a-zA-Z]+)")

func ParseCommand(cmd string) func(s *Simulation) {
	switch cmd {
	case "MOVE":
		return func(s *Simulation) {
			s.Move()
		}
	case "LEFT":
		return func(s *Simulation) {
			s.Left()
		}
	case "RIGHT":
		return func(s *Simulation) {
			s.Right()
		}
	case "REPORT":
		return func(s *Simulation) {
			println(s.Report())
		}
	default:
		return func(s *Simulation) {
			matches := placeRegex.FindStringSubmatch(cmd)
			if matches != nil {
				x, y, f := matches[1], matches[2], matches[3]
				xx, err := strconv.ParseUint(x, 10, 64)
				Check(err, "Invalid x-coordinate for PLACE command: %v", x)
				yy, err := strconv.ParseUint(y, 10, 64)
				Check(err, "Invalid y-coordinate for PLACE command: %v", y)
				ff, err := NewDirection(f)
				Check(err, "Invalid direction for PLACE command: %v", f)
				s.Place(Robot{Position{uint(xx), uint(yy), *ff}})
			} else {
				fmt.Printf("Bad command: %v\n", cmd)
			}
		}
	}
}
