package api

import "wasaphoto.uniroma1.it/wasaphoto/service/database"

// Fountain struct represent a fountain in every data exchange with the external world via REST API. JSON tags have been
// added to the struct to conform to the OpenAPI specifications regarding JSON key names.
// Note: there is a similar struct in the database package. See Fountain.FromDatabase (below) to understand why.
type User struct {
	ID        uint64 `json:"id"`
	Username  string `json:"username"`
	Follower  uint64 `json:"follower"`
	Following uint64 `json:"following"`
	Banned    uint64 `json:"banned"`
	Photos    uint64 `json:"photos"`
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
// func (f *Fountain) IsValid() bool {
//	return -90 <= f.Latitude && f.Latitude <= 90 &&
//		-180 <= f.Longitude && f.Longitude <= 180 &&
//		(f.Status == FountainStatusGood || f.Status == FountainStatusFaulty)
// }

type Photo struct {
	ID        uint64 `json:"id"`
	User_id   uint64 `json:"user_id"`
	Picture   []byte `json:"picture"`
	Likes     uint64 `json:"likes"`
	Date_time string `json:"date_time"`
	Comments  uint64 `json:"comments"`
}

func (p *Photo) FromDatabase(photo database.Photo) {
	p.ID = photo.ID
	p.User_id = photo.User_id
	p.Picture = photo.Picture
	p.Likes = photo.Likes
	p.Date_time = photo.Date_time
	p.Comments = photo.Comments
}

func (p *Photo) ToDatabase() database.Photo {
	return database.Photo{
		ID:        p.ID,
		User_id:   p.User_id,
		Picture:   p.Picture,
		Likes:     p.Likes,
		Date_time: p.Date_time,
		Comments:  p.Comments,
	}
}

type Comment struct {
	ID      uint64 `json:"id"`
	PhotoId uint64 `json:"photoid"`
	Text    string `json:"text"`
	UserId  uint64 `json:"userid"`
}

func (c *Comment) FromDatabase(comment database.Comment) {
	c.ID = comment.ID
	c.PhotoId = comment.PhotoId
	c.Text = comment.Text
	c.UserId = comment.UserId
}

func (c *Comment) ToDatabase() database.Comment {
	return database.Comment{
		ID:      c.ID,
		PhotoId: c.PhotoId,
		Text:    c.Text,
		UserId:  c.UserId,
	}
}

type Like struct {
	PhotoId uint64 `json:"photoid"`
	UserId  uint64 `json:"userid"`
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

type Follow struct {
	FollowerId uint64 `json:"followerid"`
	FollowedId uint64 `json:"followedid"`
}

func (f *Follow) FromDatabase(follow database.Follow) {
	f.FollowerId = follow.FollowerId
	f.FollowedId = follow.FollowedId
}

func (f *Follow) ToDatabase() database.Follow {
	return database.Follow{
		FollowerId: f.FollowerId,
		FollowedId: f.FollowedId,
	}
}

type Ban struct {
	UserId   uint64 `json:"userid"`
	BannedId uint64 `json:"bannedid"`
}

func (b *Ban) FromDatabase(ban database.Ban) {
	b.UserId = ban.UserId
	b.BannedId = ban.BannedId
}

func (b *Ban) ToDatabase() database.Ban {
	return database.Ban{
		UserId:   b.UserId,
		BannedId: b.BannedId,
	}
}
