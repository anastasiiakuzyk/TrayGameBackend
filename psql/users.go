package psql

import (
	"context"
	"strings"
)

type User struct {
	AuthDate   uint   `json:"auth_date"`
	FirstName  string `json:"first_name,omitempty"`
	SecondName string `json:"second_name,omitempty"`
	Hash       string `json:"-"`
	Id         uint   `json:"id"`
	TelegramId uint   `json:"telegramId"`
	GameUUID   string `json:"game_uuid"`
	PhotoUrl   string `json:"photo_url"`
	Username   string `json:"username"`
}

func (u *User) GetNickname() string {
	if strings.Trim(u.Username, " ") != "" {
		return u.Username
	}
	if strings.Trim(u.FirstName, " ") != "" && strings.Trim(u.SecondName, " ") != "" {
		return u.FirstName + " " + u.SecondName
	}
	if strings.Trim(u.FirstName, " ") != "" {
		return u.FirstName
	}
	if strings.Trim(u.SecondName, " ") != "" {
		return u.SecondName
	}

	return "Unnamed User"
}

func (u *User) AddKill() error {
	return GetInstance().AddKill(u.Id)
}

func (u *User) AddDeath() error {
	return GetInstance().AddDeath(u.Id)
}

func (u *User) NewGame() error {
	return GetInstance().AddGamesPlayed(u.Id)
}

func (cont *PGXContainer) GetUserByStartRequest(telegramId uint) (*User, error) {
	cont.mu.Lock()
	defer cont.mu.Unlock()
	row := cont.conn.QueryRow(context.Background(), "SELECT * FROM users WHERE telegram_id=$1", telegramId)
	var user User
	err := row.Scan(&user.Id, &user.Username, &user.FirstName, &user.SecondName, &user.Hash, &user.TelegramId, &user.AuthDate, &user.PhotoUrl, &user.GameUUID)
	return &user, err
}

func (cont *PGXContainer) GetUserByGameUUID(gameUUID string) (*User, error) {
	cont.mu.Lock()
	defer cont.mu.Unlock()
	row := cont.conn.QueryRow(context.Background(), "SELECT * FROM users WHERE game_uuid=$1", gameUUID)
	var user User
	err := row.Scan(&user.Id, &user.Username, &user.FirstName, &user.SecondName, &user.Hash, &user.TelegramId, &user.AuthDate, &user.PhotoUrl, &user.GameUUID)
	return &user, err
}

func (cont *PGXContainer) CreateUser(authDate uint, firstName string, secondName string, hash string, telegramId uint, photoUrl string, username string, gameUUID string) error {
	cont.mu.Lock()
	defer cont.mu.Unlock()

	_, err := cont.conn.Query(context.Background(), "insert into users(username, first_name, second_name, telegram_hash, telegram_id, auth_date, photo_url, game_uuid) VALUES($1, $2, $3, $4, $5, $6, $7, $8);",
		username, firstName, secondName, hash, telegramId, authDate, photoUrl, gameUUID)

	return err
}
