package main

const (
	tableWidth  int = 5
	tableHeight int = 5
)

func main() {
	t := Table{width: tableWidth, height: tableHeight}
	r := NewRobot(&t)
	r.Report()
}
