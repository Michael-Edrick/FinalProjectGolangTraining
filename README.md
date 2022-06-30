This is a Golang project using PostgreSQL for the database and Postman for testing

## Getting Started

First, create a PostgreSQL database :

```bash
CREATE DATABASE databaseName;
```

or using GUI to create the database.

Create an `.env` by copying the format from file `.env.example` , you can set the value of `DB_USER`, `DB_PASSWORD`, and `DB_NAME` as you like.

After that, using migrate command in your terminal
to create all tables :

```bash
migrate -database "postgres://[DB_USER]:[DB_PASSWORD]@[HOST]:[PORT]/[DB_NAME]?sslmode=disable" -path db/migrations up
```

or this command to delete all tables

```bash
migrate -database "postgres://[DB_USER]:[DB_PASSWORD]@[HOST]:[PORT]/[DB_NAME]?sslmode=disable" -path db/migrations down
```

After creating all tables, `go` run `main.go` and open [http://localhost:8088](http://localhost:8088) with your Postman app.
