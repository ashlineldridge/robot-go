package main

import (
	"bufio"
	"bytes"
	"os"
	"testing"

	"github.com/ashlineldridge/robot-go/robot"
	"github.com/ashlineldridge/robot-go/table"
	"github.com/stretchr/testify/assert"
)

func TestApplication(t *testing.T) {

	// Make the sample input file available as a Reader.
	inputFile, err := os.Open("test_resources/input.txt")
	if err != nil {
		panic(err)
	}
	inputReader := bufio.NewReader(inputFile)

	// Make sure the file handle is closed.
	defer func() {
		if err := inputFile.Close(); err != nil {
			panic(err)
		}
	}()

	// Read the expected output into a string.
	b, err := ioutil.readfile("test_resources/output.txt")
	if err != nil {
		panic(err)
	}
	expectedoutput := string(b)

	// Create the robot with a buffered writer so we can test its output.
	table := table.Table{Width: tableWidth, Height: tableHeight}
	var outputBuffer bytes.Buffer
	robot := robot.NewRobot(table, &outputBuffer)

	// Run the robot against the sample input file and check that its output is as expected.
	processInput(&robot, inputReader)
	assert.Equal(t, expectedOutput, outputBuffer.String(),
		"The robot's output was did not match the expected test_resources/output.txt file.")
}
