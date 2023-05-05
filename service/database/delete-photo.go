package database

func (db *appdbimpl) DeletePhoto(id uint64, user_id uint64) error {

	rows, err := db.c.Query(`SELECT userid FROM photos WHERE id=?`, id)
	if err != nil {
		return err
	}

	var actual_user_id uint64

	for rows.Next() {
		err = rows.Scan(&actual_user_id)
		if err != nil {
			return err
		}
	}

	err = rows.Err()
	if err != nil {
		return err
	}

	if actual_user_id == user_id {
		res, err := db.c.Exec(`DELETE FROM photos WHERE id=?`, id)
		if err != nil {
			return err
		}

		affected, err := res.RowsAffected()
		if err != nil {
			return err
		} else if affected == 0 {
			return ErrPhotoNotFound
		}

		res, err = db.c.Exec(`DELETE FROM likes WHERE photoid=?`, id)
		if err != nil {
			return err
		}

		res, err = db.c.Exec(`DELETE FROM comments WHERE photoid=?`, id)
		if err != nil {
			return err
		}

		// Update photos value of the user
		_, err = db.c.Exec(`UPDATE users SET photos=photos-1 WHERE id=?`, user_id)

		if err != nil {
			return err
		}

		defer func() { _ = rows.Close() }()

		return nil
	}

	defer func() { _ = rows.Close() }()
	return nil

}
