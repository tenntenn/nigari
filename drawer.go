package nigari

import (
	"image"
	"image/draw"

	"github.com/rivo/uniseg"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

type Drawer struct {
	Base    font.Face
	Emoji   font.Face
	Spacing float64
	Width   fixed.Int26_6
}

func (d *Drawer) spacing() float64 {
	if d.Spacing < 0 {
		return 1.5
	}
	return d.Spacing
}

func (d *Drawer) Draw(s string, x, y int, dst draw.Image, fg image.Image) {

	s = d.removeMultiCodepoints(s)
	ww := &WordWrapper{
		Measurer: MeasurerFunc(func(c, prevC rune) fixed.Int26_6 {
			advance, ok := d.face(c).GlyphAdvance(c)
			if !ok {
				return 0
			}
			return d.face(prevC).Kern(prevC, c) + advance
		}),
		Width: d.Width,
	}
	lines := ww.Do(s)

	fd := &font.Drawer{
		Dst:  dst,
		Src:  fg,
		Face: d.Base,
		Dot:  fixed.P(x, y),
	}

	dy := max(d.Base.Metrics().Height.Mul(fixed.Int26_6(d.spacing()*64)),
		d.Emoji.Metrics().Height.Mul(fixed.Int26_6(d.spacing()*64)))

	var (
		prevC    = rune(-1)
		prevFace font.Face
	)

	for _, line := range lines {
		for _, c := range line {

			if prevC >= 0 && prevFace != nil {
				fd.Dot.X += prevFace.Kern(prevC, c)
			}

			fd.Face = d.face(c)
			dr, mask, maskp, advance, ok := fd.Face.Glyph(fd.Dot, c)
			if !ok {
				continue
			}

			draw.DrawMask(fd.Dst, dr, fd.Src, image.Point{}, mask, maskp, draw.Over)
			fd.Dot.X += advance
			prevC = c
			prevFace = fd.Face
		}
		fd.Dot.X = fixed.I(x)
		fd.Dot.Y += dy
	}
}

func (d *Drawer) face(c rune) font.Face {
	if IsEmoji(c) {
		return d.Emoji
	}
	return d.Base
}

func (d *Drawer) removeMultiCodepoints(s string) string {
	var result []rune
	gr := uniseg.NewGraphemes(s)
	for gr.Next() {
		rs := gr.Runes()
		if len(rs) == 1 {
			result = append(result, rs[0])
		}
	}
	return string(result)
}

func max(x, y fixed.Int26_6) fixed.Int26_6 {
	if x > y {
		return x
	}
	return y
}
