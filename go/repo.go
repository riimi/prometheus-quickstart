package main

import (
	"database/sql"
	"errors"
)

type GreetingRepo interface {
	SelectGreetingsLimit10() ([]*Greeting, error)
	InsertGreeting(g Greeting) error
}

type Greeting struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type GreetingRepoMySQL struct {
	db *sql.DB
}

var (
	ErrNil = errors.New("nil")
)

func (repo *GreetingRepoMySQL) SelectGreetingsLimit10() ([]*Greeting, error) {
	if repo.db == nil {
		return nil, ErrNil
	}
	greetings := make([]*Greeting, 0, 10)

	rows, err := repo.db.Query("SELECT * FROM greetings LIMIT 10")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		g := &Greeting{}
		err := rows.Scan(&g.ID, &g.Text)
		if err != nil {
			return nil, err
		}
		greetings = append(greetings, g)
	}

	return greetings, nil
}

func (repo *GreetingRepoMySQL) InsertGreeting(g Greeting) error {
	if repo.db == nil {
		return ErrNil
	}

	_, err := repo.db.Exec("INSERT INTO greetings (text) VALUES (?)", g.Text)
	if err != nil {
		return err
	}

	return nil
}
