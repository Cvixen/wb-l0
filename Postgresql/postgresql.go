package Postgresql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "test"
	password = "test"
	dbname   = "wb-l0"
)

/**
Подключение к базе данных
*/
func NewPostgresDB() *sqlx.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sqlx.Open("postgres", psqlconn)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
