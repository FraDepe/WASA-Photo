package database

func (db *appdbimpl) ListComments(id string) ([]Comment error) {
	var stream_comments []Comment

	rows, err := db.c.Query(`SELECT id, text, userid FROM comments WHERE photoid=?`, id) //Con il cambio del db per i commenti andrà aggiunto nella select photoid
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var comment Comment
		err = rows.Scan(&comment.ID, &comment.Text, &comment.UserId) //Con il cambio del db per i commenti andrà aggiunto nella select photoid

		if err != nil {
			return nil, err
		}
		stream_comments = append(stream_comments, comment)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	defer func() { _ = rows.Close() }()
	
	return stream_comments, nil
}
