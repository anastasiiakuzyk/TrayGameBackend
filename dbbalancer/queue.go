package dbbalancer

import "sync"

type queue struct {
	start    *queueItem
	lastItem *queueItem
	itemNum  int
	mu       sync.Mutex
}

func NewDBQueue() queue {
	return queue{
		start:    nil,
		lastItem: nil,
		itemNum:  0,
		mu:       sync.Mutex{},
	}
}

func (q *queue) Get(inx int) *queueItem {
	if inx > q.itemNum {
		return nil
	}
	q.mu.Lock()
	defer q.mu.Unlock()

	curItem := q.start
	for i := 0; i < q.itemNum; i++ {
		if i == inx {
			break
		}

		curItem = curItem.next
	}

	return curItem
}

func (q *queue) Add(item queueItem) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.itemNum == 0 {
		q.start = &item
		q.lastItem = &item
	} else {
		q.lastItem.next = &item
	}

	q.itemNum++
}

func (q *queue) Remove(inx int) {
	if q.itemNum == 0 {
		return
	}
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.itemNum == 1 {
		q.start = nil
		q.lastItem = nil
		q.itemNum = 0
	} else {
		curItem := q.start
		var prevItem *queueItem = nil
		for i := 0; i < q.itemNum; i++ {
			if i == inx {
				if prevItem == nil {
					q.start = nil
				}
				prevItem.next = curItem.next
				q.itemNum--
				return
			}

			prevItem = curItem
			curItem = curItem.next
		}
	}
}

func (q *queue) GetAll() []*queueItem {
	q.mu.Lock()
	defer q.mu.Unlock()

	var res []*queueItem
	curItem := q.start
	for i := 0; i < q.itemNum; i++ {
		res = append(res, curItem)
		curItem = curItem.next
	}

	return res
}

func (q *queue) GetAllAndClear() []*queueItem {
	q.mu.Lock()
	defer q.mu.Unlock()

	var res []*queueItem
	curItem := q.start
	for i := 0; i < q.itemNum; i++ {
		res = append(res, curItem)
		if curItem.next == nil {
			q.itemNum = i + 1
			break
		}
		curItem = curItem.next
	}

	q.start = nil
	q.lastItem = nil
	q.itemNum = 0

	return res
}

func (q *queue) RemoveAll() {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.start = nil
	q.lastItem = nil
	q.itemNum = 0
}

type QueueItemActionType string

const (
	QueueItemActionKill  QueueItemActionType = "kill"
	QueueItemActionDeath QueueItemActionType = "death"
	QueueItemActionGame  QueueItemActionType = "game"
)

type queueItem struct {
	next     *queueItem
	action   QueueItemActionType
	gameUUID string
}

func NewQueueItem(action QueueItemActionType, gameUUID string) queueItem {
	return queueItem{
		next:     nil,
		action:   action,
		gameUUID: gameUUID,
	}
}
