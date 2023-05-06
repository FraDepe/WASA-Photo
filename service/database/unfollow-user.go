package database

func (db *appdbimpl) UnfollowUser(logged_user uint64, user_id uint64) error {

	// If exist array is not empty the guy is following so he can unfollow
	if db.isFollowing(user_id, logged_user) {
		_, err := db.c.Exec(`DELETE FROM follows WHERE followedid=? and followerid=?`,
			user_id, logged_user)
		if err != nil {
			return err
		}

		// Update followed e follower values of the users
		_, err = db.c.Exec(`UPDATE users SET following=following-1 WHERE id=?`, logged_user)
		if err != nil {
			return err
		}
		_, err = db.c.Exec(`UPDATE users SET follower=follower-1 WHERE id=?`, user_id)
		if err != nil {
			return err
		}
		return nil

	} else {
		return ErrMustFollow
	}
}
