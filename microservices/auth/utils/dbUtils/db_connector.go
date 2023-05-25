package dbUtils

import (
	"database/sql"
	"fmt"
)

func NewDBConnector(connectionString string) *DBConnector {
	return &DBConnector{connectionString: connectionString}
}

type DBConnector struct {
	connectionString string
	db               *sql.DB
}

func (c *DBConnector) Open() error {
	db, err := sql.Open("sqlserver", c.connectionString)
	if err != nil {
		return err
	}
	c.db = db
	fmt.Printf("Connected to DB: %s\n", c.connectionString)
	return nil
}

func (c *DBConnector) GetDB() (*sql.DB, error) {
	if c.db == nil {
		return nil, fmt.Errorf("DB not opened; connection string is '%s'\n", c.connectionString)
	}
	return c.db, nil
}

func (c *DBConnector) Close() error {
	if err := c.db.Close(); err != nil {
		return err
	}
	fmt.Printf("Disctonnected from DB\n")
	return nil
}
