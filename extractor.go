package gotypes

import (
	"reflect"
	"strings"
)

// ExtractJson extracts the JSON attributes from a struct
func ExtractJson(s interface{}) (string, map[string]string) {
	// Use reflection to inspect the struct
	val := reflect.ValueOf(s)
	// Extract the type of the struct
	typ := val.Type()

	var jsonAttributes map[string]string = make(map[string]string)

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		tag := field.Tag.Get("json")

		// If the tag is empty, we use the field name as the JSON attribute
		if tag != "" {
			jsonAttributes[tag] = field.Type.String()
		} else {
			jsonAttributes[convertNameToJson(field.Name)] = field.Type.String()
		}
	}

	// get struct name
	name := typ.Name()

	return name, jsonAttributes
}

// convertNameToJson converts a struct field name to JSON format
func convertNameToJson(name string) string {
	lower := strings.ToLower(name)
	res := strings.Replace(lower, " ", "_", -1)
	res = strings.Replace(res, "-", "_", -1)

	return res
}

// GetStructName gets the name of a struct
func GetStructName(s interface{}) string {
	val := reflect.ValueOf(s)
	typ := val.Type()

	return typ.Name()
}
