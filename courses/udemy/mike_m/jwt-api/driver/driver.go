package driver

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
)

func logFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// ConnectDB stablishes a database connection
func ConnectDB() *sql.DB {
	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	logFatal(err)

	db, err := sql.Open("postgres", pgUrl)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	return db
}
