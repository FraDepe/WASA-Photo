package database

func (db *appdbimpl) UnbanUser(logged_user uint64, user_id uint64) error {

	// Check if the guy exists
	if !db.existence(user_id) {
		return ErrUserDoesNotExist
	}

	// Check if the guy is banned or no
	if db.isBanned(logged_user, user_id) {
		_, err := db.c.Exec(`DELETE FROM bans WHERE bannedid=? and userid=?`,
			user_id, logged_user)
		if err != nil {
			return err
		}

		// Update banned value of the user who's unbanning
		_, err = db.c.Exec(`UPDATE users SET banned=banned-1 WHERE id=?`, logged_user)
		if err != nil {
			return err
		}

		return nil

	} else {

		return ErrUserNotFound
	}

}
