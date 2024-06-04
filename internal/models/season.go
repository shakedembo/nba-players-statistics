package models

import "github.com/uptrace/bun"

type Season struct {
	bun.BaseModel `bun:"table:seasons,alias:s"`
	StartYear     uint16 `bun:",pk"`
	EndYear       uint16 `bun:",pk"`
}
