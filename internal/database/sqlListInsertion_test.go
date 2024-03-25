package database

import (
	"testing"
)

func TestSQLListInsertion(t *testing.T) {
	type args struct {
		sql              string
		numElementInList int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test", args{sql: "SELECT * FROM my_tbl WHERE uuid IN (%s)", numElementInList: 4}, "SELECT * FROM my_tbl WHERE uuid IN (?, ?, ?, ?)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SQLListInsertion(tt.args.sql, tt.args.numElementInList); got != tt.want {
				t.Errorf("SQLListInsertion() = %v, want %v", got, tt.want)
			}
		})
	}
}
