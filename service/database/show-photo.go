package database

import "wasaphoto.uniroma1.it/wasaphoto/service/utils"

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

	err = rows.Err()
	if err != nil {
		return photo, err
	}

	// Check if the guy who is asking for the photo is banned or no
	if !db.isBanned(user_id_p, userid) {

		rows, err = db.c.Query(`SELECT id, userid, picture, likes, date_time, comments FROM photos WHERE id=?`, photoid)
		if err != nil {
			return photo, err
		}

		for rows.Next() {
			err = rows.Scan(&photo.ID, &photo.User_id, &photo.Picture, &photo.Likes, &photo.Date_time, &photo.Comments)
			if err != nil {
				return photo, err
			}
		}

		err = rows.Err()
		if err != nil {
			return photo, err
		}

		if photo.Picture == nil {
			return photo, utils.ErrPhotoNotFound
		}

		defer func() { _ = rows.Close() }()
		return photo, nil
	} else {
		defer func() { _ = rows.Close() }()
		return photo, utils.ErrBanned
	}

}
