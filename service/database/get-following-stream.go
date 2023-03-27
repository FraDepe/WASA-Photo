package database

func (db *appdbimpl) GetFollowingStream(u User) ([]Photo, error) {
	var stream_photo []Photo

	rows, err := db.c.Query(`SELECT id, picture, likes, date_time, comments FROM photos WHERE id=?`, u.ID)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var ph Photo
		err = rows.Scan(&ph.ID, &ph.Picture, &ph.Likes, &ph.Date_time, &ph.Comments)

		if err != nil {
			return nil, err
		}
		stream_photo = append(stream_photo, ph)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return nil, err
}
