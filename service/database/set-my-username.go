package database

import "wasaphoto.uniroma1.it/wasaphoto/service/utils"

func (db *appdbimpl) SetMyUsername(n_u string, id uint64) error {
	res, err := db.c.Exec(`UPDATE users SET username=? WHERE id=?`, n_u, id)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return utils.ErrUserDoesNotExist
	}
	return nil
}
