package database

import "database/sql"
// I don't know if I'll actually use this at all. I'm very hesitant to write mocks for my database calls.

// MockDBOperations is a mock implementation of the database.DbOperations interface for testing purposes.
type MockDBOperations struct {
	ExecuteFn func(sql string, args ...any) error
	ExecuteGetLastfn func(sql string, args ...any) (int, error)
	Queryfn func(sql string, args ...any) (*sql.Rows, error)
	QueryRowFn func(sql string, args ...any) *sql.Row
}

func (m *MockDBOperations) Execute(sql string, args ...any) error {
	if m.ExecuteFn != nil {
		return m.ExecuteFn(sql, args...)
	}
	return nil
}
func (m *MockDBOperations) ExecuteGetLast(sql string, args ...any) (int, error) {
	if m.ExecuteGetLastfn != nil {
		return m.ExecuteGetLastfn(sql, args...)
	}
	return -1, nil
}
func (m *MockDBOperations) Query(sql string, args ...any) (*sql.Rows, error) {
	if m.Queryfn != nil {
		return m.Queryfn(sql, args...)
	}
	return nil, nil
}
func (m *MockDBOperations) QueryRow(query string, args ...any) *sql.Row {
	if m.QueryRowFn != nil {
		return m.QueryRowFn(query, args...)
	}
	return nil
}