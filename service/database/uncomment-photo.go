package database

func (db *appdbimpl) UncommentPhoto(id uint64, userid uint64) error {

	rows, err := db.c.Query(`SELECT userid FROM comments WHERE id=?`, id)
	if err != nil {
		return nil
	}

	var actual_user_id uint64

	for rows.Next() {
		err = rows.Scan(&actual_user_id)
		if err != nil {
			return nil
		}
	}

	if actual_user_id == userid {
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

		return nil
	}
	return nil
}
