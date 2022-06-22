package controllers

type StartRequest struct {
	AuthDate   uint   `json:"auth_date"`
	FirstName  string `json:"first_name,omitempty"`
	SecondName string `json:"second_name,omitempty"`
	Hash       string `json:"hash"`
	Id         uint   `json:"id"`
	PhotoUrl   string `json:"photo_url"`
	Username   string `json:"username"`
}

type UserResponse struct {
	Id       uint   `json:"id"`
	Nickname string `json:"nickname"`
	PhotoUrl string `json:"photoUrl"`
}

type AnalyticsResponse struct {
	Id          uint `json:"id"`
	UserId      uint `json:"userId"`
	TotalKills  uint `json:"totalKills"`
	TotalDeath  uint `json:"totalDeath"`
	GamesPlayed uint `json:"gamesPlayed"`
}
