package util

import (
	"strings"
)

// Cn merges multiple class strings together, removing duplicates and empty strings
func Cn(classes ...string) string {
	classMap := make(map[string]bool)
	var result []string

	for _, class := range classes {
		if class == "" {
			continue
		}

		// Split each class string by spaces and add to map to remove duplicates
		parts := strings.Fields(class)
		for _, part := range parts {
			if part != "" && !classMap[part] {
				classMap[part] = true
				result = append(result, part)
			}
		}
	}

	return strings.Join(result, " ")
}
