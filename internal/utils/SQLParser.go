package utils

import (
	"fmt"
	"log/slog"
	"os"
)

// Return a string representation of the relative file path to the /sql folder
//
// Because this is for interal use, it will panic if it cannot read the file
func SQLReader(path string) string {
	sqlFile, err := os.ReadFile(path)
	if err != nil {
		slog.Error(fmt.Sprint(err))
		panic(err)
	}
	query := string(sqlFile)

	slog.Debug("Query: \n" + query)

	return query
}