package pixelflut

// PositionParser wraps a function Parse to parse an x and y coordinate.
type PositionParser interface {
	// Parses the string into an x and y coordinate.
	Parse(xs, ys string) (x, y uint32, err error)
}
