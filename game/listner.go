package game

import (
	"github.com/gorilla/websocket"
	"net/http"
	"untitledGameBackend/psql"
	"untitledGameBackend/utils"
)

var (
	upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}
)

func Listen(response http.ResponseWriter, request *http.Request, engine *Engine) {
	response.Header().Set("Access-Control-Allow-Origin", "*")
	uuid := request.URL.Query().Get("uuid")

	if uuid == "" {
		badRequestBytes, _ := utils.NewErrorResponse(http.StatusBadRequest, "Bad request").MarshalJSON()
		_, err := response.Write(badRequestBytes)
		if err != nil {
			return
		}
		return
	}

	user, err := psql.GetInstance().GetUserByGameUUID(uuid)
	if err != nil {
		forbiddenBytes, _ := utils.NewErrorResponse(http.StatusForbidden, "Forbidden").MarshalJSON()
		_, err := response.Write(forbiddenBytes)
		if err != nil {
			return
		}
		return
	}

	conn, err := upgrader.Upgrade(response, request, nil)
	if err != nil {
		return
	}

	engine.AddPlayer(uuid, user.GetNickname())

	defer func() {
		engine.Map.Players.Remove(uuid)
		err := conn.Close()
		if err != nil {
			return
		}
	}()

	err = engine.ConnectPlayer(conn, uuid)
	if err != nil {
		err := conn.Close()
		if err != nil {
			return
		}
		return
	}
	update := UpdateClientPackage{
		Alive: true,
		Map:   engine.Map.ToSyncMap(),
	}

	err = conn.WriteJSON(update)
	if err != nil {
		return
	}

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		var pack UpdatePosPackage
		err = pack.UnmarshalJSON(msg)
		if err != nil {
			continue
		}

		if pack.UUID != uuid {
			continue
		}

		engine.UpdatePlayerPos(pack)
	}
}
