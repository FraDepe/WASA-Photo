package database

func (db *appdbimpl) GetFollowed(userIdFollowed uint64, userId uint64) (Follow, error) {

	var follow Follow
	follow.FollowerId = 0
	follow.FollowedId = 0

	rows, err := db.c.Query(`SELECT followerid, followedid FROM follows WHERE followerid=? and followedid=?`, userId, userIdFollowed)
	if err != nil {
		return follow, err
	}

	for rows.Next() {
		err = rows.Scan(&follow.FollowerId, &follow.FollowedId)
		if err != nil {
			return follow, err
		}
	}

	err = rows.Err()
	if err != nil {
		return follow, err
	}

	defer func() { _ = rows.Close() }()
	return follow, nil

}
