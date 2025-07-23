package main

import (
	"main/internal/routes"
	"net/http"

	"github.com/urfave/negroni/v3"
)

func main() {
  mux := http.NewServeMux()
  routes.SearchRoutes(mux)

  n := negroni.Classic() // Includes some default middlewares
  n.UseHandler(mux)
  println("Server is running on port 3000")
  http.ListenAndServe(":3000", n)
}