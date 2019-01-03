package coord

type Coord struct {
	X, Y        int
	initialized bool
}

func New(x, y int) Coord {
	return Coord{X: x, Y: y, initialized: true}
}

func (c Coord) IsNull() bool {
	return !c.initialized
}

func (c Coord) DistanceTo(x, y int) int {
	return abs(c.X-x) + abs(c.Y-y)
}

func Find(coords []Coord, coord Coord) (int, Coord) {
	for index, c := range coords {
		if c.X == coord.X && c.Y == coord.Y && c.initialized {
			return index, c
		}
	}
	return -1, Coord{0, 0, false}
}

func Contains(coords []Coord, coord Coord) bool {
	index, _ := Find(coords, coord)
	return index >= 0
}

func (c Coord) DistanceToCoord(other Coord) int {
	return Distance(c, other)
}

func Distance(c1, c2 Coord) int {
	return c1.DistanceTo(c2.X, c2.Y)
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return a * -1
}
