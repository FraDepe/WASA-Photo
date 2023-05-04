package database

func (db *appdbimpl) ShowPhoto(photoid uint64, userid uint64) (Photo, error) {
	var photo Photo

	// Get userid of the guy who uploaded the photo
	rows, err := db.c.Query(`SELECT userid FROM photos WHERE id=?`, photoid)
	if err != nil {
		return photo, err
	}
	var user_id_p uint64
	for rows.Next() {
		err = rows.Scan(&user_id_p)
		if err != nil {
			return photo, err
		}
	}

	// Check if the guy who is asking for the photo is banned or no
	rows, err = db.c.Query(`SELECT userid FROM bans WHERE userid=? and bannedid=?`, user_id_p, userid)
	if err != nil {
		return photo, err
	}
	var exist []uint64
	for rows.Next() {
		var id uint64
		err = rows.Scan(&id)
		if err != nil {
			return photo, err
		}
		exist = append(exist, id)
	}

	// If exist array is empty, the guy who's asking for the photo, can receive it
	if len(exist) == 0 {
		rows, err = db.c.Query(`SELECT id, userid, picture, likes, date_time FROM photos WHERE id=?`, photoid)
		if err != nil {
			return photo, err
		}

		for rows.Next() {
			err = rows.Scan(&photo.ID, &photo.User_id, &photo.Picture, &photo.Likes, &photo.Date_time, &photo.Comments)
			if err != nil {
				return photo, err
			}
		}

		if photo.Picture == nil {
			return photo, ErrPhotoNotFound
		}

		err = rows.Err()
		if err != nil {
			return photo, err
		}

		defer func() { _ = rows.Close() }()
		return photo, nil
	}

	err = rows.Err()
	if err != nil {
		return photo, err
	}

	defer func() { _ = rows.Close() }()
	return photo, nil
}
