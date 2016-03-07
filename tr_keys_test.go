package maputil

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMultiReplace(t *testing.T) {
	Convey("Multiple matching runes all get replaced", t, func() {
		ourMap := map[string]interface{}{
			".i.have.many.dots.": true,
		}

		So(TrKeys(ourMap, '.', '_'), ShouldContainKey, "_i_have_many_dots_")
	})
}

func TestRecursion(t *testing.T) {
	Convey("Nested keys all get transliterated", t, func() {
		input := map[string]interface{}{
			"sub.tree": map[string]interface{}{
				"leaf.node": true,
			},
		}
		result := TrKeys(input, '.', '_')

		Convey("Branch nodes are transliterated", func() {
			So(result, ShouldContainKey, "sub_tree")
		})

		Convey("Leaf nodes are transliterated", func() {
			So(result["sub_tree"], ShouldContainKey, "leaf_node")
		})
	})
}
