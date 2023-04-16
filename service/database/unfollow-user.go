package database

func (db *appdbimpl) UnfollowUser(logged_user uint64, user_id uint64) error {

	// Check if the guy is following or no
	rows, err := db.c.Query(`SELECT userid FROM follows WHERE followedid=? and followerid=?`, user_id, logged_user)
	if err != nil {
		return nil
	}
	var exist []uint64
	for rows.Next() {
		var id uint64
		err = rows.Scan(&id)
		exist = append(exist, id)
		if err != nil {
			return nil
		}
	}

	// If exist array is not empty the guy is following so he can unfollow
	if len(exist) != 0 {
		_, err = db.c.Exec(`DELETE FROM follows WHERE followedid=? and followerid=?`,
			user_id, logged_user)

		if err != nil {
			return err
		}
	}

	defer func() { _ = rows.Close() }()
	return nil
}
