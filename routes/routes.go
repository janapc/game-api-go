package routes

import (
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/janapc/game-api-go/controllers"
	"github.com/janapc/game-api-go/utils"
)

func RouteGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	method := r.Method
	switch method {
	case "POST":
		controllers.SaveGame(w, r)
	case "GET":
		controllers.ListGames(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write(utils.HandleResponseJSON("route is not found"))
	}
}

func RouteGameByParam(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id := getUuidByUrl(r)
	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write(utils.HandleResponseJSON("id invalid"))
		return
	}
	method := r.Method
	switch method {
	case "DELETE":
		controllers.RemoveGame(w, r, id)
	case "GET":
		controllers.FindGame(w, r, id)
	case "PATCH":
		controllers.UpdateGame(w, r, id)
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write(utils.HandleResponseJSON("route is not found"))
	}
}

func getUuidByUrl(r *http.Request) string {
	id := strings.TrimPrefix(r.URL.Path, "/game/")
	_, err := uuid.Parse(id)
	if err != nil {
		return ""
	}
	return id
}
