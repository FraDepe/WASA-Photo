package database

func (db *appdbimpl) CommentPhoto(c Comment) (Comment, error) {

	user_id := c.UserId
	photo_id := c.PhotoId

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

	// Check if the guy who is commenting is banned or no
	rows, err = db.c.Query(`SELECT userid FROM bans WHERE userid=? and bannedid=?`, user_id_p, user_id)
	if err != nil {
		return c, err
	}
	var exist []uint64
	for rows.Next() {
		var id uint64
		err = rows.Scan(&id)
		if err != nil {
			return c, err
		}
		exist = append(exist, id)
	}

	err = rows.Err()
	if err != nil {
		return c, err
	}

	// If exist array is empty the guy who's commenting is not banned and so he can comments
	if len(exist) == 0 {
		_, err := db.c.Exec(`INSERT INTO comments (id, photoid, text, userid) VALUES (NULL, ?, ?, ?)`,
			c.PhotoId, c.Text, c.UserId)

		if err != nil {
			return c, err
		}

		// Update comment value of the photo
		_, err = db.c.Exec(`UPDATE photos SET comments=comments+1 WHERE id=?`, photo_id)
		if err != nil {
			return c, err
		}

		defer func() { _ = rows.Close() }()
		return c, nil
	}

	defer func() { _ = rows.Close() }()

	return c, nil
}
