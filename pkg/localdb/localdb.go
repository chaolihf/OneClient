package localdb

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"github.com/kardianos/service"
)

var dbConnection *sql.DB
var logger service.Logger

func InitDb(service_logger service.Logger) {
	logger = service_logger
	var err error
	dbConnection, err = sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.Error(err)
	}
	sqlStmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`
	_, err = dbConnection.Exec(sqlStmt)
	if err != nil {
		logger.Errorf("%q: %s\n", err, sqlStmt)
		return
	}
}

func CloseDb() {
	dbConnection.Close()

}
func PingNetwork() {
	tx, err := dbConnection.Begin()
	if err != nil {
		logger.Error(err)
	}
	stmt, err := tx.Prepare("insert into foo(id, name) values(?, ?)")
	if err != nil {
		logger.Error(err)
	}
	defer stmt.Close()
	for i := 0; i < 100; i++ {
		_, err = stmt.Exec(i, fmt.Sprintf("test %d", i))
		if err != nil {
			logger.Error(err)
		}
	}
	err = tx.Commit()
	if err != nil {
		logger.Error(err)
	}
}

func GetMonitorData() {
	dbTransaction, err := dbConnection.Begin()
	if err != nil {
		logger.Error(err)
	}
	dbCommand, err := dbTransaction.Prepare("select count(1) counter from foo")
	if err != nil {
		logger.Error(err)
	}
	defer dbCommand.Close()
	var count int
	dbCommand.QueryRow().Scan(&count)
	dbTransaction.Commit()
	logger.Info(fmt.Sprintf("count: %d", count))
}
