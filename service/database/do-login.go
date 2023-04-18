package database

func (db *appdbimpl) DoLogin(u string) (User, error) {

	var user User
	var avaible uint64

	rows, err := db.c.Query(`SELECT IFNULL(id, 0) FROM users WHERE username=?`, u)
	if err != nil {
		return user, err
	}

	for rows.Next() {
		err = rows.Scan(&avaible)
		if err != nil {
			return user, err
		}
	}

	if avaible == 0 {
		user.Follower = 0
		user.Following = 0
		user.Banned = 0
		user.Photos = 0

		res, err := db.c.Exec(`INSERT INTO users (id, username, follower, following, banned, photos) VALUES (NULL, ?, ?, ?, ?, ?)`, u, user.Follower, user.Following, user.Banned, user.Photos)
		if err != nil {
			return user, err
		}

		signed_user_id, err := res.LastInsertId()
		if err != nil {
			return user, err
		}
		user.ID = uint64(signed_user_id)

		defer func() { _ = rows.Close() }()
		return user, nil
	} else {
		user.Username = u

		defer func() { _ = rows.Close() }()
		return user, nil
	}
}
