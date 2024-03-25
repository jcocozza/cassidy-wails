package database

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	"github.com/jcocozza/cassidy-wails/internal/utils"
	_ "github.com/mattn/go-sqlite3"
)

const (
	sqliteTestDB = "/Users/josephcocozza/Repositories/unnamed-app/cassidy-wails/internal/test/cassidy-wails-test.db"
)

//var SQLiteDB *sql.DB

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
// Connect to a SQLite database.
func ConnectToSQLite(path string) (*sql.DB, error) {
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
// Create the test database.
func InitTestDB() *Database {
	os.Remove(sqliteTestDB)
	DB, err := ConnectToSQLite(sqliteTestDB)
	if err != nil {
		panic("unable to create test db: " + fmt.Sprint(err))
	}
	TableCreate(DB)
	RunInserts(DB)
	return &Database{DB: DB}
}