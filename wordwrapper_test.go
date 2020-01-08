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
