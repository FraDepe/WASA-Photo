package database

func (db *appdbimpl) FollowUser(logged_user uint64, user_id uint64) (User, error) {

	var user User

	// Check if the guy who is following is banned or no
	rows, err := db.c.Query(`SELECT userid FROM bans WHERE userid=? and bannedid=?`, user_id, logged_user)
	if err != nil {

		return user, err
	}
	var exist []uint64
	for rows.Next() {
		var id uint64
		err = rows.Scan(&id)
		if err != nil {
			return user, err
		}
		exist = append(exist, id)
	}

	// If exist array is empty the guy who's following is not banned and so he can follows
	if len(exist) == 0 {
		_, err = db.c.Exec(`INSERT INTO follows (followerid, followedid) VALUES (?, ?)`,
			logged_user, user_id)

		if err != nil {
			return user, err
		}

		// Update followed e follower values of the users
		_, err = db.c.Exec(`UPDATE users SET following=following+1 WHERE id=?`, logged_user)
		if err != nil {
			return user, err
		}
		_, err = db.c.Exec(`UPDATE users SET follower=follower+1 WHERE id=?`, user_id)
		if err != nil {
			return user, err
		}

		err = rows.Err()
		if err != nil {
			return user, err
		}

	}

	err = rows.Err()
	if err != nil {
		return user, err
	}

	defer func() { _ = rows.Close() }()
	return user, nil
}
