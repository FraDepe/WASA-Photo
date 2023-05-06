package database

func (db *appdbimpl) FollowUser(logged_user uint64, user_id uint64) (User, error) {

	var user User

	// Check if the users to follow exists
	if !db.existence(user_id) {
		return user, ErrUserDoesNotExist
	}

	// Check if guy is banned before to follow
	if !db.isBanned(user_id, logged_user) {
		_, err := db.c.Exec(`INSERT INTO follows (followerid, followedid) VALUES (?, ?)`,
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

		return user, nil
	} else {

		return user, ErrBanned
	}
}
