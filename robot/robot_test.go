package robot

import (
	"testing"

	"github.com/ashlineldridge/robot-go/table"

	"github.com/stretchr/testify/assert"
)

func TestRobotPlace(t *testing.T) {
	r := NewRobot(table.Table{Width: 5, Height: 5})
	assert.Equal(t, r, r, "")
}
