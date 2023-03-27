package database

func (db *appdbimpl) GetUserProfile(username string) (User, error) {

	var user User

	rows, err := db.c.Query(`SELECT id, username, follower, following, banned, photos FROM users WHERE username=?`, username)
	if err != nil {
		return user, err
	}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.Follower, &user.Following, &user.Banned, &user.Photos) // niente photos e banned????????
		if err != nil {
			return user, err
		}
	}

	if user.Username != username {
		return user, err
	}

	return user, nil

}
