package entity

import (
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
)

type Game struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ReleaseDate string `json:"release_date"`
	CreatedAt   string `json:"created_at"`
}

type GameProps struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ReleaseDate string `json:"release_date"`
}

const createdAtFormat = "2006-01-02 15:04:05"

func NewGame(game GameProps) *Game {
	return &Game{
		ID:          uuid.New().String(),
		Name:        game.Name,
		Description: game.Description,
		ReleaseDate: game.ReleaseDate,
		CreatedAt:   time.Now().Format(createdAtFormat),
	}
}

func (game *Game) Validate() error {
	if game.Name == "" || game.Description == "" || game.ReleaseDate == "" {
		return errors.New("the fields name, description and release_date is mandatory")
	}
	match, err := regexp.MatchString(`\d{2}\/\d{2}\/\d{4}`, game.ReleaseDate)
	if !match || err != nil {
		return errors.New("the field release_date should is this format DD/MM/YYYY")
	}
	return nil
}

func (game *Game) MergeFieldsGame(input GameProps) {
	if input.Name != "" {
		game.Name = input.Name
	}
	if input.Description != "" {
		game.Description = input.Description
	}
	if input.ReleaseDate != "" {
		game.ReleaseDate = input.ReleaseDate
	}
}
