// Module do formating and transmit data to MSSQL module for write to DB
// Формирует и передает данные MSSQL-модулю для записи в БД.
package sensors2mssql

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/blablatov/scada4sensors2mssql/mssqldsn"
	"github.com/blablatov/scada4sensors2mssql/smssqlinsert"
)

type SensMssqler interface {
	SensMssql() bool
}

type ReqOperators struct {
	InsertReqSql string
	SelectReqSql string
	UpdateReqSql string
	Create       string
}

func (rq ReqOperators) SensMssql() bool {
	cs := make(chan string) // Channel of function sql-client. Канал функции sql-клиента
	// Structure DSN. Структура DSN.
	dd := mssqldsn.DataDsn{
		Debug:    true,
		User:     "user",
		Password: "password",
		Port:     1433,
		Server:   "dbserver",
		Database: "sensordb",
	}
	// Calling an interface method for formating DSN and connect to the DB.
	// Вызов метода интерфейса, для формирования DSN и создания подключения к СУБД
	var d mssqldsn.ConDsner = dd
	db := d.SqlConDsn()
	defer db.Close()

	start := time.Now()
	// Calling the data copy method in the goroutine. Вызов метода копирования данных в горутине.
	var wg sync.WaitGroup // Synchronization of goroutines. Синхронизация горутин.
	go smssqlinsert.SqlInserTrs(rq.InsertReqSql, db, cs, wg)
	// Getting data from channel cs goroutine. Получение данных из канала cs горутины
	log.Println("\nResult of request Insert to MSSQL via goroutine: ", <-cs)
	// Wait of counter. Ожидание счетчика
	go func() {
		wg.Wait()
		close(cs)
	}()
	secs := time.Since(start).Seconds()
	fmt.Printf("%.2fs Request execution time Insert via goroutine\n", secs)
	return true
}
