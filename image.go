package pixelflut

// Image stores RGB image information of a given size in OpenGL compatible
// manner.
type Image struct {
	SizeX uint32
	SizeY uint32
	Pix   []uint8
}

// NewImage creates a new Image of a given size.
func NewImage(sizeX, sizeY uint32) *Image {
	pix := make([]uint8, 3*sizeX*sizeY)
	return &Image{
		SizeX: sizeX,
		SizeY: sizeY,
		Pix:   pix,
	}
}

// PixOffset calculates the offset of a pixel in the Pix buffer. This will
// only yield sensible values if the pixel lies on the image.
func (img *Image) PixOffset(x, y uint32) uint32 {
	return 3 * (img.SizeX*y + x)
}

// Draw draws a Pixel onto the Image. This function takes the alpha values
// into account, painting the new Color over the older one.
//
// Be informed, that Draw does not check if the Pixel lies on the Image. If
// it does not this results in undefined behaviour and can lead to crashes.
func (img *Image) Draw(pixel Pixel) {
	pos := img.PixOffset(pixel.X, pixel.Y)

	r, g, b, a := pixel.Color.RGBA()
	a = 0xff - a

	r += uint32(img.Pix[pos]) * a
	g += uint32(img.Pix[pos+1]) * a
	b += uint32(img.Pix[pos+2]) * a

	img.Pix[pos] = uint8(r >> 8)
	img.Pix[pos+1] = uint8(g >> 8)
	img.Pix[pos+2] = uint8(b >> 8)
}
