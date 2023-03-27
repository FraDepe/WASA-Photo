/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

var ErrUserDoesNotExist = errors.New("User does not exist")
var ErrInternalServer = errors.New("Internal server error")
var ErrPhotoNotFound = errors.New("Photo not found")

// Fountain struct represent a fountain in every API call between this package and the outside world.
// Note that the internal representation of fountain in the database might be different.
//type Fountain struct {
//	ID        uint64
//	Latitude  float64
//	Longitude float64
//	Status    string
//}

type User struct {
	ID        string
	Username  string
	Follower  float64
	Following float64
	Banned    float64
	Photos    float64
}

type Photo struct {
	ID        string
	Picture   []byte
	Likes     int
	Date_time string
	Comments  []Comment
}

type Comment struct {
	ID     string
	Text   string
	UserId string
}

type Like struct {
	PhotoId string
	UserId  string
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	GetFollowingStream(User) ([]Photo, error)
	GetMyStream(User) ([]Photo, error)
	SetMyUsername(string, string) error
	UploadPhoto(Photo) (Photo, error)
	ShowPhoto(string) (Photo, error)
	DeletePhoto(string) error
	GetUserProfile(string) (User, error)
	ListComments(string) ([]Comment, error)
	CommentPhoto(string) (Comment, error)
	GetComment(string, string) (Comment, error)
	ModifyComment(string, string) error
	UncommentPhoto(string, string) error
	ListLikes(string) ([]Like, error)
	LikePhoto(string, string) (Like, error)
	UnlikePhoto(string, string) error
	ListFollowed(string) ([]string, error)     //forse ritorna tanti user boh
	FollowerUser(string, string) (User, error) //porcata????
	UnfollowUser(string, string) error
	ListBanned(string) ([]string, error) //forse ritorna tanti user boh
	BanUser(string, string) (User, error)
	UnbanUser(string, string) error

	// Ping checks whether the database is available or not (in that case, an error will be returned)
	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='users';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE users (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL,
			follower FLOAT NOT NULL,
			following FLOAT NOT NULL,
			banned FLOAT NOT NULL,
			photos FLOAT NOT NULL
			);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='photos';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE photos (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			picture BLOB NOT NULL,
			likes FLOAT NOT NULL,
			date_time TEXT NOT NULL
			);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='comments';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE comments (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			text TEXT NOT NULL,
			userid INTEGER NOT NULL
			);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='likes';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE likes (
			photoid INTEGER NOT NULL,
			userid INTEGER NOT NULL,
			FOREIGN KEY (photoid) REFERENCES photos(id),
			FOREIGN KEY (userid) REFERENCES users(id)
			);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
