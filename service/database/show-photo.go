package database

func (db *appdbimpl) ShowPhoto(id string) (Photo, error) {
	var photo Photo

	rows, err := db.c.Query(`SELECT id, picture, likes, date_time FROM photos WHERE id=?`, id)
	if err != nil {
		return photo, err
	}

	for rows.Next() {
		err = rows.Scan(&photo.ID, &photo.Picture, &photo.Likes, &photo.Date_time, &photo.Comments) // niente comments???????
		if err != nil {
			return photo, err
		}
	}

	if photo.Picture == nil {
		return photo, ErrPhotoNotFound
	}

	return photo, nil
}
