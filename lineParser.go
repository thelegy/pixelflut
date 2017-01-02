package pixelflut

// Lineparser warps a function Parse to parse a line of commands and returns
// the proper value with the proper type.
type LineParser interface {
	// Parses a line and returns a value of the proper type.
	//
	// A line starting with "PX" is a pixel and will return a Pixel type.
	Parse(s string) (interface{}, error)
}
