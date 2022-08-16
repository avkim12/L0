package postgres

import "fmt"

var (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "123"
	dbname   = "postgres"

	psqlInfo  = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)
)