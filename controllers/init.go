package controllers

import (
	"database/sql"
)

var (
	dbPool *sql.DB
	DBPool *sql.DB
	UserID int64 = 0
)

func Init(db *sql.DB) {
	dbPool = db
	DBPool = db
}
