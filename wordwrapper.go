package nigari

import (
	"unicode"

	"golang.org/x/image/math/fixed"
)

type WordWrapper struct {
	Measurer Measurer
	Width    fixed.Int26_6
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
		width    fixed.Int26_6
		start, i int
	)

	for i < len(rs) {
		prevC := rune(-1)
		if i-1 >= 0 {
			prevC = rs[i-1]
		}
		dw := w.Measurer.Do(rs[i], prevC)
		if width+dw <= w.Width {
			width += dw
			i++
			continue
		}

		for i-start > 0 && i-1 > 0 && gyomatsuKinsoku[rs[i-1]] {
			prevC := rune(-1)
			if i-1 >= 0 {
				prevC = rs[i-1]
			}
			width -= w.Measurer.Do(rs[i], prevC)
			if width < 0 {
				width = 0
			}
			i--
		}

		for i-start > 0 && gyotouKinsoku[rs[i]] {
			prevC := rune(-1)
			if i-1 >= 0 {
				prevC = rs[i-1]
			}

			width -= w.Measurer.Do(rs[i], prevC)
			if width < 0 {
				width = 0
			}
			i--
		}

		// hyphenation
		ws, we := word(rs, i-1)
		if ws >= 0 && we > i-1 {
			i--
			// TODO: measure size of "-"
			lines = append(lines, string(rs[start:i])+"-")
		} else {
			lines = append(lines, string(rs[start:i]))
		}

		start = i
		width = 0
	}

	if i-start > 0 {
		lines = append(lines, string(rs[start:i]))
	}

	return lines
}

func word(rs []rune, i int) (start, end int) {

	if i < 0 || i >= len(rs) {
		return -1, -1
	}

	if !isAlpha(rs[i]) {
		return -1, -1
	}

	start = i
	for start > 0 {
		if isAlpha(rs[start-1]) {
			start--
		} else if unicode.IsSpace(rs[start-1]) {
			break
		} else {
			return -1, -1
		}
	}

	end = i
	for end < len(rs) {
		if isAlpha(rs[end]) {
			end++
		} else if unicode.IsSpace(rs[end]) {
			break
		} else {
			return -1, -1
		}
	}

	return start, end
}

func isAlpha(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}
