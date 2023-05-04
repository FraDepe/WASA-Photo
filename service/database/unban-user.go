package database

func (db *appdbimpl) UnbanUser(logged_user uint64, user_id uint64) error {

	// Check if the guy is banned or no
	rows, err := db.c.Query(`SELECT userid FROM bans WHERE bannedid=? and userid=?`, user_id, logged_user)
	if err != nil {
		return err
	}
	var exist []uint64
	for rows.Next() {
		var id uint64
		err = rows.Scan(&id)
		if err != nil {
			return err
		}
		exist = append(exist, id)
	}

	// If exist array is not empty the guy is banned so we can unban him
	if len(exist) != 0 {
		_, err = db.c.Exec(`DELETE FROM bans WHERE bannedid=? and userid=?`,
			user_id, logged_user)
		if err != nil {
			return err
		}

		// Update banned value of the user who's unbanning
		_, err = db.c.Exec(`UPDATE users SET banned=banned-1 WHERE id=?`, logged_user)
		if err != nil {
			return err
		}
	}

	defer func() { _ = rows.Close() }()
	return nil
}
