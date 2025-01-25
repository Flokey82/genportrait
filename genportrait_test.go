package genportrait

import (
	"fmt"
	"image/png"
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {
	g, err := LoadSprites()
	if err != nil {
		t.Errorf("LoadSprites() error = %v", err)
		return
	}
	for i := 0; i < 4; i++ {
		img := g.Random()

		// Write to file for manual inspection
		f, err := os.Create(fmt.Sprintf("portrait_%d.png", i))
		if err != nil {
			t.Errorf("os.Create() error = %v", err)
			return
		}
		defer f.Close()
		if err := png.Encode(f, img); err != nil {
			t.Errorf("png.Encode() error = %v", err)
			return
		}
	}
}
