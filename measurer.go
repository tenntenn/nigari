package nigari

import (
	"golang.org/x/image/math/fixed"
)

type Measurer interface {
	Do(c, prevC rune) fixed.Int26_6
}

type MeasurerFunc func(c, prevC rune) fixed.Int26_6

var _ Measurer = MeasurerFunc(nil)

func (m MeasurerFunc) Do(c, prevC rune) fixed.Int26_6 {
	return m(c, prevC)
}
