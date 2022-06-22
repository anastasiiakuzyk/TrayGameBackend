package game

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"math/rand"
	"sync"
	"time"
	"untitledGameBackend/dbbalancer"
)

type Engine struct {
	sync.RWMutex
	Map                 *Map
	addPlayerChan       chan *Player
	removePlayerChan    chan string
	updatePlayerPosChan chan *UpdatePosPackage
	errChan             chan error
}

func NewDefaultEngine() *Engine {
	return &Engine{Map: NewDefaultMap(), addPlayerChan: make(chan *Player, 10), removePlayerChan: make(chan string, 10), updatePlayerPosChan: make(chan *UpdatePosPackage, 1000), errChan: make(chan error, 10)}
}

func (e *Engine) ConnectPlayer(conn *websocket.Conn, uuid string) error {
	player, err := e.Map.Players.Get(uuid)
	if err != nil {
		return err
	}

	player.Conn = &WebsocketConn{
		Conn:    *conn,
		RWMutex: sync.RWMutex{},
	}

	return nil
}

type SyncEngine struct {
	Map SyncMap
}

func (e *Engine) ToSyncStruct() SyncEngine {
	return SyncEngine{Map: e.Map.ToSyncMap()}
}

func (e *Engine) AddPlayer(uuid string, nickname string) string {
	p := &Player{
		UUID:     uuid,
		Nickname: nickname,
		Alive:    true,
		Kills:    0,
		KeyBoard: NewKeyboard(),
		Position: &Position{
			X: uint(rand.Intn(int(e.Map.Width))),
			Y: uint(rand.Intn(int(e.Map.Height))),
		},
		Direction: 0,
	}
	p.Trajectory = append(p.Trajectory, *p.Position)
	e.Map.Players.Append(p)
	dbbalancer.GetBalancer().GetQueue().Add(dbbalancer.NewQueueItem(dbbalancer.QueueItemActionGame, uuid))

	return p.UUID
}

func (e *Engine) GamePlay() {
	ticker := time.NewTicker(time.Second / 128)
	helpersTicker := time.NewTicker(time.Second * 15)
	helpersCleanTicker := time.NewTicker(time.Minute)

	for {
		select {
		case <-ticker.C:
			e.Lock()
			e.Map.Players.ForEach(func(uuid string, player *Player) bool {
				if player.Trajectory[0].X != player.Position.X || player.Trajectory[0].Y != player.Position.Y {
					player.Trajectory = append(player.Trajectory, *player.Position)
					if len(player.Trajectory) > 120 {
						player.Trajectory = player.Trajectory[len(player.Trajectory)-120:]
					}
				}

				const defaultSpeed = 3
				var speed = uint(defaultSpeed)
				for _, eff := range player.Effects {
					if eff == "booster" {
						speed *= 2
						break
					}
				}

				if player.KeyBoard.Left && !player.KeyBoard.Right {
					player.Position.X -= speed
				} else if !player.KeyBoard.Left && player.KeyBoard.Right {
					player.Position.X += speed
				}
				if player.KeyBoard.Top && !player.KeyBoard.Bottom {
					player.Position.Y -= speed
				} else if !player.KeyBoard.Top && player.KeyBoard.Bottom {
					player.Position.Y += speed
				}

				if (player.Position.X > e.Map.Width) && player.KeyBoard.Right {
					player.Position.X = 1
				}
				if (player.Position.X < speed) && player.KeyBoard.Left {
					player.Position.X = e.Map.Width
				}

				if (player.Position.Y > e.Map.Height) && player.KeyBoard.Bottom {
					player.Position.Y = 1
				}
				if (player.Position.Y < speed) && player.KeyBoard.Top {
					player.Position.Y = e.Map.Height
				}

				e.Map.Helpers.ForEach(func(uuid string, Helper *Helper) bool {
					if int(player.Position.X)+13 >= int(Helper.Position.X) && int(player.Position.X)-13 <= int(Helper.Position.X) && int(player.Position.Y)+13 >= int(Helper.Position.Y) && int(player.Position.Y)-13 <= int(Helper.Position.Y) {
						player.Effects = append(player.Effects, Helper.Type)
						go func() {
							time.Sleep(5 * time.Second)
							player.Effects = []string{}
						}()
						e.Map.Helpers.Remove(uuid)
					}
					return true
				})

				e.Map.Players.Update(player.UUID, player)

				e.Map.Players.ForEach(func(uuid2 string, p2 *Player) bool {
					if uuid != uuid2 {
						if p2.Alive && player.Alive {
							for i := 0; i < len(p2.Trajectory); i++ {
								if int(player.Position.X)+13 >= int(p2.Trajectory[i].X) && int(player.Position.X)-13 <= int(p2.Trajectory[i].X) && int(player.Position.Y)+13 >= int(p2.Trajectory[i].Y) && int(player.Position.Y)-13 <= int(p2.Trajectory[i].Y) {
									player.Alive = false
									e.Map.Players.Update(player.UUID, player)
									fmt.Println("Killed", player.Nickname)
									p2.Kills += 1
									e.Map.Players.Update(p2.UUID, p2)
									dbbalancer.GetBalancer().GetQueue().Add(dbbalancer.NewQueueItem(dbbalancer.QueueItemActionDeath, player.UUID))
									dbbalancer.GetBalancer().GetQueue().Add(dbbalancer.NewQueueItem(dbbalancer.QueueItemActionKill, p2.UUID))

									return false
								}
							}
						}
					}
					return true
				})

				return true
			})
			e.Unlock()
		case <-helpersTicker.C:
			e.Lock()
			if len(e.Map.Players.ToArray()) != 0 {
				helper := &Helper{
					Position: Position{
						X: uint(rand.Intn(int(e.Map.Width))),
						Y: uint(rand.Intn(int(e.Map.Height))),
					},
					Type: GetRandomHelperType(),
					UUID: uuid.New().String(),
				}

				e.Map.Helpers.Append(helper)
			}
			e.Unlock()
		case <-helpersCleanTicker.C:

		}
	}
}

func (e *Engine) Start() {
	go e.GamePlay()

	ticker := time.NewTicker(time.Second / 64)
	ticker3 := time.NewTicker(time.Second * 60)
	packs := 0
	for {
		select {
		case <-ticker.C:
			e.Lock()
			e.Map.Players.ForEach(func(uuid string, player *Player) bool {
				if player.Conn == nil {
					return true
				}

				update := UpdateClientPackage{
					Alive: player.Alive,
					Map:   e.Map.ToSyncMap(),
				}

				player.Lock()
				err := player.Conn.WriteMutableJSON(update)
				player.Unlock()
				if err != nil {
					e.errChan <- err
				}

				if !player.Alive {
					err := player.Conn.Close()
					if err != nil {
						return false
					}
				}

				return true
			})
			packs++
			e.Unlock()
		case <-ticker3.C:
			fmt.Println(time.Now().String(), " Avg tick of last 60 seconds ", packs/60)
			packs = 0
		}
	}
}

func (e *Engine) UpdatePlayerPos(pos UpdatePosPackage) {
	e.Lock()
	player, err := e.Map.Players.Get(pos.UUID)
	if err != nil {
		e.errChan <- err
	}
	player.KeyBoard.Top = pos.Top
	player.KeyBoard.Bottom = pos.Bottom
	player.KeyBoard.Right = pos.Right
	player.KeyBoard.Left = pos.Left

	e.Map.Players.Update(player.UUID, player)
	e.Unlock()
}
