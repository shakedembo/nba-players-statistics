package models

import "github.com/uptrace/bun"

type Season struct {
	bun.BaseModel `bun:"table:seasons,alias:s"`

	Id    int     `bun:",pk,autoincrement"`
	Games []*Game `bun:"rel:has-many,join:id=season_id"`
}
