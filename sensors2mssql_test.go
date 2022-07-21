package sensors2mssql

import (
	"fmt"
	"sync"
	"testing"

	"github.com/blablatov/scada4sensors2mssql/mssqldsn"
	"github.com/blablatov/scada4sensors2mssql/smssqlinsert"
)

const (
	dBaseName    = "mssqlserver"
	SensorType   = "dallas"
	InsertReqSql = "INSERT " + dBaseName + " VALUES (SensorType, " + SensorType + ")"
)

func Test(t *testing.T) {
	var tests = []struct {
		InsertReqSql string
		SelectReqSql string
		UpdateReqSql string
		Create       string
	}{
		{"INSERT " + dBaseName + " VALUES (SensorType, " + SensorType + ")",
			"SELECT Name FROM TableDB WHERE BusUnit = 65;",
			"UPDATE TableDB SET Id=Id, Name=Name, BusUnit=BusUnit, ItemNumber=ItemNumber, ItemName=ItemName SELECT Id, Name, BusUnit, ItemNumber, ItemName FROM AnyTable WHERE BusUnit = 65;",
			"CREATE TABLE MyTable(Id int, Name nvarchar(250), BusUnit int, ItemNumber nvarchar(50), ItemName nvarchar(100));"},
	}

	var prevInsertReqSql string
	for _, test := range tests {
		if test.InsertReqSql != prevInsertReqSql {
			fmt.Printf("\n%s\n", test.InsertReqSql)
			prevInsertReqSql = test.InsertReqSql
		}
	}

	var prevSelectReqSql string
	for _, test := range tests {
		if test.SelectReqSql != prevSelectReqSql {
			fmt.Printf("\n%s\n", test.SelectReqSql)
			prevSelectReqSql = test.SelectReqSql
		}
	}

	var prevUpdateReqSql string
	for _, test := range tests {
		if test.UpdateReqSql != prevUpdateReqSql {
			fmt.Printf("\n%s\n", test.UpdateReqSql)
			prevUpdateReqSql = test.UpdateReqSql
		}
	}

	var prevCreate string
	for _, test := range tests {
		if test.Create != prevCreate {
			fmt.Printf("\n%s\n", test.Create)
			prevCreate = test.Create
		}
	}
}

func BenchmarkInterfaceDsn(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < 3; i++ {
		dd := mssqldsn.DataDsn{
			Debug:    true,
			User:     "user",
			Password: "password",
			Port:     1433,
			Server:   "dbserver",
			Database: "sensordb",
		}
		var d mssqldsn.ConDsner = dd
		db := d.SqlConDsn()
		defer db.Close()
		fmt.Println("DSN is: ", db)
	}
}

func BenchmarkGoroutineSqlInserTrs(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < 3; i++ {
		// Structure DSN. Структура DSN.
		dd := mssqldsn.DataDsn{
			Debug:    true,
			User:     "user",
			Password: "password",
			Port:     1433,
			Server:   "dbserver",
			Database: "sensordb",
		}
		var d mssqldsn.ConDsner = dd
		db := d.SqlConDsn()
		cs := make(chan string) // Send result to the main programm. Передача результата в main-программу.
		var wg sync.WaitGroup
		go smssqlinsert.SqlInserTrs(InsertReqSql, db, cs, wg)
		// Getting data from goroutine. Получение данных из канала горутины.
		fmt.Println("\nResponse of goroutine: ", <-cs)
		go func() {
			wg.Wait()
			close(cs)
		}()
	}
}
