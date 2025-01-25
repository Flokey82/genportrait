package genportrait

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"math/rand"

	_ "embed"
)

//go:embed sprites/ears_32_1x3.png
var ears_png []byte

//go:embed sprites/eyes_32_1x3.png
var eyes_png []byte

//go:embed sprites/eyebrows_32_1x4.png
var eyebrows_png []byte

//go:embed sprites/heads_32_1x4.png
var heads_png []byte

//go:embed sprites/mouths_32_1x4.png
var mouths_png []byte

//go:embed sprites/noses_32_1x2.png
var noses_png []byte

func LoadSprites() (*Generator, error) {
	const tileSize = 32

	// Load sprites
	ears, err := newSpritesheet(ears_png, tileSize)
	if err != nil {
		return nil, err
	}
	eyes, err := newSpritesheet(eyes_png, tileSize)
	if err != nil {
		return nil, err
	}
	eyebrows, err := newSpritesheet(eyebrows_png, tileSize)
	if err != nil {
		return nil, err
	}
	heads, err := newSpritesheet(heads_png, tileSize)
	if err != nil {
		return nil, err
	}
	mouths, err := newSpritesheet(mouths_png, tileSize)
	if err != nil {
		return nil, err
	}
	noses, err := newSpritesheet(noses_png, tileSize)
	if err != nil {
		return nil, err
	}

	return &Generator{
		heads:    heads,
		eyes:     eyes,
		eyebrows: eyebrows,
		ears:     ears,
		noses:    noses,
		mouths:   mouths,
	}, nil
}

type Generator struct {
	heads    *spritesheet
	eyes     *spritesheet
	eyebrows *spritesheet
	ears     *spritesheet
	noses    *spritesheet
	mouths   *spritesheet
}

// SkinColors is a list of skin colors.
// NOTE: Copilot came up with those colors, no idea if they are accurate or exhaustive.
var SkinColors = []color.RGBA{
	{255, 219, 172, 255}, // PeachPuff
	{255, 205, 148, 255}, // BurlyWood
	{255, 192, 203, 255}, // Pink
	{255, 228, 196, 255}, // Bisque
	{255, 228, 225, 255}, // MistyRose
	{139, 69, 19, 255},   // SaddleBrown
	{160, 82, 45, 255},   // Sienna
	{210, 105, 30, 255},  // Chocolate
	{205, 133, 63, 255},  // Peru
	{244, 164, 96, 255},  // SandyBrown
	{222, 184, 135, 255}, // BurlyWood
	{210, 180, 140, 255}, // Tan
	{188, 143, 143, 255}, // RosyBrown
	{205, 92, 92, 255},   // Reddish
	{165, 42, 42, 255},   // Brown
}

// EyeColors is a list of eye colors that can be used in the portrait generator.
var EyeColors = []color.RGBA{
	{0, 0, 0, 255},     // Black
	{0, 0, 255, 255},   // Blue
	{0, 128, 0, 255},   // Green
	{128, 0, 0, 255},   // Maroon
	{128, 0, 128, 255}, // Purple
	{255, 0, 0, 255},   // Red
}

// Random generates a random portrait.
func (g *Generator) Random() image.Image {
	// Draw head. (pick random head)
	headIdx := rand.Intn(g.heads.numTiles())

	// Draw eyes. (pick random eyes)
	eyesIdx := rand.Intn(g.eyes.numTiles())

	// Draw eyebrows. (pick random eyebrows)
	eyebrowsIdx := rand.Intn(g.eyebrows.numTiles())

	// Draw ears. (pick random ears)
	earsIdx := rand.Intn(g.ears.numTiles())

	// Draw nose. (pick random nose)
	noseIdx := rand.Intn(g.noses.numTiles())

	// Draw mouth. (pick random mouth)
	mouthIdx := rand.Intn(g.mouths.numTiles())

	// Pick a new skin color for the portrait.
	skinColor := SkinColors[rand.Intn(len(SkinColors))]

	// Pick a random eye color
	eyeColor := EyeColors[rand.Intn(len(EyeColors))]

	return g.Generate(eyeColor, skinColor, headIdx, eyesIdx, eyebrowsIdx, earsIdx, noseIdx, mouthIdx)
}

var defaultSkinColor = color.RGBA{0xee, 0xc3, 0x9a, 0xff}
var defaultEyeColor = color.RGBA{0x8f, 0x56, 0x3b, 0xff}

func (g *Generator) Generate(eyeColor, skinColor color.RGBA, headIndex, eyeIndex, eyebrowIndex, earIndex, noseIndex, mouthIndex int) image.Image {
	portrait := image.NewRGBA(image.Rect(0, 0, 32, 32))
	// Draw head.
	headImg := g.heads.TileImage(headIndex)
	draw.Draw(portrait, portrait.Bounds(), headImg, image.Point{0, 0}, draw.Over)

	// Draw eyes.
	eyesImg := g.eyes.TileImage(eyeIndex)
	draw.Draw(portrait, portrait.Bounds(), eyesImg, image.Point{0, 0}, draw.Over)

	// Draw eyebrows.
	eyebrowsImg := g.eyebrows.TileImage(eyebrowIndex)
	draw.Draw(portrait, portrait.Bounds(), eyebrowsImg, image.Point{0, 0}, draw.Over)

	// Draw ears.
	earsImg := g.ears.TileImage(earIndex)
	draw.Draw(portrait, portrait.Bounds(), earsImg, image.Point{0, 0}, draw.Over)

	// Draw nose.
	noseImg := g.noses.TileImage(noseIndex)
	draw.Draw(portrait, portrait.Bounds(), noseImg, image.Point{0, 0}, draw.Over)

	// Draw mouth.
	mouthImg := g.mouths.TileImage(mouthIndex)
	draw.Draw(portrait, portrait.Bounds(), mouthImg, image.Point{0, 0}, draw.Over)

	// Replace
	portrait = replaceColor(portrait, defaultSkinColor, skinColor)
	portrait = replaceColor(portrait, defaultEyeColor, eyeColor)

	return portrait
}

func replaceColor(img image.Image, from, to color.Color) *image.RGBA {
	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)
	draw.Draw(newImg, bounds, img, bounds.Min, draw.Src)

	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			if img.At(x, y) == from {
				newImg.Set(x, y, to)
			}
		}
	}

	return newImg
}

// spritesheet is a convenience wrapper around locating sprites in a spritesheet.
type spritesheet struct {
	image    image.Image
	tileSize int // Size of each tile in the spritesheet
	xCount   int // Number of tiles in the x direction
	yCount   int // Number of tiles in the y direction
	x        int // Width of the spritesheet
	y        int // Height of the spritesheet
}

func newSpritesheet(imgData []byte, tileSize int) (*spritesheet, error) {
	// Decode an image from the image file's byte slice.
	// Now the byte slice is generated with //go:generate for Go 1.15 or older.
	// If you use Go 1.16 or newer, it is strongly recommended to use //go:embed to embed the image file.
	// See https://pkg.go.dev/embed for more details.
	img, _, err := image.Decode(bytes.NewReader(imgData))
	if err != nil {
		return nil, err
	}

	// Get the size of the image
	bounds := img.Bounds()
	x := bounds.Dx()
	y := bounds.Dy()

	// Calculate the number of tiles in the x and y directions
	xCount := x / tileSize
	yCount := y / tileSize

	return &spritesheet{
		image:    img,
		tileSize: tileSize,
		xCount:   xCount,
		yCount:   yCount,
		x:        x,
		y:        y,
	}, nil
}

// numTiles returns the number of tiles in the spritesheet.
func (s *spritesheet) numTiles() int {
	return s.xCount * s.yCount
}

// TileImage returns an image.Image of the tile at the given index.
// TODO: This should maybe take an image (and maybe offset) to draw on
// instead of returning a new image. Also the color replacement could be
// done here.
func (s *spritesheet) TileImage(index int) image.Image {
	// Calculate the x and y position of the tile in the spritesheet
	x := (index % s.xCount) * s.tileSize
	y := (index / s.xCount) * s.tileSize

	// Create a new RGBA image for the tile
	tile := image.NewRGBA(image.Rect(0, 0, s.tileSize, s.tileSize))

	// Copy the tile from the spritesheet to the new image
	for i := 0; i < s.tileSize; i++ {
		for j := 0; j < s.tileSize; j++ {
			tile.Set(i, j, s.image.At(x+i, y+j))
		}
	}

	return tile
}
