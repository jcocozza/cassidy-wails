package database

import (
	"database/sql"
	"fmt"
	"log/slog"
)

// A database object that implements DbOperations interface
type Database struct {
	DB *sql.DB
}
// Execute sql without returning rows.
//
// Pass in args for placeholders in query.
func (d *Database) Execute(sql string, args ...any) error {
	slog.Debug(runningSql + sql)
    slog.Debug("Args: " + fmt.Sprint(args...))

	_, err := d.DB.Exec(sql, args...)
	if err != nil {
		return err
	}
	slog.Debug(completedSql)
	return nil
}
// Execute sql and return the inserted row id.
//
// Return -1 for row id if there is an error
//
// Pass in args for placeholders in query.
func (d *Database) ExecuteGetLast(sql string, args ...any) (int, error) {
	slog.Debug(runningSql + sql)
    slog.Debug("Args: " + fmt.Sprint(args...))

	result, err := d.DB.Exec(sql, args...)
    if err != nil {
        slog.Error(fmt.Sprint(err))
		return -1, err
    }

    lastRow, err2 := result.LastInsertId()
    if err2 != nil {
        slog.Error(fmt.Sprint(err))
		return -1, err
    }
	slog.Debug(completedSql)
    return int(lastRow), nil
}
// Execute sql and return (several) rows.
//
// Pass in args for placeholders in query.
//
// Make sure to use defer rows.Close().
func (d *Database) Query(sql string, args ...any) (*sql.Rows, error) {
	slog.Debug(runningSql + sql)
    slog.Debug("Args: " + fmt.Sprint(args...))

    result, err := d.DB.Query(sql, args...)
    if err != nil {
        slog.Error(fmt.Sprint(err))
        return nil, err
    }

    slog.Debug(completedSql)
    return result, nil
}
// Execute sql that is expected to return at most 1 row.
//
// Errors are deferred until Row's Scan method is called.
// If the query selects no rows, the *Row's Scan will return ErrNoRows.
func (d *Database) QueryRow(sql string, args ...any) *sql.Row {
	slog.Debug(runningSql + sql)
    slog.Debug("Args: " + fmt.Sprint(args...))

	result := d.DB.QueryRow(sql, args...)

	slog.Debug(completedSql)
    return result
}
