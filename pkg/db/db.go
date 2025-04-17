package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

var db *sql.DB

const (
	schema = `
	CREATE TABLE scheduler (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		date CHAR(8) NOT NULL DEFAULT,
		title TEXT NOT NULL DEFAULT,
		comment TEXT,
		repeat VARCHAR(128) NOT NULL DEFAULT
	);
	
	CREATE INDEX idx_date ON scheduler(date);
	`
)

func Init(dbFile string) error {

	fmt.Printf("Путь к файлу базы данных: %s\n", dbFile)

	_, err := os.Stat(dbFile)

	var install bool
	if err != nil {
		install = true
	}

	db, err = sql.Open("sqlite", dbFile)
	if err != nil {
		return err
	}
	defer db.Close()

	if install {
		_, err = db.Exec(schema)
		if err != nil {
			return err
		}
		fmt.Println("Файл базы данных создан")
	} else {
		fmt.Println("Файл базы данных уже существует")
	}

	return nil
}
