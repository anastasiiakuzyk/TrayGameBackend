package controllers

import (
	"net/http"
	"untitledGameBackend/psql"
	"untitledGameBackend/utils"
)

func GetUserHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "*")

	if req.Method == http.MethodOptions {
		res.WriteHeader(http.StatusOK)
		return
	}
	res.Header().Set("Access-Control-Allow-Origin", "*")
	if req.Method != http.MethodGet {
		methodNotAllowedResponse, _ := utils.NewErrorResponse(http.StatusMethodNotAllowed, "Method not allowed").MarshalJSON()
		_, err := res.Write(methodNotAllowedResponse)
		if err != nil {
			return
		}
	}

	gameUUID := req.URL.Query().Get("uuid")
	if len(gameUUID) == 0 {
		badRequestResponse, _ := utils.NewErrorResponse(http.StatusBadRequest, "Bad request").MarshalJSON()
		res.WriteHeader(http.StatusBadRequest)
		_, err := res.Write(badRequestResponse)
		if err != nil {
			return
		}
		return
	}

	user, err := psql.GetInstance().GetUserByGameUUID(gameUUID)
	if err != nil {
		internalServerErrorResponse, _ := utils.NewErrorResponse(http.StatusInternalServerError, "Internal Server Error").MarshalJSON()
		res.WriteHeader(http.StatusInternalServerError)
		_, err := res.Write(internalServerErrorResponse)
		if err != nil {
			return
		}
		return
	}

	userResponse := &UserResponse{
		Id:       user.Id,
		Nickname: user.GetNickname(),
		PhotoUrl: user.PhotoUrl,
	}
	userResponseBytes, err := userResponse.MarshalJSON()
	if err != nil {
		internalServerErrorResponse, _ := utils.NewErrorResponse(http.StatusInternalServerError, "Internal Server Error").MarshalJSON()
		res.WriteHeader(http.StatusInternalServerError)
		_, err := res.Write(internalServerErrorResponse)
		if err != nil {
			return
		}
		return
	}

	_, err = res.Write(userResponseBytes)
	if err != nil {
		internalServerErrorResponse, _ := utils.NewErrorResponse(http.StatusInternalServerError, "Internal Server Error").MarshalJSON()
		res.WriteHeader(http.StatusInternalServerError)
		_, err := res.Write(internalServerErrorResponse)
		if err != nil {
			return
		}
		return
	}
}
