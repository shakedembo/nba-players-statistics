package models

import "github.com/uptrace/bun"

type Player struct {
	bun.BaseModel `bun:"table:players,alias:p"`

	Id     int    `bun:",pk,autoincrement"`
	Name   string `bun:",notnull"`
	TeamId int    `bun:",notnull"`
	Team   *Team  `bun:"rel:belongs-to,join:team_id=id"`
}
