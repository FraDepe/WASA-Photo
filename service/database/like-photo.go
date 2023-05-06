package database

func (db *appdbimpl) LikePhoto(l Like) (Like, error) {

	if !db.existence(l.UserId) {
		return l, ErrUserDoesNotExist
	}

	// Get userid of the guy who uploaded the photo
	rows, err := db.c.Query(`SELECT userid FROM photos WHERE id=?`, l.PhotoId)
	if err != nil {
		return l, err
	}
	var user_id_p uint64
	for rows.Next() {
		err = rows.Scan(&user_id_p)
		if err != nil {
			return l, err
		}
	}

	err = rows.Err()
	if err != nil {
		return l, err
	}

	if user_id_p == 0 {
		defer func() { _ = rows.Close() }()
		return l, ErrPhotoNotFound
	}

	// Check if the guy who is liking is banned or no
	if !db.isBanned(user_id_p, l.UserId) {
		_, err = db.c.Exec(`INSERT INTO likes (photoid, userid) VALUES (?, ?)`,
			l.PhotoId, l.UserId)
		if err != nil {
			return l, err
		}

		// Update likes value of the photo
		_, err = db.c.Exec(`UPDATE photos SET likes=likes+1 WHERE id=?`, l.PhotoId)
		if err != nil {
			return l, err
		}
	}

	defer func() { _ = rows.Close() }()
	return l, nil
}
