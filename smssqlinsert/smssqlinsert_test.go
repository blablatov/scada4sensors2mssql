package smssqlinsert

import (
	"fmt"
	"testing"
)

const (
	dBaseName  = "mssqlserver"
	SensorType = "dallas_1"
)

func Test(t *testing.T) {
	var tests = []struct {
		insertIntegrTableSql string
	}{
		{"INSERT " + dBaseName + " VALUES (SensorType, " + SensorType + ")"},
		{"INSERT dBase VALUES (SensorType, dallas_2)"},
	}

	var previnsertIntegrTableSql string
	for _, test := range tests {
		if test.insertIntegrTableSql != previnsertIntegrTableSql {
			fmt.Printf("\n%s\n", test.insertIntegrTableSql)
			previnsertIntegrTableSql = test.insertIntegrTableSql
		}
	}
}
