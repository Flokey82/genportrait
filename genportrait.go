package genportrait

import (
	"image"
	"image/color"
	"image/draw"
	"math/rand"

	_ "embed"

	spritesheet "github.com/Flokey82/go_spritesheet"
)

//go:embed sprites/ears_32_1x3.png
var ears_png []byte

//go:embed sprites/eyes_32_1x3.png
var eyes_png []byte

//go:embed sprites/eyebrows_32_1x4.png
var eyebrows_png []byte

//go:embed sprites/heads_32_1x4.png
var heads_png []byte

//go:embed sprites/noses_32_1x2.png
var noses_png []byte

//go:embed sprites/beards_32_1x5.png
var beards_png []byte

//go:embed sprites/hair_32_1x7.png
var hair_png []byte

//go:embed sprites/hair_curly_32_1x7.png
var hair_curly_png []byte

//go:embed sprites/mouths_32_1x4.png
var mouths_png []byte

func LoadSprites() (*Generator, error) {
	const tileSize = 32

	// Load sprites
	ears, err := spritesheet.New(ears_png, tileSize)
	if err != nil {
		return nil, err
	}
	eyes, err := spritesheet.New(eyes_png, tileSize)
	if err != nil {
		return nil, err
	}
	eyebrows, err := spritesheet.New(eyebrows_png, tileSize)
	if err != nil {
		return nil, err
	}
	heads, err := spritesheet.New(heads_png, tileSize)
	if err != nil {
		return nil, err
	}
	noses, err := spritesheet.New(noses_png, tileSize)
	if err != nil {
		return nil, err
	}
	beards, err := spritesheet.New(beards_png, tileSize)
	if err != nil {
		return nil, err
	}
	hair, err := spritesheet.New(hair_png, tileSize)
	if err != nil {
		return nil, err
	}
	hairCurly, err := spritesheet.New(hair_curly_png, tileSize)
	if err != nil {
		return nil, err
	}
	mouths, err := spritesheet.New(mouths_png, tileSize)
	if err != nil {
		return nil, err
	}

	return &Generator{
		heads:     heads,
		eyes:      eyes,
		eyebrows:  eyebrows,
		ears:      ears,
		beards:    beards,
		noses:     noses,
		hair:      hair,
		hairCurly: hairCurly,
		mouths:    mouths,
	}, nil
}

type Generator struct {
	heads     *spritesheet.Spritesheet
	eyes      *spritesheet.Spritesheet
	eyebrows  *spritesheet.Spritesheet
	ears      *spritesheet.Spritesheet
	beards    *spritesheet.Spritesheet
	noses     *spritesheet.Spritesheet
	hair      *spritesheet.Spritesheet
	hairCurly *spritesheet.Spritesheet
	mouths    *spritesheet.Spritesheet
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

// HairColors is a list of hair colors that can be used in the portrait generator.
var HairColors = []color.RGBA{
	{0, 0, 0, 255},       // Black
	{0, 0, 255, 255},     // Blue
	{0, 128, 0, 255},     // Green
	{128, 0, 0, 255},     // Maroon
	{128, 0, 128, 255},   // Purple
	{255, 0, 0, 255},     // Red
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
	{255, 255, 0, 255},   // Yellow (Blonde)
	{255, 215, 0, 255},   // Gold

	// Gray, White, Silver
	{128, 128, 128, 255}, // Gray
	{192, 192, 192, 255}, // Silver
	{211, 211, 211, 255}, // LightGray
	{220, 220, 220, 255}, // Gainsboro
	{245, 245, 245, 255}, // WhiteSmoke
	{255, 250, 250, 255}, // Snow
}

// Random generates a random portrait.
func (g *Generator) Random() image.Image {
	// Draw head. (pick random head)
	headIdx := rand.Intn(g.heads.NumTiles())

	// Draw eyes. (pick random eyes)
	eyesIdx := rand.Intn(g.eyes.NumTiles())

	// Draw eyebrows. (pick random eyebrows)
	eyebrowsIdx := rand.Intn(g.eyebrows.NumTiles())

	// Draw ears. (pick random ears)
	earsIdx := rand.Intn(g.ears.NumTiles())

	// Draw beard. (pick random beard)
	beardIdx := rand.Intn(g.beards.NumTiles())

	// Draw nose. (pick random nose)
	noseIdx := rand.Intn(g.noses.NumTiles())

	// Draw hair. (pick random hair)
	hairIdx := rand.Intn(g.hair.NumTiles())

	// Draw mouth. (pick random mouth)
	mouthIdx := rand.Intn(g.mouths.NumTiles())

	// Pick a new skin color for the portrait.
	skinColor := SkinColors[rand.Intn(len(SkinColors))]

	// Pick a random eye color
	eyeColor := EyeColors[rand.Intn(len(EyeColors))]

	// Pick a random hair color
	hairColor := HairColors[rand.Intn(len(HairColors))]

	// Pick if the hair should be curly
	hairCurly := rand.Intn(2) == 1

	return g.Generate(eyeColor, skinColor, hairColor, headIdx, eyesIdx, eyebrowsIdx, earsIdx, beardIdx, noseIdx, hairIdx, mouthIdx, hairCurly)
}

var defaultSkinColor = color.RGBA{0xee, 0xc3, 0x9a, 0xff}
var defaultEyeColor = color.RGBA{0x8f, 0x56, 0x3b, 0xff}
var defaultHairColor = color.RGBA{0xfb, 0xf2, 0x36, 0xff}

func (g *Generator) Generate(eyeColor, skinColor, hairColor color.RGBA, headIndex, eyeIndex, eyebrowIndex, earIndex, beardIdx, noseIndex, hairIndex, mouthIndex int, hairCurly bool) image.Image {
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

	// Draw beard.
	beardImg := g.beards.TileImage(beardIdx)
	draw.Draw(portrait, portrait.Bounds(), beardImg, image.Point{0, 0}, draw.Over)

	// Draw nose.
	noseImg := g.noses.TileImage(noseIndex)
	draw.Draw(portrait, portrait.Bounds(), noseImg, image.Point{0, 0}, draw.Over)

	// Draw hair.
	if hairCurly {
		hairImg := g.hairCurly.TileImage(hairIndex)
		draw.Draw(portrait, portrait.Bounds(), hairImg, image.Point{0, 0}, draw.Over)
	} else {
		hairImg := g.hair.TileImage(hairIndex)
		draw.Draw(portrait, portrait.Bounds(), hairImg, image.Point{0, 0}, draw.Over)

	}

	// Draw mouth.
	mouthImg := g.mouths.TileImage(mouthIndex)
	draw.Draw(portrait, portrait.Bounds(), mouthImg, image.Point{0, 0}, draw.Over)

	// Replace
	portrait = spritesheet.ReplaceColor(portrait, defaultSkinColor, skinColor)
	portrait = spritesheet.ReplaceColor(portrait, defaultEyeColor, eyeColor)
	portrait = spritesheet.ReplaceColor(portrait, defaultHairColor, hairColor)

	return portrait
}
