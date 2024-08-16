package internal

import (
	"context"
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"log"
	"nba-players-statistics/config"
	"nba-players-statistics/internal/migrations"
	"nba-players-statistics/internal/models"
	"nba-players-statistics/pkg"
)

type StatisticsDao interface {
	Disconnect()
	Log(ctx context.Context, data pkg.LogRequest) error
	CreateSeason() error
	GetCurrentSeason() (*models.Season, error)
}

type PsqlStatisticsDao struct {
	logger *log.Logger
	db     *bun.DB
}

func NewPsqlStatisticsDao(ctx context.Context, config *config.DbConfig, logger *log.Logger) (*PsqlStatisticsDao, error) {
	logger.Printf("trying to connect to psql db with user %s", config.User)
	dsn := "postgres://postgres:postgres@localhost:5432/?sslmode=disable"
	sqlDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqlDb, pgdialect.New())

	err := migrations.Initialize(ctx, db)
	if err != nil {
		return nil, err
	}

	logger.Printf("connected to psql db successfully!")
	return &PsqlStatisticsDao{
		logger: logger,
		db:     db,
	}, nil
}

func (p *PsqlStatisticsDao) Disconnect() {
	if err := p.db.Close(); err != nil {
		p.logger.Printf("error occurred trying to close connection to psql db. err: `%v`", err)
		return
	}
	p.logger.Printf("closed psql db connection successfully")
}

func (p *PsqlStatisticsDao) Log(ctx context.Context, data pkg.LogRequest) error {
	panic("not implemented")
}

func (p *PsqlStatisticsDao) CreateSeason() error {
	//TODO implement me
	panic("implement me")
}

func (p *PsqlStatisticsDao) GetCurrentSeason() (*models.Season, error) {
	//TODO implement me
	panic("implement me")
}
