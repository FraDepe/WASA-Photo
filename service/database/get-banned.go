package database

func (db *appdbimpl) GetBanned(userIdBanned uint64, userId uint64) (Ban, error) {

	var ban Ban
	ban.UserId = 0
	ban.BannedId = 0

	rows, err := db.c.Query(`SELECT userid, bannedid FROM bans WHERE userid=? and bannedid=?`, userId, userIdBanned)
	if err != nil {
		return ban, err
	}

	for rows.Next() {
		err = rows.Scan(&ban.UserId, &ban.BannedId)
		if err != nil {
			return ban, err
		}
	}

	err = rows.Err()
	if err != nil {
		return ban, err
	}

	defer func() { _ = rows.Close() }()
	return ban, nil

}
