package pgsql

import (
	"database/sql"
	"errors"
	"time"

	"github.com/phazon85/mangos-account-registration/pkg/acct"
)

const (
	activeAccount = `SELECT username FROM account WHERE username=?;`
	createEntry   = `INSERT INTO account (username, sha_pass_hash) VALUES (?, SHA1(CONCAT(UPPER(?),':',UPPER(?))));`
	updateEntry   = `UPDATE account SET sha_pass_hash=SHA1(CONCAT(UPPER(?),':',UPPER(?))), v=0, s=0 where username='?';`
)

var (
	// ErrAccountDoesNotExist ...
	ErrAccountDoesNotExist = errors.New("Account does not exist")
)

// Account ...
type Account struct {
	ID            string
	Username      string
	SHAPassHash   string
	GMLevel       int
	SessionKey    string
	V             string
	S             string
	Email         string
	JoinDate      time.Time
	LastIP        string
	FailedLogins  int
	Locked        int
	LastLogin     time.Time
	ActiveRealmID int
	Expansion     int
	MuteTime      int
	Locale        int
	OS            string
	PlayerBot     int
}

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

// CreateAccount ...
func (c *Client) CreateAccount(req *acct.CreateRequest) error {
	newAccount := newPlayerAccount(req.Name, req.Password)
	_, err := c.db.Exec(
		createEntry,
		newAccount.Username,
		newAccount.Username,
		req.Password,
	)
	if err != nil {
		return err
	}

	return nil
}

func newPlayerAccount(uname, pass string) *Account {
	return &Account{
		Username:    uname,
		SHAPassHash: pass,
		V:           "0",
		S:           "0",
	}
}

//ResetPassword ...
func (c *Client) ResetPassword(req *acct.CreateRequest) error {
	err := c.checkIfAccountExists(req.Name)
	if err != nil {
		return err
	}

	newAccount := newPlayerAccount(req.Name, req.Password)

	stmt, err := c.db.Prepare("UPDATE account SET sha_pass_hash=SHA1(CONCAT(UPPER(?),':',UPPER(?))), v=0, s=0 where username=?;")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(newAccount.Username, newAccount.SHAPassHash, newAccount.Username)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) checkIfAccountExists(name string) error {
	result := &Account{}
	row := c.db.QueryRow(activeAccount, name)
	if _ = row.Scan(&result.Username); result.Username != name {
		return ErrAccountDoesNotExist
	}
	return nil
}
