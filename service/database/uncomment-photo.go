package database

func (db *appdbimpl) UncommentPhoto(id uint64, userid uint64) error {

	// Query to get userid of user who commented
	rows, err := db.c.Query(`SELECT userid FROM comments WHERE id=?`, id)
	if err != nil {
		return err
	}

	var comment_user_id uint64

	for rows.Next() {
		err = rows.Scan(&comment_user_id)
		if err != nil {
			return err
		}
	}

	err = rows.Err()
	if err != nil {
		return err
	}

	// Query to get photoid who got commented
	rows, err = db.c.Query(`SELECT photoid FROM comments WHERE id=?`, id)
	if err != nil {
		return err
	}

	var photoid uint64

	for rows.Next() {
		err = rows.Scan(&photoid)
		if err != nil {
			return err
		}
	}

	err = rows.Err()
	if err != nil {
		return err
	}

	if comment_user_id == userid {
		res, err := db.c.Exec(`DELETE FROM comments WHERE id=?`, id)
		if err != nil {
			return err
		}

		affected, err := res.RowsAffected()
		if err != nil {
			return err
		} else if affected == 0 {
			return ErrPhotoNotFound
		}

		err = rows.Err()
		if err != nil {
			return err
		}

		// Update comment value of the photo
		_, err = db.c.Exec(`UPDATE photos SET comments=comments-1 WHERE id=?`, photoid)

		if err != nil {
			return err
		}

		return nil
	}

	return nil
}
