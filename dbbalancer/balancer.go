package dbbalancer

import (
	"sync"
	"time"
	"untitledGameBackend/psql"
)

type balancer struct {
	q  queue
	mu sync.Mutex
}

var (
	b    *balancer
	once sync.Once
)

func GetBalancer() *balancer {
	if b == nil {
		once.Do(func() {
			b = newBalancer()
		})
	}
	b.mu.Lock()
	defer b.mu.Unlock()

	return b
}

func newBalancer() *balancer {
	return &balancer{q: NewDBQueue(), mu: sync.Mutex{}}
}

func (ba *balancer) Init() {
	ti := time.NewTicker(time.Second)

	for _ = range ti.C {
		items := GetBalancer().GetQueue().GetAllAndClear()
		if len(items) != 0 {
			for _, items := range items {
				user, err := psql.GetInstance().GetUserByGameUUID(items.gameUUID)
				if err != nil {
					continue
				}

				if items.action == QueueItemActionKill {
					err := user.AddKill()
					if err != nil {
						continue
					}
				}
				if items.action == QueueItemActionDeath {
					err := user.AddDeath()
					if err != nil {
						continue
					}
				}
				if items.action == QueueItemActionGame {
					err := user.NewGame()
					if err != nil {
						continue
					}
				}
			}
		}
	}
}

func (ba *balancer) GetQueue() *queue {
	ba.mu.Lock()
	defer ba.mu.Unlock()

	return &ba.q
}
