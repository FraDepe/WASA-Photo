package database

func (db *appdbimpl) LikePhoto(l Like) (Like, error) {

	// Get userid of the guy who uploaded the photo
	rows, err := db.c.Query(`SELECT userid FROM photos WHERE id=?`, l.PhotoId)
	if err != nil {
		return l, nil
	}
	var user_id_p uint64
	for rows.Next() {
		err = rows.Scan(&user_id_p)
		if err != nil {
			return l, nil
		}
	}

	// Check if the guy who is liking is banned or no
	rows, err = db.c.Query(`SELECT userid FROM bans WHERE userid=? and bannedid=?`, user_id_p, l.UserId)
	if err != nil {
		return l, nil
	}
	var exist []uint64
	for rows.Next() {
		var id uint64
		err = rows.Scan(&id)
		if err != nil {
			return l, nil
		}
		exist = append(exist, id)
	}

	// If exist array is empty the guy who's liking is not banned and so he can likes
	if len(exist) == 0 {
		_, err = db.c.Exec(`INSERT INTO likes (photoid userid) VALUES (?, ?)`,
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
