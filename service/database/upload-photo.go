package database

func (db *appdbimpl) UploadPhoto(p Photo) (Photo, error) {

	p.Comments = "[]" // Si risolverebbe dividendo i commenti dalle foto e legasrli con foreign key

	_, err := db.c.Exec(`INSERT INTO photos (id, picture, like, date_time, comments) VALUES (NULL, ?, ?, ?, ?)`,
		p.Picture, p.Likes, p.Date_time, p.Comments)

	if err != nil {
		return p, err
	}

	return p, nil
}
