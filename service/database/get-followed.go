package database

func (db *appdbimpl) GetFollowed(userIdFollowed uint64, userId uint64) (User, error) {

	var user User
	user.ID = 0
	user.Username = ""
	user.Follower = 0
	user.Following = 0
	user.Banned = 0
	user.Photos = 0

	rows, err := db.c.Query(`SELECT u.id, u.username, u.follower, u.following, u.banned, u.photos FROM users u, follows f WHERE u.id=? and f.followedid=u.id and f.followerid=?`, userIdFollowed, userId)
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
