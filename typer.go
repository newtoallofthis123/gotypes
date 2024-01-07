package gotypes

func GenerateType(filename string, s interface{}) {
	// Generate the TypeScript type
	tsType := GenerateTS(s)

	// Save the TypeScript type to a file
	saveToFile(filename, tsType)
}

// GenerateTypes generates TypeScript interfaces from Go structs
func GenerateTypes(filename string, structs *[]interface{}) {
	types := make([]string, len(*structs))

	for i, s := range *structs {
		types[i] = GenerateTS(s)
	}

	// Save the TypeScript types to a file
	saveToFile(filename, combineTypes(&types))
}
