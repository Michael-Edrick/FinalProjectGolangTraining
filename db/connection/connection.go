package connection

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "final_project"
)

func InitDatabase() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host= %s port= %d user= %s "+" password= %s dbname= %s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("%v\n", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("%v\n", err)
		return nil, err
	}
	log.Println("successfully Connect to Database")
	return db, nil
}
