package entities

type Mini struct {
	AccessToken string
	ExpiresIn   int64
	AppID       string
	AppSecret   string
}

func (m *Mini) FindAccessTokenByAppID() (*Mini, error) {
	mini := &Mini{}
	err := db.QueryRow("SELECT access_token, expires_in, app_id, secret FROM mini WHERE app_id = $1", m.AppID).Scan(&mini.AccessToken, &mini.ExpiresIn, &mini.AppID, &mini.AppSecret)
	return mini, err
}

func (m *Mini) UpdateAccessTokenAndExpiresIn() (err error) {
	_, err = db.Exec("UPDATE mini SET access_token = $1, expires_in = $2 WHERE app_id = $3", m.AccessToken, m.ExpiresIn, m.AppID)
	return
}
