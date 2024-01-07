package gotypes

// ConvertToTSType converts a Go type to a TypeScript type
//
// This is a very basic implementation and does not support
// all of TypeScript's types
// TODO: Add support for more types
func ConvertToTSType(s string) string {
	isArray := false
	if len(s) > 2 && s[0:2] == "[]" {
		isArray = true
		s = s[2:]
	}

	tsType := ""
	switch s {
	case "string":
		tsType = "string"
	case "int", "int8", "int16", "int32", "uint", "uint8", "uint16":
		tsType = "number"
	case "int64", "uint32", "uint64":
		tsType = "bigint"
	case "float32", "float64":
		tsType = "number"
	case "bool":
		tsType = "boolean"
	default:
		tsType = "any"
	}

	if isArray {
		tsType = tsType + "[]"
	}

	return tsType
}

// Get TS type template
func getTSTypeTemplate(name string) string {
	return "export type " + name + " = {\n"
}

func addTSAttribute(name string, tsType string) string {
	return "\t" + name + ": " + tsType + ";\n"
}

// GenerateTSFromStructs generates TypeScript interfaces from Go structs
func GenerateTSFromStructs(structs *[]interface{}) string {
	types := make([]string, len(*structs))

	for i, s := range *structs {
		types[i] = GenerateTS(s)
	}

	return combineTypes(&types)
}

// GenerateTS generates a TypeScript interface from a Go struct
func GenerateTS(s interface{}) string {
	// Extract the JSON attributes from the struct
	name, jsonAttributes := ExtractJson(s)

	// Get the TS type template
	tsType := getTSTypeTemplate(name)

	// Add the attributes to the TS type
	for key, value := range jsonAttributes {
		tsType = tsType + addTSAttribute(key, ConvertToTSType(value))
	}

	// Close the TS type
	tsType = tsType + "}"

	return tsType
}
