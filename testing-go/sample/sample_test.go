package sample

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_privateFunc(t *testing.T) {
	expected := Sample{
		FirstName: "first",
		LastName:  "first",
	}
	comparer := cmp.Comparer(func(x, y Sample) bool {
		return x.FirstName == y.FirstName
	})
	res := privateFunc()
	if diff := cmp.Diff(expected, res, comparer); diff != "" {
		t.Error(res)
	}
}
