package models

import "github.com/uptrace/bun"

type PlayerLog struct {
	bun.BaseModel `bun:"table:players_logs,alias:pl"`

	PlayerId      int     `bun:",pk,notnull"`
	GameNumber    uint8   `bun:",pk,notnull"`
	SeasonId      int     `bun:",pk,notnull"`
	Player        *Player `bun:"rel:belongs-to,join:player_id=id"`
	Game          *Game   `bun:"rel:belongs-to,join:game_number=number,join:season_id=season_id"`
	Points        uint8
	Rebounds      uint8
	Assists       uint8
	Steals        uint8
	Blocks        uint8
	Turnovers     uint8
	Fouls         uint8
	MinutesPlayed float32
}
