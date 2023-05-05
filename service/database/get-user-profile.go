package database

func (db *appdbimpl) GetUserProfile(userid uint64, loggedUser uint64) (User, error) {

	var user User

	// Check if the guy is banned or no
	rows, err := db.c.Query(`SELECT userid FROM bans WHERE bannedid=? and userid=?`, loggedUser, userid)
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

	err = rows.Err()
	if err != nil {
		return user, err
	}

	if len(exist) != 0 { // You are banned

		defer func() { _ = rows.Close() }()
		return user, ErrBanned

	}

	// Fetch infos of the profile
	rows, err = db.c.Query(`SELECT id, username, follower, following, banned, photos FROM users WHERE id=?`, userid)
	if err != nil {
		return user, err
	}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.Follower, &user.Following, &user.Banned, &user.Photos)
		if err != nil {
			return user, err
		}
	}

	if user.ID == 0 {
		return user, ErrUserDoesNotExist
	}

	err = rows.Err()
	if err != nil {
		return user, err
	}

	if user.ID != userid {
		return user, nil
	}

	defer func() { _ = rows.Close() }()
	return user, nil

}
