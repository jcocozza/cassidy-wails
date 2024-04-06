package database

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/jcocozza/cassidy-wails/internal/utils"
	_ "github.com/mattn/go-sqlite3"
)

const (
	sqliteTestDB = "/Users/josephcocozza/Repositories/unnamed-app/cassidy-wails/internal/test/cassidy-wails-test.db"
	cassidyDB = ".cassidy.db"
)
// Connect to a SQLite database.
func connectToSQLite(path string) (*sql.DB, error) {
    db, err := sql.Open("sqlite3", path)
    if err != nil {
        slog.Error("Failed to open database")
        return nil, err
    }

    // Ping the database to ensure connectivity
    err = db.Ping()
    if err != nil {
        slog.Error("Failed to connect to database - ping failed")
        return nil, err
    }
    slog.Debug("Connected to SQLite Database")
    return db, nil
}
func createDatabaseFile(dbPath string) error {
    // Ensure the directory structure exists
    dir := filepath.Dir(dbPath)
    if err := os.MkdirAll(dir, 0755); err != nil {
        return err
    }
    // Create the database file
    _, err := os.Create(dbPath)
    if err != nil {
        return err
    }
    fmt.Println("Database file created at:", dbPath)
    return nil
}
// Will attempt to connect to the application database that is packaged with the app.
func ConnectToCassidyDB() (*Database, error) {
	exePath, err := os.Executable()
	fmt.Println("os exe dir: " + exePath)
	if err != nil {
		return nil, err
	}
	exeDir := filepath.Dir(filepath.Dir(exePath))
	dbPath := filepath.Join(exeDir,"Resources", cassidyDB)
	fmt.Println("database dir: " + dbPath)

	// Check if the database file exists, if not, create it
    if _, err := os.Stat(dbPath); os.IsNotExist(err) {
        if err := createDatabaseFile(dbPath); err != nil {
            return nil, err
        }
    }

	db, err1 := connectToSQLite(dbPath)
	if err1 != nil {
		return nil, err1
	}

	return &Database{DB: db}, nil
}

// The following functions are for testing and development purposes. They should not be used in any application deployment

// Run schema creation for the database.
func TableCreate(db *sql.DB) {
	dir := "/Users/josephcocozza/Repositories/unnamed-app/cassidy-wails/internal/sqlcode/schema/"

	// Open the directory
	d, err := os.Open(dir)
	if err != nil {
		panic("Unable to open directory!")
	}
	defer d.Close()

	// Get a list of file names in the directory
	files, err := d.Readdirnames(-1)
	if err != nil {
		panic("Unable to read directory")
	}

	for _, file := range files {
		sql := utils.SQLReader(dir + file)
		_, err := db.Exec(sql)
		if err != nil {
			panic("sql schema creation failed: " + fmt.Sprint(err))
		}
	}
}
// Run the test_inserts.sql file for the passed database.
func RunInserts(db *sql.DB) {
	sql := utils.SQLReader("/Users/josephcocozza/Repositories/unnamed-app/cassidy-wails/internal/test/test_inserts.sql")
	_, err := db.Exec(sql)
	if err != nil {
		panic("test inserts failed: " + fmt.Sprint(err))
	}
}
// Create the test database.
func InitTestDB() *Database {
	os.Remove(sqliteTestDB)
	DB, err := connectToSQLite(sqliteTestDB)
	if err != nil {
		panic("unable to create test db: " + fmt.Sprint(err))
	}
	TableCreate(DB)
	RunInserts(DB)
	return &Database{DB: DB}
}