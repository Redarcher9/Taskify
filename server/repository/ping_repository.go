package repository

import (
	"context"
	"database/sql"
	"taskify/domain"
)

type PingRepository struct {
	db      *sql.DB
	timeout int
}

func NewPingRepository(database *sql.DB, timeout int) domain.PingRepository {
	return &PingRepository{
		db:      database,
		timeout: timeout,
	}
}

func (pr *PingRepository) CheckPing(c context.Context) error {
	err := pr.db.Ping()
	if err != nil {
		return err
	}
	return nil
}
