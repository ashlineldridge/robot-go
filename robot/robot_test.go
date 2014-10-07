package robot

import (
	"os"
	"testing"

	"github.com/ashlineldridge/robot-go/table"

	"github.com/stretchr/testify/assert"
)

func TestRobotPlaceValid(t *testing.T) {
	r := NewRobot(table.Table{Width: 5, Height: 5}, os.Stdout)

	bottomLeft := table.Position{0, 0, table.North}
	assert.Equal(t, *r.Place(&bottomLeft), bottomLeft, "{0,0,North} should be valid.")

	bottomRight := table.Position{4, 0, table.North}
	assert.Equal(t, *r.Place(&bottomRight), bottomRight, "{4,0,North} should be valid.")

	topRight := table.Position{4, 4, table.North}
	assert.Equal(t, *r.Place(&topRight), topRight, "{4,4,North} should be valid.")

	topLeft := table.Position{0, 4, table.North}
	assert.Equal(t, *r.Place(&topLeft), topLeft, "{0,4,North} should be valid.")
}

func TestRobotPlaceInvalid(t *testing.T) {
	r := NewRobot(table.Table{Width: 5, Height: 5}, os.Stdout)

	outOfBounds := table.Position{-1, -1, table.North}
	assert.Nil(t, r.Place(&outOfBounds), "{-1,-1,North} should not be valid.")

	outOfBounds = table.Position{0, 5, table.North}
	assert.Nil(t, r.Place(&outOfBounds), "{0,5,North} should not be valid.")

	outOfBounds = table.Position{5, 0, table.North}
	assert.Nil(t, r.Place(&outOfBounds), "{5,0,North} should not be valid.")

	outOfBounds = table.Position{0, 0, table.Invalid}
	assert.Nil(t, r.Place(&outOfBounds), "{0,0,Invalid} should not be valid.")
}

func TestRobotRotation(t *testing.T) {
	r := NewRobot(table.Table{Width: 5, Height: 5}, os.Stdout)
	r.Place(&table.Position{0, 0, table.North})

	assert.Equal(t, r.Right().D, table.East, "Robot should be facing East")
	assert.Equal(t, r.Right().D, table.South, "Robot should be facing South")
	assert.Equal(t, r.Right().D, table.West, "Robot should be facing West")
	assert.Equal(t, r.Right().D, table.North, "Robot should be facing North")

	assert.Equal(t, r.Left().D, table.West, "Robot should be facing North")
	assert.Equal(t, r.Left().D, table.South, "Robot should be facing South")
	assert.Equal(t, r.Left().D, table.East, "Robot should be facing East")
	assert.Equal(t, r.Left().D, table.North, "Robot should be facing North")
}

func TestRobotMove(t *testing.T) {
	r := NewRobot(table.Table{Width: 5, Height: 5}, os.Stdout)
	r.Place(&table.Position{0, 0, table.North})

	// Move the robot to the top left corner
	r.Move()
	r.Move()
	r.Move()
	p := r.Move()
	assert.Equal(t, *p, table.Position{0, 4, table.North})

	// Check that the robot can't be moved off the top edge
	p = r.Move()
	assert.Equal(t, *p, table.Position{0, 4, table.North})

	// Move the robot to the top right corner
	r.Right()
	r.Move()
	r.Move()
	r.Move()
	p = r.Move()
	assert.Equal(t, *p, table.Position{4, 4, table.East})

	// Check that the robot can't be moved off the right edge
	p = r.Move()
	assert.Equal(t, *p, table.Position{4, 4, table.East})
}
