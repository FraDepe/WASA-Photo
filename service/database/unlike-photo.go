package database

func (db *appdbimpl) UnlikePhoto(u_id uint64, p_id uint64) error {

	res, err := db.c.Exec(`DELETE FROM likes WHERE photoid=?, userid=?`,
		u_id, p_id)

	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return ErrLikeNotFound
	}

	return nil
}
