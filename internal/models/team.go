package models

import "github.com/uptrace/bun"

type Team struct {
	bun.BaseModel `bun:"table:teams,alias:t"`

	Id      int `bun:",pk,autoincrement"`
	Name    string
	Players []*Player `bun:"rel:has-many,join:id=team_id"`
}
