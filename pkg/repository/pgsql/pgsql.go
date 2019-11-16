package pgsql

import "database/sql"

// Client ...
type Client struct {
	db *sql.DB
}

// New ...
func New(db *sql.DB) *Client {
	return &Client{
		db: db,
	}
}

// Register ...
func (c *Client) Register() {}
