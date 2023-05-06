package database

func (db *appdbimpl) CommentPhoto(c Comment) (Comment, error) {

	user_id := c.UserId
	photo_id := c.PhotoId

	if !db.existence(user_id) {
		return c, ErrUserDoesNotExist
	}

	// Get userid of the guy who uploaded the photo
	rows, err := db.c.Query(`SELECT userid FROM photos WHERE id=?`, photo_id)
	if err != nil {
		return c, err
	}
	var user_id_p uint64
	for rows.Next() {
		if err != nil {
			return c, err
		}
		err = rows.Scan(&user_id_p)
	}

	err = rows.Err()
	if err != nil {
		return c, err
	}

	if user_id_p == 0 {
		return c, ErrPhotoNotFound
	}

	// Check if the guy who is commenting is banned or no
	if !db.isBanned(user_id_p, user_id) {
		res, err := db.c.Exec(`INSERT INTO comments (id, photoid, text, userid) VALUES (NULL, ?, ?, ?)`,
			c.PhotoId, c.Text, c.UserId)

		if err != nil {
			return c, err
		}

		comment_id, err := res.LastInsertId()

		// Update comment value of the photo
		_, err = db.c.Exec(`UPDATE photos SET comments=comments+1 WHERE id=?`, photo_id)
		if err != nil {
			return c, err
		}

		c.ID = uint64(comment_id)

		defer func() { _ = rows.Close() }()
		return c, nil
	} else {
		defer func() { _ = rows.Close() }()
		return c, ErrBanned
	}

}
