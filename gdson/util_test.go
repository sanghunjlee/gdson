package gdson

import "testing"

func TestGetStructFieldNames(t *testing.T) {
	type testStruct struct {
		test1 string
		test2 []string
		test3 int
		test4 bool
	}
	var expected = []string{"test1", "test2", "test3", "test4"}
	test := &testStruct{test1: "", test2: nil, test3: -1, test4: false}
	testResult := getStructFieldNames(test)
	if len(expected) != len(testResult) {
		t.Errorf(`length mismatch: (expected) %v != (result) %v`, len(expected), len(testResult))
	} else {
		var found bool
		for i := 0; i < len(expected); i++ {
			found = false
			for j := 0; j < len(testResult); j++ {
				if expected[i] == testResult[j] {
					found = true
					break
				}
			}
			if !found {
				t.Errorf(`The expected value "%s" not found in the result`, expected[i])
			}
		}
	}
}
