package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/janapc/game-api-go/database"
	"github.com/janapc/game-api-go/entity"
	"github.com/janapc/game-api-go/repository"
	"github.com/janapc/game-api-go/utils"
)

func SaveGame(w http.ResponseWriter, r *http.Request) {
	var g entity.GameProps
	err := json.NewDecoder(r.Body).Decode(&g)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.HandleResponseJSON("the body should have name, description and release_date"))
		return
	}
	game := entity.NewGame(g)
	err = game.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.HandleResponseJSON(err.Error()))
		return
	}
	err = repository.InsertGame(database.Database, game)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func ListGames(w http.ResponseWriter, r *http.Request) {
	games, err := repository.SelectAllGames(database.Database)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	result, err := json.Marshal(games)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func RemoveGame(w http.ResponseWriter, r *http.Request, id string) {
	_, err := repository.FindGameById(database.Database, id)
	if err != nil {
		if err.Error() == "register not found" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = repository.DeleteGameById(database.Database, id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func FindGame(w http.ResponseWriter, r *http.Request, id string) {
	game, err := repository.FindGameById(database.Database, id)
	if err != nil {
		if err.Error() == "register not found" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	result, err := json.Marshal(game)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func UpdateGame(w http.ResponseWriter, r *http.Request, id string) {
	var input entity.GameProps
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.HandleResponseJSON("the body should have name, description and release_date"))
		return
	}
	game, err := repository.FindGameById(database.Database, id)
	if err != nil {
		if err.Error() == "register not found" {
			w.WriteHeader(http.StatusNotFound)
			w.Write(utils.HandleResponseJSON(err.Error()))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	game.MergeFieldsGame(input)
	err = game.Validate()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(utils.HandleResponseJSON(err.Error()))
		return
	}
	repository.UpdateGameById(database.Database, game)
	w.WriteHeader(http.StatusNoContent)
}
