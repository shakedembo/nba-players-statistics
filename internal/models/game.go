package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Game struct {
	bun.BaseModel `bun:"table:games,alias:g"`

	Id   int `bun:",pk,autoincrement"`
	Date time.Time
}
