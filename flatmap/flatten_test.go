package flatmap

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFlatten(t *testing.T) {
	cases := []struct {
		Input  map[string]interface{}
		Output Map
	}{
		{
			Input: map[string]interface{}{
				"foo": "bar",
				"bar": "baz",
			},
			Output: Map{
				"foo": "bar",
				"bar": "baz",
			},
		},

		{
			Input: map[string]interface{}{
				"foo": []string{
					"one",
					"two",
				},
			},
			Output: Map{
				"foo.#": 2,
				"foo.0": "one",
				"foo.1": "two",
			},
		},

		{
			Input: map[string]interface{}{
				"foo": []map[interface{}]interface{}{
					map[interface{}]interface{}{
						"name":    "bar",
						"port":    3000,
						"enabled": true,
					},
				},
			},
			Output: Map{
				"foo.#":         1,
				"foo.0.name":    "bar",
				"foo.0.port":    int64(3000),
				"foo.0.enabled": true,
			},
		},

		{
			Input: map[string]interface{}{
				"foo": []map[interface{}]interface{}{
					map[interface{}]interface{}{
						"name": "bar",
						"ports": []string{
							"1",
							"2",
						},
					},
				},
			},
			Output: Map{
				"foo.#":         1,
				"foo.0.name":    "bar",
				"foo.0.ports.#": 2,
				"foo.0.ports.0": "1",
				"foo.0.ports.1": "2",
			},
		},
	}

	for _, tc := range cases {
		actual := Flatten(tc.Input)
		if !reflect.DeepEqual(actual, tc.Output) {
			t.Fatal(cmp.Diff(actual, tc.Output))
		}
	}
}
