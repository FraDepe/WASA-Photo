package database

func (db *appdbimpl) UploadPhoto(p Photo) (Photo, error) {

	_, err := db.c.Exec(`INSERT INTO photos (id, userid, picture, like, date_time, comments) VALUES (NULL, ?, ?, ?, ?, ?)`,
		p.Comments, p.Picture, p.Likes, p.Date_time, p.Comments)

	if err != nil {
		return p, err
	}

	return p, nil
}