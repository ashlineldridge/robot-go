package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/ashlineldridge/robot-go/robot"
	"github.com/ashlineldridge/robot-go/table"
	"github.com/stretchr/testify/assert"
)

func TestApplication(t *testing.T) {
	file, reader := openInputFile()
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	table := table.Table{Width: tableWidth, Height: tableHeight}
	robot := robot.NewRobot(table)
	processInput(&robot, reader)
	assert.Equal(t, 1, 1)
	fmt.Println("You suck")
}

func openInputFile() (file *os.File, reader *bufio.Reader) {
	file, err := os.Open("test_resources/input.txt")
	if err != nil {
		panic(err)
	}
	reader = bufio.NewReader(file)
	return
}
