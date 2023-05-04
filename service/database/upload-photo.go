package database

func (db *appdbimpl) UploadPhoto(p Photo) (Photo, error) {

	_, err := db.c.Exec(`INSERT INTO photos (id, userid, picture, likes, date_time, comments) VALUES (NULL, ?, ?, ?, ?, ?)`,
		p.Comments, p.Picture, p.Likes, p.Date_time, p.Comments)
	if err != nil {
		return p, err
	}

	// Update photos value of the user
	_, err = db.c.Exec(`UPDATE users SET photos=photos+1 WHERE id=?`, p.User_id)
	if err != nil {
		return p, err
	}

	return p, nil
}
