package markdown

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Markdown_ToOutliners(t *testing.T) {
	var cases = []struct {
		in   Markdown
		want []*Outliner
	}{
		{
			Markdown{Path: "a.md", Contents: []string{"some desc", "# o1", "hoge", "", "##  o2", "foo", "bar", "#    o3   ", "", "#### "}},
			[]*Outliner{
				{"o1", 1, []string{"hoge"}},
				{"o2", 2, []string{"foo", "bar"}},
				{"o3", 1, []string{}},
				{"", 4, []string{}},
			},
		},
	}
	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			got := tt.in.ToOutliners()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToOutliners()=%#v; want %#v\ncmp.Diff(got,want)=%s", got, tt.want, cmp.Diff(got, tt.want))
			}
		})
	}
}
