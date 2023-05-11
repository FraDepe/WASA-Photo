package database

// Check if the users to follow exists
func (db *appdbimpl) existence(userId uint64) bool {
	rows, err := db.c.Query(`SELECT id FROM users WHERE id=?`, userId)
	if err != nil {
		return false
	}
	var tmp uint64
	for rows.Next() {
		err = rows.Scan(&tmp)
		if err != nil {
			return false
		}
	}
	err = rows.Err()
	if err != nil {
		return false
	}
	if tmp == 0 {
		defer func() { _ = rows.Close() }()
		return false
	}

	defer func() { _ = rows.Close() }()
	return true
}

// Check if the user is banned
func (db *appdbimpl) isBanned(userId uint64, loggedId uint64) bool {
	rows, err := db.c.Query(`SELECT userid FROM bans WHERE userid=? and bannedid=?`, userId, loggedId)
	if err != nil {
		return false
	}
	var exist []uint64
	for rows.Next() {
		var id uint64
		err = rows.Scan(&id)
		if err != nil {
			return false
		}
		exist = append(exist, id)
	}

	err = rows.Err()
	if err != nil {
		return false
	}

	defer func() { _ = rows.Close() }()

	if len(exist) == 0 {
		return false
	} else {
		return true
	}
}

// Check if the user is following
func (db *appdbimpl) isFollowing(userId uint64, loggedId uint64) bool {
	rows, err := db.c.Query(`SELECT followerid FROM follows WHERE followedid=? and followerid=?`, userId, loggedId)
	if err != nil {
		return false
	}
	var exist []uint64
	for rows.Next() {
		var id uint64
		err = rows.Scan(&id)
		if err != nil {
			return false
		}
		exist = append(exist, id)
	}

	err = rows.Err()
	if err != nil {
		return false
	}

	defer func() { _ = rows.Close() }()

	if len(exist) == 0 {
		return false
	} else {
		return true
	}
}
