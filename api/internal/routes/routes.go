package routes

import (
	"main/internal/controllers"
	"net/http"
)


func SearchRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", controllers.FetchBook())
}