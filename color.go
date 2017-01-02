package pixelflut

// Color stores an RGBA Color
type Color interface {

	// RGBA returns the color with premultiplied alpha values
	// You might want to bit-shift the r, g and b values >> 8 before use.
	RGBA() (r, g, b, a uint32)

	// HexString returns the hex representation of the color
	HexString() string
}
