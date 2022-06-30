This is a Golang project using PostgreSQL for database and Postman for testing

## Getting Started

First, create a PostgreSQL database :

```bash
CREATE DATABASE databaseName;
```

or use the GUI to create a database.

Create `.env` file by copying the format from the `.env.example` file, you can set `DB_USER`, `DB_PASSWORD` and `DB_NAME` values ​​as you like.

After that, use the migration command in your terminal
to create all tables:

```bash
migrate -database "postgres://[DB_USER]:[DB_PASSWORD]@[HOST]:[PORT]/[DB_NAME]?sslmode=disable" -path db/migrations up
```

or this command to delete all tables

```bash
migrate -database "postgres://[DB_USER]:[DB_PASSWORD]@[HOST]:[PORT]/[DB_NAME]?sslmode=disable" -path db/migrations down
```

After creating all tables, `go` run `main.go` and open [http://localhost:8088](http://localhost:8088) with your Postman application.

Import `FinalProject.postman_collection.json` file into your Postman application and use the collection to try to hit the endpoint.
