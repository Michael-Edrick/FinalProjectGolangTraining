package connection

import (
	"FinalProject/utils"
	"database/sql"
	"fmt"
	"log"
)

func InitDatabase(config *utils.Config) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host= %s port= %s user= %s "+" password= %s dbname= %s sslmode=disable", config.PostgresHost, config.PostgresPort, config.PostgresUser, config.PostgresPassword, config.PostgresName)
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
