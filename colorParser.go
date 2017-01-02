package pixelflut

// ColorParser wraps a function Parse to parse a color string into a value of
// Color.
type ColorParser interface {
	// Parses the string and tries to identify a Color from it.
	Parse(s string) (col Color, err error)
}
