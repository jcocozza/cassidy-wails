package database

import (
	"database/sql"
	"fmt"
	"log/slog"
)

// A transaction object object that implements DbOperations interface
type Transaction struct {
	TX *sql.Tx
}
// Execute sql in a transaction without returning any rows.
//
// Pass in args for placeholders in query.
//
// Execute does not commit/rollback the transaction.
// This is designed to be part of a longer chain of sql.
func (t *Transaction) Execute(sql string, args ...any) error {
	slog.Debug(runningSqlTransactionComponent + sql)
    slog.Debug("Args: " + fmt.Sprint(args...))
	_, err := t.TX.Exec(sql, args)
	if err != nil {
		return err
	}
	slog.Debug(completedSqlTransactionComponent)
	return nil
}
// Execute sql in a transaction and return the inserted row id.
//
// Return -1 for row id if there is an error
//
// Pass in args for placeholders in query.
//
// ExecuteGetLast does not commit/rollback the transaction.
// This is designed to be part of a longer chain of sql.
func (t *Transaction) ExecuteGetLast(sql string, args ...any) (int, error) {
	slog.Debug(runningSqlTransactionComponent + sql)
    slog.Debug("Args: " + fmt.Sprint(args...))

	result, err := t.TX.Exec(sql, args...)
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
// Execute sql in a transaction and return (several) rows.
//
// Pass in args for placeholders in query.
//
// Make sure to use defer rows.Close().
//
// Query does not commit/rollback the transaction.
// This is designed to be part of a longer chain of sql.
func (t *Transaction) Query(sql string, args ...any) (*sql.Rows, error) {
	slog.Debug(runningSqlTransactionComponent + sql)
    slog.Debug("Args: " + fmt.Sprint(args...))

	result, err := t.TX.Query(sql, args)
    if err != nil {
        slog.Error(fmt.Sprint(err))
        return nil, err
    }

	slog.Debug(completedSql)
	return result, nil
}
// Execute sql in a transaction is expected to return at most 1 row.
//
// Pass in args for placeholders in query.
//
// Errors are deferred until Row's Scan method is called.
// If the query selects no rows, the *Row's Scan will return ErrNoRows.
//
// QueryRow does not commit/rollback the transaction.
// This is designed to be part of a longer chain of sql.
func (t *Transaction) QueryRow(sql string, args ...any) *sql.Row {
	slog.Debug(runningSqlTransactionComponent + sql)
    slog.Debug("Args: " + fmt.Sprint(args...))

	result := t.TX.QueryRow(sql, args...)

	slog.Debug(completedSql)
	return result
}