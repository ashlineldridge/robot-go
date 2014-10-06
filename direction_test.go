package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidDirections(t *testing.T) {
	assert.True(t, North.IsValid(), "North should be valid.")
	assert.True(t, East.IsValid(), "East should be valid.")
	assert.True(t, South.IsValid(), "South should be valid.")
	assert.True(t, West.IsValid(), "West should be valid.")
}

func TestInvalidDirections(t *testing.T) {
	assert.False(t, Invalid.IsValid(), "Invalid should not be valid.")
	assert.False(t, Direction(Invalid+1).IsValid(), "Values beyond the max enum should not be valid.")
	assert.False(t, Direction(North-1).IsValid(), "Values before the min enum should not be valid.")
}

func TestNewDirection(t *testing.T) {
	assert.Equal(t, North, NewDirection("NORTH"), "NORTH could not be correctly parsed.")
	assert.Equal(t, East, NewDirection("EAST"), "EAST could not be correctly parsed.")
	assert.Equal(t, South, NewDirection("SOUTH"), "SOUTH could not be correctly parsed.")
	assert.Equal(t, West, NewDirection("WEST"), "WEST could not be correctly parsed.")
	assert.Equal(t, Invalid, NewDirection("FOOBAR"), "Invalid direction was not detected.")
}

func TestDirectionRotation(t *testing.T) {
	assert.Equal(t, East, North.Rotate(1), "North rotated by 1 should be East")
	assert.Equal(t, South, East.Rotate(1), "East rotated by 1 should be South")
	assert.Equal(t, West, South.Rotate(1), "South rotated by 1 should be West")
	assert.Equal(t, North, West.Rotate(1), "West rotated by 1 should be North")

	assert.Equal(t, West, North.Rotate(-1), "North rotated by -1 should be East")
	assert.Equal(t, South, West.Rotate(-1), "East rotated by -1 should be South")
	assert.Equal(t, East, South.Rotate(-1), "South rotated by -1 should be West")
	assert.Equal(t, North, East.Rotate(-1), "West rotated by -1 should be North")

	assert.Equal(t, North, North.Rotate(4), "North rotated by 4 should be North")
	assert.Equal(t, North, North.Rotate(-4), "North rotated by -4 should be North")
}
