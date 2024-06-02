package internal

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"log"
	"nba-players-statistics/config"
	"nba-players-statistics/internal/models"
)

type StatisticsDao interface {
	Disconnect()
}

type PsqlStatisticsDao struct {
	logger *log.Logger
	db     *pg.DB
}

func (p *PsqlStatisticsDao) Disconnect() {
	if err := p.db.Close(); err != nil {
		p.logger.Printf("error occurred trying to close connection to psql db. err: `%v`", err)
		return
	}
	p.logger.Printf("closed psql db connection successfully")
}

func NewPsqlStatisticsDao(config *config.DbConfig, logger *log.Logger) (*PsqlStatisticsDao, error) {
	logger.Printf("trying to connect to psql db with user %s", config.User)
	db := pg.Connect(&pg.Options{
		User:     config.User,
		Password: config.Password,
	})

	var p *models.Player
	if err := db.Model(p).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	}); err != nil {
		return nil, err
	}

	var t *models.Team
	if err := db.Model(t).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	}); err != nil {
		return nil, err
	}

	logger.Printf("connected to psql db successfully!")
	return &PsqlStatisticsDao{
		logger: logger,
		db:     db,
	}, nil
}
