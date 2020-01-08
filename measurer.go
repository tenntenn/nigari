package nigari

type Measurer interface {
	Do(r rune) (w, h float64)
}

type MeasurerFunc func(r rune) (w, h float64)

var _ Measurer = MeasurerFunc(nil)

func (m MeasurerFunc) Do(r rune) (w, h float64) {
	return m(r)
}
