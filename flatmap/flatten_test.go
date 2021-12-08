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

func TestFlattenAllReflectKinds(t *testing.T) {
	defer func() {
		r := recover()
		if r != nil {
			t.Fatal("flattening null panicked:", r)
		}
	}()
	var emptyInterface interface{}
	m := make(map[string]interface{})
	m["nil"] = nil
	m["Bool"] = true
	m["Int"] = int(123)
	m["Int8"] = int8(123)
	m["Int16"] = int16(123)
	m["Int32"] = int32(123)
	m["Int64"] = int64(123)
	m["Uint"] = uint(123)
	m["Uint8"] = uint8(123)
	m["Uint16"] = uint16(123)
	m["Uint32"] = uint32(123)
	m["Uint64"] = uint64(123)
	m["Uintptr"] = uintptr(123)
	m["Float32"] = float32(12.3)
	m["Float64"] = float64(12.3)
	m["Complex64"] = complex64(12 + 3i)
	m["Complex128"] = complex128(12 + 3i)
	m["Array"] = []int{1, 2, 3}
	m["Chan"] = make(chan<- int)
	m["Func"] = func() int { return 123 }
	m["Interface"] = emptyInterface
	m["Map"] = map[string]int{"what": 123}
	//m["Ptr"] = 123
	m["Slice"] = []int{1, 2, 3}[1:3]
	m["String"] = "123"
	m["Struct"] = struct{}{}
	//m["UnsafePointer"] = 123
	Flatten(m)
}
