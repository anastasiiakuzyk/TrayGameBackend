package game

import (
	"github.com/gorilla/websocket"
	"sync"
)

type Player struct {
	sync.RWMutex
	UUID       string         `json:"uuid"`
	Conn       *WebsocketConn `json:"-"`
	Nickname   string         `json:"nickname"`
	Alive      bool           `json:"alive"`
	Effects    []string       `json:"effects"`
	Kills      uint           `json:"kills"`
	KeyBoard   *KeyBoard      `json:"-"`
	Position   *Position      `json:"position"`
	Direction  uint           `json:"-"`
	Trajectory []Position     `json:"trajectory"`
}

type WebsocketConn struct {
	websocket.Conn
	sync.RWMutex
}

func (conn *WebsocketConn) WriteMutableJSON(v interface{}) error {
	conn.Lock()
	defer conn.Unlock()
	return conn.WriteJSON(v)
}

type KeyBoard struct {
	Top    bool
	Left   bool
	Right  bool
	Bottom bool
}

func NewKeyboard() *KeyBoard {
	return &KeyBoard{
		Top:    false,
		Left:   false,
		Right:  false,
		Bottom: false,
	}
}

type Position struct {
	X uint `json:"x"`
	Y uint `json:"y"`
}
