package pgsql

import (
	"crypto/sha1"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/phazon85/mangos-account-registration/pkg/acct"
)

const (
	activeAccount = `SELECT username FROM account WHERE username=?;`
	createEntry   = `INSERT INTO account (username, sha_pass_hash) VALUES (?, SHA1(CONCAT(UPPER(?),':',UPPER(?))));`
	updateEntry   = `UPDATRE account SET ?, v=0, s=0 where username='?';`
)

var (
	// ErrAccountExists ...
	ErrAccountExists = errors.New("Account already exists")
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
	sha := createSHA(req.Name, req.Password)
	newAccount := newPlayerAccount(req.Name, sha)

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

func createSHA(uname, pass string) string {
	hash := sha1.New()
	io.WriteString(hash, fmt.Sprintf("%s:%s", strings.ToUpper(uname), strings.ToUpper(pass)))
	sha := base64.URLEncoding.EncodeToString(hash.Sum(nil))

	return sha
}

func newPlayerAccount(uname, shaPass string) *Account {
	return &Account{
		Username:    uname,
		SHAPassHash: shaPass,
		V:           "0",
		S:           "0",
	}
}

//ResetPassword ...
func (c *Client) ResetPassword(name string) error {
	err := c.checkIfAccountExists(name)
	if err != nil {
		return err
	}

	_, err = c.db.Exec(updateEntry, createSHA(name, "Test"), name)

	return nil
}

func (c *Client) checkIfAccountExists(name string) error {
	result := &Account{}
	row := c.db.QueryRow(activeAccount, name)
	if err := row.Scan(result.Username); err == nil && result.Username != name {
		return ErrAccountExists
	}
	return nil
}
