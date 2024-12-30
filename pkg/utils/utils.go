package utils

import "strings"

// ToUpperCase converts a string to uppercase
// @Description Converts the given input string to uppercase.
// @Param input string The string to convert.
// @Return string The uppercase version of the input string.
func ToUpperCase(input string) string {
	return strings.ToUpper(input)
}

// Contains checks if a slice contains a specific element
// @Description Checks whether the given slice contains a specific string element.
// @Param slice []string The slice of strings to search within.
// @Param item string The string element to search for.
// @Return bool True if the slice contains the item, false otherwise.
func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
