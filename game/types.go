package game

type UpdatePosPackage struct {
	UUID string `json:"uuid"`
	Left bool `json:"left"`
	Right bool `json:"right"`
	Top bool `json:"top"`
	Bottom bool `json:"bottom"`
}

type UpdateClientPackage struct {
	Alive bool	`json:"alive"`
	Map SyncMap `json:"map"`
}