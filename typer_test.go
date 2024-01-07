package gotypes_test

import (
	"os"
	"reflect"
	"testing"

	gotypes "github.com/newtoallofthis123/go_types"
)

func TestGenerateType(t *testing.T) {
	gotypes.GenerateType("test.ts", UserRequest{})

	// read the file
	// check if the file is equal to the expected
	content, err := os.ReadFile("test.ts")
	if err != nil {
		panic(err)
	}

	req := "export type UserRequest = {\n\tusername: string;\n\tpassword: string;\n\temail: string;\n\ttimes: number[];\n}\n"
	got := string(content)

	if reflect.DeepEqual(got, req) {
		t.Errorf("Expected %s, got %s", req, got)
	}
}

func TestGenerateTypes(t *testing.T) {
	gotypes.GenerateTypes("test.ts", &[]interface{}{UserRequest{}, UserResponse{}})

	// read the file
	// check if the file is equal to the expected
	content, err := os.ReadFile("test.ts")
	if err != nil {
		panic(err)
	}

	req := "export type UserRequest = {\n\tusername: string;\n\tpassword: string;\n\temail: string;\n\ttimes: number[];\n}\n\nexport type UserResponse = {\n\tusername: string;\n\tpassword: string;\n\temail: string;\n\ttimes: number[];\n}\n"
	got := string(content)

	if reflect.DeepEqual(got, req) {
		t.Errorf("Expected %s, got %s", req, got)
	}
}
