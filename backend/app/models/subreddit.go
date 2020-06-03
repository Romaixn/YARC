package models

import (
	"database/sql"
	"errors"
	"regexp"
	"strings"

	"github.com/lib/pq"
)

const (
	subNameMinLen     = 1
	subNameMaxLen     = 16
	descriptionMaxLen = 512
)

// SubredditModel defines the database which the functions operate on.
type SubredditModel struct {
	DB *sql.DB
}

// ErrSubNameInvalid means the subreddit name contains invalid characters or is too long/short.
var ErrSubNameInvalid = errors.New("the subreddit name should be 1-16 alphanumerical or underscore characters")

// ErrSubNameDup means the subreddit name already exists in the database.
var ErrSubNameDup = errors.New("the subreddit name is used")

// ErrDescriptionInvalid means the description of the subreddit is too long.
var ErrDescriptionInvalid = errors.New("the description should be at most 512 characters")

// ErrSubredditNotExist means the subreddit doesn't exist in the database.
var ErrSubredditNotExist = errors.New("the subreddit does not exist")

var subnameRegExp = regexp.MustCompile("^[a-zA-Z0-9_]*$")

// SubredditInfo is the public info of a subreddit.
type SubredditInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Insert adds a new subreddit to the database.
func (m *SubredditModel) Insert(name, description string) error {
	// Validate the subreddit name.
	nameLen := len(name)
	if nameLen < subNameMinLen || nameLen > subNameMaxLen || !subnameRegExp.MatchString(name) {
		return ErrSubNameInvalid
	}

	// Validate the description.
	if len(description) > descriptionMaxLen {
		return ErrDescriptionInvalid
	}

	// Insert into database.
	stmt := `INSERT INTO subreddit (sub_name, description) VALUES ($1, $2)`
	_, err := m.DB.Exec(stmt, name, description)
	if err, ok := err.(*pq.Error); ok {
		if strings.Contains(err.Message, "sub_name_unique") {
			return ErrSubNameDup
		}
		return err
	}

	// Successfully inserted to the database.
	return nil
}

// Get selects a subreddit in the database and returns its info.
func (m *SubredditModel) Get(name string) (SubredditInfo, error) {
	subInfo := SubredditInfo{}

	stmt := `SELECT sub_name, description FROM subreddit WHERE sub_name = $1`
	row := m.DB.QueryRow(stmt, name)

	err := row.Scan(&subInfo.Name, &subInfo.Description)
	if err != nil {
		return subInfo, ErrSubredditNotExist
	}

	return subInfo, nil
}

// GetTrending returns a list of trending subreddits limited by a number.
func (m *SubredditModel) GetTrending(limit int) ([]SubredditInfo, error) {
	// TODO, probably use Redis.
	return []SubredditInfo{}, nil
}
