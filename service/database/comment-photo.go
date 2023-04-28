package database

func (db *appdbimpl) CommentPhoto(c Comment) (Comment, error) {

	user_id := c.UserId
	photo_id := c.PhotoId

	// Get userid of the guy who uploaded the photo
	rows, err := db.c.Query(`SELECT userid FROM photos WHERE id=?`, photo_id)
	if err != nil {
		return c, nil
	}
	var user_id_p uint64
	for rows.Next() {
		if err != nil {
			return c, nil
		}
		err = rows.Scan(&user_id_p)
	}

	// Check if the guy who is commenting is banned or no
	rows, err = db.c.Query(`SELECT userid FROM bans WHERE userid=? and bannedid=?`, user_id_p, user_id)
	if err != nil {
		return c, nil
	}
	var exist []uint64
	for rows.Next() {
		var id uint64
		err = rows.Scan(&id)
		if err != nil {
			return c, nil
		}
		exist = append(exist, id)
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
