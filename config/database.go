package config

import (
	"database/sql"
	"log"
)

type Database struct {
	*sql.DB
}

func (d *Database) Db() (*Database, error) {
	db, err := sql.Open("mysql", "root:@/laravel_blog")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &Database{db}, nil
}
