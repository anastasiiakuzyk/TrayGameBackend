package controllers

import (
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"net/http"
	"untitledGameBackend/game"
	"untitledGameBackend/psql"
	"untitledGameBackend/utils"
)

func StartHandler(res http.ResponseWriter, req *http.Request, engine *game.Engine) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Headers", "*")

	if req.Method == http.MethodOptions {
		res.WriteHeader(http.StatusOK)
		return
	}
	res.Header().Set("Access-Control-Allow-Origin", "*")
	if req.Method != http.MethodPost {
		methodNotAllowedResponse, _ := utils.NewErrorResponse(http.StatusMethodNotAllowed, "Method not allowed").MarshalJSON()
		_, err := res.Write(methodNotAllowedResponse)
		if err != nil {
			return
		}
	}
	var user StartRequest
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		badRequestResponse, _ := utils.NewErrorResponse(http.StatusBadRequest, "Bad request").MarshalJSON()
		_, err := res.Write(badRequestResponse)
		if err != nil {
			return
		}
		return
	}

	var uuidStringed string
	dbUser, err := psql.GetInstance().GetUserByStartRequest(user.Id)
	if err != nil {
		log.Println("Error on get existing user" + err.Error())
		uuidStringed = uuid.New().String()
		err = psql.GetInstance().CreateUser(user.AuthDate, user.FirstName, user.SecondName, user.Hash, user.Id, user.PhotoUrl, user.Username, uuidStringed)
		if err != nil {
			log.Println("Error on create user" + err.Error())
			internalErrorResponse, _ := utils.NewErrorResponse(http.StatusInternalServerError, "On create user: Internal server error"+err.Error()).MarshalJSON()
			_, err := res.Write(internalErrorResponse)
			if err != nil {
				return
			}
			return
		}
		dbUser, err = psql.GetInstance().GetUserByGameUUID(uuidStringed)
		if err != nil {
			log.Println("Error on get created user" + err.Error())
			internalErrorResponse, _ := utils.NewErrorResponse(http.StatusInternalServerError, "Internal server error").MarshalJSON()
			_, err := res.Write(internalErrorResponse)
			if err != nil {
				return
			}
			return
		}
		err = psql.GetInstance().CreateAnalyticsByUserId(dbUser.Id)
		if err != nil {
			log.Println("Error on create analytics" + err.Error())
			internalErrorResponse, _ := utils.NewErrorResponse(http.StatusInternalServerError, "Internal server error").MarshalJSON()
			_, err := res.Write(internalErrorResponse)
			if err != nil {
				return
			}
			return
		}
	} else {
		uuidStringed = dbUser.GameUUID
		_, err := psql.GetInstance().GetAnalyticsByUserId(dbUser.Id)
		if err != nil {
			log.Println("Error on get analytics for existing user" + err.Error())

			err := psql.GetInstance().CreateAnalyticsByUserId(dbUser.Id)
			if err != nil {
				log.Println("Error on create analytics for existing user" + err.Error())
				internalErrorResponse, _ := utils.NewErrorResponse(http.StatusInternalServerError, "Internal server error").MarshalJSON()
				_, err := res.Write(internalErrorResponse)
				if err != nil {
					return
				}
				return
			}
		}
	}

	statusOkResponse, err := utils.NewStartResponse(http.StatusOK, uuidStringed).MarshalJSON()
	if err != nil {
		log.Println("Error on response generation" + err.Error())

		internalServerErrorResponse, _ := utils.NewErrorResponse(http.StatusInternalServerError, "Internal Server Error").MarshalJSON()
		res.WriteHeader(http.StatusInternalServerError)
		_, err := res.Write(internalServerErrorResponse)
		if err != nil {
			return
		}
		return
	}
	_, err = res.Write(statusOkResponse)
	if err != nil {
		log.Println("Error on response write" + err.Error())

		internalServerErrorResponse, _ := utils.NewErrorResponse(http.StatusInternalServerError, "Internal Server Error").MarshalJSON()
		res.WriteHeader(http.StatusInternalServerError)
		_, err := res.Write(internalServerErrorResponse)
		if err != nil {
			return
		}
		return
	}
}
