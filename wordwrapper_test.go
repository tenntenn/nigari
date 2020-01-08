package nigari_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/tenntenn/nigari"
)

func TestWordWrapper(t *testing.T) {
	cases := []struct {
		s    string
		w    float64
		want []string
	}{
		{"123。56789", 3, []string{"12", "3。5", "678", "9"}},
		{"abcd efgh", 3, []string{"ab-", "cd ", "ef-", "gh"}},
		{"abcdefgh", 3, []string{"ab-", "cd-", "ef-", "gh"}},
	}

	for _, tt := range cases {
		tt := tt
		name := fmt.Sprintf("%s-%.0f", tt.s, tt.w)
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			w := &nigari.WordWrapper{
				Measurer: nigari.MeasurerFunc(func(r rune) (w, h float64) {
					return 1, 1
				}),
				Width: tt.w,
			}
			got := w.Do(tt.s)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Error("unexpected result:", diff)
			}
		})
	}
}

func Test_word(t *testing.T) {
	cases := []struct {
		s     string
		i     int
		start int
		end   int
	}{
		{"cat", 1, 0, 3},
		{" cat", 1, 1, 4},
		{"@at", 1, -1, -1},
		{"cat dog", 3, -1, -1},
		{"cat  dog", 4, -1, -1},
		{"cat  dog", 5, 5, 8},
		{"cat@", 1, -1, -1},
	}

	for _, tt := range cases {
		tt := tt
		name := fmt.Sprintf("%s-%d", tt.s, tt.i)
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			start, end := nigari.ExportWord([]rune(tt.s), tt.i)
			if start != tt.start || end != tt.end {
				t.Errorf("want %d,%d but got %d,%d", tt.start, tt.end, start, end)
			}
		})
	}
}
