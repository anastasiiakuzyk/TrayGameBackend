package game

import (
	"errors"
	"sync"
)

type Players struct {
	mp sync.Map
}

func NewPlayers() *Players {
	return &Players{sync.Map{}}
}

func (p *Players) Append(player *Player) {
	p.mp.Store(player.UUID, player)
}
func (p *Players) Get(uuid string) (*Player, error) {
	in, ok := p.mp.Load(uuid)
	if !ok {
		return nil, errors.New("Player didn't exists")
	}

	return in.(*Player), nil
}
func (p *Players) Remove(uuid string) {
	p.mp.Delete(uuid)
}

func (p *Players) Update(uuid string, player *Player) {
	p.Remove(uuid)
	p.Append(player)
}

func (p *Players) ForEach(callback func(uuid string, player *Player) bool) {
	p.mp.Range(func(key, value interface{}) bool {
		player := value.(*Player)
		uuid := key.(string)
		return callback(uuid, player)
	})
}

func (p *Players) ToArray() []*Player {
	var players []*Player
	p.ForEach(func(uuid string, player *Player) bool {
		players = append(players, player)

		return true
	})

	return players
}