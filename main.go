package main

import (
	"log"
	"net/http"
	"untitledGameBackend/controllers"
	"untitledGameBackend/dbbalancer"
	"untitledGameBackend/game"
)

func main() {
	log.SetFlags(0)

	engine := game.NewDefaultEngine()
	go engine.Start()
	go dbbalancer.GetBalancer().Init()

	http.HandleFunc("/ws", func(writer http.ResponseWriter, request *http.Request) {
		game.Listen(writer, request, engine)
	})

	http.HandleFunc("/start", func(writer http.ResponseWriter, request *http.Request) {
		controllers.StartHandler(writer, request, engine)
	})

	http.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request) {
		controllers.GetUserHandler(writer, request)
	})

	http.HandleFunc("/stats", func(writer http.ResponseWriter, request *http.Request) {
		controllers.GetStatsHandler(writer, request)
	})

	http.Handle("/", http.FileServer(http.Dir("./public")))

	errors := make(chan error)
	go func() {
		errors <- http.ListenAndServe(":8080", nil)
	}()

	log.Println("Server started on port: :8080")

	for err := range errors {
		log.Fatal(err)
	}
}
