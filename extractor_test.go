package gotypes_test

import (
	"reflect"
	"testing"

	. "github.com/newtoallofthis123/go_types"
)

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Times    []int  `json:"times"`
}

type UserResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Times    []int  `json:"times"`
}

func TestJSONAttributes(t *testing.T) {
	req := map[string]string{
		"username": "string",
		"password": "string",
		"email":    "string",
		"times":    "[]int",
	}

	name, got := ExtractJson(UserRequest{})

	// check equality of the two maps
	if !reflect.DeepEqual(got, req) {
		t.Errorf("Expected %v, got %v", req, got)
	}

	if name != "UserRequest" {
		t.Errorf("Expected %s, got %s", "UserRequest", name)
	}
}
