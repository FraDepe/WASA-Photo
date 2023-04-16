package database

func (db *appdbimpl) GetFollowingStream(u User, id uint64) ([]Photo, error) {
	var stream_photo []Photo

	// Check if the guy who is asking for stream of photo is banned or no
	rows, err := db.c.Query(`SELECT userid FROM bans WHERE userid=? and bannedid=?`, u.ID, id)
	if err != nil {
		return nil, nil
	}
	var exist []uint64
	for rows.Next() {
		var id uint64
		err = rows.Scan(&id)
		exist = append(exist, id)
		if err != nil {
			return nil, nil
		}
	}

	// Check if the guy who is asking for stream of photo is following or no
	rows, err = db.c.Query(`SELECT userid FROM follows WHERE folloedid=? and followerid=?`, u.ID, id)
	if err != nil {
		return nil, nil
	}
	var isFollowing []uint64
	for rows.Next() {
		var id uint64
		err = rows.Scan(&id)
		exist = append(exist, id)
		if err != nil {
			return nil, nil
		}
	}

	// If exist array is empty the guy who's asking for stream of photo, can receive it
	if (len(exist) == 0) && (len(isFollowing) > 0) {
		rows, err = db.c.Query(`SELECT id, userid, picture, likes, date_time, comments FROM photos WHERE id=?`, u.ID)
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var ph Photo
			err = rows.Scan(&ph.ID, &ph.User_id, &ph.Picture, &ph.Likes, &ph.Date_time, &ph.Comments)

			if err != nil {
				return nil, err
			}
			stream_photo = append(stream_photo, ph)
		}
		if err = rows.Err(); err != nil {
			return nil, err
		}

		defer func() { _ = rows.Close() }()
		return nil, err

	}

	defer func() { _ = rows.Close() }()
	return stream_photo, nil
}
