package main

import (
	"fmt"
	"net/http"

	"github.com/EzraChai/goapi/internal/handlers"
	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting Go API service")

	fmt.Println(`
+------------------------------------------------------------------------------+
|  Welcome to Go API service.                                                  |
|  Video source: https://youtu.be/8uiZC0l4Ajw?si=nczuIACnXNHBhmdt              |
+------------------------------------------------------------------------------+
	`)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Error(err)
	}
}
