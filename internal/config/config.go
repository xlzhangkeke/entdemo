package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

var MySqlDataSource = `root:pandora@tcp(10.95.84.100:33306)/ent_examples?parseTime=True`

func Init() {
	db, err := sql.Open("mysql", MySqlDataSource)
	if err != nil {
		log.Fatalf("failed openning connection to mysql: %+v", err)
	}
	dropDatabase := `drop database if exists ent_examples`
	_, err = db.Exec(dropDatabase)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed exec drop database"))
	}
	createDatabase := `create database if not exists ent_examples`
	_, err = db.Exec(createDatabase)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed create database"))
	}
}
