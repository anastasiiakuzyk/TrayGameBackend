package psql

import "context"

type Analytics struct {
	Id          uint `json:"id"`
	UserId      uint `json:"userId"`
	TotalKills  uint `json:"totalKills"`
	TotalDeath  uint `json:"totalDeath"`
	GamesPlayed uint `json:"gamesPlayed"`
}

func (cont *PGXContainer) GetAnalyticsByUserId(userId uint) (*Analytics, error) {
	cont.mu.Lock()
	defer cont.mu.Unlock()
	row := cont.conn.QueryRow(context.Background(), "SELECT * FROM analytics WHERE user_id=$1", userId)
	var analytics Analytics
	err := row.Scan(&analytics.Id, &analytics.UserId, &analytics.TotalKills, &analytics.TotalDeath, &analytics.GamesPlayed)
	return &analytics, err
}

func (cont *PGXContainer) CreateAnalyticsByUserId(userId uint) error {
	cont.mu.Lock()
	defer cont.mu.Unlock()
	_, err := cont.conn.Query(context.Background(), "insert into analytics(user_id, total_kills, total_death, games_played) VALUES($1, $2, $3, $4);", userId, 0, 0, 0)
	return err
}

func (cont *PGXContainer) AddKill(userId uint) error {
	cont.mu.Lock()
	defer cont.mu.Unlock()
	_, err := cont.conn.Exec(context.Background(), "UPDATE analytics SET total_kills=total_kills+1 WHERE user_id=$1", userId)
	return err
}

func (cont *PGXContainer) AddDeath(userId uint) error {
	cont.mu.Lock()
	defer cont.mu.Unlock()
	_, err := cont.conn.Exec(context.Background(), "UPDATE analytics SET total_death=total_death+1 WHERE user_id=$1", userId)
	return err
}

func (cont *PGXContainer) AddGamesPlayed(userId uint) error {
	cont.mu.Lock()
	defer cont.mu.Unlock()
	_, err := cont.conn.Exec(context.Background(), "UPDATE analytics SET games_played=games_played+1 WHERE user_id=$1", userId)
	return err
}
