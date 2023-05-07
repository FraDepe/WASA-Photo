package database

import "wasaphoto.uniroma1.it/wasaphoto/service/utils"

func (db *appdbimpl) ListFollowed(userId uint64, loggedUser uint64) ([]User, error) {
	var user_list []User

	// If who's asking is following can receive it
	if db.isFollowing(userId, loggedUser) || userId == loggedUser {

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

		err = rows.Err()
		if err != nil {
			return nil, err
		}

		defer func() { _ = rows.Close() }()
		return user_list, nil
	} else {

		return user_list, utils.ErrMustFollow
	}
}
