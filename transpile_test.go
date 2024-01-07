package gotypes_test

import (
	"reflect"
	"testing"

	gotypes "github.com/newtoallofthis123/gotypes"
)

func TestTypeTest(t *testing.T) {
	req := "number[]"
	got := gotypes.ConvertToTSType("[]float32")

	if got != req {
		t.Errorf("Expected %s, got %s", req, got)
	}
}

func TestSingleTest(t *testing.T) {
	req := "export type UserRequest = {\n\tusername: string;\n\tpassword: string;\n\temail: string;\n\ttimes: number[];\n}\n"
	got := gotypes.GenerateTS(UserRequest{})

	if reflect.DeepEqual(got, req) {
		t.Errorf("Expected %s, got %s", req, got)
	}
}
