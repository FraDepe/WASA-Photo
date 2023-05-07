package database

import "wasaphoto.uniroma1.it/wasaphoto/service/utils"

func (db *appdbimpl) UnlikePhoto(u_id uint64, p_id uint64) error {

	if !db.existence(u_id) {
		return utils.ErrUserDoesNotExist
	}

	res, err := db.c.Exec(`DELETE FROM likes WHERE photoid=? and userid=?`,
		p_id, u_id)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return utils.ErrLikeNotFound
	}

	// Update likes value of the photo
	_, err = db.c.Exec(`UPDATE photos SET likes=likes-1 WHERE id=?`, p_id)
	if err != nil {
		return err
	}

	return nil
}
