package repository

import (
	"database/sql"
	"errors"
	"log"

	"github.com/janapc/game-api-go/entity"
)

func InsertGame(db *sql.DB, game *entity.Game) error {
	stmt, err := db.Prepare("insert into games(id, name, description, release_date, created_at) values(?,?,?,?,?)")
	if err != nil {
		log.Println(err)
		return errors.New("internal server error")
	}
	defer stmt.Close()
	_, err = stmt.Exec(game.ID, game.Name, game.Description, game.ReleaseDate, game.CreatedAt)
	if err != nil {
		log.Println(err)
		return errors.New("internal server error")
	}
	return nil
}

func SelectAllGames(db *sql.DB) ([]entity.Game, error) {
	rows, err := db.Query("select id, name, description, release_date, created_at from games")
	if err != nil {
		log.Println(err)
		return nil, errors.New("internal server error")
	}
	defer rows.Close()
	var games []entity.Game
	for rows.Next() {
		var game entity.Game
		err := rows.Scan(&game.ID, &game.Name, &game.Description, &game.ReleaseDate, &game.CreatedAt)
		if err != nil {
			log.Println(err)
			return nil, errors.New("internal server error")
		}
		games = append(games, game)
	}
	if len(games) == 0 {
		games = []entity.Game{}
	}
	return games, nil
}

func FindGameById(db *sql.DB, id string) (*entity.Game, error) {
	stmt, err := db.Prepare("select id, name, description, release_date, created_at from games where id = ?")
	if err != nil {
		log.Println(err)
		return nil, errors.New("internal server error")
	}
	defer stmt.Close()
	var game entity.Game
	err = stmt.QueryRow(id).Scan(&game.ID, &game.Name, &game.Description, &game.ReleaseDate, &game.CreatedAt)
	if err != nil {
		log.Println(err)
		return nil, errors.New("register not found")
	}
	return &game, nil
}

func DeleteGameById(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from games where id = ?")
	if err != nil {
		log.Println(err)
		return errors.New("internal server error")
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		log.Println(err)
		return errors.New("internal server error")
	}
	return nil
}

func UpdateGameById(db *sql.DB, game *entity.Game) error {
	stmt, err := db.Prepare("update games set name = ?, description = ?, release_date  = ? where id = ?")
	if err != nil {
		log.Println(err)
		return errors.New("internal server error")
	}
	defer stmt.Close()
	_, err = stmt.Exec(game.Name, game.Description, game.ReleaseDate, game.ID)
	if err != nil {
		log.Println(err)
		return errors.New("internal server error")
	}
	return nil
}
