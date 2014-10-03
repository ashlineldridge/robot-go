package main

import "testing"

func TestValidDirections(t *testing.T) {
	if !North.IsValid() {
		t.Log("North should be valid.")
		t.Fail()
	}
	if !East.IsValid() {
		t.Log("East should be valid.")
		t.Fail()
	}
	if !South.IsValid() {
		t.Log("South should be valid.")
		t.Fail()
	}
	if !West.IsValid() {
		t.Log("West should be valid.")
		t.Fail()
	}
}

func TestInvalidDirections(t *testing.T) {
	if Invalid.IsValid() {
		t.Log("Invalid should not be valid.")
		t.Fail()
	}
	if Direction(Invalid + 1).IsValid() {
		t.Log("Values beyond the max enumerated value should not be valid.")
		t.Fail()
	}
	if Direction(North - 1).IsValid() {
		t.Log("Values before the min enumerated value should not be valid.")
		t.Fail()
	}
}

func TestNewDirection(t *testing.T) {
	if NewDirection("NORTH") != North {
		t.Log("NORTH could not be correctly parsed.")
		t.Fail()
	}
	if NewDirection("EAST") != East {
		t.Log("EAST could not be correctly parsed.")
		t.Fail()
	}
	if NewDirection("SOUTH") != South {
		t.Log("SOUTH could not be correctly parsed.")
		t.Fail()
	}
	if NewDirection("WEST") != West {
		t.Log("WEST could not be correctly parsed.")
		t.Fail()
	}
	if NewDirection("FOOBAR") != Invalid {
		t.Log("Invalid direction was not detected.")
		t.Fail()
	}
}
