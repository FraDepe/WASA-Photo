package database

func (db *appdbimpl) GetLike(photoId uint64, userId uint64) (Like, error) {

	var like Like
	like.PhotoId = 0
	like.UserId = 0

	rows, err := db.c.Query(`SELECT photoid, userid FROM likes WHERE photoid=? and userid=?`, photoId, userId)
	if err != nil {
		return like, err
	}

	for rows.Next() {
		err = rows.Scan(&like.PhotoId, &like.UserId)
		if err != nil {
			return like, err
		}
	}

	err = rows.Err()
	if err != nil {
		return like, err
	}

	defer func() { _ = rows.Close() }()
	return like, nil

}
