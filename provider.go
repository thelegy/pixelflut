package pixelflut

// Provider wraps logic to obtain Pixel information.
type Provider interface {
	// Provide returns a channel to provide pixels.
	Provide() <-chan (Pixel)
}
