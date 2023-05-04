package database

func (db *appdbimpl) DoLogin(u string) (User, error) {

	var user User
	var avaible []uint64

	rows, err := db.c.Query(`SELECT id FROM users WHERE username=?`, u)
	if err != nil {
		return user, err
	}

	for rows.Next() {
		var tmp uint64
		err = rows.Scan(&tmp)
		avaible = append(avaible, tmp)
		if err != nil {
			return user, err
		}
	}

	if len(avaible) == 0 {
		user.Follower = 0
		user.Following = 0
		user.Banned = 0
		user.Photos = 0
		user.Username = u

		res, err := db.c.Exec(`INSERT INTO users (id, username, follower, following, banned, photos) VALUES (NULL, ?, ?, ?, ?, ?)`, u, user.Follower, user.Following, user.Banned, user.Photos)
		if err != nil {
			return user, err
		}

		signed_user_id, err := res.LastInsertId()
		if err != nil {
			return user, err
		}
		user.ID = uint64(signed_user_id)

		err = rows.Err()
		if err != nil {
			return user, err
		}

		defer func() { _ = rows.Close() }()
		return user, nil
	} else {

		rows, err := db.c.Query(`SELECT id, username, follower, following, banned, photos FROM users WHERE username=?`, u)
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
	}
}
