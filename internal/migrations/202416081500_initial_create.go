package migrations

import (
	"context"
	"database/sql"
	"log"
	"os"
	"reflect"

	"nba-players-statistics/internal/models"

	"github.com/onsi/gomega/gstruct/errors"

	"github.com/uptrace/bun"
)

var logger = log.New(os.Stdout, "nba-player-statistics-migrations ", log.LstdFlags|log.Lshortfile)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		logger.Print(" [init migration up] ")
		dbModels := []any{
			&models.Game{},
			&models.Player{},
			&models.PlayerLog{},
			&models.Season{},
			&models.Team{},
			&models.TeamLog{},
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
			return errs
		}

		return nil
	}, func(ctx context.Context, db *bun.DB) error {
		logger.Print(" [init migration down] ")
		dbModels := []any{
			&models.Game{},
			&models.Player{},
			&models.PlayerLog{},
			&models.Season{},
			&models.Team{},
			&models.TeamLog{},
		}
		var errs = errors.AggregateError{}
		for _, model := range dbModels {
			if res, err := dropTable(ctx, db, model); err != nil {
				logger.Printf("error occurred trying to drop the db model `%s`. Response: `%v`. Error: `%s`",
					reflect.TypeOf(model).Name(), res, err.Error())
				errs = append(errs, err)
			}
		}

		if len(errs) != 0 {
			return errs
		}

		return nil
	})
}

func createTable[T any](ctx context.Context, db *bun.DB, model T) (sql.Result, error) {
	return db.NewCreateTable().IfNotExists().Model(model).Exec(ctx)
}

func dropTable[T any](ctx context.Context, db *bun.DB, model T) (sql.Result, error) {
	return db.NewDropTable().Model(model).Exec(ctx)
}
