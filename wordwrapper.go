package nigari

type WordWrapper struct {
	Measurer Measurer
	Width    float64
}

func (w *WordWrapper) Do(s string) []string {

	if w.Width <= 0 {
		return []string{s}
	}

	if s == "" {
		return []string{""}
	}

	rs := []rune(s)

	var (
		lines    []string
		width    float64
		start, i int
	)

	for i < len(rs) {
		_, dw := w.Measurer.Do(rs[i])
		if width+dw <= w.Width {
			width += dw
			i++
			continue
		}

		for i-start > 0 && i-1 > 0 && gyomatsuKinsoku[rs[i-1]] {
			_, dw := w.Measurer.Do(rs[i])
			width -= dw
			if width < 0 {
				width = 0
			}
			i--
		}

		for i-start > 0 && gyotouKinsoku[rs[i]] {
			_, dw := w.Measurer.Do(rs[i])
			width -= dw
			if width < 0 {
				width = 0
			}
			i--
		}

		lines = append(lines, string(rs[start:i]))
		start = i
		width = 0
	}

	if i-start > 0 {
		lines = append(lines, string(rs[start:i]))
	}

	return lines
}
