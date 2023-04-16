package database

func (db *appdbimpl) BanUser(logged_user uint64, user_id uint64) (User, error) {

	var user User

	// Check if the guy is already banned
	rows, err := db.c.Query(`SELECT userid FROM bans WHERE userid=? and bannedid=?`, logged_user, user_id)
	if err != nil {
		return user, nil
	}
	var exist []uint64
	for rows.Next() {
		var id uint64
		err = rows.Scan(&id)
		exist = append(exist, id)
		if err != nil {
			return user, nil
		}
	}

	// If exist array is empty the guy is not already banned so he can ban him
	if len(exist) == 0 {
		_, err = db.c.Exec(`INSERT INTO bans (userid bannedid) VALUES (?, ?)`,
			logged_user, user_id)

		if err != nil {
			return user, err
		}
	}

	defer func() { _ = rows.Close() }()
	return user, nil
}