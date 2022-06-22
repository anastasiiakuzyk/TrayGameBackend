package game

import "math/rand"

var types = []string{"booster", "saver"}

func GetRandomHelperType() string {
	return types[rand.Intn(1)]
}

type Helper struct {
	Position Position `json:"position"`
	Type string `json:"type"`
	UUID string `json:"uuid"`
}