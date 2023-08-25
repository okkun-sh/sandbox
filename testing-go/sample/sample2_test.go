package sample_test

import (
	"fmt"
	"testing"
	"testing-go/sample"

	"github.com/google/go-cmp/cmp"
)

func TestMain(m *testing.M) {
	fmt.Println("before")

	m.Run()

	fmt.Println("after")
}

func TestPublicFunc(t *testing.T) {
	expected := sample.Sample{
		FirstName: "first",
		LastName:  "last",
	}
	res := sample.PublicFunc()
	if diff := cmp.Diff(expected, res); diff != "" {
		t.Error(res)
	}
}

func TestPublicFunc2(t *testing.T) {
	expected := sample.Sample{
		FirstName: "first",
		LastName:  "last",
	}
	res := sample.PublicFunc()
	if diff := cmp.Diff(expected, res); diff != "" {
		t.Error(res)
	}
}

var blackhole sample.Sample

func BenchmarkPublicFunc(b *testing.B) {
	res := sample.PublicFunc()
	blackhole = res
}
