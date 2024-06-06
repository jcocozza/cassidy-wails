package database

import (
	"io"
	"os"
    "path/filepath"
)

// Copy file data from the source to the destination
func copy(source, destination string) error {
    file, err := os.Open(source)
    if err != nil {
        return err
    }
    defer file.Close()

    duplicate, err := os.OpenFile(destination, os.O_RDWR|os.O_CREATE, 0666)
    if err != nil {
        return err
    }
    defer duplicate.Close()
    _, err1 := io.Copy(duplicate, file)
    if err1 != nil {
        return err
    }
    return nil
}
// Get the user downloads directory
func getDownloadDir() (string, error) {
    homeDir, err := os.UserHomeDir()
    if err != nil {
        return "", err
    }
    // Construct the path to the Downloads folder
    downloadsPath := filepath.Join(homeDir, "Downloads")
    return downloadsPath, nil
}
// copy the app database to the downloads folder
func ExportDatabase() error {
    downloadDir, err := getDownloadDir()
    if err != nil {
        return err
    }
    downloadsFile := filepath.Join(downloadDir, "CassidyDB")

    dbPath, err := getCassidyDBPath()
    if err != nil {
        return err
    }
    err1 := copy(dbPath, downloadsFile)
    if err1 != nil {
        return err1
    }
    return nil
}
