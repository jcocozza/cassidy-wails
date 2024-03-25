package database

import (
	"fmt"
	"strings"
)
const placeholder = "?"

// Used for inserting a list into sql statements used by the database/sql package.
//
// e.g. SELECT * FROM mytbl WHERE id IN (%s) => SELECT * FROM mytbl WHERE id IN (?, ?, ?, ?).
//
// Note: it's important to use the '%s' syntax for the list so that the funtion knows where to replace.
func SQLListInsertion(sql string, numElementInList int) string {
	placeholders := make([]string, numElementInList)

	// Generate placeholders (?, ?, ?, ...)
	for i := 0; i < numElementInList; i++ {
		placeholders[i] = placeholder
	}

	placeholderString := strings.Join(placeholders, ", ")

	query := fmt.Sprintf(sql, placeholderString)
	return query
}