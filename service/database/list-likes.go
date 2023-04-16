package database

func (db *appdbimpl) ListLikes(photoid uint64, userid uint64) ([]Like, error) {
	var stream_like []Like

	// Get userid of the guy who uploaded the photo
	rows, err := db.c.Query(`SELECT userid FROM photos WHERE id=?`, photoid)
	if err != nil {
		return nil, nil
	}
	var user_id_p uint64
	for rows.Next() {
		err = rows.Scan(&user_id_p)
		if err != nil {
			return nil, nil
		}
	}

	// Check if the guy who is asking for stream of like is the same who uploaded the photo
	if user_id_p == userid {
		rows, err := db.c.Query(`SELECT photoid, userid FROM likes WHERE photoid=?`, photoid)
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var like Like
			err = rows.Scan(&like.PhotoId, &like.UserId)

			if err != nil {
				return nil, err
			}
			stream_like = append(stream_like, like)
		}
		if err = rows.Err(); err != nil {
			return nil, err
		}

		defer func() { _ = rows.Close() }()
		return stream_like, nil
	}

	defer func() { _ = rows.Close() }()
	return stream_like, nil
}
