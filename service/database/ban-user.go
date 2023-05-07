package database

import "wasaphoto.uniroma1.it/wasaphoto/service/utils"

func (db *appdbimpl) BanUser(logged_user uint64, user_id uint64) (User, error) {

	var user User

	// Check if the guy exists
	if !db.existence(user_id) {
		return user, utils.ErrUserDoesNotExist
	}

	// Check if the guy is already banned
	if !db.isBanned(logged_user, user_id) {
		_, err := db.c.Exec(`INSERT INTO bans (userid, bannedid) VALUES (?, ?)`,
			logged_user, user_id)

		if err != nil {
			return user, err
		}

		// Update banned value of the user who's banning
		_, err = db.c.Exec(`UPDATE users SET banned=banned+1 WHERE id=?`, logged_user)

		if err != nil {
			return user, err
		}

		rows, err := db.c.Query(`SELECT id, username, follower, following, banned, photos FROM users WHERE id=?`, user_id)
		if err != nil {
			return user, err
		}
		for rows.Next() {
			err = rows.Scan(&user.ID, &user.Username, &user.Follower, &user.Following, &user.Banned, &user.Photos)
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

	} else {
		return user, utils.ErrUserAlreadyBanned
	}

}
