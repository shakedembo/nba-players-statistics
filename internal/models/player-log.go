package models

import "github.com/uptrace/bun"

type PlayerLog struct {
	bun.BaseModel `bun:"table:players_logs,alias:pl"`
	Id            int     `bun:",pk,autoincrement"`
	Player        *Player `bun:"rel:has-one,join:id=team_id"`
	Game          *Game   `bun:"rel:has-one,join:id=id"`
	Points        uint8
	Rebounds      uint8
	Assists       uint8
	Steals        uint8
	Blocks        uint8
	Turnovers     uint8
	Fouls         uint8
	MinutesPlayed float32
}
