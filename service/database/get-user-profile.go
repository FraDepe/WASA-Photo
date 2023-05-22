package database

import "wasaphoto.uniroma1.it/wasaphoto/service/utils"

func (db *appdbimpl) GetUserProfile(username string, loggedUser uint64) (User, error) {

	var user User

	// Fetch infos of the profile
	rows, err := db.c.Query(`SELECT id, username, follower, following, banned, photos FROM users WHERE username=?`, username)
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
		return user, utils.ErrUserDoesNotExist
	}

	err = rows.Err()
	if err != nil {
		return user, err
	}

	// Check if the guy is banned or no
	if db.isBanned(user.ID, loggedUser) {
		return user, utils.ErrBanned

	}

	defer func() { _ = rows.Close() }()
	return user, nil

}
