package database

import "database/sql"

const (
	runningSql = "********** Running sql: **********\n"
	completedSql = "********** Sql completed **********"

	runningSqlTransactionComponent = "********** Running sql transaction component: **********\n"
	completedSqlTransactionComponent = "********** Sql transaction component completed **********"
)

var BaseDir string = "/Users/josephcocozza/Repositories/unnamed-app/backend"

// Methods for interacting with a database.
type DbOperations interface {
	// Execute sql without returning rows.
	//
	// Pass in args for placeholders in query.
	Execute(sql string, args ...any) error
	// Execute sql and return the inserted row id.
	// Return -1 for row id if there is an error.
	//
	// Pass in args for placeholders in query.
	ExecuteGetLast(sql string, args ...any) (int, error)
	// Execute sql and return (several) rows.
	//
	// Pass in args for placeholders in query.
	//
	// Make sure to use defer rows.Close().
	Query(sql string, args ...any) (*sql.Rows, error)
	// Execute sql that is expected to return at most 1 row.
	//
	// Errors are deferred until Row's Scan method is called.
	// If the query selects no rows, the *Row's Scan will return ErrNoRows.
	QueryRow(sql string, args ...any) *sql.Row
}