package gdson

import (
	"testing"
)

func TestAdd(t *testing.T) {
	test := &Gdson{}
	var testArgs = []string{"condition", "day=1"}
	if err := test.Add(testArgs); err != nil {
		t.Error(err.Error())
	}
}
