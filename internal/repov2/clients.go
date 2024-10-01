package repov2

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	_ "modernc.org/sqlite"
)

type Client struct {
	connStr string
	Driver  *sql.DB
}

//func New(path string) (*Client, error) {
//	client := &Client{
//		connStr: path,
//	}
//
//	if err := client.Conn(); err != nil {
//		return nil, err
//	}
//
//	if err := client.InitDb(); err != nil {
//		return nil, fmt.Errorf("fail to init db: %s", err.Error())
//	}
//
//	return client, nil
//}

func Conn(connStr string) (*sql.DB, error) {
	//if c.Driver != nil {
	//	if err := c.Driver.Ping(); err == nil {
	//		return nil
	//	}
	//}

	db, err := sql.Open("sqlite3", connStr)
	if err != nil {
		return nil, fmt.Errorf("fail to open db: %s", err.Error())
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("fail to connect db: %s", err.Error())
	}

	if _, err := db.Exec(`PRAGMA foreign_keys=ON`); err != nil {
		return nil, err
	}

	return db, nil
}

func (c *Client) Close() {
	if c.Driver != nil {
		_ = c.Driver.Close()
	}
}
