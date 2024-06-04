package internal

import (
	"context"
	"database/sql"
	"github.com/onsi/gomega/gstruct/errors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"log"
	"nba-players-statistics/config"
	"nba-players-statistics/internal/models"
	"nba-players-statistics/pkg"
	"reflect"
)

type StatisticsDao interface {
	Disconnect()
	Log(ctx context.Context, data pkg.LogRequest) error
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

	dbModels := []any{
		&models.Player{},
		&models.PlayerLog{},
		&models.Team{},
		&models.TeamLog{},
		&models.Game{},
		&models.Season{},
	}
	var errs = errors.AggregateError{}
	for _, model := range dbModels {
		if res, err := createTable(ctx, db, model); err != nil {
			logger.Printf("error occurred trying to create the db model `%s`. Response: `%v`. Error: `%s`",
				reflect.TypeOf(model).Name(), res, err.Error())
			errs = append(errs, err)
		}
	}

	if len(errs) != 0 {
		return nil, errs
	}

	logger.Printf("connected to psql db successfully!")
	return &PsqlStatisticsDao{
		logger: logger,
		db:     db,
	}, nil
}

func createTable[T any](ctx context.Context, db *bun.DB, model T) (sql.Result, error) {
	return db.NewCreateTable().IfNotExists().Model(model).Exec(ctx)
}

func (p *PsqlStatisticsDao) Log(ctx context.Context, data pkg.LogRequest) error {
	panic("not implemented")
}

func (p *PsqlStatisticsDao) Disconnect() {
	if err := p.db.Close(); err != nil {
		p.logger.Printf("error occurred trying to close connection to psql db. err: `%v`", err)
		return
	}
	p.logger.Printf("closed psql db connection successfully")
}
