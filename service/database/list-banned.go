package database

func (db *appdbimpl) ListBanned(loggedUser uint64) ([]User, error) {
	var user_list []User

	rows, err := db.c.Query(`	SELECT u.id, u.username, u.follower, u.following, u.banned, u.photos  
								FROM users u, bans b
								WHERE b.userid=? and b.bannedid=u.id`, loggedUser)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Username, &user.Follower, &user.Following, &user.Banned, &user.Photos)

		if err != nil {
			return nil, err
		}
		user_list = append(user_list, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	defer func() { _ = rows.Close() }()
	return user_list, nil

}
