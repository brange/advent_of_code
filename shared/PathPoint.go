package shared

// PathPoint is a Point with a distance field
type PathPoint struct {
	Point    Point
	Distance int
}

// FindPathPoint ..
func FindPathPoint(pathPoints []PathPoint, point Point) (PathPoint, bool) {
	var pp PathPoint

	for _, pathPoint := range pathPoints {
		if Equals(pathPoint.Point, point) {
			return pathPoint, true
		}
	}

	return pp, false
}
