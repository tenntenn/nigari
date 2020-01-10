package main

import (
	"image"
	"image/png"
	"io/ioutil"
	"os"

	"github.com/golang/freetype/truetype"
	"github.com/tenntenn/nigari"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

func main() {
	baseFont, err := loadFont(os.Args[1], 10)
	if err != nil {
		panic(err)
	}
	emojiFont, err := loadFont(os.Args[2], 10)
	if err != nil {
		panic(err)
	}

	d := &nigari.Drawer{
		Base:    baseFont,
		Emoji:   emojiFont,
		Spacing: 1.5,
		Width:   fixed.I(10), // 非常に小さな幅
	}

	dst := image.NewCMYK(image.Rect(0, 0, 200, 200))
	fg := image.Black
	d.Draw("Hello😁🍣🍺😺🐑ほげほげほげほげほげほげほげ👁️ふが", 10, 20, dst, fg)

	file, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	if err := png.Encode(file, dst); err != nil {
		panic(err)
	}
}

func loadFont(path string, size float64) (font.Face, error) {
	fontBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	f, err := truetype.Parse(fontBytes)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(f, &truetype.Options{
		Size:    size,
		DPI:     128,
		Hinting: font.HintingNone,
	}), nil
}
