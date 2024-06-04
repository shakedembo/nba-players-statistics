package models

import "github.com/uptrace/bun"

type TeamLog struct {
	bun.BaseModel `bun:"table:teams_logs,alias:tl"`
	Id            int  `bun:",pk,autoincrement"`
	Team          Team `bun:"rel:has-one,join:id=id"`
	Game          Game `bun:"rel:has-one,join:id=id"`
	Points        uint8
	Rebounds      uint8
	Assists       uint8
	Steals        uint8
	Blocks        uint8
	Turnovers     uint8
	Fouls         uint8
}
