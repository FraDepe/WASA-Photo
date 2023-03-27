package api

import "git.sapienzaapps.it/gamificationlab/wasa-fontanelle/service/database"

// Fountain struct represent a fountain in every data exchange with the external world via REST API. JSON tags have been
// added to the struct to conform to the OpenAPI specifications regarding JSON key names.
// Note: there is a similar struct in the database package. See Fountain.FromDatabase (below) to understand why.
type User struct {
	ID        string  `json:"id"`
	Username  string  `json:"username"`
	Follower  float64 `json:"follower"`
	Following float64 `json:"following"`
	Banned    float64 `json:"banned"`
	Photos    float64 `json:"photos"`
}

// FromDatabase populates the struct with data from the database, overwriting all values.
// You might think this is code duplication, which is correct. However, it's "good" code duplication because it allows
// us to uncouple the database and API packages.
// Suppose we were using the "database.Fountain" struct inside the API package; in that case, we were forced to conform
// either the API specifications to the database package or the other way around. However, very often, the database
// structure is different from the structure of the REST API.
// Also, in this way the database package is freely usable by other packages without the assumption that structs from
// the database should somehow be JSON-serializable (or, in general, serializable).
func (u *User) FromDatabase(user database.User) {
	u.ID = user.ID
	u.Username = user.Username
	u.Follower = user.Follower
	u.Following = user.Following
	u.Banned = user.Banned
	u.Photos = user.Photos
}

// ToDatabase returns the fountain in a database-compatible representation
func (u *User) ToDatabase() database.User {
	return database.User{
		ID:        u.ID,
		Username:  u.Username,
		Follower:  u.Follower,
		Following: u.Following,
		Banned:    u.Banned,
		Photos:    u.Photos,
	}
}

// IsValid checks the validity of the content. In particular, coordinates should be in their range of validity, and the
// status should be either FountainStatusGood or FountainStatusFaulty. Note that the ID is not checked, as fountains
// read from requests have zero IDs as the user won't send us the ID in that way.
//func (f *Fountain) IsValid() bool {
//	return -90 <= f.Latitude && f.Latitude <= 90 &&
//		-180 <= f.Longitude && f.Longitude <= 180 &&
//		(f.Status == FountainStatusGood || f.Status == FountainStatusFaulty)
//}

type Photo struct {
	ID        string    `json:"id"`
	Picture   []byte    `json:"picture"`
	Likes     int       `json:"likes"`
	Date_time string    `json:"date_time"`
	Comments  []Comment `json:"comments"`
}

func (p *Photo) FromDatabase(photo database.Photo) {
	p.ID = photo.ID
	p.Picture = photo.Picture
	p.Likes = photo.Likes
	p.Date_time = photo.Date_time
	p.Comments = photo.Comments
}

func (p *Photo) ToDatabase() database.Photo {
	return database.Photo{
		ID:        p.ID,
		Picture:   p.Picture,
		Likes:     p.Likes,
		Date_time: p.Date_time,
		Comments:  p.Comments,
	}
}

type Comment struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	UserId string `json:"userid"`
}

func (c *Comment) FromDatabase(comment database.Comment) {
	c.ID = comment.ID
	c.Text = comment.Text
	c.UserId = comment.UserId
}

func (c *Comment) ToDatabase() database.Comment {
	return database.Comment{
		ID:     c.ID,
		Text:   c.Text,
		UserId: c.UserId,
	}
}

type Like struct {
	PhotoId string `json:"photoid"`
	UserId  string `json:"userid"`
}

func (l *Like) FromDatabase(like database.Like) {
	l.PhotoId = like.PhotoId
	l.UserId = like.UserId
}

func (l *Like) ToDatabase() database.Like {
	return database.Like{
		PhotoId: l.PhotoId,
		UserId:  l.UserId,
	}
}
