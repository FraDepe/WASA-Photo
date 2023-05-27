package database

func (db *appdbimpl) GetUserProfileStream(username string, loggedUser uint64) ([]Photo, error) {

	var stream_photo []Photo

	rows, err := db.c.Query(`SELECT p.id, p.userid, p.picture, p.likes, p.date_time, p.comments FROM photos p, users u 
							 WHERE ?=u.username and u.id=p.userid ORDER BY p.id DESC`, username)
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
