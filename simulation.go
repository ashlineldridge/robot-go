package main

type Simulation struct {
	Robot *Robot
	Table Table
}

func (s *Simulation) Place(r Robot) {
	s.Robot = &r
}

func (s *Simulation) Move() {
	if s.Robot != nil {
		next := s.Robot.Position.Next()
		if s.Table.IsValid(next) {
			s.Robot.Position = next
		}
	}
}

func (s *Simulation) Right() {
	s.rotate(1)
}

func (s *Simulation) Left() {
	s.rotate(-1)
}

func (s *Simulation) Report() string {
	if s.Robot != nil {
		return s.Robot.Position.String()
	}
	return "Robot has not been placed"
}

func (s *Simulation) rotate(units int) {
	if s.Robot != nil {
		s.Robot.Position.Facing = s.Robot.Position.Facing.Rotate(units)
	}
}
