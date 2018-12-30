package shared

// Point is a x,y coordinate
type Point struct {
	X, Y int
}

// Equals returns true if A and B is the same point (same x and y)
func Equals(pointA, pointB Point) bool {
	return pointA.X == pointB.X && pointA.Y == pointB.Y
}

// FirstInReadingOrder checks if pointA is before pointB "in reading order" (see Day15 2018)
func FirstInReadingOrder(pointA, pointB Point) bool {
	if pointA.Y != pointB.Y {
		return pointA.Y < pointB.Y
	}
	return pointA.X < pointB.X
}

// Distance returns the distance between A and B
func Distance(pointA, pointB Point) int {
	return abs(pointA.X-pointB.X) + abs(pointA.Y-pointB.Y)
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return a * -1
}
