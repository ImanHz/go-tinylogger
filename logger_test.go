package gotinylogger

import "testing"

func PrettyTest(t *testing.T) {

	logger := New()
	testSlice := []int{1, 2, 3}
	logger.Pretty("int slice", testSlice)

	var testStruct = struct {
		fruit  string
		weight float64
	}{fruit: "apple", weight: 12.4}

	logger.Pretty("fruit struct", testStruct)

	logger.SetColor(false)
	var mp = map[string]int{"first": 1, "second": 2}
	logger.Pretty("map", mp)
}
