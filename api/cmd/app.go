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

  http.ListenAndServe(":3000", n)
}