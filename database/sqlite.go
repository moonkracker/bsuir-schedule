package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	Db *sql.DB
}

func NewDatabase(path string) (*Database, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	return &Database{Db: db}, nil
}

func (db *Database) GetDatabase() *sql.DB {
	return db.Db
}

func (db *Database) CloseDatabase() error {
	return db.Db.Close()
}

func (db *Database) CreateTable(tableName string, columns map[string]string) error {
	fields := ""
	for k, v := range columns {
		fields += fmt.Sprintf("%s %s, ", k, v)
	}
	fields = fields[:len(fields)-2]
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s)", tableName, fields)
	_, err := db.Db.Exec(query)
	return err
}

func (db *Database) InsertInTable(tableName string, columns map[string]string) error {
	var fields, values string
	for k, v := range columns {
		fields += fmt.Sprintf("%s, ", k)
		values += fmt.Sprintf("%s, ", v)
	}
	fields = fields[:len(fields)-2]
	values = values[:len(values)-2]
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tableName, fields, values)
	_, err := db.Db.Exec(query)
	return err
}

func (db *Database) Select(tableName string, columns []string, where string) ([]string, error) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s", columns[0], tableName, where)
	for _, v := range columns[1:] {
		query += fmt.Sprintf(", %s", v)
	}
	rows, err := db.Db.Query(query)
	if err != nil {
		return nil, err
	}
	var result []string
	for rows.Next() {
		var value string
		err = rows.Scan(&value)
		if err != nil {
			return nil, err
		}
		result = append(result, value)
	}
	return result, nil
}

func (db *Database) Update(tableName string, columns []string, values []string, where string) error {
	query := fmt.Sprintf("UPDATE %s SET %s = %s WHERE %s", tableName, columns[0], values[0], where)
	for _, v := range columns[1:] {
		query += fmt.Sprintf(", %s = %s", v, values[0])
	}
	_, err := db.Db.Exec(query)
	return err
}

func (db *Database) Delete(tableName string, where string) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s", tableName, where)
	_, err := db.Db.Exec(query)
	return err
}

func (db *Database) DropTable(tableName string) error {
	query := fmt.Sprintf("DROP TABLE %s", tableName)
	_, err := db.Db.Exec(query)
	return err
}

func (db *Database) CreateIndex(tableName string, column string) error {
	query := fmt.Sprintf("CREATE INDEX %s_%s ON %s(%s)", tableName, column, tableName, column)
	_, err := db.Db.Exec(query)
	return err
}
