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

var ErrUserDoesNotExist = errors.New("user does not exist")
var ErrInternalServer = errors.New("internal server error")
var ErrPhotoNotFound = errors.New("photo not found")
var ErrLikeNotFound = errors.New("like not found")

// Fountain struct represent a fountain in every API call between this package and the outside world.
// Note that the internal representation of fountain in the database might be different.
// type Fountain struct {
//	ID        uint64
//	Latitude  float64
//	Longitude float64
//	Status    string
// }

type User struct {
	ID        uint64
	Username  string
	Follower  uint64
	Following uint64
	Banned    uint64
	Photos    uint64
}

type Photo struct {
	ID        uint64
	User_id   uint64
	Picture   []byte
	Likes     uint64
	Date_time string
	Comments  uint64
}

type Comment struct {
	ID      uint64
	PhotoId uint64
	Text    string
	UserId  uint64
}

type Like struct {
	PhotoId uint64
	UserId  uint64
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	DoLogin(string) (User, error)
	//GetFollowingStream(User, uint64) ([]Photo, error)
	GetMyStream(User) ([]Photo, error)
	SetMyUsername(string, uint64) error
	UploadPhoto(Photo) (Photo, error)
	ShowPhoto(uint64, uint64) (Photo, error)
	DeletePhoto(uint64, uint64) error
	GetUserProfile(uint64) (User, error)
	ListComments(uint64, uint64) ([]Comment, error)
	CommentPhoto(Comment) (Comment, error)
	GetComment(uint64) (Comment, error)
	// ModifyComment(string, string) error
	UncommentPhoto(uint64, uint64) error
	ListLikes(uint64, uint64) ([]Like, error)
	LikePhoto(Like) (Like, error)
	UnlikePhoto(uint64, uint64) error
	ListFollowed(uint64, uint64) ([]User, error)
	FollowUser(uint64, uint64) (User, error) // Perchè torno un user?
	UnfollowUser(uint64, uint64) error
	ListBanned(uint64) ([]User, error)
	BanUser(uint64, uint64) (User, error) // Perchè torno un user?
	UnbanUser(uint64, uint64) error

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

	// USER
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='users';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE users (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL UNIQUE,
			follower INTEGER NOT NULL,
			following INTEGER NOT NULL,
			banned INTEGER NOT NULL,
			photos INTEGER NOT NULL
			);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	// PHOTOS
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='photos';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE photos (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			userid INTEGER NOT NULL,
			picture BLOB NOT NULL,
			likes INTEGER NOT NULL,
			date_time TEXT NOT NULL,
			comments INTEGER NOT NULL,
			FOREIGN KEY (userid) REFERENCES users(id)
			);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	// COMMENTS
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='comments';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE comments (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			photoid INTEGER NOT NULL,
			text TEXT NOT NULL,
			userid INTEGER NOT NULL,
			FOREIGN KEY (photoid) REFERENCES photos(id),
			FOREIGN KEY (userid) REFERENCES users(id)
			);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	// LIKES
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='likes';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE likes (
			photoid INTEGER NOT NULL,
			userid INTEGER NOT NULL,
			FOREIGN KEY (photoid) REFERENCES photos(id),
			FOREIGN KEY (userid) REFERENCES users(id),
			PRIMARY KEY("photoid", "userid")
			);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	// FOLLOWS
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='follows';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE follows (
			followerid INTEGER NOT NULL,
			followedid INTEGER NOT NULL,
			FOREIGN KEY (followerid) REFERENCES users(id),
			FOREIGN KEY (followedid) REFERENCES users(id),
			PRIMARY KEY("followerid", "followedid")
			);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	// BANS
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='bans';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE bans (
			userid INTEGER NOT NULL,
			bannedid INTEGER NOT NULL,
			FOREIGN KEY (userid) REFERENCES users(id),
			FOREIGN KEY (bannedid) REFERENCES users(id),
			PRIMARY KEY("userid", "bannedid")
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
