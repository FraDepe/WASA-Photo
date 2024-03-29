package database

func (db *appdbimpl) GetMyStream(user uint64) ([]Photo, error) {
	var stream_photo []Photo

	rows, err := db.c.Query(`SELECT id, userid, picture, likes, date_time, comments FROM photos p, follows f 
							 WHERE ?=f.followerid and f.followedid=p.userid ORDER BY id DESC`, user)
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

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	defer func() { _ = rows.Close() }()
	return stream_photo, nil
}
