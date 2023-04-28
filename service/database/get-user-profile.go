package database

func (db *appdbimpl) GetUserProfile(userid uint64) (User, error) {

	var user User

	rows, err := db.c.Query(`SELECT id, username, follower, following, banned, photos FROM users WHERE id=?`, userid)
	if err != nil {
		return user, err
	}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.Follower, &user.Following, &user.Banned, &user.Photos)
		if err != nil {
			return user, err
		}
	}

	if user.ID != userid {
		return user, err
	}

	defer func() { _ = rows.Close() }()
	return user, nil

}
