package repository

import "database/sql"

type GlobalRepo struct {
	db *sql.DB
}

func NewGlobalRepo(db *sql.DB) *GlobalRepo {
	return &GlobalRepo{db: db}
}
