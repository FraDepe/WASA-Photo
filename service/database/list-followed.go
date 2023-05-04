package database

func (db *appdbimpl) ListFollowed(userId uint64, loggedUser uint64) ([]User, error) {
	var user_list []User

	// Check if the guy who is asking for user list is following the user
	rows, err := db.c.Query(`SELECT followerid FROM follows WHERE followerid=? and followedid=?`, loggedUser, userId)
	if err != nil {
		return nil, err
	}
	var exist []uint64
	for rows.Next() {
		var id uint64
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		exist = append(exist, id)
	}

	// If exist array is not empty, the guy who's asking for user list, can receive it
	if len(exist) != 0 || userId == loggedUser {

		rows, err := db.c.Query(`	SELECT u.id, u.username, u.follower, u.following, u.banned, u.photos  
									FROM users u, follows f
									WHERE f.followerid=? and f.followedid=u.id`, userId)
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

	defer func() { _ = rows.Close() }()
	return user_list, nil
}
