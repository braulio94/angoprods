package main

import (
	"github.com/braulio94/angoprods/core"
	"github.com/rs/cors"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	app := routes(core.Launch())
	handler := cors.Default().Handler(app.Router)
	_ = http.ListenAndServe(port, handler)
}
