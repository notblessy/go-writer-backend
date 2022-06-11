package configs

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func CreateConnection() (*sqlx.DB, error) {
	dsnString := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?multiStatements=true&parseTime=true",
		os.Getenv("RDBMS_USERNAME"),
		os.Getenv("RDBMS_PASSWORD"),
		os.Getenv("RDBMS_HOST")+":"+os.Getenv("RDBMS_PORT"),
		os.Getenv("RDBMS_DATABASE"),
	)

	db, errConn := sqlx.Open("mysql", dsnString)
	if errConn != nil {
		return nil, errConn
	}

	errPing := db.Ping()
	if errPing != nil {
		return nil, errConn
	}

	return db, nil
}
