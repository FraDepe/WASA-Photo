package database

func (db *appdbimpl) DeletePhoto(id string) error {

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
