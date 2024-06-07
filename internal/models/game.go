package models

import (
	"github.com/uptrace/bun"
)

type Game struct {
	bun.BaseModel `bun:"table:games,alias:g"`

	Number   uint8   `bun:",pk"`
	SeasonId int     `bun:",pk,notnull"`
	Season   *Season `bun:"rel:belongs-to,join:season_id=id"`
}
