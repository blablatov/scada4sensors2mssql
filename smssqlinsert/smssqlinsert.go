// Demo MSSQL module for writing data to a table per transaction.
// Демо MSSQL-модуль для записи данных в таблицу за транзакцию.

package smssqlinsert

import (
	"database/sql"
	"log"
	"sync"
)

func SqlInserTrs(InsertReqSql string, db *sql.DB, cs chan string, wg sync.WaitGroup) {
	defer wg.Done()
	// Starting the transaction. Начало транзакции
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	// The rollback will be ignored if the tx has been committed later in the function.
	// Откат действий будет проигнорирован, если позже выполнение транзакции будет зафиксировано.
	defer tx.Rollback()

	stmt, err := tx.Prepare(InsertReqSql)
	if err != nil {
		log.Fatal(err)
	}
	// Prepared statements take up server resources and should be closed after use.
	// Операторы занимают ресурсы сервера и должны быть закрыты, после выполнения.
	defer stmt.Close()

	if _, err := stmt.Exec("open source"); err != nil {
		log.Fatal(err)
	}
	// Finishing the transaction. Завершение транзакции
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
	// Send result to the main programm. Передача результата в main-программу.
	cs <- "Data was written to the DB"
}
