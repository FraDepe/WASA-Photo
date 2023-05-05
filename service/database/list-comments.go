package database

func (db *appdbimpl) ListComments(photoid uint64, userid uint64) ([]Comment, error) {
	var stream_comments []Comment

	// Get userid of the guy who uploaded the photo
	rows, err := db.c.Query(`SELECT userid FROM photos WHERE id=?`, photoid)
	if err != nil {
		return nil, err
	}
	var user_id_p uint64
	for rows.Next() {
		err = rows.Scan(&user_id_p)
		if err != nil {
			return nil, err
		}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	// Check if the guy who is asking for stream of comment is banned or no
	rows, err = db.c.Query(`SELECT userid FROM bans WHERE userid=? and bannedid=?`, user_id_p, userid)
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

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	// If exist array is empty, the guy who's asking for stream of comment, can receive it
	if len(exist) == 0 {
		rows, err := db.c.Query(`SELECT id, photoid, text, userid FROM comments WHERE photoid=?`, photoid)
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			var comment Comment
			err = rows.Scan(&comment.ID, &comment.PhotoId, &comment.Text, &comment.UserId)

			if err != nil {
				return nil, err
			}
			stream_comments = append(stream_comments, comment)
		}

		err = rows.Err()
		if err != nil {
			return nil, err
		}

		defer func() { _ = rows.Close() }()
		return stream_comments, nil
	}

	defer func() { _ = rows.Close() }()
	return stream_comments, nil
}
