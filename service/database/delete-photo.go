package database

func (db *appdbimpl) DeletePhoto(id uint64, user_id uint64) error {

	rows, err := db.c.Query(`SELECT userid FROM photos WHERE id=?`, id)
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
		return nil
	}

	// Forse dovrei gestire il caso in cui gli id non
	// corrispondono

	defer func() { _ = rows.Close() }()
	return nil

}