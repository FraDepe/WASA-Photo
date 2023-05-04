package database

func (db *appdbimpl) GetComment(commentId uint64) (Comment, error) {

	var comment Comment

	rows, err := db.c.Query(`SELECT id, photoid, text, userid FROM comments WHERE id=?`, commentId)
	if err != nil {
		return comment, err
	}

	for rows.Next() {
		err = rows.Scan(&comment.ID, &comment.PhotoId, &comment.Text, &comment.UserId)
		if err != nil {
			return comment, err
		}
	}

	err = rows.Err()
	if err != nil {
		return comment, err
	}

	defer func() { _ = rows.Close() }()
	return comment, nil

}
