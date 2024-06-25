package gotinylogger

import "testing"

func Test(t *testing.T) {

	testSlice := []int{1, 2, 3}
	Pretty("int slice", &testSlice)

	var testStruct = struct {
		fruit  string
		weight float64
	}{fruit: "apple", weight: 12.4}

	Pretty("fruit struct", testStruct)

	var mp = map[string]int{"first": 1, "second": 2}
	Pretty("map", mp)
}
