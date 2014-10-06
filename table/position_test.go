package table

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPositionStringification(t *testing.T) {
	p := Position{X: 1, Y: 1, D: North}
	assert.Equal(t, "1,1,NORTH", p.String(), "Position should be represented as '1,1,NORTH'.")
}

func TestPositionRotation(t *testing.T) {
	p := Position{X: 1, Y: 1, D: North}
	rotated := p.Rotate(1)
	assert.Equal(t, Position{X: 1, Y: 1, D: East}, rotated, "Rotated position should be '1,1,EAST'.")
}
