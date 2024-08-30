package sqllite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	_ "modernc.org/sqlite"
)

type Client struct {
	connStr string
	Driver  *sql.DB

	UsersRepo   *UsersRepo
	CirclesRepo *CirclesRepo
	LabelsRepo  *AreasRepo
}

func New(path string) (*Client, error) {
	client := &Client{
		connStr: path,
	}

	if err := client.conn(); err != nil {
		return nil, err
	}

	if err := client.InitDb(); err != nil {
		return nil, fmt.Errorf("fail to init db: %s", err.Error())
	}

	client.UsersRepo = NewUsersRepo(client.Driver)
	client.CirclesRepo = NewCirclesRepo(client.Driver)
	client.LabelsRepo = NewAreasRepo(client.Driver)

	return client, nil
}

func (c *Client) conn() error {
	if c.Driver != nil {
		if err := c.Driver.Ping(); err == nil {
			return nil
		}
	}

	db, err := sql.Open("sqlite3", c.connStr)
	if err != nil {
		return fmt.Errorf("fail to open db: %s", err.Error())
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("fail to connect db: %s", err.Error())
	}

	if _, err := db.Exec(`PRAGMA foreign_keys=ON`); err != nil {
		return err
	}

	c.Driver = db

	return nil
}

func (c *Client) Close() {
	if c.Driver != nil {
		_ = c.Driver.Close()
	}
}
