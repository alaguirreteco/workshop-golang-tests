package calc

import "testing"

type TestDataItem struct {
	input    float64 // inputs to `raiz` function
	result   float64 // result of `raiz` function
	hasError bool    // does `raiz` function returns error
}

func TestRaiz(t *testing.T) {
	dataItems := []TestDataItem{
		{9, 3, false},
		{400, 20, false},
		{0, 0, false},
		{-1, 0, true},
	}
	for _, item := range dataItems {
		err, result := raiz(item.input)

		if item.hasError {
			if err == nil {
				t.Errorf("raiz() with args %v : FAILED, expected an error but got value '%v'", item.input, result)
			} else {
				t.Logf("raiz() with args %v : PASSED, expected an error and got an error '%v'", item.input, err)
			}
		} else {
			if result != item.result {
				t.Errorf("raiz() with args %v : FAILED, expected %v but got value '%v'", item.input, item.result, result)
			} else {
				t.Logf("raiz() with args %v : PASSED, expected %v and got value '%v'", item.input, item.result, result)
			}
		}
	}
}
