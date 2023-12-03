package catsql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sql.DB
	filename string
}

func OpenCatDB() (*DB, error) {
	dbTempFile, err := os.CreateTemp("", "catdb")
	if err != nil {
		return nil, err
	}
	dbTempFile.Close()
	db, err := sql.Open("sqlite3", dbTempFile.Name())
	if err != nil {
		return nil, err
	}
	createItemTable := `CREATE TABLE item (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,	
		"name" TEXT NOT NULL,
		"description" TEXT NOT NULL,
		"price" INTEGER NOT NULL
	);`

	statement, err := db.Prepare(createItemTable)
	if err != nil {
		return nil, err
	}
	if _, err := statement.Exec(); err != nil {
		return nil, err
	}

	if _, err := db.Exec(`INSERT INTO item (id, name, description, price) VALUES(1, "Scrumpy Numplings","Scrumpy Numplings are stanky, and kitties love stanky things!", 9);`); err != nil {
		return nil, fmt.Errorf("couldn't create sample item: %w", err)
	}

	return &DB{
		DB:       db,
		filename: dbTempFile.Name(),
	}, nil
}

func (c *DB) Close() error {
	if err := c.DB.Close(); err != nil {
		return err
	}
	return os.Remove(c.filename)
}
