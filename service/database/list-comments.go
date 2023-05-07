package database

import "wasaphoto.uniroma1.it/wasaphoto/service/utils"

func (db *appdbimpl) ListComments(photoid uint64, userid uint64) ([]Comment, error) {
	var stream_comments []Comment

	if !db.existence(userid) {
		return nil, utils.ErrUserDoesNotExist
	}

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

	if user_id_p == 0 {
		return nil, utils.ErrPhotoNotFound
	}

	// Check if who's asking is banned or no
	if !db.isBanned(user_id_p, userid) {
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
	} else {
		defer func() { _ = rows.Close() }()
		return stream_comments, utils.ErrBanned
	}

}
